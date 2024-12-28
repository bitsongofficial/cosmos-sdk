package lockup

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/header"
	"cosmossdk.io/math"
	lockupaccount "cosmossdk.io/x/accounts/defaults/lockup"
	types "cosmossdk.io/x/accounts/defaults/lockup/v1"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *IntegrationTestSuite) TestContinuousLockingAccount() {
	t := s.T()
	app := setupApp(t)
	currentTime := time.Now()
	ctx := sdk.NewContext(app.CommitMultiStore(), false, app.Logger()).WithHeaderInfo(header.Info{
		Time: currentTime,
	})
	s.setupStakingParams(ctx, app)
	ownerAddrStr, err := app.AuthKeeper.AddressCodec().BytesToString(accOwner)
	require.NoError(t, err)
	s.fundAccount(app, ctx, accOwner, sdk.Coins{sdk.NewCoin("stake", math.NewInt(1000000))})
	randAcc := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())

	_, accountAddr, err := app.AccountsKeeper.Init(ctx, lockupaccount.CONTINUOUS_LOCKING_ACCOUNT, accOwner, &types.MsgInitLockupAccount{
		Owner:     ownerAddrStr,
		StartTime: currentTime,
		// end time in 1 minutes
		EndTime: currentTime.Add(time.Minute),
	}, sdk.Coins{sdk.NewCoin("stake", math.NewInt(1000))}, nil)
	require.NoError(t, err)

	addr, err := app.AuthKeeper.AddressCodec().BytesToString(randAcc)
	require.NoError(t, err)

	vals, err := app.StakingKeeper.GetAllValidators(ctx)
	require.NoError(t, err)
	val := vals[0]

	t.Run("error - execute message, wrong sender", func(t *testing.T) {
		msg := &types.MsgSend{
			Sender:    addr,
			ToAddress: addr,
			Amount:    sdk.Coins{sdk.NewCoin("stake", math.NewInt(100))},
		}
		err := s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NotNil(t, err)
	})
	t.Run("error - execute send message, insufficient fund", func(t *testing.T) {
		msg := &types.MsgSend{
			Sender:    ownerAddrStr,
			ToAddress: addr,
			Amount:    sdk.Coins{sdk.NewCoin("stake", math.NewInt(100))},
		}
		err := s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NotNil(t, err)
	})

	// Update context time
	// 12 sec = 1/5 of a minute so 200stake should be released
	ctx = ctx.WithHeaderInfo(header.Info{
		Time: currentTime.Add(time.Second * 12),
	})

	// Check if token is sendable
	t.Run("ok - execute send message", func(t *testing.T) {
		msg := &types.MsgSend{
			Sender:    ownerAddrStr,
			ToAddress: addr,
			Amount:    sdk.Coins{sdk.NewCoin("stake", math.NewInt(100))},
		}
		err := s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NoError(t, err)

		balance := app.BankKeeper.GetBalance(ctx, randAcc, "stake")
		require.True(t, balance.Amount.Equal(math.NewInt(100)))
	})

	t.Run("ok - execute delegate message", func(t *testing.T) {
		msg := &types.MsgDelegate{
			Sender:           ownerAddrStr,
			ValidatorAddress: val.OperatorAddress,
			Amount:           sdk.NewCoin("stake", math.NewInt(100)),
		}
		err = s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NoError(t, err)

		valbz, err := app.StakingKeeper.ValidatorAddressCodec().StringToBytes(val.OperatorAddress)
		require.NoError(t, err)

		del, err := app.StakingKeeper.Delegations.Get(
			ctx, collections.Join(sdk.AccAddress(accountAddr), sdk.ValAddress(valbz)),
		)
		require.NoError(t, err)
		require.NotNil(t, del)

		// check if tracking is updated accordingly
		lockupAccountInfoResponse := s.queryLockupAccInfo(ctx, app, accountAddr)
		delLocking := lockupAccountInfoResponse.DelegatedLocking
		require.True(t, delLocking.AmountOf("stake").Equal(math.NewInt(100)))
	})
	t.Run("ok - execute withdraw reward message", func(t *testing.T) {
		msg := &types.MsgWithdrawReward{
			Sender:           ownerAddrStr,
			ValidatorAddress: val.OperatorAddress,
		}
		err = s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NoError(t, err)
	})
	t.Run("ok - execute undelegate message", func(t *testing.T) {
		vals, err := app.StakingKeeper.GetAllValidators(ctx)
		require.NoError(t, err)
		val := vals[0]
		msg := &types.MsgUndelegate{
			Sender:           ownerAddrStr,
			ValidatorAddress: val.OperatorAddress,
			Amount:           sdk.NewCoin("stake", math.NewInt(100)),
		}
		err = s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NoError(t, err)
		valbz, err := app.StakingKeeper.ValidatorAddressCodec().StringToBytes(val.OperatorAddress)
		require.NoError(t, err)

		ubd, err := app.StakingKeeper.GetUnbondingDelegation(
			ctx, sdk.AccAddress(accountAddr), sdk.ValAddress(valbz),
		)
		require.NoError(t, err)
		require.Equal(t, len(ubd.Entries), 1)

		// check if an entry is added
		unbondingEntriesResponse := s.queryUnbondingEntries(ctx, app, accountAddr, val.OperatorAddress)
		entries := unbondingEntriesResponse.UnbondingEntries
		require.True(t, entries[0].Amount.Amount.Equal(math.NewInt(100)))
		require.True(t, entries[0].ValidatorAddress == val.OperatorAddress)
	})

	// Update context time to end time
	ctx = ctx.WithHeaderInfo(header.Info{
		Time: currentTime.Add(time.Minute),
	})

	// trigger endblock for staking to handle matured unbonding delegation
	_, err = app.StakingKeeper.EndBlocker(ctx)
	require.NoError(t, err)

	// test if tracking delegate work perfectly
	t.Run("ok - execute delegate message", func(t *testing.T) {
		msg := &types.MsgDelegate{
			Sender:           ownerAddrStr,
			ValidatorAddress: val.OperatorAddress,
			Amount:           sdk.NewCoin("stake", math.NewInt(100)),
		}
		err = s.executeTx(ctx, msg, app, accountAddr, accOwner)
		require.NoError(t, err)

		valbz, err := app.StakingKeeper.ValidatorAddressCodec().StringToBytes(val.OperatorAddress)
		require.NoError(t, err)

		del, err := app.StakingKeeper.Delegations.Get(
			ctx, collections.Join(sdk.AccAddress(accountAddr), sdk.ValAddress(valbz)),
		)
		require.NoError(t, err)
		require.NotNil(t, del)

		// check if tracking is updated accordingly
		lockupAccountInfoResponse := s.queryLockupAccInfo(ctx, app, accountAddr)
		delLocking := lockupAccountInfoResponse.DelegatedLocking
		// should be update as ubd entry is matured
		require.True(t, delLocking.AmountOf("stake").Equal(math.ZeroInt()))
		delFree := lockupAccountInfoResponse.DelegatedFree
		require.True(t, delFree.AmountOf("stake").Equal(math.NewInt(100)))

		// check if the entry is removed
		unbondingEntriesResponse := s.queryUnbondingEntries(ctx, app, accountAddr, val.OperatorAddress)
		entries := unbondingEntriesResponse.UnbondingEntries
		require.Len(t, entries, 0)
	})
}
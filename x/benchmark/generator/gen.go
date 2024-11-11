package gen

import (
	"fmt"
	"math/rand/v2"

	"github.com/cespare/xxhash/v2"

	"cosmossdk.io/x/benchmark"
)

type Options struct {
	Seed uint64

	KeyMean     uint64
	KeyStdDev   uint64
	ValueMean   uint64
	ValueStdDev uint64

	DeleteFraction float64

	BucketCount uint64
}

type State struct {
	Keys map[int]map[int]struct{}
}

type Generator struct {
	Options

	rand *rand.Rand
}

func NewGenerator(opts Options) *Generator {
	return &Generator{
		Options: opts,
		rand:    rand.New(rand.NewPCG(opts.Seed, opts.Seed>>32)),
	}
}

func (g *Generator) Next() (*benchmark.Op, uint64) {
	op := &benchmark.Op{
		Seed:        g.rand.Uint64(),
		KeyLength:   g.NormUint64(g.KeyMean, g.KeyStdDev),
		ValueLength: g.NormUint64(g.ValueMean, g.ValueStdDev),
	}
	if op.KeyLength == 0 {
		op.KeyLength = 1
	}

	return op, g.rand.Uint64N(g.BucketCount)
}

func (g *Generator) Bytes(seed, length uint64) []byte {
	b := make([]byte, length)
	rounds := length / 8
	remainder := length % 8
	var h uint64
	for i := uint64(0); i < rounds; i++ {
		h = xxhash.Sum64(encodeUint64(seed + i))
		for j := uint64(0); j < 8; j++ {
			b[i*8+j] = byte(h >> (8 * j))
		}
	}
	if remainder > 0 {
		h = xxhash.Sum64(encodeUint64(seed + rounds))
		for j := uint64(0); j < remainder; j++ {
			b[rounds*8+j] = byte(h >> (8 * j))
		}
	}
	return b
}

func (g *Generator) NormUint64(mean, stdDev uint64) uint64 {
	return uint64(g.rand.NormFloat64()*float64(stdDev) + float64(mean))
}

func (g *Generator) Key() []byte {
	length := g.NormUint64(g.KeyMean, g.KeyStdDev)
	if length == 0 {
		length = 1
	}
	seed := g.rand.Uint64()
	return g.Bytes(seed, length)
}

func (g *Generator) Value() []byte {
	length := g.NormUint64(g.ValueMean, g.ValueStdDev)
	if length == 0 {
		length = 1
	}
	seed := g.rand.Uint64()
	return g.Bytes(seed, length)
}

func (g *Generator) UintN(n uint64) uint64 {
	return g.rand.Uint64N(n)
}

func encodeUint64(x uint64) []byte {
	var b [8]byte
	b[0] = byte(x)
	b[1] = byte(x >> 8)
	b[2] = byte(x >> 16)
	b[3] = byte(x >> 24)
	b[4] = byte(x >> 32)
	b[5] = byte(x >> 40)
	b[6] = byte(x >> 48)
	b[7] = byte(x >> 56)
	return b[:]
}

const maxStoreKeyGenIterations = 100

func StoreKeys(prefix string, seed, count uint64) ([]string, error) {
	g := NewGenerator(Options{})
	r := rand.New(rand.NewPCG(seed, seed>>32))
	keys := make([]string, count)
	seen := make(map[string]struct{})

	var i, j uint64
	for i < count {
		if j > maxStoreKeyGenIterations {
			return nil, fmt.Errorf("failed to generate %d unique store keys", count)
		}
		sk := fmt.Sprintf("%s_%x", prefix, g.Bytes(r.Uint64(), 8))
		if _, ok := seen[sk]; ok {
			j++
			continue
		}
		keys[i] = sk
		seen[sk] = struct{}{}
		i++
	}
	return keys, nil
}
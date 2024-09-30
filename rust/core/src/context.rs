use crate::message::Message;
use crate::response::{Response, ResponseBody};
use bump_scope::{Bump, BumpBox};
use ixc_message_api::handler::{HandlerErrorCode, HostBackend};
use ixc_message_api::header::{MessageHeader, MESSAGE_HEADER_SIZE};
use ixc_message_api::packet::MessagePacket;
use ixc_message_api::AccountID;
use ixc_message_api::code::ErrorCode;
use ixc_schema::codec::Codec;
use ixc_schema::mem::MemoryManager;
use ixc_schema::value::ResponseValue;
use crate::error::Error;

/// Context wraps a single message request (and possibly response as well) along with
/// the router callbacks necessary for making nested message calls.
pub struct Context<'a> {
    message_packet: &'a mut MessagePacket,
    host_callbacks: &'a dyn HostBackend,
    memory_manager: &'a MemoryManager<'a, 'a>,
}

impl<'a> Context<'a> {
    /// Create a new context from a message packet and host callbacks.
    pub fn new(message_packet: &'a mut MessagePacket, host_callbacks: &'a dyn HostBackend, memory_manager: &'a MemoryManager<'a, 'a>) -> Self {
        Self {
            message_packet,
            host_callbacks,
            memory_manager,
        }
    }

    /// This is the address of the account that is getting called.
    /// In a receiving account, this is the account's own address.
    pub fn account_id(&self) -> AccountID {
        self.message_packet.header().account
    }

    /// This is the address of the account which is making the message call.
    pub fn caller(&self) -> AccountID {
        self.message_packet.header().sender_account
    }

    // /// Returns a new response with the given value.
    // pub fn ok<R: ResponseValue, E: ResponseValue>(&self, res: <R as ResponseValue>::MaybeBorrowed<'a>) -> Response<'a, R, E> {
    //     Ok(res.to_owned())
    // }

    // /// Dynamically invokes a module message.
    // /// Static module client instances should be preferred wherever possible,
    // /// so that static dependency analysis can be performed.
    // pub fn dynamic_invoke_module<'b, M: Message<'b, true>>(&self, message: M) -> Response<M::Response, M::Error>
    // {
    //     unimplemented!()
    // }

    /// Dynamically invokes an account message.
    /// Static account client instances should be preferred wherever possible,
    /// so that static dependency analysis can be performed.
    pub unsafe fn dynamic_invoke<'b, M: Message<'b, false>>(&mut self, account: AccountID, message: M) -> Response<M::Response, M::Error> {
        // create a new bump scope and memory manager
        // let mut guard = self.memory_manager.scope().scope_guard();
        // let scope = guard.scope();
        let bump = Bump::new();
        let scope = bump.as_scope();
        let mem_mgr = MemoryManager::new(scope);

        // encode the message body
        let msg_body = M::Codec::encode_value(&message, mem_mgr.scope()).
            map_err(|_| Error::KnownHandlerError(HandlerErrorCode::EncodingError))?;

        // create the message header and fill in call details
        let mut header = scope.alloc_default::<MessageHeader>();
        header.sender_account = self.message_packet.header().account;
        header.account = account;
        header.in_pointer1.set_slice(msg_body);
        header.message_selector = M::SELECTOR;

        // package the header in a packet
        let header_ptr = header.into_mut();
        let header_ptr: *mut MessageHeader = header_ptr;
        let mut packet = unsafe { MessagePacket::new(header_ptr, MESSAGE_HEADER_SIZE) };

        // invoke the message
        let res = self.host_callbacks.invoke(&mut packet);

        // parse the response data
        // TODO how can we have some simple case for () responses and errors?
        match res {
            Ok(_) => {
                // let out_data = packet.header().out_pointer1.get(&packet);
                // let res = <M::Response as ResponseValue>::decode_value::<M::Codec>(out_data, &mem_mgr).
                //     map_err(|_| Error::KnownHandlerError(HandlerErrorCode::EncodingError))?;
                // Ok(ResponseBody::new(bump, packet, res))
                todo!()
            }
            Err(_) => {
                todo!()
            }
        }
    }

    /// Dynamically invokes a message that does not modify state.
    pub fn dynamic_invoke_readonly<'b, M: Message<'b, false>>(&self, account: &AccountID, message: M) -> Response<M::Response, M::Error> {
        todo!()
    }

    /// Dynamically invokes a message that does not read or write state.
    pub fn dynamic_invoke_pure<'b, M: Message<'b, false>>(&self, account: &AccountID, message: M) -> Response<M::Response, M::Error> {
        todo!()
    }

    /// Get the host backend.
    pub unsafe fn get_host_backend(&self) -> &dyn HostBackend {
        self.host_callbacks
    }

    // /// Get the address of the module implementing the given trait, client type or module message, if any.
    // pub fn get_module_address<T: ModuleAPI>(&self) -> Response<Address> {
    //     unimplemented!()
    // }
    //
    // /// Create a new account with the given initialization data.
    // pub fn new_account<H: AccountHandler>(&mut self, init: H::Init) -> Result<<<H as AccountAPI>::ClientFactory as AccountClientFactory>::Client, ()> {
    //     unimplemented!()
    // }
    //
    // /// Create a temporary account with the given initialization data.
    // /// Its address will be empty from the perspective of all observers,
    // /// and it will not be persisted.
    // pub fn new_temp_account<H: AccountHandler>(&mut self, init: H::Init) -> Result<<<H as AccountAPI>::ClientFactory as AccountClientFactory>::Client, ()> {
    //     unimplemented!()
    // }
    //
    // /// Returns a deterministic ID unique to the message call which this context pertains to.
    // /// Such IDs can be used to generate unique IDs for state objects.
    // /// The index parameter can be used to generate up to 256 such unique IDs per message call.
    // pub fn unique_id(&mut self, index: u8) -> Result<u128, ()> {
    //     unimplemented!()
    // }
}

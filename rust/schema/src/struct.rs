use crate::decoder::{DecodeError, Decoder};

/// StructCodec is the trait that should be derived to encode and decode a struct.
///
/// It should generally be used in conjunction with the `#[derive(StructCodec)]` attribute
/// attached to a struct definition.
/// It is unsafe to implement this trait manually, because the compiler cannot
///  guarantee correct implementations.
///
/// Any struct which contains fields which implement [`value::MaybeBorrowed`] can be derived
/// to implement this trait.
/// Structs and their fields may optionally contain a single lifetime parameter, in which
/// case decoded values will be borrowed from the input data wherever possible.
///
/// Example:
/// ```
/// #[derive(StructCodec)]
/// pub struct MyStruct<'a> {
///   pub field1: u8,
///   pub field2: &'a str,
/// }
///
///
/// #[derive(StructCodec)]
/// pub struct MyStruct2 {
///   pub field1: simple_time::Time,
///   pub field2: interchain_message_api::Address,
/// }
/// ```
pub unsafe trait StructCodec {
    /// A dummy function for derived macro type checking.
    fn dummy(&self);
}


pub unsafe trait StructVisitor<'a> {
    fn visit_field<D: Decoder<'a>>(&mut self, index: usize, decoder: &mut D) -> Result<(), DecodeError>;
}
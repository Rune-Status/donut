package message

import (
    "errors"
    "github.com/sprinkle-it/donut/pkg/buffer"
)

var (
    ErrEncodeNotSupported = errors.New("message: encoding is not supported for this message")
)

func Singleton(msg Message) func() Message { return func() Message { return msg } }

type Message interface {
    Config() Config
}

type Inbound interface {
    Message

    // Decodes the message from the buffer. It is expected that the buffer contains the number of bytes required to
    // decode the message, if not then this function will return an error. If this function returns no error then
    // decoding the message was successful. Implementations of this method are expected not to validate the decoded
    // data to any contextual source. Exceptions to this include where there are checks to assure that the message is
    // valid. Any other validation should be done outside of the message by receivers of the message.
    Decode(buf *buffer.ByteBuffer, length int) error
}

type Outbound interface {
    Message

    // Encodes the outbound message to a byte buffer. It is expected that the buffer capacity is enough to contain the
    // entire encoded message, if not then this function will return an error.
    Encode(*buffer.ByteBuffer) error
}
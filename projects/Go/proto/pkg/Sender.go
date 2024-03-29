//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: pkg.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package pkg

import "errors"
import "fbeproj/proto/fbe"
import "fbeproj/proto/osa"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = osa.Version

// Fast Binary Encoding pkg sender
type Sender struct {
    *fbe.Sender
    osaSender *osa.Sender
}

// Create a new pkg sender with an empty buffer
func NewSender() *Sender {
    return NewSenderWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new pkg sender with the given buffer
func NewSenderWithBuffer(buffer *fbe.Buffer) *Sender {
    return &Sender{
        fbe.NewSender(buffer, false),
        osa.NewSenderWithBuffer(buffer),
    }
}

// Imported senders

func (s *Sender) OsaSender() *osa.Sender { return s.osaSender }

// Sender models accessors


// Send methods

func (s *Sender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    default:
        _ = value
        break
    }
    if result, err := s.osaSender.Send(value); (result > 0) || (err != nil) {
        return result, err
    }
    return 0, nil
}

//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: osa.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package osa

import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding osa final sender
type FinalSender struct {
    *fbe.Sender
}

// Create a new osa final sender with an empty buffer
func NewFinalSender() *FinalSender {
    return NewFinalSenderWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new osa final sender with the given buffer
func NewFinalSenderWithBuffer(buffer *fbe.Buffer) *FinalSender {
    return &FinalSender{
        fbe.NewSender(buffer, true),
    }
}

// Sender models accessors


// Send methods

func (s *FinalSender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    default:
        _ = value
        break
    }
    return 0, nil
}

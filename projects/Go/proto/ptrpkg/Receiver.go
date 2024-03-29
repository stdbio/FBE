//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: ptrpkg.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package ptrpkg

import "errors"
import "fbeproj/proto/fbe"
import "fbeproj/proto/variants_ptr"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = variants_ptr.Version

// Fast Binary Encoding ptrpkg receiver
type Receiver struct {
    *fbe.Receiver
    variants_ptrReceiver *variants_ptr.Receiver

}

// Create a new ptrpkg receiver with an empty buffer
func NewReceiver() *Receiver {
    return NewReceiverWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new ptrpkg receiver with the given buffer
func NewReceiverWithBuffer(buffer *fbe.Buffer) *Receiver {
    receiver := &Receiver{
        fbe.NewReceiver(buffer, false),
        variants_ptr.NewReceiverWithBuffer(buffer),
    }
    receiver.SetupHandlerOnReceive(receiver)
    return receiver
}

// Imported receivers

// Get the variants_ptr receiver
func (r *Receiver) Variants_ptrReceiver() *variants_ptr.Receiver { return r.variants_ptrReceiver }
// Set the variants_ptr receiver
func (r *Receiver) SetVariants_ptrReceiver(receiver *variants_ptr.Receiver) { r.variants_ptrReceiver = receiver }

// Setup handlers
func (r *Receiver) SetupHandlers(handlers interface{}) {
    r.Receiver.SetupHandlers(handlers)
    r.variants_ptrReceiver.SetupHandlers(handlers)
}


// Receive message handler
func (r *Receiver) OnReceive(fbeType int, buffer []byte) (bool, error) {
    switch fbeType {
    default:
        _ = fbeType
        break
    }

    if r.variants_ptrReceiver != nil {
        if ok, err := r.variants_ptrReceiver.OnReceive(fbeType, buffer); ok {
            return ok, err
        }
    }

    return false, nil
}

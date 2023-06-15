//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package protoex

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding OrderMessage final model
type FinalModelOrderMessage struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    Body *FinalModelOrder
}

// Create a new OrderMessage final model
func NewFinalModelOrderMessage(buffer *fbe.Buffer, offset int) *FinalModelOrderMessage {
    fbeResult := FinalModelOrderMessage{buffer: buffer, offset: offset}
    fbeResult.Body = NewFinalModelOrder(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelOrderMessage) FBEAllocationSize(fbeValue *OrderMessage) int {
    fbeResult := 0 +
        fm.Body.FBEAllocationSize(&fbeValue.Body) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelOrderMessage) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelOrderMessage) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelOrderMessage) FBEType() int { return 11 }

// Get the final offset
func (fm *FinalModelOrderMessage) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelOrderMessage) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelOrderMessage) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelOrderMessage) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelOrderMessage) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelOrderMessage) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.Body.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.Body.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelOrderMessage) Get() (*OrderMessage, int, error) {
    fbeResult := NewOrderMessage()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelOrderMessage) GetValue(fbeValue *OrderMessage) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelOrderMessage) GetFields(fbeValue *OrderMessage) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.Body.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.Body.GetValue(&fbeValue.Body); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

// Set the struct value
func (fm *FinalModelOrderMessage) Set(fbeValue *OrderMessage) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelOrderMessage) SetFields(fbeValue *OrderMessage) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.Body.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.Body.Set(&fbeValue.Body); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

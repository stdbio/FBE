//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package test

import "errors"
import "../fbe"

// Fast Binary Encoding FlagsEmpty field model
type FieldModelFlagsEmpty struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new FlagsEmpty field model
func NewFieldModelFlagsEmpty(buffer *fbe.Buffer, offset int) *FieldModelFlagsEmpty {
    return &FieldModelFlagsEmpty{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelFlagsEmpty) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelFlagsEmpty) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelFlagsEmpty) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelFlagsEmpty) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelFlagsEmpty) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelFlagsEmpty) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelFlagsEmpty) Verify() bool { return true }

// Get the value
func (fm *FieldModelFlagsEmpty) Get() (*FlagsEmpty, error) {
    var value FlagsEmpty
    return &value, fm.GetValueDefault(&value, FlagsEmpty(0))
}

// Get the value with provided default value
func (fm *FieldModelFlagsEmpty) GetDefault(defaults FlagsEmpty) (*FlagsEmpty, error) {
    var value FlagsEmpty
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelFlagsEmpty) GetValue(value *FlagsEmpty) error {
    return fm.GetValueDefault(value, FlagsEmpty(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelFlagsEmpty) GetValueDefault(value *FlagsEmpty, defaults FlagsEmpty) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = FlagsEmpty(fbe.ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelFlagsEmpty) Set(value *FlagsEmpty) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelFlagsEmpty) SetValue(value FlagsEmpty) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int32(value))
    return nil
}

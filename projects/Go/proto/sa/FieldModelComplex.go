//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: sa.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package sa

import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding Complex field model
type FieldModelComplex struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    Name *fbe.FieldModelString
    Sex *FieldModelOptionalSex
    Flag *FieldModelOptionalMyFLags
    Extra *FieldModelOptionalExtra
    Nums *FieldModelVectorInt64
}

// Create a new Complex field model
func NewFieldModelComplex(buffer *fbe.Buffer, offset int) *FieldModelComplex {
    fbeResult := FieldModelComplex{buffer: buffer, offset: offset}
    fbeResult.Name = fbe.NewFieldModelString(buffer, 4 + 4)
    fbeResult.Sex = NewFieldModelOptionalSex(buffer, fbeResult.Name.FBEOffset() + fbeResult.Name.FBESize())
    fbeResult.Flag = NewFieldModelOptionalMyFLags(buffer, fbeResult.Sex.FBEOffset() + fbeResult.Sex.FBESize())
    fbeResult.Extra = NewFieldModelOptionalExtra(buffer, fbeResult.Flag.FBEOffset() + fbeResult.Flag.FBESize())
    fbeResult.Nums = NewFieldModelVectorInt64(buffer, fbeResult.Extra.FBEOffset() + fbeResult.Extra.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelComplex) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelComplex) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.Name.FBESize() +
        fm.Sex.FBESize() +
        fm.Flag.FBESize() +
        fm.Extra.FBESize() +
        fm.Nums.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelComplex) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.Name.FBEExtra() +
        fm.Sex.FBEExtra() +
        fm.Flag.FBEExtra() +
        fm.Extra.FBEExtra() +
        fm.Nums.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelComplex) FBEType() int { return 3 }

// Get the field offset
func (fm *FieldModelComplex) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelComplex) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelComplex) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelComplex) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelComplex) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelComplex) VerifyType(fbeVerifyType bool) bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return false
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return false
    }

    fbeStructType := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4))
    if fbeVerifyType && (fbeStructType != fm.FBEType()) {
        return false
    }

    fm.buffer.Shift(fbeStructOffset)
    fbeResult := fm.VerifyFields(fbeStructSize)
    fm.buffer.Unshift(fbeStructOffset)
    return fbeResult
}

// // Check if the struct value fields are valid
func (fm *FieldModelComplex) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Name.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Name.Verify() {
        return false
    }
    fbeCurrentSize += fm.Name.FBESize()

    if (fbeCurrentSize + fm.Sex.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Sex.Verify() {
        return false
    }
    fbeCurrentSize += fm.Sex.FBESize()

    if (fbeCurrentSize + fm.Flag.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Flag.Verify() {
        return false
    }
    fbeCurrentSize += fm.Flag.FBESize()

    if (fbeCurrentSize + fm.Extra.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Extra.Verify() {
        return false
    }
    fbeCurrentSize += fm.Extra.FBESize()

    if (fbeCurrentSize + fm.Nums.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Nums.Verify() {
        return false
    }
    fbeCurrentSize += fm.Nums.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelComplex) GetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, nil
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return 0, errors.New("model is broken")
    }

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Get the struct value (end phase)
func (fm *FieldModelComplex) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelComplex) Get() (*Complex, error) {
    fbeResult := NewComplex()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelComplex) GetValue(fbeValue *Complex) error {
    fbeBegin, err := fm.GetBegin()
    if fbeBegin == 0 {
        return err
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset()))
    fm.GetFields(fbeValue, fbeStructSize)
    fm.GetEnd(fbeBegin)
    return nil
}

// Get the struct fields values
func (fm *FieldModelComplex) GetFields(fbeValue *Complex, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Name.FBESize()) <= fbeStructSize {
        fbeValue.Name, _ = fm.Name.Get()
    } else {
        fbeValue.Name = ""
    }
    fbeCurrentSize += fm.Name.FBESize()

    if (fbeCurrentSize + fm.Sex.FBESize()) <= fbeStructSize {
        fbeValue.Sex, _ = fm.Sex.Get()
    } else {
        fbeValue.Sex = nil
    }
    fbeCurrentSize += fm.Sex.FBESize()

    if (fbeCurrentSize + fm.Flag.FBESize()) <= fbeStructSize {
        fbeValue.Flag, _ = fm.Flag.Get()
    } else {
        fbeValue.Flag = nil
    }
    fbeCurrentSize += fm.Flag.FBESize()

    if (fbeCurrentSize + fm.Extra.FBESize()) <= fbeStructSize {
        fbeValue.Extra, _ = fm.Extra.Get()
    } else {
        fbeValue.Extra = nil
    }
    fbeCurrentSize += fm.Extra.FBESize()

    if (fbeCurrentSize + fm.Nums.FBESize()) <= fbeStructSize {
        fbeValue.Nums, _ = fm.Nums.Get()
    } else {
        fbeValue.Nums = make([]int64, 0)
    }
    fbeCurrentSize += fm.Nums.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelComplex) SetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := fm.FBEBody()
    fbeStructOffset := fm.buffer.Allocate(fbeStructSize) - fm.buffer.Offset()
    if (fbeStructOffset <= 0) || ((fm.buffer.Offset() + fbeStructOffset + fbeStructSize) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeStructOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset, uint32(fbeStructSize))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4, uint32(fm.FBEType()))

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Set the struct value (end phase)
func (fm *FieldModelComplex) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelComplex) Set(fbeValue *Complex) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelComplex) SetFields(fbeValue *Complex) error {
    var err error = nil

    if err = fm.Name.Set(fbeValue.Name); err != nil {
        return err
    }
    if err = fm.Sex.Set(fbeValue.Sex); err != nil {
        return err
    }
    if err = fm.Flag.Set(fbeValue.Flag); err != nil {
        return err
    }
    if err = fm.Extra.Set(fbeValue.Extra); err != nil {
        return err
    }
    if err = fm.Nums.Set(fbeValue.Nums); err != nil {
        return err
    }
    return err
}

//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: extra.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

package extra

import "errors"
import "fbeproj/proto/fbe"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version

// Fast Binary Encoding Extra field model
type FieldModelExtra struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    Num *fbe.FieldModelInt64
    Data *fbe.FieldModelString
    Info *FieldModelPtrInfo
    Info2 *FieldModelPtrInfo
    Info3 *FieldModelInfo
    Infov *FieldModelVectorInfo
    Infopv *FieldModelVectorPtrInfo
    Infol *FieldModelVectorInfo
    Infopl *FieldModelVectorPtrInfo
}

// Create a new Extra field model
func NewFieldModelExtra(buffer *fbe.Buffer, offset int) *FieldModelExtra {
    fbeResult := FieldModelExtra{buffer: buffer, offset: offset}
    fbeResult.Num = fbe.NewFieldModelInt64(buffer, 4 + 4)
    fbeResult.Data = fbe.NewFieldModelString(buffer, fbeResult.Num.FBEOffset() + fbeResult.Num.FBESize())
    fbeResult.Info = NewFieldModelPtrInfo(buffer, fbeResult.Data.FBEOffset() + fbeResult.Data.FBESize())
    fbeResult.Info2 = NewFieldModelPtrInfo(buffer, fbeResult.Info.FBEOffset() + fbeResult.Info.FBESize())
    fbeResult.Info3 = NewFieldModelInfo(buffer, fbeResult.Info2.FBEOffset() + fbeResult.Info2.FBESize())
    fbeResult.Infov = NewFieldModelVectorInfo(buffer, fbeResult.Info3.FBEOffset() + fbeResult.Info3.FBESize())
    fbeResult.Infopv = NewFieldModelVectorPtrInfo(buffer, fbeResult.Infov.FBEOffset() + fbeResult.Infov.FBESize())
    fbeResult.Infol = NewFieldModelVectorInfo(buffer, fbeResult.Infopv.FBEOffset() + fbeResult.Infopv.FBESize())
    fbeResult.Infopl = NewFieldModelVectorPtrInfo(buffer, fbeResult.Infol.FBEOffset() + fbeResult.Infol.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelExtra) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelExtra) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.Num.FBESize() +
        fm.Data.FBESize() +
        fm.Info.FBESize() +
        fm.Info2.FBESize() +
        fm.Info3.FBESize() +
        fm.Infov.FBESize() +
        fm.Infopv.FBESize() +
        fm.Infol.FBESize() +
        fm.Infopl.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelExtra) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.Num.FBEExtra() +
        fm.Data.FBEExtra() +
        fm.Info.FBEExtra() +
        fm.Info2.FBEExtra() +
        fm.Info3.FBEExtra() +
        fm.Infov.FBEExtra() +
        fm.Infopv.FBEExtra() +
        fm.Infol.FBEExtra() +
        fm.Infopl.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelExtra) FBEType() int { return 2 }

// Get the field offset
func (fm *FieldModelExtra) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelExtra) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelExtra) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelExtra) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelExtra) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelExtra) VerifyType(fbeVerifyType bool) bool {
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
func (fm *FieldModelExtra) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Num.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Num.Verify() {
        return false
    }
    fbeCurrentSize += fm.Num.FBESize()

    if (fbeCurrentSize + fm.Data.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Data.Verify() {
        return false
    }
    fbeCurrentSize += fm.Data.FBESize()

    if (fbeCurrentSize + fm.Info.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Info.Verify() {
        return false
    }
    fbeCurrentSize += fm.Info.FBESize()

    if (fbeCurrentSize + fm.Info2.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Info2.Verify() {
        return false
    }
    fbeCurrentSize += fm.Info2.FBESize()

    if (fbeCurrentSize + fm.Info3.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Info3.Verify() {
        return false
    }
    fbeCurrentSize += fm.Info3.FBESize()

    if (fbeCurrentSize + fm.Infov.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Infov.Verify() {
        return false
    }
    fbeCurrentSize += fm.Infov.FBESize()

    if (fbeCurrentSize + fm.Infopv.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Infopv.Verify() {
        return false
    }
    fbeCurrentSize += fm.Infopv.FBESize()

    if (fbeCurrentSize + fm.Infol.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Infol.Verify() {
        return false
    }
    fbeCurrentSize += fm.Infol.FBESize()

    if (fbeCurrentSize + fm.Infopl.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Infopl.Verify() {
        return false
    }
    fbeCurrentSize += fm.Infopl.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelExtra) GetBegin() (int, error) {
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
func (fm *FieldModelExtra) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelExtra) Get() (*Extra, error) {
    fbeResult := NewExtra()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelExtra) GetValue(fbeValue *Extra) error {
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
func (fm *FieldModelExtra) GetFields(fbeValue *Extra, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Num.FBESize()) <= fbeStructSize {
        fbeValue.Num, _ = fm.Num.Get()
    } else {
        fbeValue.Num = 0
    }
    fbeCurrentSize += fm.Num.FBESize()

    if (fbeCurrentSize + fm.Data.FBESize()) <= fbeStructSize {
        fbeValue.Data, _ = fm.Data.Get()
    } else {
        fbeValue.Data = ""
    }
    fbeCurrentSize += fm.Data.FBESize()

    if (fbeCurrentSize + fm.Info.FBESize()) <= fbeStructSize {
        fbeValue.Info, _ = fm.Info.Get()
    } else {
        fbeValue.Info = nil
    }
    fbeCurrentSize += fm.Info.FBESize()

    if (fbeCurrentSize + fm.Info2.FBESize()) <= fbeStructSize {
        fbeValue.Info2, _ = fm.Info2.Get()
    } else {
        fbeValue.Info2 = nil
    }
    fbeCurrentSize += fm.Info2.FBESize()

    if (fbeCurrentSize + fm.Info3.FBESize()) <= fbeStructSize {
        _ = fm.Info3.GetValue(&fbeValue.Info3)
    } else {
        fbeValue.Info3 = *NewInfo()
    }
    fbeCurrentSize += fm.Info3.FBESize()

    if (fbeCurrentSize + fm.Infov.FBESize()) <= fbeStructSize {
        fbeValue.Infov, _ = fm.Infov.Get()
    } else {
        fbeValue.Infov = make([]Info, 0)
    }
    fbeCurrentSize += fm.Infov.FBESize()

    if (fbeCurrentSize + fm.Infopv.FBESize()) <= fbeStructSize {
        fbeValue.Infopv, _ = fm.Infopv.Get()
    } else {
        fbeValue.Infopv = make([]*Info, 0)
    }
    fbeCurrentSize += fm.Infopv.FBESize()

    if (fbeCurrentSize + fm.Infol.FBESize()) <= fbeStructSize {
        fbeValue.Infol, _ = fm.Infol.Get()
    } else {
        fbeValue.Infol = make([]Info, 0)
    }
    fbeCurrentSize += fm.Infol.FBESize()

    if (fbeCurrentSize + fm.Infopl.FBESize()) <= fbeStructSize {
        fbeValue.Infopl, _ = fm.Infopl.Get()
    } else {
        fbeValue.Infopl = make([]*Info, 0)
    }
    fbeCurrentSize += fm.Infopl.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelExtra) SetBegin() (int, error) {
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
func (fm *FieldModelExtra) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelExtra) Set(fbeValue *Extra) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelExtra) SetFields(fbeValue *Extra) error {
    var err error = nil

    if err = fm.Num.Set(fbeValue.Num); err != nil {
        return err
    }
    if err = fm.Data.Set(fbeValue.Data); err != nil {
        return err
    }
    if err = fm.Info.Set(fbeValue.Info); err != nil {
        return err
    }
    if err = fm.Info2.Set(fbeValue.Info2); err != nil {
        return err
    }
    if err = fm.Info3.Set(&fbeValue.Info3); err != nil {
        return err
    }
    if err = fm.Infov.Set(fbeValue.Infov); err != nil {
        return err
    }
    if err = fm.Infopv.Set(fbeValue.Infopv); err != nil {
        return err
    }
    if err = fm.Infol.Set(fbeValue.Infol); err != nil {
        return err
    }
    if err = fm.Infopl.Set(fbeValue.Infopl); err != nil {
        return err
    }
    return err
}

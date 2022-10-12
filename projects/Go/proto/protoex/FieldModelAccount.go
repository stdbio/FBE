//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

package protoex

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding Account field model
type FieldModelAccount struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    Id *fbe.FieldModelInt32
    Name *fbe.FieldModelString
    State *FieldModelStateEx
    Wallet *FieldModelBalance
    Asset *FieldModelOptionalBalance
    Orders *FieldModelVectorOrder
}

// Create a new Account field model
func NewFieldModelAccount(buffer *fbe.Buffer, offset int) *FieldModelAccount {
    fbeResult := FieldModelAccount{buffer: buffer, offset: offset}
    fbeResult.Id = fbe.NewFieldModelInt32(buffer, 4 + 4)
    fbeResult.Name = fbe.NewFieldModelString(buffer, fbeResult.Id.FBEOffset() + fbeResult.Id.FBESize())
    fbeResult.State = NewFieldModelStateEx(buffer, fbeResult.Name.FBEOffset() + fbeResult.Name.FBESize())
    fbeResult.Wallet = NewFieldModelBalance(buffer, fbeResult.State.FBEOffset() + fbeResult.State.FBESize())
    fbeResult.Asset = NewFieldModelOptionalBalance(buffer, fbeResult.Wallet.FBEOffset() + fbeResult.Wallet.FBESize())
    fbeResult.Orders = NewFieldModelVectorOrder(buffer, fbeResult.Asset.FBEOffset() + fbeResult.Asset.FBESize())
    return &fbeResult
}

// Get the field size
func (fm *FieldModelAccount) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelAccount) FBEBody() int {
    fbeResult := 4 + 4 +
        fm.Id.FBESize() +
        fm.Name.FBESize() +
        fm.State.FBESize() +
        fm.Wallet.FBESize() +
        fm.Asset.FBESize() +
        fm.Orders.FBESize() +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelAccount) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        fm.Id.FBEExtra() +
        fm.Name.FBEExtra() +
        fm.State.FBEExtra() +
        fm.Wallet.FBEExtra() +
        fm.Asset.FBEExtra() +
        fm.Orders.FBEExtra() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelAccount) FBEType() int { return 3 }

// Get the field offset
func (fm *FieldModelAccount) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelAccount) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelAccount) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelAccount) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelAccount) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelAccount) VerifyType(fbeVerifyType bool) bool {
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
func (fm *FieldModelAccount) VerifyFields(fbeStructSize int) bool {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Id.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Id.Verify() {
        return false
    }
    fbeCurrentSize += fm.Id.FBESize()

    if (fbeCurrentSize + fm.Name.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Name.Verify() {
        return false
    }
    fbeCurrentSize += fm.Name.FBESize()

    if (fbeCurrentSize + fm.State.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.State.Verify() {
        return false
    }
    fbeCurrentSize += fm.State.FBESize()

    if (fbeCurrentSize + fm.Wallet.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Wallet.Verify() {
        return false
    }
    fbeCurrentSize += fm.Wallet.FBESize()

    if (fbeCurrentSize + fm.Asset.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Asset.Verify() {
        return false
    }
    fbeCurrentSize += fm.Asset.FBESize()

    if (fbeCurrentSize + fm.Orders.FBESize()) > fbeStructSize {
        return true
    }
    if !fm.Orders.Verify() {
        return false
    }
    fbeCurrentSize += fm.Orders.FBESize()

    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelAccount) GetBegin() (int, error) {
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
func (fm *FieldModelAccount) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelAccount) Get() (*Account, error) {
    fbeResult := NewAccount()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelAccount) GetValue(fbeValue *Account) error {
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
func (fm *FieldModelAccount) GetFields(fbeValue *Account, fbeStructSize int) {
    fbeCurrentSize := 4 + 4

    if (fbeCurrentSize + fm.Id.FBESize()) <= fbeStructSize {
        fbeValue.Id, _ = fm.Id.Get()
    } else {
        fbeValue.Id = 0
    }
    fbeCurrentSize += fm.Id.FBESize()

    if (fbeCurrentSize + fm.Name.FBESize()) <= fbeStructSize {
        fbeValue.Name, _ = fm.Name.Get()
    } else {
        fbeValue.Name = ""
    }
    fbeCurrentSize += fm.Name.FBESize()

    if (fbeCurrentSize + fm.State.FBESize()) <= fbeStructSize {
        _ = fm.State.GetValueDefault(&fbeValue.State, StateEx_initialized | StateEx_bad | StateEx_sad)
    } else {
        fbeValue.State = StateEx_initialized | StateEx_bad | StateEx_sad
    }
    fbeCurrentSize += fm.State.FBESize()

    if (fbeCurrentSize + fm.Wallet.FBESize()) <= fbeStructSize {
        _ = fm.Wallet.GetValue(&fbeValue.Wallet)
    } else {
        fbeValue.Wallet = *NewBalance()
    }
    fbeCurrentSize += fm.Wallet.FBESize()

    if (fbeCurrentSize + fm.Asset.FBESize()) <= fbeStructSize {
        fbeValue.Asset, _ = fm.Asset.Get()
    } else {
        fbeValue.Asset = nil
    }
    fbeCurrentSize += fm.Asset.FBESize()

    if (fbeCurrentSize + fm.Orders.FBESize()) <= fbeStructSize {
        fbeValue.Orders, _ = fm.Orders.Get()
    } else {
        fbeValue.Orders = make([]Order, 0)
    }
    fbeCurrentSize += fm.Orders.FBESize()
}

// Set the struct value (begin phase)
func (fm *FieldModelAccount) SetBegin() (int, error) {
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
func (fm *FieldModelAccount) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelAccount) Set(fbeValue *Account) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelAccount) SetFields(fbeValue *Account) error {
    var err error = nil

    if err = fm.Id.Set(fbeValue.Id); err != nil {
        return err
    }
    if err = fm.Name.Set(fbeValue.Name); err != nil {
        return err
    }
    if err = fm.State.Set(&fbeValue.State); err != nil {
        return err
    }
    if err = fm.Wallet.Set(&fbeValue.Wallet); err != nil {
        return err
    }
    if err = fm.Asset.Set(fbeValue.Asset); err != nil {
        return err
    }
    if err = fm.Orders.Set(fbeValue.Orders); err != nil {
        return err
    }
    return err
}
//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: template_variant.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

#include "template_variant_ptr_models.h"

namespace FBE {

FieldModelPtr_template_variant_Line::FieldModelPtr_template_variant_Line(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
{}

FieldModelPtr_template_variant_Line::~FieldModelPtr_template_variant_Line()
{
    if (ptr) delete ptr;
}

size_t FieldModelPtr_template_variant_Line::fbe_extra() const noexcept
{
    if (!ptr) return 0;

    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_ptr_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    if ((fbe_ptr_offset == 0) || ((_buffer.offset() + fbe_ptr_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_ptr_offset);
    size_t fbe_result = ptr->fbe_size() + ptr->fbe_extra();
    _buffer.unshift(fbe_ptr_offset);

    return fbe_result;
}

bool FieldModelPtr_template_variant_Line::verify() const noexcept
{
    if (!ptr) return true;

    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return true;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    if (fbe_has_value == 0)
        return true;

    uint32_t fbe_optional_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    if (fbe_optional_offset == 0)
        return false;

    _buffer.shift(fbe_optional_offset);
    bool fbe_result = ptr->verify();
    _buffer.unshift(fbe_optional_offset);
    return fbe_result;
}

bool FieldModelPtr_template_variant_Line::has_value() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return false;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    return (fbe_has_value != 0);
}

size_t FieldModelPtr_template_variant_Line::get_begin() const noexcept
{
    if (!has_value())
        return 0;

    uint32_t fbe_ptr_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    assert((fbe_ptr_offset > 0) && "Model is broken!");
    if (fbe_ptr_offset == 0)
        return 0;

    _buffer.shift(fbe_ptr_offset);
    return fbe_ptr_offset;
}

void FieldModelPtr_template_variant_Line::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPtr_template_variant_Line::get(::template_variant::Line** fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    if (ptr) delete ptr;
    ptr = new FieldModel_template_variant_Line(_buffer, 0);

    ::template_variant::Line *tempModel = new ::template_variant::Line();
    ptr->get(*tempModel);
    *fbe_value = tempModel;

    get_end(fbe_begin);
}

size_t FieldModelPtr_template_variant_Line::set_begin(bool has_value)
{
    assert(((_buffer.offset() + fbe_offset() + fbe_size()) <= _buffer.size()) && "Model is broken!");
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint8_t fbe_has_value = has_value ? 1 : 0;
    *((uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset())) = fbe_has_value;
    if (fbe_has_value == 0)
        return 0;

    uint32_t fbe_ptr_size = 4;
    uint32_t fbe_ptr_offset = (uint32_t)(_buffer.allocate(fbe_ptr_size) - _buffer.offset());
    assert(((fbe_ptr_offset > 0) && ((_buffer.offset() + fbe_ptr_offset + fbe_ptr_size) <= _buffer.size())) && "Model is broken!");
    if ((fbe_ptr_offset == 0) || ((_buffer.offset() + fbe_ptr_offset + fbe_ptr_size) > _buffer.size()))
        return 0;

    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1, fbe_ptr_offset);

    _buffer.shift(fbe_ptr_offset);
    return fbe_ptr_offset;
}

void FieldModelPtr_template_variant_Line::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPtr_template_variant_Line::set(const ::template_variant::Line* fbe_value) noexcept
{
    size_t fbe_begin = set_begin(fbe_value != nullptr);
    if (fbe_begin == 0)
        return;

    if (fbe_value != nullptr) {
        BaseFieldModel* temp = new FieldModel_template_variant_Line(_buffer, 0);
        if (ptr) delete ptr;
        ptr = temp;
        ptr->set(*fbe_value);
    }

    set_end(fbe_begin);
}

FieldModel_template_variant_Line::FieldModel_template_variant_Line(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
    , v(buffer, 4 + 4)
    , vv(buffer, v.fbe_offset() + v.fbe_size())
    , vm(buffer, vv.fbe_offset() + vv.fbe_size())
    , vo(buffer, vm.fbe_offset() + vm.fbe_size())
{}

size_t FieldModel_template_variant_Line::fbe_body() const noexcept
{
    size_t fbe_result = 4 + 4
        + v.fbe_size()
        + vv.fbe_size()
        + vm.fbe_size()
        + vo.fbe_size()
        ;
    return fbe_result;
}

size_t FieldModel_template_variant_Line::fbe_extra() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_struct_offset);

    size_t fbe_result = fbe_body()
        + v.fbe_extra()
        + vv.fbe_extra()
        + vm.fbe_extra()
        + vo.fbe_extra()
        ;

    _buffer.unshift(fbe_struct_offset);

    return fbe_result;
}

bool FieldModel_template_variant_Line::verify(bool fbe_verify_type) const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return true;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4 + 4) > _buffer.size()))
        return false;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset);
    if (fbe_struct_size < (4 + 4))
        return false;

    uint32_t fbe_struct_type = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset + 4);
    if (fbe_verify_type && (fbe_struct_type != fbe_type()))
        return false;

    _buffer.shift(fbe_struct_offset);
    bool fbe_result = verify_fields(fbe_struct_size);
    _buffer.unshift(fbe_struct_offset);
    return fbe_result;
}

bool FieldModel_template_variant_Line::verify_fields([[maybe_unused]] size_t fbe_struct_size) const noexcept
{
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + v.fbe_size()) > fbe_struct_size)
        return true;
    if (!v.verify())
        return false;
    fbe_current_size += v.fbe_size();

    if ((fbe_current_size + vv.fbe_size()) > fbe_struct_size)
        return true;
    if (!vv.verify())
        return false;
    fbe_current_size += vv.fbe_size();

    if ((fbe_current_size + vm.fbe_size()) > fbe_struct_size)
        return true;
    if (!vm.verify())
        return false;
    fbe_current_size += vm.fbe_size();

    if ((fbe_current_size + vo.fbe_size()) > fbe_struct_size)
        return true;
    if (!vo.verify())
        return false;
    fbe_current_size += vo.fbe_size();

    return true;
}

size_t FieldModel_template_variant_Line::get_begin() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    assert(((fbe_struct_offset > 0) && ((_buffer.offset() + fbe_struct_offset + 4 + 4) <= _buffer.size())) && "Model is broken!");
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4 + 4) > _buffer.size()))
        return 0;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset);
    assert((fbe_struct_size >= (4 + 4)) && "Model is broken!");
    if (fbe_struct_size < (4 + 4))
        return 0;

    _buffer.shift(fbe_struct_offset);
    return fbe_struct_offset;
}

void FieldModel_template_variant_Line::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModel_template_variant_Line::get(::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset());
    get_fields(fbe_value, fbe_struct_size);
    get_end(fbe_begin);
}

void FieldModel_template_variant_Line::get_fields([[maybe_unused]] ::FBE::Base& base_fbe_value, [[maybe_unused]] size_t fbe_struct_size) noexcept
{
    ::template_variant::Line& fbe_value = static_cast<::template_variant::Line&>(base_fbe_value);
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + v.fbe_size()) <= fbe_struct_size)
        {
            v.get(fbe_value.v);
        }
    else
        fbe_value.v = ::variants::V();
    fbe_current_size += v.fbe_size();

    if ((fbe_current_size + vv.fbe_size()) <= fbe_struct_size)
        {
            vv.get(fbe_value.vv);
        }
    else
        fbe_value.vv.clear();
    fbe_current_size += vv.fbe_size();

    if ((fbe_current_size + vm.fbe_size()) <= fbe_struct_size)
        {
            vm.get(fbe_value.vm);
        }
    else
        fbe_value.vm.clear();
    fbe_current_size += vm.fbe_size();

    if ((fbe_current_size + vo.fbe_size()) <= fbe_struct_size)
        {
            vo.get(fbe_value.vo);
        }
    else
        fbe_value.vo = std::nullopt;
    fbe_current_size += vo.fbe_size();
}

size_t FieldModel_template_variant_Line::set_begin()
{
    assert(((_buffer.offset() + fbe_offset() + fbe_size()) <= _buffer.size()) && "Model is broken!");
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_size = (uint32_t)fbe_body();
    uint32_t fbe_struct_offset = (uint32_t)(_buffer.allocate(fbe_struct_size) - _buffer.offset());
    assert(((fbe_struct_offset > 0) && ((_buffer.offset() + fbe_struct_offset + fbe_struct_size) <= _buffer.size())) && "Model is broken!");
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + fbe_struct_size) > _buffer.size()))
        return 0;

    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset(), fbe_struct_offset);
    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset, fbe_struct_size);
    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset + 4, fbe_type());

    _buffer.shift(fbe_struct_offset);
    return fbe_struct_offset;
}

void FieldModel_template_variant_Line::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModel_template_variant_Line::set(const ::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = set_begin();
    if (fbe_begin == 0)
        return;

    set_fields(fbe_value);
    set_end(fbe_begin);
}

void FieldModel_template_variant_Line::set_fields([[maybe_unused]] const ::FBE::Base& base_fbe_value) noexcept
{
    [[maybe_unused]] const ::template_variant::Line& fbe_value = static_cast<const ::template_variant::Line&>(base_fbe_value);
    v.set(fbe_value.v);
    vv.set(fbe_value.vv);
    vm.set(fbe_value.vm);
    vo.set(fbe_value.vo);
}

namespace template_variant {

bool LineModel::verify()
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return false;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    if (fbe_full_size < model.fbe_size())
        return false;

    return model.verify();
}

size_t LineModel::create_begin()
{
    size_t fbe_begin = this->buffer().allocate(4 + model.fbe_size());
    return fbe_begin;
}

size_t LineModel::create_end(size_t fbe_begin)
{
    size_t fbe_end = this->buffer().size();
    uint32_t fbe_full_size = (uint32_t)(fbe_end - fbe_begin);
    *((uint32_t*)(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4)) = fbe_full_size;
    return fbe_full_size;
}

size_t LineModel::serialize(const ::template_variant::Line& value)
{
    size_t fbe_begin = create_begin();
    model.set(value);
    size_t fbe_full_size = create_end(fbe_begin);
    return fbe_full_size;
}

size_t LineModel::deserialize(::template_variant::Line& value) noexcept
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return 0;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    assert((fbe_full_size >= model.fbe_size()) && "Model is broken!");
    if (fbe_full_size < model.fbe_size())
        return 0;

    model.get(value);
    return fbe_full_size;
}

} // namespace template_variant

FieldModelPtr_template_variant_Line2::FieldModelPtr_template_variant_Line2(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
{}

FieldModelPtr_template_variant_Line2::~FieldModelPtr_template_variant_Line2()
{
    if (ptr) delete ptr;
}

size_t FieldModelPtr_template_variant_Line2::fbe_extra() const noexcept
{
    if (!ptr) return 0;

    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_ptr_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    if ((fbe_ptr_offset == 0) || ((_buffer.offset() + fbe_ptr_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_ptr_offset);
    size_t fbe_result = ptr->fbe_size() + ptr->fbe_extra();
    _buffer.unshift(fbe_ptr_offset);

    return fbe_result;
}

bool FieldModelPtr_template_variant_Line2::verify() const noexcept
{
    if (!ptr) return true;

    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return true;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    if (fbe_has_value == 0)
        return true;

    uint32_t fbe_optional_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    if (fbe_optional_offset == 0)
        return false;

    _buffer.shift(fbe_optional_offset);
    bool fbe_result = ptr->verify();
    _buffer.unshift(fbe_optional_offset);
    return fbe_result;
}

bool FieldModelPtr_template_variant_Line2::has_value() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return false;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    return (fbe_has_value != 0);
}

size_t FieldModelPtr_template_variant_Line2::get_begin() const noexcept
{
    if (!has_value())
        return 0;

    uint32_t fbe_ptr_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    assert((fbe_ptr_offset > 0) && "Model is broken!");
    if (fbe_ptr_offset == 0)
        return 0;

    _buffer.shift(fbe_ptr_offset);
    return fbe_ptr_offset;
}

void FieldModelPtr_template_variant_Line2::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPtr_template_variant_Line2::get(::template_variant::Line2** fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    if (ptr) delete ptr;
    ptr = new FieldModel_template_variant_Line2(_buffer, 0);

    ::template_variant::Line2 *tempModel = new ::template_variant::Line2();
    ptr->get(*tempModel);
    *fbe_value = tempModel;

    get_end(fbe_begin);
}

size_t FieldModelPtr_template_variant_Line2::set_begin(bool has_value)
{
    assert(((_buffer.offset() + fbe_offset() + fbe_size()) <= _buffer.size()) && "Model is broken!");
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint8_t fbe_has_value = has_value ? 1 : 0;
    *((uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset())) = fbe_has_value;
    if (fbe_has_value == 0)
        return 0;

    uint32_t fbe_ptr_size = 4;
    uint32_t fbe_ptr_offset = (uint32_t)(_buffer.allocate(fbe_ptr_size) - _buffer.offset());
    assert(((fbe_ptr_offset > 0) && ((_buffer.offset() + fbe_ptr_offset + fbe_ptr_size) <= _buffer.size())) && "Model is broken!");
    if ((fbe_ptr_offset == 0) || ((_buffer.offset() + fbe_ptr_offset + fbe_ptr_size) > _buffer.size()))
        return 0;

    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1, fbe_ptr_offset);

    _buffer.shift(fbe_ptr_offset);
    return fbe_ptr_offset;
}

void FieldModelPtr_template_variant_Line2::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPtr_template_variant_Line2::set(const ::template_variant::Line2* fbe_value) noexcept
{
    size_t fbe_begin = set_begin(fbe_value != nullptr);
    if (fbe_begin == 0)
        return;

    if (fbe_value != nullptr) {
        BaseFieldModel* temp = new FieldModel_template_variant_Line2(_buffer, 0);
        if (ptr) delete ptr;
        ptr = temp;
        ptr->set(*fbe_value);
    }

    set_end(fbe_begin);
}

FieldModel_template_variant_Line2::FieldModel_template_variant_Line2(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
    , vm(buffer, 4 + 4)
{}

size_t FieldModel_template_variant_Line2::fbe_body() const noexcept
{
    size_t fbe_result = 4 + 4
        + vm.fbe_size()
        ;
    return fbe_result;
}

size_t FieldModel_template_variant_Line2::fbe_extra() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_struct_offset);

    size_t fbe_result = fbe_body()
        + vm.fbe_extra()
        ;

    _buffer.unshift(fbe_struct_offset);

    return fbe_result;
}

bool FieldModel_template_variant_Line2::verify(bool fbe_verify_type) const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return true;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4 + 4) > _buffer.size()))
        return false;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset);
    if (fbe_struct_size < (4 + 4))
        return false;

    uint32_t fbe_struct_type = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset + 4);
    if (fbe_verify_type && (fbe_struct_type != fbe_type()))
        return false;

    _buffer.shift(fbe_struct_offset);
    bool fbe_result = verify_fields(fbe_struct_size);
    _buffer.unshift(fbe_struct_offset);
    return fbe_result;
}

bool FieldModel_template_variant_Line2::verify_fields([[maybe_unused]] size_t fbe_struct_size) const noexcept
{
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + vm.fbe_size()) > fbe_struct_size)
        return true;
    if (!vm.verify())
        return false;
    fbe_current_size += vm.fbe_size();

    return true;
}

size_t FieldModel_template_variant_Line2::get_begin() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    assert(((fbe_struct_offset > 0) && ((_buffer.offset() + fbe_struct_offset + 4 + 4) <= _buffer.size())) && "Model is broken!");
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4 + 4) > _buffer.size()))
        return 0;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset);
    assert((fbe_struct_size >= (4 + 4)) && "Model is broken!");
    if (fbe_struct_size < (4 + 4))
        return 0;

    _buffer.shift(fbe_struct_offset);
    return fbe_struct_offset;
}

void FieldModel_template_variant_Line2::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModel_template_variant_Line2::get(::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset());
    get_fields(fbe_value, fbe_struct_size);
    get_end(fbe_begin);
}

void FieldModel_template_variant_Line2::get_fields([[maybe_unused]] ::FBE::Base& base_fbe_value, [[maybe_unused]] size_t fbe_struct_size) noexcept
{
    ::template_variant::Line2& fbe_value = static_cast<::template_variant::Line2&>(base_fbe_value);
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + vm.fbe_size()) <= fbe_struct_size)
        {
            vm.get(fbe_value.vm);
        }
    else
        fbe_value.vm.clear();
    fbe_current_size += vm.fbe_size();
}

size_t FieldModel_template_variant_Line2::set_begin()
{
    assert(((_buffer.offset() + fbe_offset() + fbe_size()) <= _buffer.size()) && "Model is broken!");
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_size = (uint32_t)fbe_body();
    uint32_t fbe_struct_offset = (uint32_t)(_buffer.allocate(fbe_struct_size) - _buffer.offset());
    assert(((fbe_struct_offset > 0) && ((_buffer.offset() + fbe_struct_offset + fbe_struct_size) <= _buffer.size())) && "Model is broken!");
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + fbe_struct_size) > _buffer.size()))
        return 0;

    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset(), fbe_struct_offset);
    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset, fbe_struct_size);
    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset + 4, fbe_type());

    _buffer.shift(fbe_struct_offset);
    return fbe_struct_offset;
}

void FieldModel_template_variant_Line2::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModel_template_variant_Line2::set(const ::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = set_begin();
    if (fbe_begin == 0)
        return;

    set_fields(fbe_value);
    set_end(fbe_begin);
}

void FieldModel_template_variant_Line2::set_fields([[maybe_unused]] const ::FBE::Base& base_fbe_value) noexcept
{
    [[maybe_unused]] const ::template_variant::Line2& fbe_value = static_cast<const ::template_variant::Line2&>(base_fbe_value);
    vm.set(fbe_value.vm);
}

namespace template_variant {

bool Line2Model::verify()
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return false;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    if (fbe_full_size < model.fbe_size())
        return false;

    return model.verify();
}

size_t Line2Model::create_begin()
{
    size_t fbe_begin = this->buffer().allocate(4 + model.fbe_size());
    return fbe_begin;
}

size_t Line2Model::create_end(size_t fbe_begin)
{
    size_t fbe_end = this->buffer().size();
    uint32_t fbe_full_size = (uint32_t)(fbe_end - fbe_begin);
    *((uint32_t*)(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4)) = fbe_full_size;
    return fbe_full_size;
}

size_t Line2Model::serialize(const ::template_variant::Line2& value)
{
    size_t fbe_begin = create_begin();
    model.set(value);
    size_t fbe_full_size = create_end(fbe_begin);
    return fbe_full_size;
}

size_t Line2Model::deserialize(::template_variant::Line2& value) noexcept
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return 0;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    assert((fbe_full_size >= model.fbe_size()) && "Model is broken!");
    if (fbe_full_size < model.fbe_size())
        return 0;

    model.get(value);
    return fbe_full_size;
}

} // namespace template_variant

FieldModelPtr_template_variant_Line3::FieldModelPtr_template_variant_Line3(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
{}

FieldModelPtr_template_variant_Line3::~FieldModelPtr_template_variant_Line3()
{
    if (ptr) delete ptr;
}

size_t FieldModelPtr_template_variant_Line3::fbe_extra() const noexcept
{
    if (!ptr) return 0;

    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_ptr_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    if ((fbe_ptr_offset == 0) || ((_buffer.offset() + fbe_ptr_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_ptr_offset);
    size_t fbe_result = ptr->fbe_size() + ptr->fbe_extra();
    _buffer.unshift(fbe_ptr_offset);

    return fbe_result;
}

bool FieldModelPtr_template_variant_Line3::verify() const noexcept
{
    if (!ptr) return true;

    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return true;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    if (fbe_has_value == 0)
        return true;

    uint32_t fbe_optional_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    if (fbe_optional_offset == 0)
        return false;

    _buffer.shift(fbe_optional_offset);
    bool fbe_result = ptr->verify();
    _buffer.unshift(fbe_optional_offset);
    return fbe_result;
}

bool FieldModelPtr_template_variant_Line3::has_value() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return false;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    return (fbe_has_value != 0);
}

size_t FieldModelPtr_template_variant_Line3::get_begin() const noexcept
{
    if (!has_value())
        return 0;

    uint32_t fbe_ptr_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1);
    assert((fbe_ptr_offset > 0) && "Model is broken!");
    if (fbe_ptr_offset == 0)
        return 0;

    _buffer.shift(fbe_ptr_offset);
    return fbe_ptr_offset;
}

void FieldModelPtr_template_variant_Line3::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPtr_template_variant_Line3::get(::template_variant::Line3** fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    if (ptr) delete ptr;
    ptr = new FieldModel_template_variant_Line3(_buffer, 0);

    ::template_variant::Line3 *tempModel = new ::template_variant::Line3();
    ptr->get(*tempModel);
    *fbe_value = tempModel;

    get_end(fbe_begin);
}

size_t FieldModelPtr_template_variant_Line3::set_begin(bool has_value)
{
    assert(((_buffer.offset() + fbe_offset() + fbe_size()) <= _buffer.size()) && "Model is broken!");
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint8_t fbe_has_value = has_value ? 1 : 0;
    *((uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset())) = fbe_has_value;
    if (fbe_has_value == 0)
        return 0;

    uint32_t fbe_ptr_size = 4;
    uint32_t fbe_ptr_offset = (uint32_t)(_buffer.allocate(fbe_ptr_size) - _buffer.offset());
    assert(((fbe_ptr_offset > 0) && ((_buffer.offset() + fbe_ptr_offset + fbe_ptr_size) <= _buffer.size())) && "Model is broken!");
    if ((fbe_ptr_offset == 0) || ((_buffer.offset() + fbe_ptr_offset + fbe_ptr_size) > _buffer.size()))
        return 0;

    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset() + 1, fbe_ptr_offset);

    _buffer.shift(fbe_ptr_offset);
    return fbe_ptr_offset;
}

void FieldModelPtr_template_variant_Line3::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPtr_template_variant_Line3::set(const ::template_variant::Line3* fbe_value) noexcept
{
    size_t fbe_begin = set_begin(fbe_value != nullptr);
    if (fbe_begin == 0)
        return;

    if (fbe_value != nullptr) {
        BaseFieldModel* temp = new FieldModel_template_variant_Line3(_buffer, 0);
        if (ptr) delete ptr;
        ptr = temp;
        ptr->set(*fbe_value);
    }

    set_end(fbe_begin);
}

FieldModel_template_variant_Line3::FieldModel_template_variant_Line3(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
    , value(buffer, 4 + 4)
{}

size_t FieldModel_template_variant_Line3::fbe_body() const noexcept
{
    size_t fbe_result = 4 + 4
        + value.fbe_size()
        ;
    return fbe_result;
}

size_t FieldModel_template_variant_Line3::fbe_extra() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_struct_offset);

    size_t fbe_result = fbe_body()
        + value.fbe_extra()
        ;

    _buffer.unshift(fbe_struct_offset);

    return fbe_result;
}

bool FieldModel_template_variant_Line3::verify(bool fbe_verify_type) const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return true;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4 + 4) > _buffer.size()))
        return false;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset);
    if (fbe_struct_size < (4 + 4))
        return false;

    uint32_t fbe_struct_type = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset + 4);
    if (fbe_verify_type && (fbe_struct_type != fbe_type()))
        return false;

    _buffer.shift(fbe_struct_offset);
    bool fbe_result = verify_fields(fbe_struct_size);
    _buffer.unshift(fbe_struct_offset);
    return fbe_result;
}

bool FieldModel_template_variant_Line3::verify_fields([[maybe_unused]] size_t fbe_struct_size) const noexcept
{
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + value.fbe_size()) > fbe_struct_size)
        return true;
    if (!value.verify())
        return false;
    fbe_current_size += value.fbe_size();

    return true;
}

size_t FieldModel_template_variant_Line3::get_begin() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    assert(((fbe_struct_offset > 0) && ((_buffer.offset() + fbe_struct_offset + 4 + 4) <= _buffer.size())) && "Model is broken!");
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4 + 4) > _buffer.size()))
        return 0;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset);
    assert((fbe_struct_size >= (4 + 4)) && "Model is broken!");
    if (fbe_struct_size < (4 + 4))
        return 0;

    _buffer.shift(fbe_struct_offset);
    return fbe_struct_offset;
}

void FieldModel_template_variant_Line3::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModel_template_variant_Line3::get(::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset());
    get_fields(fbe_value, fbe_struct_size);
    get_end(fbe_begin);
}

void FieldModel_template_variant_Line3::get_fields([[maybe_unused]] ::FBE::Base& base_fbe_value, [[maybe_unused]] size_t fbe_struct_size) noexcept
{
    ::template_variant::Line3& fbe_value = static_cast<::template_variant::Line3&>(base_fbe_value);
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + value.fbe_size()) <= fbe_struct_size)
        {
            value.get(fbe_value.value);
        }
    else
        fbe_value.value = ::variants::Value();
    fbe_current_size += value.fbe_size();
}

size_t FieldModel_template_variant_Line3::set_begin()
{
    assert(((_buffer.offset() + fbe_offset() + fbe_size()) <= _buffer.size()) && "Model is broken!");
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_size = (uint32_t)fbe_body();
    uint32_t fbe_struct_offset = (uint32_t)(_buffer.allocate(fbe_struct_size) - _buffer.offset());
    assert(((fbe_struct_offset > 0) && ((_buffer.offset() + fbe_struct_offset + fbe_struct_size) <= _buffer.size())) && "Model is broken!");
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + fbe_struct_size) > _buffer.size()))
        return 0;

    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset(), fbe_struct_offset);
    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset, fbe_struct_size);
    unaligned_store<uint32_t>(_buffer.data() + _buffer.offset() + fbe_struct_offset + 4, fbe_type());

    _buffer.shift(fbe_struct_offset);
    return fbe_struct_offset;
}

void FieldModel_template_variant_Line3::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModel_template_variant_Line3::set(const ::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = set_begin();
    if (fbe_begin == 0)
        return;

    set_fields(fbe_value);
    set_end(fbe_begin);
}

void FieldModel_template_variant_Line3::set_fields([[maybe_unused]] const ::FBE::Base& base_fbe_value) noexcept
{
    [[maybe_unused]] const ::template_variant::Line3& fbe_value = static_cast<const ::template_variant::Line3&>(base_fbe_value);
    value.set(fbe_value.value);
}

namespace template_variant {

bool Line3Model::verify()
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return false;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    if (fbe_full_size < model.fbe_size())
        return false;

    return model.verify();
}

size_t Line3Model::create_begin()
{
    size_t fbe_begin = this->buffer().allocate(4 + model.fbe_size());
    return fbe_begin;
}

size_t Line3Model::create_end(size_t fbe_begin)
{
    size_t fbe_end = this->buffer().size();
    uint32_t fbe_full_size = (uint32_t)(fbe_end - fbe_begin);
    *((uint32_t*)(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4)) = fbe_full_size;
    return fbe_full_size;
}

size_t Line3Model::serialize(const ::template_variant::Line3& value)
{
    size_t fbe_begin = create_begin();
    model.set(value);
    size_t fbe_full_size = create_end(fbe_begin);
    return fbe_full_size;
}

size_t Line3Model::deserialize(::template_variant::Line3& value) noexcept
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return 0;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    assert((fbe_full_size >= model.fbe_size()) && "Model is broken!");
    if (fbe_full_size < model.fbe_size())
        return 0;

    model.get(value);
    return fbe_full_size;
}

} // namespace template_variant

} // namespace FBE

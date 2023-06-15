//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: pkg.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

#include "pkg_ptr_pmr_models.h"

namespace FBE {

FieldModelPMRPtr_pkg_Info::FieldModelPMRPtr_pkg_Info(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
{}

FieldModelPMRPtr_pkg_Info::~FieldModelPMRPtr_pkg_Info()
{
    if (ptr) delete ptr;
}

size_t FieldModelPMRPtr_pkg_Info::fbe_extra() const noexcept
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

bool FieldModelPMRPtr_pkg_Info::verify() const noexcept
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

bool FieldModelPMRPtr_pkg_Info::has_value() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return false;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    return (fbe_has_value != 0);
}

size_t FieldModelPMRPtr_pkg_Info::get_begin() const noexcept
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

void FieldModelPMRPtr_pkg_Info::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMRPtr_pkg_Info::get(::pkg_pmr::Info** fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    if (ptr) delete ptr;
    ptr = new FieldModelPMR_pkg_Info(_buffer, 0);

    ::pkg_pmr::Info *tempModel = new ::pkg_pmr::Info();
    ptr->get(*tempModel);
    *fbe_value = tempModel;

    get_end(fbe_begin);
}

size_t FieldModelPMRPtr_pkg_Info::set_begin(bool has_value)
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

void FieldModelPMRPtr_pkg_Info::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMRPtr_pkg_Info::set(const ::pkg_pmr::Info* fbe_value) noexcept
{
    size_t fbe_begin = set_begin(fbe_value != nullptr);
    if (fbe_begin == 0)
        return;

    if (fbe_value != nullptr) {
        BaseFieldModel* temp = new FieldModelPMR_pkg_Info(_buffer, 0);
        if (ptr) delete ptr;
        ptr = temp;
        ptr->set(*fbe_value);
    }

    set_end(fbe_begin);
}

FieldModelPMR_pkg_Info::FieldModelPMR_pkg_Info(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
    , info(buffer, 4 + 4)
    , sex(buffer, info.fbe_offset() + info.fbe_size())
    , flag(buffer, sex.fbe_offset() + sex.fbe_size())
    , extra(buffer, flag.fbe_offset() + flag.fbe_size())
{}

size_t FieldModelPMR_pkg_Info::fbe_body() const noexcept
{
    size_t fbe_result = 4 + 4
        + info.fbe_size()
        + sex.fbe_size()
        + flag.fbe_size()
        + extra.fbe_size()
        ;
    return fbe_result;
}

size_t FieldModelPMR_pkg_Info::fbe_extra() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_struct_offset);

    size_t fbe_result = fbe_body()
        + info.fbe_extra()
        + sex.fbe_extra()
        + flag.fbe_extra()
        + extra.fbe_extra()
        ;

    _buffer.unshift(fbe_struct_offset);

    return fbe_result;
}

bool FieldModelPMR_pkg_Info::verify(bool fbe_verify_type) const noexcept
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

bool FieldModelPMR_pkg_Info::verify_fields([[maybe_unused]] size_t fbe_struct_size) const noexcept
{
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + info.fbe_size()) > fbe_struct_size)
        return true;
    if (!info.verify())
        return false;
    fbe_current_size += info.fbe_size();

    if ((fbe_current_size + sex.fbe_size()) > fbe_struct_size)
        return true;
    if (!sex.verify())
        return false;
    fbe_current_size += sex.fbe_size();

    if ((fbe_current_size + flag.fbe_size()) > fbe_struct_size)
        return true;
    if (!flag.verify())
        return false;
    fbe_current_size += flag.fbe_size();

    if ((fbe_current_size + extra.fbe_size()) > fbe_struct_size)
        return true;
    if (!extra.verify())
        return false;
    fbe_current_size += extra.fbe_size();

    return true;
}

size_t FieldModelPMR_pkg_Info::get_begin() const noexcept
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

void FieldModelPMR_pkg_Info::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMR_pkg_Info::get(::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset());
    get_fields(fbe_value, fbe_struct_size);
    get_end(fbe_begin);
}

void FieldModelPMR_pkg_Info::get_fields([[maybe_unused]] ::FBE::Base& base_fbe_value, [[maybe_unused]] size_t fbe_struct_size) noexcept
{
    ::pkg_pmr::Info& fbe_value = static_cast<::pkg_pmr::Info&>(base_fbe_value);
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + info.fbe_size()) <= fbe_struct_size)
        {
            info.get(fbe_value.info);
        }
    else
        fbe_value.info = "";
    fbe_current_size += info.fbe_size();

    if ((fbe_current_size + sex.fbe_size()) <= fbe_struct_size)
        {
            sex.get(fbe_value.sex);
        }
    else
        fbe_value.sex = ::osa_pmr::Sex();
    fbe_current_size += sex.fbe_size();

    if ((fbe_current_size + flag.fbe_size()) <= fbe_struct_size)
        {
            flag.get(fbe_value.flag);
        }
    else
        fbe_value.flag = ::osa_pmr::MyFLags();
    fbe_current_size += flag.fbe_size();

    if ((fbe_current_size + extra.fbe_size()) <= fbe_struct_size)
        {
            extra.get(fbe_value.extra);
        }
    else
        fbe_value.extra = ::osa_pmr::Extra();
    fbe_current_size += extra.fbe_size();
}

size_t FieldModelPMR_pkg_Info::set_begin()
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

void FieldModelPMR_pkg_Info::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMR_pkg_Info::set(const ::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = set_begin();
    if (fbe_begin == 0)
        return;

    set_fields(fbe_value);
    set_end(fbe_begin);
}

void FieldModelPMR_pkg_Info::set_fields([[maybe_unused]] const ::FBE::Base& base_fbe_value) noexcept
{
    [[maybe_unused]] const ::pkg_pmr::Info& fbe_value = static_cast<const ::pkg_pmr::Info&>(base_fbe_value);
    info.set(fbe_value.info);
    sex.set(fbe_value.sex);
    flag.set(fbe_value.flag);
    extra.set(fbe_value.extra);
}

namespace pkg_pmr {

bool InfoModel::verify()
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return false;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    if (fbe_full_size < model.fbe_size())
        return false;

    return model.verify();
}

size_t InfoModel::create_begin()
{
    size_t fbe_begin = this->buffer().allocate(4 + model.fbe_size());
    return fbe_begin;
}

size_t InfoModel::create_end(size_t fbe_begin)
{
    size_t fbe_end = this->buffer().size();
    uint32_t fbe_full_size = (uint32_t)(fbe_end - fbe_begin);
    *((uint32_t*)(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4)) = fbe_full_size;
    return fbe_full_size;
}

size_t InfoModel::serialize(const ::pkg_pmr::Info& value)
{
    size_t fbe_begin = create_begin();
    model.set(value);
    size_t fbe_full_size = create_end(fbe_begin);
    return fbe_full_size;
}

size_t InfoModel::deserialize(::pkg_pmr::Info& value) noexcept
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

} // namespace pkg_pmr

FieldModelPMRPtr_pkg_Detail::FieldModelPMRPtr_pkg_Detail(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
{}

FieldModelPMRPtr_pkg_Detail::~FieldModelPMRPtr_pkg_Detail()
{
    if (ptr) delete ptr;
}

size_t FieldModelPMRPtr_pkg_Detail::fbe_extra() const noexcept
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

bool FieldModelPMRPtr_pkg_Detail::verify() const noexcept
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

bool FieldModelPMRPtr_pkg_Detail::has_value() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return false;

    uint8_t fbe_has_value = *((const uint8_t *)(_buffer.data() + _buffer.offset() + fbe_offset()));
    return (fbe_has_value != 0);
}

size_t FieldModelPMRPtr_pkg_Detail::get_begin() const noexcept
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

void FieldModelPMRPtr_pkg_Detail::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMRPtr_pkg_Detail::get(::pkg_pmr::Detail** fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    if (ptr) delete ptr;
    ptr = new FieldModelPMR_pkg_Detail(_buffer, 0);

    ::pkg_pmr::Detail *tempModel = new ::pkg_pmr::Detail();
    ptr->get(*tempModel);
    *fbe_value = tempModel;

    get_end(fbe_begin);
}

size_t FieldModelPMRPtr_pkg_Detail::set_begin(bool has_value)
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

void FieldModelPMRPtr_pkg_Detail::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMRPtr_pkg_Detail::set(const ::pkg_pmr::Detail* fbe_value) noexcept
{
    size_t fbe_begin = set_begin(fbe_value != nullptr);
    if (fbe_begin == 0)
        return;

    if (fbe_value != nullptr) {
        BaseFieldModel* temp = new FieldModelPMR_pkg_Detail(_buffer, 0);
        if (ptr) delete ptr;
        ptr = temp;
        ptr->set(*fbe_value);
    }

    set_end(fbe_begin);
}

FieldModelPMR_pkg_Detail::FieldModelPMR_pkg_Detail(FBEBuffer& buffer, size_t offset) noexcept : _buffer(buffer), _offset(offset)
    , extrav(buffer, 4 + 4)
    , extram(buffer, extrav.fbe_offset() + extrav.fbe_size())
{}

size_t FieldModelPMR_pkg_Detail::fbe_body() const noexcept
{
    size_t fbe_result = 4 + 4
        + extrav.fbe_size()
        + extram.fbe_size()
        ;
    return fbe_result;
}

size_t FieldModelPMR_pkg_Detail::fbe_extra() const noexcept
{
    if ((_buffer.offset() + fbe_offset() + fbe_size()) > _buffer.size())
        return 0;

    uint32_t fbe_struct_offset = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset() + fbe_offset());
    if ((fbe_struct_offset == 0) || ((_buffer.offset() + fbe_struct_offset + 4) > _buffer.size()))
        return 0;

    _buffer.shift(fbe_struct_offset);

    size_t fbe_result = fbe_body()
        + extrav.fbe_extra()
        + extram.fbe_extra()
        ;

    _buffer.unshift(fbe_struct_offset);

    return fbe_result;
}

bool FieldModelPMR_pkg_Detail::verify(bool fbe_verify_type) const noexcept
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

bool FieldModelPMR_pkg_Detail::verify_fields([[maybe_unused]] size_t fbe_struct_size) const noexcept
{
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + extrav.fbe_size()) > fbe_struct_size)
        return true;
    if (!extrav.verify())
        return false;
    fbe_current_size += extrav.fbe_size();

    if ((fbe_current_size + extram.fbe_size()) > fbe_struct_size)
        return true;
    if (!extram.verify())
        return false;
    fbe_current_size += extram.fbe_size();

    return true;
}

size_t FieldModelPMR_pkg_Detail::get_begin() const noexcept
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

void FieldModelPMR_pkg_Detail::get_end(size_t fbe_begin) const noexcept
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMR_pkg_Detail::get(::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = get_begin();
    if (fbe_begin == 0)
        return;

    uint32_t fbe_struct_size = unaligned_load<uint32_t>(_buffer.data() + _buffer.offset());
    get_fields(fbe_value, fbe_struct_size);
    get_end(fbe_begin);
}

void FieldModelPMR_pkg_Detail::get_fields([[maybe_unused]] ::FBE::Base& base_fbe_value, [[maybe_unused]] size_t fbe_struct_size) noexcept
{
    ::pkg_pmr::Detail& fbe_value = static_cast<::pkg_pmr::Detail&>(base_fbe_value);
    size_t fbe_current_size = 4 + 4;

    if ((fbe_current_size + extrav.fbe_size()) <= fbe_struct_size)
        {
            extrav.get(fbe_value.extrav);
        }
    else
        fbe_value.extrav.clear();
    fbe_current_size += extrav.fbe_size();

    if ((fbe_current_size + extram.fbe_size()) <= fbe_struct_size)
        {
            extram.get(fbe_value.extram);
        }
    else
        fbe_value.extram.clear();
    fbe_current_size += extram.fbe_size();
}

size_t FieldModelPMR_pkg_Detail::set_begin()
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

void FieldModelPMR_pkg_Detail::set_end(size_t fbe_begin)
{
    _buffer.unshift(fbe_begin);
}

void FieldModelPMR_pkg_Detail::set(const ::FBE::Base& fbe_value) noexcept
{
    size_t fbe_begin = set_begin();
    if (fbe_begin == 0)
        return;

    set_fields(fbe_value);
    set_end(fbe_begin);
}

void FieldModelPMR_pkg_Detail::set_fields([[maybe_unused]] const ::FBE::Base& base_fbe_value) noexcept
{
    [[maybe_unused]] const ::pkg_pmr::Detail& fbe_value = static_cast<const ::pkg_pmr::Detail&>(base_fbe_value);
    extrav.set(fbe_value.extrav);
    extram.set(fbe_value.extram);
}

namespace pkg_pmr {

bool DetailModel::verify()
{
    if ((this->buffer().offset() + model.fbe_offset() - 4) > this->buffer().size())
        return false;

    uint32_t fbe_full_size = unaligned_load<uint32_t>(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4);
    if (fbe_full_size < model.fbe_size())
        return false;

    return model.verify();
}

size_t DetailModel::create_begin()
{
    size_t fbe_begin = this->buffer().allocate(4 + model.fbe_size());
    return fbe_begin;
}

size_t DetailModel::create_end(size_t fbe_begin)
{
    size_t fbe_end = this->buffer().size();
    uint32_t fbe_full_size = (uint32_t)(fbe_end - fbe_begin);
    *((uint32_t*)(this->buffer().data() + this->buffer().offset() + model.fbe_offset() - 4)) = fbe_full_size;
    return fbe_full_size;
}

size_t DetailModel::serialize(const ::pkg_pmr::Detail& value)
{
    size_t fbe_begin = create_begin();
    model.set(value);
    size_t fbe_full_size = create_end(fbe_begin);
    return fbe_full_size;
}

size_t DetailModel::deserialize(::pkg_pmr::Detail& value) noexcept
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

} // namespace pkg_pmr

} // namespace FBE

//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: pkg.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

#pragma once

#ifdef isset
#undef isset
#endif

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe_custom_models.h"

#include "pkg_ptr.h"

#include "osa_models.h"

namespace FBE {

class FieldModelPtr_pkg_Info
{
public:
    FieldModelPtr_pkg_Info(FBEBuffer& buffer, size_t offset) noexcept;
    ~FieldModelPtr_pkg_Info();

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 5; }
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 1; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify() const noexcept;
    // Check if the struct ptr is nullptr
    bool has_value() const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::pkg::Info** fbe_value) noexcept;
    // Get the struct fields values
    void get_fields(::pkg::Info& fbe_value, size_t fbe_struct_size) noexcept;

    // Set the struct value (begin phase)
    size_t set_begin(bool has_value);
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::pkg::Info* fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::pkg::Info* fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    BaseFieldModel* ptr{nullptr};
};

class FieldModelPtr_pkg_Detail
{
public:
    FieldModelPtr_pkg_Detail(FBEBuffer& buffer, size_t offset) noexcept;
    ~FieldModelPtr_pkg_Detail();

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 5; }
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 2; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify() const noexcept;
    // Check if the struct ptr is nullptr
    bool has_value() const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::pkg::Detail** fbe_value) noexcept;
    // Get the struct fields values
    void get_fields(::pkg::Detail& fbe_value, size_t fbe_struct_size) noexcept;

    // Set the struct value (begin phase)
    size_t set_begin(bool has_value);
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::pkg::Detail* fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::pkg::Detail* fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    BaseFieldModel* ptr{nullptr};
};

// Fast Binary Encoding ::pkg::Info field model
class FieldModel_pkg_Info : public BaseFieldModel
{
public:
    FieldModel_pkg_Info(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept override { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept override { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept override;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 1; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept override { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept override { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept override;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept override;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept override;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept override;

    // Get the struct value
    void get(::FBE::Base& fbe_value) noexcept override;
    // Get the struct fields values
    void get_fields(::FBE::Base& fbe_value, size_t fbe_struct_size) noexcept override;

    // Set the struct value (begin phase)
    size_t set_begin() override;
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin) override;

    // Set the struct value
    void set(const ::FBE::Base& fbe_value) noexcept override;
    // Set the struct fields values
    void set_fields(const ::FBE::Base& fbe_value) noexcept override;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<stdb::memory::string> info;
    FieldModel<::osa::Sex> sex;
    FieldModel<::osa::MyFLags> flag;
    FieldModel<::osa::Extra> extra;
};

namespace pkg {

// Fast Binary Encoding Info model
class InfoModel : public FBE::Model
{
public:
    InfoModel() : model(this->buffer(), 4) {}
    InfoModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel_pkg_Info::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::pkg::Info& value);
    // Deserialize the struct value
    size_t deserialize(::pkg::Info& value) noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel_pkg_Info model;
};

} // namespace pkg

// Fast Binary Encoding ::pkg::Detail field model
class FieldModel_pkg_Detail : public BaseFieldModel
{
public:
    FieldModel_pkg_Detail(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept override { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept override { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept override;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 2; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept override { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept override { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept override;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept override;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept override;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept override;

    // Get the struct value
    void get(::FBE::Base& fbe_value) noexcept override;
    // Get the struct fields values
    void get_fields(::FBE::Base& fbe_value, size_t fbe_struct_size) noexcept override;

    // Set the struct value (begin phase)
    size_t set_begin() override;
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin) override;

    // Set the struct value
    void set(const ::FBE::Base& fbe_value) noexcept override;
    // Set the struct fields values
    void set_fields(const ::FBE::Base& fbe_value) noexcept override;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModelVector<::osa::Extra> extrav;
    FieldModelMap<int32_t, ::osa::Extra> extram;
};

namespace pkg {

// Fast Binary Encoding Detail model
class DetailModel : public FBE::Model
{
public:
    DetailModel() : model(this->buffer(), 4) {}
    DetailModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel_pkg_Detail::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::pkg::Detail& value);
    // Deserialize the struct value
    size_t deserialize(::pkg::Detail& value) noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel_pkg_Detail model;
};

} // namespace pkg

} // namespace FBE

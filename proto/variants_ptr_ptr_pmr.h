//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants_ptr.fbe
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

#include "fbe.h"
#include "arena/arena.hpp"

namespace variants_ptr_pmr {
using namespace FBE;
using allocator_type = pmr::polymorphic_allocator<char>;
} // namespace variants_ptr_pmr

namespace FBE {
using namespace ::variants_ptr_pmr;
} // namespace FBE

#include "fbe_ptr.h"

namespace variants_ptr_pmr {
// forward declaration
struct Simple;
struct Value;
struct ValueContainer;

using Expr = std::variant<bool, stdb::memory::arena_string, int32_t>;
std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Expr& value);
auto is_equal(const Expr& lhs, const Expr& rhs) -> bool;

using V = std::variant<int32_t, stdb::memory::arena_string, double, ::variants_ptr_pmr::Simple, ::variants_ptr_pmr::Simple*, pmr::vector<::variants_ptr_pmr::Simple>, pmr::vector<int32_t>, pmr::unordered_map<int32_t, ::variants_ptr_pmr::Simple>, pmr::vector<FBE::pmr_buffer_t>, pmr::vector<stdb::memory::arena_string>, pmr::unordered_map<int32_t, FBE::pmr_buffer_t>, pmr::unordered_map<stdb::memory::arena_string, FBE::pmr_buffer_t>, pmr::vector<::variants_ptr_pmr::Simple*>, ::variants_ptr_pmr::Expr>;
std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const V& value);
auto is_equal(const V& lhs, const V& rhs) -> bool;

struct Simple : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    stdb::memory::arena_string name;

    size_t fbe_type() const noexcept { return 1; }

    Simple();
    explicit Simple(allocator_type alloc);
    explicit Simple(const stdb::memory::arena_string& arg_name);
    Simple(const Simple& other) = default;
    Simple(Simple&& other) noexcept;
    ~Simple() override;

    Simple& operator=(const Simple& other) = default;
    Simple& operator=(Simple&& other) noexcept;

    bool operator==(const Simple& other) const noexcept;
    bool operator!=(const Simple& other) const noexcept { return !operator==(other); }
    bool operator<(const Simple& other) const noexcept;
    bool operator<=(const Simple& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Simple& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Simple& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const Simple& value);

    void swap(Simple& other) noexcept;
    friend void swap(Simple& value1, Simple& value2) noexcept { value1.swap(value2); }
};

} // namespace variants_ptr_pmr

template<>
struct std::hash<variants_ptr_pmr::Simple>
{
    typedef variants_ptr_pmr::Simple argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace variants_ptr_pmr {

struct Value : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    ::variants_ptr_pmr::V v;
    std::optional<::variants_ptr_pmr::V> vo;
    std::optional<::variants_ptr_pmr::V> vo2;

    size_t fbe_type() const noexcept { return 2; }

    Value();
    explicit Value(allocator_type alloc);
    Value(::variants_ptr_pmr::V arg_v, std::optional<::variants_ptr_pmr::V> arg_vo, std::optional<::variants_ptr_pmr::V> arg_vo2);
    Value(const Value& other) = default;
    Value(Value&& other) noexcept;
    ~Value() override;

    Value& operator=(const Value& other) = default;
    Value& operator=(Value&& other) noexcept;

    bool operator==(const Value& other) const noexcept;
    bool operator!=(const Value& other) const noexcept { return !operator==(other); }
    bool operator<(const Value& other) const noexcept;
    bool operator<=(const Value& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Value& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Value& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const Value& value);

    void swap(Value& other) noexcept;
    friend void swap(Value& value1, Value& value2) noexcept { value1.swap(value2); }
};

} // namespace variants_ptr_pmr

template<>
struct std::hash<variants_ptr_pmr::Value>
{
    typedef variants_ptr_pmr::Value argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace variants_ptr_pmr {

struct ValueContainer : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    pmr::vector<::variants_ptr_pmr::V> vv;
    pmr::unordered_map<int32_t, ::variants_ptr_pmr::V> vm;

    size_t fbe_type() const noexcept { return 3; }

    ValueContainer();
    explicit ValueContainer(allocator_type alloc);
    ValueContainer(pmr::vector<::variants_ptr_pmr::V> arg_vv, pmr::unordered_map<int32_t, ::variants_ptr_pmr::V> arg_vm);
    ValueContainer(const ValueContainer& other) = default;
    ValueContainer(ValueContainer&& other) noexcept;
    ~ValueContainer() override;

    ValueContainer& operator=(const ValueContainer& other) = default;
    ValueContainer& operator=(ValueContainer&& other) noexcept;

    bool operator==(const ValueContainer& other) const noexcept;
    bool operator!=(const ValueContainer& other) const noexcept { return !operator==(other); }
    bool operator<(const ValueContainer& other) const noexcept;
    bool operator<=(const ValueContainer& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const ValueContainer& other) const noexcept { return !operator<=(other); }
    bool operator>=(const ValueContainer& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const ValueContainer& value);

    void swap(ValueContainer& other) noexcept;
    friend void swap(ValueContainer& value1, ValueContainer& value2) noexcept { value1.swap(value2); }
};

} // namespace variants_ptr_pmr

template<>
struct std::hash<variants_ptr_pmr::ValueContainer>
{
    typedef variants_ptr_pmr::ValueContainer argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace variants_ptr_pmr {

} // namespace variants_ptr_pmr

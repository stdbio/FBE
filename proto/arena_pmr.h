//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: arena.fbe
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

#include "arena_common_pmr.h"

namespace arena_pmr {
using namespace FBE;
using allocator_type = pmr::polymorphic_allocator<char>;
} // namespace arena_pmr

namespace FBE {
using namespace ::arena_pmr;
} // namespace FBE

namespace arena_pmr {

struct Item
{
    ArenaManagedCreateOnlyTag;

    ::arena_common_pmr::Optr optr;
    ::arena_common_pmr::Alias alias;
    pmr::vector<::arena_common_pmr::Expression> expressions;
    pmr::map<int32_t, ::arena_common_pmr::Alias> aliases_int;

    size_t fbe_type() const noexcept { return 1; }

    Item();
    explicit Item(allocator_type alloc);
    Item(const ::arena_common_pmr::Optr& arg_optr, const ::arena_common_pmr::Alias& arg_alias, const pmr::vector<::arena_common_pmr::Expression>& arg_expressions, const pmr::map<int32_t, ::arena_common_pmr::Alias>& arg_aliases_int);
    Item(const Item& other) = default;
    Item(Item&& other) = default;
    ~Item() = default;

    Item& operator=(const Item& other) = default;
    Item& operator=(Item&& other) = default;

    bool operator==(const Item& other) const noexcept;
    bool operator!=(const Item& other) const noexcept { return !operator==(other); }
    bool operator<(const Item& other) const noexcept;
    bool operator<=(const Item& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Item& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Item& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Item& value);

    void swap(Item& other) noexcept;
    friend void swap(Item& value1, Item& value2) noexcept { value1.swap(value2); }
};

} // namespace arena_pmr

template<>
struct std::hash<arena_pmr::Item>
{
    typedef arena_pmr::Item argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace arena_pmr {

struct Item2
{
    ArenaManagedCreateOnlyTag;

    FBE::pmr_buffer_t bytes_v;

    size_t fbe_type() const noexcept { return 2; }

    Item2();
    explicit Item2(allocator_type alloc);
    explicit Item2(const FBE::pmr_buffer_t& arg_bytes_v);
    Item2(const Item2& other) = default;
    Item2(Item2&& other) = default;
    ~Item2() = default;

    Item2& operator=(const Item2& other) = default;
    Item2& operator=(Item2&& other) = default;

    bool operator==(const Item2& other) const noexcept;
    bool operator!=(const Item2& other) const noexcept { return !operator==(other); }
    bool operator<(const Item2& other) const noexcept;
    bool operator<=(const Item2& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Item2& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Item2& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Item2& value);

    void swap(Item2& other) noexcept;
    friend void swap(Item2& value1, Item2& value2) noexcept { value1.swap(value2); }
};

} // namespace arena_pmr

template<>
struct std::hash<arena_pmr::Item2>
{
    typedef arena_pmr::Item2 argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace arena_pmr {

} // namespace arena_pmr

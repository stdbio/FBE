//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: arena_ptr.fbe
// FBE version: 1.10.0.0
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

#include "arena_common.h"

namespace arena_ptr {
using namespace FBE;
using namespace ::arena_common;
} // namespace arena_ptr

namespace FBE {
using namespace ::arena_ptr;
} // namespace FBE

#include "fbe_ptr.h"

namespace arena_ptr {

struct Line : FBE::Base
{
    ::arena_common::Expression expression;

    size_t fbe_type() const noexcept { return 1; }

    Line();
    explicit Line(::arena_common::Expression arg_expression);
    Line(const Line& other) = default;
    Line(Line&& other) noexcept;
    ~Line() override;

    Line& operator=(const Line& other) = default;
    Line& operator=(Line&& other) noexcept;

    bool operator==(const Line& other) const noexcept;
    bool operator!=(const Line& other) const noexcept { return !operator==(other); }
    bool operator<(const Line& other) const noexcept;
    bool operator<=(const Line& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Line& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Line& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const Line& value);

    void swap(Line& other) noexcept;
    friend void swap(Line& value1, Line& value2) noexcept { value1.swap(value2); }
};

} // namespace arena_ptr

template<>
struct std::hash<arena_ptr::Line>
{
    typedef arena_ptr::Line argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace arena_ptr {

struct Line2 : FBE::Base
{
    FBE::buffer_t bytes_v;

    size_t fbe_type() const noexcept { return 2; }

    Line2();
    explicit Line2(const FBE::buffer_t& arg_bytes_v);
    Line2(const Line2& other) = default;
    Line2(Line2&& other) noexcept;
    ~Line2() override;

    Line2& operator=(const Line2& other) = default;
    Line2& operator=(Line2&& other) noexcept;

    bool operator==(const Line2& other) const noexcept;
    bool operator!=(const Line2& other) const noexcept { return !operator==(other); }
    bool operator<(const Line2& other) const noexcept;
    bool operator<=(const Line2& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Line2& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Line2& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const Line2& value);

    void swap(Line2& other) noexcept;
    friend void swap(Line2& value1, Line2& value2) noexcept { value1.swap(value2); }
};

} // namespace arena_ptr

template<>
struct std::hash<arena_ptr::Line2>
{
    typedef arena_ptr::Line2 argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace arena_ptr {

} // namespace arena_ptr

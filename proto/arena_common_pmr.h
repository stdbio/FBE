//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: arena_common.fbe
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

namespace arena_common_pmr {
using namespace FBE;
using allocator_type = pmr::polymorphic_allocator<char>;
} // namespace arena_common_pmr

namespace FBE {
using namespace ::arena_common_pmr;
} // namespace FBE

namespace arena_common_pmr {

enum class Optr : uint8_t
{
    EQ = (uint8_t)0u,
    GT,
    LT,
    GE,
    LE,
};

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] Optr value);

#if defined(LOGGING_PROTOCOL)
CppLogging::Record& operator<<(CppLogging::Record& record, Optr value);
#endif
// forward declaration
struct Alias;
struct Expression;

using Expr = std::variant<std::monostate, bool, int32_t, ArenaString>;
std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Expr& value);

struct Alias
{
    ArenaManagedCreateOnlyTag;

    ArenaString name;
    ::arena_common_pmr::Optr optr;
    ::arena_common_pmr::Expr expr;

    size_t fbe_type() const noexcept { return 1; }

    Alias();
    explicit Alias(allocator_type alloc);
    Alias(const ArenaString& arg_name, const ::arena_common_pmr::Optr& arg_optr, const ::arena_common_pmr::Expr& arg_expr);
    Alias(const Alias& other) = default;
    Alias(Alias&& other) = default;
    ~Alias() = default;

    Alias& operator=(const Alias& other) = default;
    Alias& operator=(Alias&& other) = default;

    bool operator==(const Alias& other) const noexcept;
    bool operator!=(const Alias& other) const noexcept { return !operator==(other); }
    bool operator<(const Alias& other) const noexcept;
    bool operator<=(const Alias& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Alias& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Alias& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Alias& value);

    void swap(Alias& other) noexcept;
    friend void swap(Alias& value1, Alias& value2) noexcept { value1.swap(value2); }
};

} // namespace arena_common_pmr

template<>
struct std::hash<arena_common_pmr::Alias>
{
    typedef arena_common_pmr::Alias argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace arena_common_pmr {

struct Expression
{
    ArenaManagedCreateOnlyTag;

    pmr::vector<ArenaString> keys;
    pmr::vector<::arena_common_pmr::Alias> aliases;
    pmr::map<int32_t, ::arena_common_pmr::Alias> alias_int;

    size_t fbe_type() const noexcept { return 2; }

    Expression();
    explicit Expression(allocator_type alloc);
    Expression(const pmr::vector<ArenaString>& arg_keys, const pmr::vector<::arena_common_pmr::Alias>& arg_aliases, const pmr::map<int32_t, ::arena_common_pmr::Alias>& arg_alias_int);
    Expression(const Expression& other) = default;
    Expression(Expression&& other) = default;
    ~Expression() = default;

    Expression& operator=(const Expression& other) = default;
    Expression& operator=(Expression&& other) = default;

    bool operator==(const Expression& other) const noexcept;
    bool operator!=(const Expression& other) const noexcept { return !operator==(other); }
    bool operator<(const Expression& other) const noexcept;
    bool operator<=(const Expression& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Expression& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Expression& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Expression& value);

    void swap(Expression& other) noexcept;
    friend void swap(Expression& value1, Expression& value2) noexcept { value1.swap(value2); }
};

} // namespace arena_common_pmr

template<>
struct std::hash<arena_common_pmr::Expression>
{
    typedef arena_common_pmr::Expression argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace arena_common_pmr {

} // namespace arena_common_pmr

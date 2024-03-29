//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: osa.fbe
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

namespace osa_pmr {
using namespace FBE;
using allocator_type = pmr::polymorphic_allocator<char>;
} // namespace osa_pmr

namespace FBE {
using namespace ::osa_pmr;
} // namespace FBE

namespace osa_pmr {

enum class Sex
{
    male,
    female,
};

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] Sex value);

#if defined(LOGGING_PROTOCOL)
CppLogging::Record& operator<<(CppLogging::Record& record, Sex value);
#endif

enum class MyFLags
{
    flag0 = (int32_t)0x00ll,
    flag1 = (int32_t)0x01ll,
    flag2 = (int32_t)0x02ll,
};

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] MyFLags value);

#if defined(LOGGING_PROTOCOL)
CppLogging::Record& operator<<(CppLogging::Record& record, MyFLags value);
#endif

struct Extra
{
    ArenaManagedCreateOnlyTag;

    ArenaString name;
    ArenaString detail;
    ::osa_pmr::Sex sex;
    ::osa_pmr::MyFLags flag;

    size_t fbe_type() const noexcept { return 1; }

    Extra();
    explicit Extra(allocator_type alloc);
    Extra(const ArenaString& arg_name, const ArenaString& arg_detail, const ::osa_pmr::Sex& arg_sex, const ::osa_pmr::MyFLags& arg_flag);
    Extra(const Extra& other) = default;
    Extra(Extra&& other) = default;
    ~Extra() = default;

    Extra& operator=(const Extra& other) = default;
    Extra& operator=(Extra&& other) = default;

    bool operator==(const Extra& other) const noexcept;
    bool operator!=(const Extra& other) const noexcept { return !operator==(other); }
    bool operator<(const Extra& other) const noexcept;
    bool operator<=(const Extra& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Extra& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Extra& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Extra& value);

    void swap(Extra& other) noexcept;
    friend void swap(Extra& value1, Extra& value2) noexcept { value1.swap(value2); }
};

} // namespace osa_pmr

template<>
struct std::hash<osa_pmr::Extra>
{
    typedef osa_pmr::Extra argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace osa_pmr {

struct Simple
{
    ArenaManagedCreateOnlyTag;

    ArenaString name;
    int32_t depth;
    std::array<::osa_pmr::Extra, 1> sa;
    ::osa_pmr::Sex sex;

    size_t fbe_type() const noexcept { return 2; }

    Simple();
    explicit Simple(allocator_type alloc);
    Simple(const ArenaString& arg_name, int32_t arg_depth, const std::array<::osa_pmr::Extra, 1>& arg_sa, const ::osa_pmr::Sex& arg_sex);
    Simple(const Simple& other) = default;
    Simple(Simple&& other) = default;
    ~Simple() = default;

    Simple& operator=(const Simple& other) = default;
    Simple& operator=(Simple&& other) = default;

    bool operator==(const Simple& other) const noexcept;
    bool operator!=(const Simple& other) const noexcept { return !operator==(other); }
    bool operator<(const Simple& other) const noexcept;
    bool operator<=(const Simple& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Simple& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Simple& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Simple& value);

    void swap(Simple& other) noexcept;
    friend void swap(Simple& value1, Simple& value2) noexcept { value1.swap(value2); }
};

} // namespace osa_pmr

template<>
struct std::hash<osa_pmr::Simple>
{
    typedef osa_pmr::Simple argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace osa_pmr {

} // namespace osa_pmr

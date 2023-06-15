//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: sa.fbe
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

namespace sa_pmr {
using namespace FBE;
using allocator_type = pmr::polymorphic_allocator<char>;
} // namespace sa_pmr

namespace FBE {
using namespace ::sa_pmr;
} // namespace FBE

#include "fbe_ptr.h"

namespace sa_pmr {

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

struct Extra : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    stdb::memory::arena_string name;
    stdb::memory::arena_string detail;
    ::sa_pmr::Sex sex;
    ::sa_pmr::MyFLags flag;

    size_t fbe_type() const noexcept { return 1; }

    Extra();
    explicit Extra(allocator_type alloc);
    Extra(const stdb::memory::arena_string& arg_name, const stdb::memory::arena_string& arg_detail, ::sa_pmr::Sex arg_sex, ::sa_pmr::MyFLags arg_flag);
    Extra(const Extra& other) = default;
    Extra(Extra&& other) noexcept;
    ~Extra() override;

    Extra& operator=(const Extra& other) = default;
    Extra& operator=(Extra&& other) noexcept;

    bool operator==(const Extra& other) const noexcept;
    bool operator!=(const Extra& other) const noexcept { return !operator==(other); }
    bool operator<(const Extra& other) const noexcept;
    bool operator<=(const Extra& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Extra& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Extra& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const Extra& value);

    void swap(Extra& other) noexcept;
    friend void swap(Extra& value1, Extra& value2) noexcept { value1.swap(value2); }
};

} // namespace sa_pmr

template<>
struct std::hash<sa_pmr::Extra>
{
    typedef sa_pmr::Extra argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace sa_pmr {

struct Simple : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    stdb::memory::arena_string name;
    int32_t depth;
    std::array<::sa_pmr::Extra, 1> sa;
    ::sa_pmr::Sex sex;

    size_t fbe_type() const noexcept { return 2; }

    Simple();
    explicit Simple(allocator_type alloc);
    Simple(const stdb::memory::arena_string& arg_name, int32_t arg_depth, std::array<::sa_pmr::Extra, 1> arg_sa, ::sa_pmr::Sex arg_sex);
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

} // namespace sa_pmr

template<>
struct std::hash<sa_pmr::Simple>
{
    typedef sa_pmr::Simple argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace sa_pmr {

struct Complex : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    stdb::memory::arena_string name;
    std::optional<::sa_pmr::Sex> sex;
    std::optional<::sa_pmr::MyFLags> flag;
    std::optional<::sa_pmr::Extra> extra;
    pmr::vector<int64_t> nums;

    size_t fbe_type() const noexcept { return 3; }

    Complex();
    explicit Complex(allocator_type alloc);
    Complex(const stdb::memory::arena_string& arg_name, std::optional<::sa_pmr::Sex> arg_sex, std::optional<::sa_pmr::MyFLags> arg_flag, std::optional<::sa_pmr::Extra> arg_extra, pmr::vector<int64_t> arg_nums);
    Complex(const Complex& other) = default;
    Complex(Complex&& other) noexcept;
    ~Complex() override;

    Complex& operator=(const Complex& other) = default;
    Complex& operator=(Complex&& other) noexcept;

    bool operator==(const Complex& other) const noexcept;
    bool operator!=(const Complex& other) const noexcept { return !operator==(other); }
    bool operator<(const Complex& other) const noexcept;
    bool operator<=(const Complex& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Complex& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Complex& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, const Complex& value);

    void swap(Complex& other) noexcept;
    friend void swap(Complex& value1, Complex& value2) noexcept { value1.swap(value2); }
};

} // namespace sa_pmr

template<>
struct std::hash<sa_pmr::Complex>
{
    typedef sa_pmr::Complex argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace sa_pmr {

} // namespace sa_pmr

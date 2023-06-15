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

namespace sa {
using namespace FBE;
} // namespace sa

namespace FBE {
using namespace ::sa;
} // namespace FBE

#include "fbe_ptr.h"

namespace sa {

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
    stdb::memory::string name;
    stdb::memory::string detail;
    ::sa::Sex sex;
    ::sa::MyFLags flag;

    size_t fbe_type() const noexcept { return 1; }

    Extra();
    Extra(const stdb::memory::string& arg_name, const stdb::memory::string& arg_detail, ::sa::Sex arg_sex, ::sa::MyFLags arg_flag);
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

} // namespace sa

template<>
struct std::hash<sa::Extra>
{
    typedef sa::Extra argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace sa {

struct Simple : FBE::Base
{
    stdb::memory::string name;
    int32_t depth;
    std::array<::sa::Extra, 1> sa;
    ::sa::Sex sex;

    size_t fbe_type() const noexcept { return 2; }

    Simple();
    Simple(const stdb::memory::string& arg_name, int32_t arg_depth, std::array<::sa::Extra, 1> arg_sa, ::sa::Sex arg_sex);
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

} // namespace sa

template<>
struct std::hash<sa::Simple>
{
    typedef sa::Simple argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace sa {

struct Complex : FBE::Base
{
    stdb::memory::string name;
    std::optional<::sa::Sex> sex;
    std::optional<::sa::MyFLags> flag;
    std::optional<::sa::Extra> extra;
    FastVec<int64_t> nums;

    size_t fbe_type() const noexcept { return 3; }

    Complex();
    Complex(const stdb::memory::string& arg_name, std::optional<::sa::Sex> arg_sex, std::optional<::sa::MyFLags> arg_flag, std::optional<::sa::Extra> arg_extra, FastVec<int64_t> arg_nums);
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

} // namespace sa

template<>
struct std::hash<sa::Complex>
{
    typedef sa::Complex argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace sa {

} // namespace sa

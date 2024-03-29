//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: cpp_only.fbe
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

namespace cpp_only {
using namespace FBE;
} // namespace cpp_only

namespace FBE {
using namespace ::cpp_only;
} // namespace FBE

namespace cpp_only {
// forward declaration
struct Struct128;

using LargeNum = std::variant<std::monostate, int64_t, __int128_t, __uint128_t, FastVec<__int128_t>, std::unordered_map<__uint128_t, __int128_t>>;
std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const LargeNum& value);

struct Struct128
{
    __int128_t f1;
    std::optional<__int128_t> f2;
    __uint128_t f3;
    std::optional<__uint128_t> f4;
    FastVec<__int128_t> f5;
    std::unordered_map<__uint128_t, __int128_t> f6;
    ::cpp_only::LargeNum f7;

    size_t fbe_type() const noexcept { return 1; }

    Struct128();
    Struct128(__int128_t arg_f1, const std::optional<__int128_t>& arg_f2, __uint128_t arg_f3, const std::optional<__uint128_t>& arg_f4, const FastVec<__int128_t>& arg_f5, const std::unordered_map<__uint128_t, __int128_t>& arg_f6, const ::cpp_only::LargeNum& arg_f7);
    Struct128(const Struct128& other) = default;
    Struct128(Struct128&& other) = default;
    ~Struct128() = default;

    Struct128& operator=(const Struct128& other) = default;
    Struct128& operator=(Struct128&& other) = default;

    bool operator==(const Struct128& other) const noexcept;
    bool operator!=(const Struct128& other) const noexcept { return !operator==(other); }
    bool operator<(const Struct128& other) const noexcept;
    bool operator<=(const Struct128& other) const noexcept { return operator<(other) || operator==(other); }
    bool operator>(const Struct128& other) const noexcept { return !operator<=(other); }
    bool operator>=(const Struct128& other) const noexcept { return !operator<(other); }

    std::string string() const;

    friend std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Struct128& value);

    void swap(Struct128& other) noexcept;
    friend void swap(Struct128& value1, Struct128& value2) noexcept { value1.swap(value2); }
};

} // namespace cpp_only

template<>
struct std::hash<cpp_only::Struct128>
{
    typedef cpp_only::Struct128 argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace cpp_only {

} // namespace cpp_only

//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: simple.fbe
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

namespace simple_pmr {
using namespace FBE;
using allocator_type = pmr::polymorphic_allocator<char>;
} // namespace simple_pmr

namespace FBE {
using namespace ::simple_pmr;
} // namespace FBE

#include "fbe_ptr.h"

namespace simple_pmr {

struct Simple : FBE::Base
{
    ArenaManagedCreateOnlyTag;

    ArenaString info;
    ::simple_pmr::Simple* simple;
    int32_t depth;
    pmr::vector<::simple_pmr::Simple*> spv;
    pmr::vector<::simple_pmr::Simple> sv;
    pmr::map<int32_t, ::simple_pmr::Simple*> spm;
    pmr::map<int32_t, ::simple_pmr::Simple> sm;

    size_t fbe_type() const noexcept { return 1; }

    Simple();
    explicit Simple(allocator_type alloc);
    Simple(const ArenaString& arg_info, std::unique_ptr<::simple_pmr::Simple> arg_simple, int32_t arg_depth, pmr::vector<std::unique_ptr<::simple_pmr::Simple>> arg_spv, pmr::vector<::simple_pmr::Simple> arg_sv, pmr::map<int32_t, std::unique_ptr<::simple_pmr::Simple>> arg_spm, pmr::map<int32_t, ::simple_pmr::Simple> arg_sm);
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

} // namespace simple_pmr

template<>
struct std::hash<simple_pmr::Simple>
{
    typedef simple_pmr::Simple argument_type;
    typedef size_t result_type;

    result_type operator() ([[maybe_unused]] const argument_type& value) const
    {
        result_type result = 17;
        return result;
    }
};

namespace simple_pmr {

} // namespace simple_pmr

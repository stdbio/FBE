//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: simple.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

#include "simple_ptr.h"

namespace simple {

Simple::Simple()
    : info()
    , simple(nullptr)
    , depth((int32_t)0ll)
    , spv()
    , sv()
    , spm()
    , sm()
{}

Simple::Simple(const FBEString& arg_info, std::unique_ptr<::simple::Simple> arg_simple, int32_t arg_depth, FastVec<std::unique_ptr<::simple::Simple>> arg_spv, FastVec<::simple::Simple> arg_sv, std::map<int32_t, std::unique_ptr<::simple::Simple>> arg_spm, std::map<int32_t, ::simple::Simple> arg_sm)
    : info(arg_info)
    , simple(arg_simple.release())
    , depth(arg_depth)
    , spv()
    , sv(std::move(arg_sv))
    , spm()
    , sm(std::move(arg_sm))
{
    spv.reserve(arg_spv.size());
    for (auto& it : arg_spv)
        #if defined(USING_STD_VECTOR)
        spv.emplace_back(it.release());
        #else
        spv.emplace_back<Safety::Unsafe>(it.release());
        #endif
    for (auto& it: arg_spm)
        spm.emplace(it.first, it.second.release());
}

Simple::Simple([[maybe_unused]] Simple&& other) noexcept
    : info(std::move(other.info))
    , simple(std::exchange(other.simple, nullptr))
    , depth(std::exchange(other.depth, (int32_t)0ll))
    , spv(std::move(other.spv))
    , sv(std::move(other.sv))
    , spm(std::move(other.spm))
    , sm(std::move(other.sm))
{}

Simple::~Simple()
{
    if (simple) delete simple;
    for (auto* it : spv)
        delete it;
    for (auto& it: spm)
        delete it.second;
}

bool Simple::operator==([[maybe_unused]] const Simple& other) const noexcept
{
    if (info != other.info)
        return false;
    // compare ptr simple
    if ((simple  == nullptr && other.simple  != nullptr) || (simple  != nullptr && other.simple  == nullptr) || (simple  != nullptr && other.simple  != nullptr && *simple != *other.simple))
        return false;
    if (depth != other.depth)
        return false;
    // compare container spv
    if (spv.size() != other.spv.size())
        return false;
    for (size_t i = 0; i < spv.size(); i++)
    {
        if (*spv[i] != *other.spv[i])
            return false;
    }
    // compare container sv
    if (sv != other.sv)
        return false;
    // compare container spm
    if (spm.size() != other.spm.size())
        return false;
    for (auto & [k, v]: spm)
    {
        if (auto pos = other.spm.find(k); pos == other.spm.end())
            return false;
        if (auto other_v = other.spm.at(k); *other_v != *v)
            return false;
    }
    // compare container sm
    if (sm != other.sm)
        return false;
    return true;
}

bool Simple::operator<([[maybe_unused]] const Simple& other) const noexcept
{
    return false;
}

Simple& Simple::operator=(Simple&& other) noexcept
{
    if (this != &other)
    {
        info = std::move(other.info);
        simple = std::exchange(other.simple, nullptr);
        depth = std::exchange(other.depth, (int32_t)0ll);
        spv = std::move(other.spv);
        sv = std::move(other.sv);
        spm = std::move(other.spm);
        sm = std::move(other.sm);
    }
    return *this;
}

std::string Simple::string() const
{
    std::stringstream ss; ss << *this; return ss.str();
}

void Simple::swap([[maybe_unused]] Simple& other) noexcept
{
    using std::swap;
    swap(info, other.info);
    swap(simple, other.simple);
    swap(depth, other.depth);
    swap(spv, other.spv);
    swap(sv, other.sv);
    swap(spm, other.spm);
    swap(sm, other.sm);
}

std::ostream& operator<<(std::ostream& stream, [[maybe_unused]] const Simple& value)
{
    stream << "Simple(";
    stream << "info="; stream << "\"" << value.info << "\"";
    stream << ",simple="; stream << "ptr of other struct: " << (value.simple == nullptr ? "nullptr" : "true");
    stream << ",depth="; stream << value.depth;
    {
        bool first = true;
        stream << ",spv=[" << value.spv.size() << "][";
        for ([[maybe_unused]] const auto& it : value.spv)
        {
            stream << std::string(first ? "" : ",") << "ptr of other struct: " << (it == nullptr ? "nullptr" : "true");
            first = false;
        }
        stream << "]";
    }
    {
        bool first = true;
        stream << ",sv=[" << value.sv.size() << "][";
        for ([[maybe_unused]] const auto& it : value.sv)
        {
            stream << std::string(first ? "" : ",") << it;
            first = false;
        }
        stream << "]";
    }
    {
        bool first = true;
        stream << ",spm=[" << value.spm.size()<< "]<{";
        for ([[maybe_unused]] const auto& it : value.spm)
        {
            stream << std::string(first ? "" : ",") << it.first;
            stream << "->";
            stream << "ptr of other struct: " << (it.second == nullptr ? "nullptr" : "true");
            first = false;
        }
        stream << "}>";
    }
    {
        bool first = true;
        stream << ",sm=[" << value.sm.size()<< "]<{";
        for ([[maybe_unused]] const auto& it : value.sm)
        {
            stream << std::string(first ? "" : ",") << it.first;
            stream << "->";
            stream << it.second;
            first = false;
        }
        stream << "}>";
    }
    stream << ")";
    return stream;
}

} // namespace simple

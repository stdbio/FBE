//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants.fbe
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

#include "fbe_json.h"

#include "variants.h"

namespace FBE {

namespace JSON {

template <class TWriter>
struct ValueWriter<TWriter, ::variants::Simple>
{
    static bool to_json(TWriter& writer, const ::variants::Simple& value, bool scope = true)
    {
        if (scope)
            if (!writer.StartObject())
                return false;
        if (!FBE::JSON::to_json_key(writer, "name") || !FBE::JSON::to_json(writer, value.name, true))
            return false;
        if (scope)
            if (!writer.EndObject())
                return false;
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, ::variants::Simple>
{
    static bool from_json(const TJson& json, ::variants::Simple& value, const char* key = nullptr)
    {
        if (key != nullptr)
            return FBE::JSON::from_json_child(json, value, key);
        bool result = true;
        result &= FBE::JSON::from_json(json, value.name, "name");
        return result;
    }
};

template <class TWriter>
struct ValueWriter<TWriter, ::variants::Value>
{
    static bool to_json(TWriter& writer, const ::variants::Value& value, bool scope = true)
    {
        if (scope)
            if (!writer.StartObject())
                return false;
        if (!FBE::JSON::to_json_key(writer, "v") || !FBE::JSON::to_json(writer, value.v, true))
            return false;
        if (scope)
            if (!writer.EndObject())
                return false;
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, ::variants::Value>
{
    static bool from_json(const TJson& json, ::variants::Value& value, const char* key = nullptr)
    {
        if (key != nullptr)
            return FBE::JSON::from_json_child(json, value, key);
        bool result = true;
        result &= FBE::JSON::from_json(json, value.v, "v");
        return result;
    }
};

} // namespace JSON

} // namespace FBE

//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums.fbe

class EnumByteJson : com.google.gson.JsonSerializer<com.chronoxor.enums.EnumByte>, com.google.gson.JsonDeserializer<com.chronoxor.enums.EnumByte>
{
    override fun serialize(src: com.chronoxor.enums.EnumByte, typeOfSrc: java.lang.reflect.Type, context: com.google.gson.JsonSerializationContext): com.google.gson.JsonElement
    {
        return com.google.gson.JsonPrimitive(src.raw)
    }

    @Throws(com.google.gson.JsonParseException::class)
    override fun deserialize(json: com.google.gson.JsonElement, type: java.lang.reflect.Type, context: com.google.gson.JsonDeserializationContext): com.chronoxor.enums.EnumByte
    {
        return com.chronoxor.enums.EnumByte(json.asJsonPrimitive.asByte)
    }
}
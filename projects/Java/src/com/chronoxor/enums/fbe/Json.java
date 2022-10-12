//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

package com.chronoxor.enums.fbe;

// Fast Binary Encoding enums JSON engine
public final class Json
{
    private static final com.google.gson.Gson _engine;

    // Get the JSON engine
    public static com.google.gson.Gson getEngine() { return _engine; }

    static
    {
        _engine = register(new com.google.gson.GsonBuilder()).create();
    }

    private Json() {}

    public static com.google.gson.GsonBuilder register(com.google.gson.GsonBuilder builder)
    {
        com.chronoxor.fbe.Json.register(builder);
        builder.registerTypeAdapter(com.chronoxor.enums.EnumByte.class, new EnumByteJson());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumChar.class, new EnumCharJson());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumWChar.class, new EnumWCharJson());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumInt8.class, new EnumInt8Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumUInt8.class, new EnumUInt8Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumInt16.class, new EnumInt16Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumUInt16.class, new EnumUInt16Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumInt32.class, new EnumInt32Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumUInt32.class, new EnumUInt32Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumInt64.class, new EnumInt64Json());
        builder.registerTypeAdapter(com.chronoxor.enums.EnumUInt64.class, new EnumUInt64Json());
        return builder;
    }
}
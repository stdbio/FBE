//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums.fbe

// Fast Binary Encoding com.chronoxor.enums final receiver
@Suppress("MemberVisibilityCanBePrivate", "PrivatePropertyName", "UNUSED_PARAMETER")
open class FinalReceiver : com.chronoxor.fbe.Receiver, IFinalReceiverListener
{
    // Receiver values accessors

    // Receiver models accessors

    constructor() : super(true)
    {
    }

    constructor(buffer: com.chronoxor.fbe.Buffer) : super(buffer, true)
    {
    }

    override fun onReceive(type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {
        return onReceiveListener(this, type, buffer, offset, size)
    }

    open fun onReceiveListener(listener: IFinalReceiverListener, type: Long, buffer: ByteArray, offset: Long, size: Long): Boolean
    {

        return false
    }
}
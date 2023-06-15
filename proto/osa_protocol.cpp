//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: osa.fbe
// FBE version: 1.11.0.0
//------------------------------------------------------------------------------

#if defined(_MSC_VER)
#pragma warning(push)
#pragma warning(disable:4065) // C4065: switch statement contains 'default' but no 'case' labels
#endif

#include "osa_protocol.h"

namespace FBE {

namespace osa {

bool Receiver::onReceive(size_t type, const void* data, size_t size)
{
    switch (type)
    {
        default: break;
    }

    return false;
}

bool Proxy::onReceive(size_t type, const void* data, size_t size)
{
    switch (type)
    {
        default: break;
    }

    return false;
}

void Client::reset_requests()
{
    Sender::reset();
    Receiver::reset();
}

void Client::watchdog_requests(uint64_t utc)
{
}

} // namespace osa

} // namespace FBE

#if defined(_MSC_VER)
#pragma warning(pop)
#endif

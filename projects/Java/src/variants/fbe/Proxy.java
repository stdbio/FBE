//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: variants.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

package variants.fbe;

// Fast Binary Encoding variants proxy
public class Proxy extends fbe.Receiver
{
    // Proxy models accessors

    public Proxy()
    {
        super(false);
    }
    public Proxy(fbe.Buffer buffer)
    {
        super(buffer, false);
    }

    // Proxy handlers

    @Override
    public boolean onReceive(long type, byte[] buffer, long offset, long size)
    {
        switch ((int)type)
        {
            default: break;
        }

        return false;
    }
}

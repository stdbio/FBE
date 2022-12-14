//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: osa.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

package osa;

public final class Sex implements Comparable<Sex>
{
    public static final Sex male = new Sex(SexEnum.male);
    public static final Sex female = new Sex(SexEnum.female);

    private SexEnum value = SexEnum.values()[0];

    public Sex() {}
    public Sex(int value) { setEnum(value); }
    public Sex(SexEnum value) { setEnum(value); }
    public Sex(Sex value) { setEnum(value); }

    public SexEnum getEnum() { return value; }
    public int getRaw() { return value.getRaw(); }

    public void setDefault() { setEnum((int)0); }

    public void setEnum(int value) { this.value = SexEnum.mapValue(value); }
    public void setEnum(SexEnum value) { this.value = value; }
    public void setEnum(Sex value) { this.value = value.value; }

    @Override
    public int compareTo(Sex other)
    {
        if (value == null)
            return -1;
        if (other.value == null)
            return 1;
        return (int)(value.getRaw() - other.value.getRaw());
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null)
            return false;

        if (!Sex.class.isAssignableFrom(other.getClass()))
            return false;

        final Sex enm = (Sex)other;

        if ((value == null) || (enm.value == null))
            return false;
        if (value.getRaw() != enm.value.getRaw())
            return false;
        return true;
    }

    @Override
    public int hashCode()
    {
        int hash = 17;
        hash = hash * 31 + ((value != null) ? value.hashCode() : 0);
        return hash;
    }

    @Override
    public String toString() { return (value != null) ? value.toString() : "<unknown>"; }
}

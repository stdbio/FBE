package tests

import (
	"fbeproj/proto/fbe"
	"fbeproj/proto/osa"
	"fbeproj/proto/pkg"
	"fbeproj/proto/proto"
	"fbeproj/proto/simple"
	"fbeproj/proto/test"
	"fbeproj/proto/variants_ptr"
	"testing"
	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestSerializationDomain(t *testing.T) {
	// Create a new account with some orders
	var account = proto.NewAccountFromFieldValues(1, "Test", proto.State_good, proto.Balance{Currency: "USD", Amount: 1000.0}, &proto.Balance{Currency: "EUR", Amount: 100.0}, make([]proto.Order, 0))
	account.Orders = append(account.Orders, proto.Order{Id: 1, Symbol: "EURUSD", Side: proto.OrderSide_buy, Type: proto.OrderType_market, Price: 1.23456, Volume: 1000.0})
	account.Orders = append(account.Orders, proto.Order{Id: 2, Symbol: "EURUSD", Side: proto.OrderSide_sell, Type: proto.OrderType_limit, Price: 1.0, Volume: 100.0})
	account.Orders = append(account.Orders, proto.Order{Id: 3, Symbol: "EURUSD", Side: proto.OrderSide_buy, Type: proto.OrderType_stop, Price: 1.5, Volume: 10.0})

	// Serialize the struct to the FBE stream
	writer := proto.NewAccountModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(account)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 252)

	// Deserialize the struct from the FBE stream
	reader := proto.NewAccountModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	account2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, account2.Id, 1)
	assert.EqualValues(t, account2.Name, "Test")
	assert.True(t, account2.State.HasFlags(proto.State_good))
	assert.EqualValues(t, account2.Wallet.Currency, "USD")
	assert.EqualValues(t, account2.Wallet.Amount, 1000.0)
	assert.NotNil(t, account2.Asset)
	assert.EqualValues(t, account2.Asset.Currency, "EUR")
	assert.EqualValues(t, account2.Asset.Amount, 100.0)
	assert.EqualValues(t, len(account2.Orders), 3)
	assert.EqualValues(t, account2.Orders[0].Id, 1)
	assert.EqualValues(t, account2.Orders[0].Symbol, "EURUSD")
	assert.EqualValues(t, account2.Orders[0].Side, proto.OrderSide_buy)
	assert.EqualValues(t, account2.Orders[0].Type, proto.OrderType_market)
	assert.EqualValues(t, account2.Orders[0].Price, 1.23456)
	assert.EqualValues(t, account2.Orders[0].Volume, 1000.0)
	assert.EqualValues(t, account2.Orders[1].Id, 2)
	assert.EqualValues(t, account2.Orders[1].Symbol, "EURUSD")
	assert.EqualValues(t, account2.Orders[1].Side, proto.OrderSide_sell)
	assert.EqualValues(t, account2.Orders[1].Type, proto.OrderType_limit)
	assert.EqualValues(t, account2.Orders[1].Price, 1.0)
	assert.EqualValues(t, account2.Orders[1].Volume, 100.0)
	assert.EqualValues(t, account2.Orders[2].Id, 3)
	assert.EqualValues(t, account2.Orders[2].Symbol, "EURUSD")
	assert.EqualValues(t, account2.Orders[2].Side, proto.OrderSide_buy)
	assert.EqualValues(t, account2.Orders[2].Type, proto.OrderType_stop)
	assert.EqualValues(t, account2.Orders[2].Price, 1.5)
	assert.EqualValues(t, account2.Orders[2].Volume, 10.0)
}

func TestSerializationStructSimple(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructSimple()

	// Serialize the struct to the FBE stream
	writer := test.NewStructSimpleModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 110)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 392)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructSimpleModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 110)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, struct2.F1, false)
	assert.EqualValues(t, struct2.F2, true)
	assert.EqualValues(t, struct2.F3, 0)
	assert.EqualValues(t, struct2.F4, 0xFF)
	assert.EqualValues(t, struct2.F5, '\000')
	assert.EqualValues(t, struct2.F6, '!')
	assert.EqualValues(t, struct2.F7, rune(0))
	assert.EqualValues(t, struct2.F8, rune(0x0444))
	assert.EqualValues(t, struct2.F9, 0)
	assert.EqualValues(t, struct2.F10, 127)
	assert.EqualValues(t, struct2.F11, 0)
	assert.EqualValues(t, struct2.F12, 0xFF)
	assert.EqualValues(t, struct2.F13, 0)
	assert.EqualValues(t, struct2.F14, 32767)
	assert.EqualValues(t, struct2.F15, 0)
	assert.EqualValues(t, struct2.F16, 0xFFFF)
	assert.EqualValues(t, struct2.F17, 0)
	assert.EqualValues(t, struct2.F18, 2147483647)
	assert.EqualValues(t, struct2.F19, 0)
	assert.EqualValues(t, struct2.F20, 0xFFFFFFFF)
	assert.EqualValues(t, struct2.F21, 0)
	assert.EqualValues(t, struct2.F22, int64(9223372036854775807))
	assert.EqualValues(t, struct2.F23, 0)
	assert.EqualValues(t, struct2.F24, uint64(0xFFFFFFFFFFFFFFFF))
	assert.EqualValues(t, struct2.F25, 0.0)
	assert.InEpsilon(t, struct2.F26, 123.456, 0.0001)
	assert.EqualValues(t, struct2.F27, 0.0)
	assert.InEpsilon(t, struct2.F28, -123.567e+123, 1e+123)
	assert.EqualValues(t, struct2.F29, fbe.DecimalFromString("0"))
	assert.EqualValues(t, struct2.F30, fbe.DecimalFromString("123456.123456"))
	assert.EqualValues(t, struct2.F31, "")
	assert.EqualValues(t, struct2.F32, "Initial string!")
	assert.EqualValues(t, struct2.F33, fbe.TimestampEpoch())
	assert.EqualValues(t, struct2.F34, fbe.TimestampEpoch())
	assert.True(t, struct2.F35.UnixNano() > fbe.TimestampFromDate(2018, 1, 1).UnixNano())
	assert.EqualValues(t, struct2.F36, fbe.UUIDNil())
	assert.NotEqual(t, struct2.F37, fbe.UUIDNil())
	assert.EqualValues(t, struct2.F38, fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000"))

	assert.EqualValues(t, struct2.F1, struct1.F1)
	assert.EqualValues(t, struct2.F2, struct1.F2)
	assert.EqualValues(t, struct2.F3, struct1.F3)
	assert.EqualValues(t, struct2.F4, struct1.F4)
	assert.EqualValues(t, struct2.F5, struct1.F5)
	assert.EqualValues(t, struct2.F6, struct1.F6)
	assert.EqualValues(t, struct2.F7, struct1.F7)
	assert.EqualValues(t, struct2.F8, struct1.F8)
	assert.EqualValues(t, struct2.F9, struct1.F9)
	assert.EqualValues(t, struct2.F10, struct1.F10)
	assert.EqualValues(t, struct2.F11, struct1.F11)
	assert.EqualValues(t, struct2.F12, struct1.F12)
	assert.EqualValues(t, struct2.F13, struct1.F13)
	assert.EqualValues(t, struct2.F14, struct1.F14)
	assert.EqualValues(t, struct2.F15, struct1.F15)
	assert.EqualValues(t, struct2.F16, struct1.F16)
	assert.EqualValues(t, struct2.F17, struct1.F17)
	assert.EqualValues(t, struct2.F18, struct1.F18)
	assert.EqualValues(t, struct2.F19, struct1.F19)
	assert.EqualValues(t, struct2.F20, struct1.F20)
	assert.EqualValues(t, struct2.F21, struct1.F21)
	assert.EqualValues(t, struct2.F22, struct1.F22)
	assert.EqualValues(t, struct2.F23, struct1.F23)
	assert.EqualValues(t, struct2.F24, struct1.F24)
	assert.EqualValues(t, struct2.F25, struct1.F25)
	assert.InEpsilon(t, struct2.F26, struct1.F26, 0.0001)
	assert.EqualValues(t, struct2.F27, struct1.F27)
	assert.InEpsilon(t, struct2.F28, struct1.F28, 1e+123)
	assert.EqualValues(t, struct2.F29, struct1.F29)
	assert.EqualValues(t, struct2.F30, struct1.F30)
	assert.EqualValues(t, struct2.F31, struct1.F31)
	assert.EqualValues(t, struct2.F32, struct1.F32)
	assert.EqualValues(t, struct2.F33, struct1.F33)
	assert.EqualValues(t, struct2.F34, struct1.F34)
	assert.EqualValues(t, struct2.F35, struct1.F35)
	assert.EqualValues(t, struct2.F36, struct1.F36)
	assert.EqualValues(t, struct2.F37, struct1.F37)
	assert.EqualValues(t, struct2.F38, struct1.F38)
	assert.EqualValues(t, struct2.F39, struct1.F39)
	assert.EqualValues(t, struct2.F40, struct1.F40)
}

func TestSerializationStructOptional(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructOptional()

	// Serialize the struct to the FBE stream
	writer := test.NewStructOptionalModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 111)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 834)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructOptionalModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 111)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, struct2.F1, false)
	assert.EqualValues(t, struct2.F2, true)
	assert.EqualValues(t, struct2.F3, 0)
	assert.EqualValues(t, struct2.F4, 0xFF)
	assert.EqualValues(t, struct2.F5, '\000')
	assert.EqualValues(t, struct2.F6, '!')
	assert.EqualValues(t, struct2.F7, rune(0))
	assert.EqualValues(t, struct2.F8, rune(0x0444))
	assert.EqualValues(t, struct2.F9, 0)
	assert.EqualValues(t, struct2.F10, 127)
	assert.EqualValues(t, struct2.F11, 0)
	assert.EqualValues(t, struct2.F12, 0xFF)
	assert.EqualValues(t, struct2.F13, 0)
	assert.EqualValues(t, struct2.F14, 32767)
	assert.EqualValues(t, struct2.F15, 0)
	assert.EqualValues(t, struct2.F16, 0xFFFF)
	assert.EqualValues(t, struct2.F17, 0)
	assert.EqualValues(t, struct2.F18, 2147483647)
	assert.EqualValues(t, struct2.F19, 0)
	assert.EqualValues(t, struct2.F20, 0xFFFFFFFF)
	assert.EqualValues(t, struct2.F21, 0)
	assert.EqualValues(t, struct2.F22, int64(9223372036854775807))
	assert.EqualValues(t, struct2.F23, 0)
	assert.EqualValues(t, struct2.F24, uint64(0xFFFFFFFFFFFFFFFF))
	assert.EqualValues(t, struct2.F25, 0.0)
	assert.InEpsilon(t, struct2.F26, 123.456, 0.0001)
	assert.EqualValues(t, struct2.F27, 0.0)
	assert.InEpsilon(t, struct2.F28, -123.567e+123, 1e+123)
	assert.EqualValues(t, struct2.F29, fbe.DecimalFromString("0"))
	assert.EqualValues(t, struct2.F30, fbe.DecimalFromString("123456.123456"))
	assert.EqualValues(t, struct2.F31, "")
	assert.EqualValues(t, struct2.F32, "Initial string!")
	assert.EqualValues(t, struct2.F33, fbe.TimestampEpoch())
	assert.EqualValues(t, struct2.F34, fbe.TimestampEpoch())
	assert.True(t, struct2.F35.UnixNano() > fbe.TimestampFromDate(2018, 1, 1).UnixNano())
	assert.EqualValues(t, struct2.F36, fbe.UUIDNil())
	assert.NotEqual(t, struct2.F37, fbe.UUIDNil())
	assert.EqualValues(t, struct2.F38, fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000"))

	assert.Nil(t, struct2.F100)
	assert.NotNil(t, struct2.F101)
	assert.EqualValues(t, *struct2.F101, true)
	assert.Nil(t, struct2.F102)
	assert.Nil(t, struct2.F103)
	assert.NotNil(t, struct2.F104)
	assert.EqualValues(t, *struct2.F104, 0xFF)
	assert.Nil(t, struct2.F105)
	assert.Nil(t, struct2.F106)
	assert.NotNil(t, struct2.F107)
	assert.EqualValues(t, *struct2.F107, '!')
	assert.Nil(t, struct2.F108)
	assert.Nil(t, struct2.F109)
	assert.NotNil(t, struct2.F110)
	assert.EqualValues(t, *struct2.F110, rune(0x0444))
	assert.Nil(t, struct2.F111)
	assert.Nil(t, struct2.F112)
	assert.NotNil(t, struct2.F113)
	assert.EqualValues(t, *struct2.F113, 127)
	assert.Nil(t, struct2.F114)
	assert.Nil(t, struct2.F115)
	assert.NotNil(t, struct2.F116)
	assert.EqualValues(t, *struct2.F116, 0xFF)
	assert.Nil(t, struct2.F117)
	assert.Nil(t, struct2.F118)
	assert.NotNil(t, struct2.F119)
	assert.EqualValues(t, *struct2.F119, 32767)
	assert.Nil(t, struct2.F120)
	assert.Nil(t, struct2.F121)
	assert.NotNil(t, struct2.F122)
	assert.EqualValues(t, *struct2.F122, 0xFFFF)
	assert.Nil(t, struct2.F123)
	assert.Nil(t, struct2.F124)
	assert.NotNil(t, struct2.F125)
	assert.EqualValues(t, *struct2.F125, 2147483647)
	assert.Nil(t, struct2.F126)
	assert.Nil(t, struct2.F127)
	assert.NotNil(t, struct2.F128)
	assert.EqualValues(t, *struct2.F128, 0xFFFFFFFF)
	assert.Nil(t, struct2.F129)
	assert.Nil(t, struct2.F130)
	assert.NotNil(t, struct2.F131)
	assert.EqualValues(t, *struct2.F131, 9223372036854775807)
	assert.Nil(t, struct2.F132)
	assert.Nil(t, struct2.F133)
	assert.NotNil(t, struct2.F131)
	assert.EqualValues(t, *struct2.F134, uint64(0xFFFFFFFFFFFFFFFF))
	assert.Nil(t, struct2.F135)
	assert.Nil(t, struct2.F136)
	assert.NotNil(t, struct2.F137)
	assert.InEpsilon(t, *struct2.F137, 123.456, 0.0001)
	assert.Nil(t, struct2.F138)
	assert.Nil(t, struct2.F139)
	assert.NotNil(t, struct2.F140)
	assert.InEpsilon(t, *struct2.F140, -123.567e+123, 1e+123)
	assert.Nil(t, struct2.F141)
	assert.Nil(t, struct2.F142)
	assert.NotNil(t, struct2.F143)
	assert.EqualValues(t, *struct2.F143, fbe.DecimalFromString("123456.123456"))
	assert.Nil(t, struct2.F144)
	assert.Nil(t, struct2.F145)
	assert.NotNil(t, struct2.F146)
	assert.EqualValues(t, *struct2.F146, "Initial string!")
	assert.Nil(t, struct2.F147)
	assert.Nil(t, struct2.F148)
	assert.NotNil(t, struct2.F149)
	assert.True(t, struct2.F149.UnixNano() > fbe.TimestampFromDate(2018, 1, 1).UnixNano())
	assert.Nil(t, struct2.F150)
	assert.Nil(t, struct2.F151)
	assert.NotNil(t, struct2.F152)
	assert.EqualValues(t, *struct2.F152, fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000"))
	assert.Nil(t, struct2.F153)
	assert.Nil(t, struct2.F154)
	assert.Nil(t, struct2.F155)
	assert.Nil(t, struct2.F156)
	assert.Nil(t, struct2.F157)
	assert.Nil(t, struct2.F158)
	assert.Nil(t, struct2.F159)
	assert.Nil(t, struct2.F160)
	assert.Nil(t, struct2.F161)
	assert.Nil(t, struct2.F162)
	assert.Nil(t, struct2.F163)
	assert.Nil(t, struct2.F164)
	assert.Nil(t, struct2.F165)

	assert.EqualValues(t, struct2.F1, struct1.F1)
	assert.EqualValues(t, struct2.F2, struct1.F2)
	assert.EqualValues(t, struct2.F3, struct1.F3)
	assert.EqualValues(t, struct2.F4, struct1.F4)
	assert.EqualValues(t, struct2.F5, struct1.F5)
	assert.EqualValues(t, struct2.F6, struct1.F6)
	assert.EqualValues(t, struct2.F7, struct1.F7)
	assert.EqualValues(t, struct2.F8, struct1.F8)
	assert.EqualValues(t, struct2.F9, struct1.F9)
	assert.EqualValues(t, struct2.F10, struct1.F10)
	assert.EqualValues(t, struct2.F11, struct1.F11)
	assert.EqualValues(t, struct2.F12, struct1.F12)
	assert.EqualValues(t, struct2.F13, struct1.F13)
	assert.EqualValues(t, struct2.F14, struct1.F14)
	assert.EqualValues(t, struct2.F15, struct1.F15)
	assert.EqualValues(t, struct2.F16, struct1.F16)
	assert.EqualValues(t, struct2.F17, struct1.F17)
	assert.EqualValues(t, struct2.F18, struct1.F18)
	assert.EqualValues(t, struct2.F19, struct1.F19)
	assert.EqualValues(t, struct2.F20, struct1.F20)
	assert.EqualValues(t, struct2.F21, struct1.F21)
	assert.EqualValues(t, struct2.F22, struct1.F22)
	assert.EqualValues(t, struct2.F23, struct1.F23)
	assert.EqualValues(t, struct2.F24, struct1.F24)
	assert.EqualValues(t, struct2.F25, struct1.F25)
	assert.InEpsilon(t, struct2.F26, struct1.F26, 0.0001)
	assert.EqualValues(t, struct2.F27, struct1.F27)
	assert.InEpsilon(t, struct2.F28, struct1.F28, 1e+123)
	assert.EqualValues(t, struct2.F29, struct1.F29)
	assert.EqualValues(t, struct2.F30, struct1.F30)
	assert.EqualValues(t, struct2.F31, struct1.F31)
	assert.EqualValues(t, struct2.F32, struct1.F32)
	assert.EqualValues(t, struct2.F33, struct1.F33)
	assert.EqualValues(t, struct2.F34, struct1.F34)
	assert.EqualValues(t, struct2.F35, struct1.F35)
	assert.EqualValues(t, struct2.F36, struct1.F36)
	assert.EqualValues(t, struct2.F37, struct1.F37)
	assert.EqualValues(t, struct2.F38, struct1.F38)
	assert.EqualValues(t, struct2.F39, struct1.F39)
	assert.EqualValues(t, struct2.F40, struct1.F40)

	assert.EqualValues(t, struct2.F100, struct1.F100)
	assert.EqualValues(t, struct2.F101, struct1.F101)
	assert.EqualValues(t, struct2.F102, struct1.F102)
	assert.EqualValues(t, struct2.F103, struct1.F103)
	assert.EqualValues(t, struct2.F104, struct1.F104)
	assert.EqualValues(t, struct2.F105, struct1.F105)
	assert.EqualValues(t, struct2.F106, struct1.F106)
	assert.EqualValues(t, struct2.F107, struct1.F107)
	assert.EqualValues(t, struct2.F108, struct1.F108)
	assert.EqualValues(t, struct2.F109, struct1.F109)
	assert.EqualValues(t, struct2.F110, struct1.F110)
	assert.EqualValues(t, struct2.F111, struct1.F111)
	assert.EqualValues(t, struct2.F112, struct1.F112)
	assert.EqualValues(t, struct2.F113, struct1.F113)
	assert.EqualValues(t, struct2.F114, struct1.F114)
	assert.EqualValues(t, struct2.F115, struct1.F115)
	assert.EqualValues(t, struct2.F116, struct1.F116)
	assert.EqualValues(t, struct2.F117, struct1.F117)
	assert.EqualValues(t, struct2.F118, struct1.F118)
	assert.EqualValues(t, struct2.F119, struct1.F119)
	assert.EqualValues(t, struct2.F120, struct1.F120)
	assert.EqualValues(t, struct2.F121, struct1.F121)
	assert.EqualValues(t, struct2.F122, struct1.F122)
	assert.EqualValues(t, struct2.F123, struct1.F123)
	assert.EqualValues(t, struct2.F124, struct1.F124)
	assert.EqualValues(t, struct2.F125, struct1.F125)
	assert.EqualValues(t, struct2.F126, struct1.F126)
	assert.EqualValues(t, struct2.F127, struct1.F127)
	assert.EqualValues(t, struct2.F128, struct1.F128)
	assert.EqualValues(t, struct2.F129, struct1.F129)
	assert.EqualValues(t, struct2.F130, struct1.F130)
	assert.EqualValues(t, struct2.F131, struct1.F131)
	assert.EqualValues(t, struct2.F132, struct1.F132)
	assert.EqualValues(t, struct2.F133, struct1.F133)
	assert.EqualValues(t, struct2.F134, struct1.F134)
	assert.EqualValues(t, struct2.F135, struct1.F135)
	assert.EqualValues(t, struct2.F136, struct1.F136)
	assert.InEpsilon(t, *struct2.F137, *struct1.F137, 0.0001)
	assert.EqualValues(t, struct2.F138, struct1.F138)
	assert.EqualValues(t, struct2.F139, struct1.F139)
	assert.InEpsilon(t, *struct2.F140, *struct1.F140, 1e+123)
	assert.EqualValues(t, struct2.F141, struct1.F141)
	assert.EqualValues(t, struct2.F142, struct1.F142)
	assert.EqualValues(t, struct2.F143, struct1.F143)
	assert.EqualValues(t, struct2.F144, struct1.F144)
	assert.EqualValues(t, struct2.F145, struct1.F145)
	assert.EqualValues(t, struct2.F146, struct1.F146)
	assert.EqualValues(t, struct2.F147, struct1.F147)
	assert.EqualValues(t, struct2.F148, struct1.F148)
	assert.EqualValues(t, struct2.F149, struct1.F149)
	assert.EqualValues(t, struct2.F150, struct1.F150)
	assert.EqualValues(t, struct2.F151, struct1.F151)
	assert.EqualValues(t, struct2.F152, struct1.F152)
	assert.EqualValues(t, struct2.F153, struct1.F153)
	assert.EqualValues(t, struct2.F154, struct1.F154)
	assert.EqualValues(t, struct2.F155, struct1.F155)
	assert.EqualValues(t, struct2.F156, struct1.F156)
	assert.EqualValues(t, struct2.F157, struct1.F157)
}

func TestSerializationStructNested(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructNested()

	// Serialize the struct to the FBE stream
	writer := test.NewStructNestedModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 112)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 2099)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructNestedModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 112)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, struct2.F1, false)
	assert.EqualValues(t, struct2.F2, true)
	assert.EqualValues(t, struct2.F3, 0)
	assert.EqualValues(t, struct2.F4, 0xFF)
	assert.EqualValues(t, struct2.F5, '\000')
	assert.EqualValues(t, struct2.F6, '!')
	assert.EqualValues(t, struct2.F7, rune(0))
	assert.EqualValues(t, struct2.F8, rune(0x0444))
	assert.EqualValues(t, struct2.F9, 0)
	assert.EqualValues(t, struct2.F10, 127)
	assert.EqualValues(t, struct2.F11, 0)
	assert.EqualValues(t, struct2.F12, 0xFF)
	assert.EqualValues(t, struct2.F13, 0)
	assert.EqualValues(t, struct2.F14, 32767)
	assert.EqualValues(t, struct2.F15, 0)
	assert.EqualValues(t, struct2.F16, 0xFFFF)
	assert.EqualValues(t, struct2.F17, 0)
	assert.EqualValues(t, struct2.F18, 2147483647)
	assert.EqualValues(t, struct2.F19, 0)
	assert.EqualValues(t, struct2.F20, 0xFFFFFFFF)
	assert.EqualValues(t, struct2.F21, 0)
	assert.EqualValues(t, struct2.F22, int64(9223372036854775807))
	assert.EqualValues(t, struct2.F23, 0)
	assert.EqualValues(t, struct2.F24, uint64(0xFFFFFFFFFFFFFFFF))
	assert.EqualValues(t, struct2.F25, 0.0)
	assert.InEpsilon(t, struct2.F26, 123.456, 0.0001)
	assert.EqualValues(t, struct2.F27, 0.0)
	assert.InEpsilon(t, struct2.F28, -123.567e+123, 1e+123)
	assert.EqualValues(t, struct2.F29, fbe.DecimalFromString("0"))
	assert.EqualValues(t, struct2.F30, fbe.DecimalFromString("123456.123456"))
	assert.EqualValues(t, struct2.F31, "")
	assert.EqualValues(t, struct2.F32, "Initial string!")
	assert.EqualValues(t, struct2.F33, fbe.TimestampEpoch())
	assert.EqualValues(t, struct2.F34, fbe.TimestampEpoch())
	assert.True(t, struct2.F35.UnixNano() > fbe.TimestampFromDate(2018, 1, 1).UnixNano())
	assert.EqualValues(t, struct2.F36, fbe.UUIDNil())
	assert.NotEqual(t, struct2.F37, fbe.UUIDNil())
	assert.EqualValues(t, struct2.F38, fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000"))

	assert.Nil(t, struct2.F100)
	assert.NotNil(t, struct2.F101)
	assert.EqualValues(t, *struct2.F101, true)
	assert.Nil(t, struct2.F102)
	assert.Nil(t, struct2.F103)
	assert.NotNil(t, struct2.F104)
	assert.EqualValues(t, *struct2.F104, 0xFF)
	assert.Nil(t, struct2.F105)
	assert.Nil(t, struct2.F106)
	assert.NotNil(t, struct2.F107)
	assert.EqualValues(t, *struct2.F107, '!')
	assert.Nil(t, struct2.F108)
	assert.Nil(t, struct2.F109)
	assert.NotNil(t, struct2.F110)
	assert.EqualValues(t, *struct2.F110, rune(0x0444))
	assert.Nil(t, struct2.F111)
	assert.Nil(t, struct2.F112)
	assert.NotNil(t, struct2.F113)
	assert.EqualValues(t, *struct2.F113, 127)
	assert.Nil(t, struct2.F114)
	assert.Nil(t, struct2.F115)
	assert.NotNil(t, struct2.F116)
	assert.EqualValues(t, *struct2.F116, 0xFF)
	assert.Nil(t, struct2.F117)
	assert.Nil(t, struct2.F118)
	assert.NotNil(t, struct2.F119)
	assert.EqualValues(t, *struct2.F119, 32767)
	assert.Nil(t, struct2.F120)
	assert.Nil(t, struct2.F121)
	assert.NotNil(t, struct2.F122)
	assert.EqualValues(t, *struct2.F122, 0xFFFF)
	assert.Nil(t, struct2.F123)
	assert.Nil(t, struct2.F124)
	assert.NotNil(t, struct2.F125)
	assert.EqualValues(t, *struct2.F125, 2147483647)
	assert.Nil(t, struct2.F126)
	assert.Nil(t, struct2.F127)
	assert.NotNil(t, struct2.F128)
	assert.EqualValues(t, *struct2.F128, 0xFFFFFFFF)
	assert.Nil(t, struct2.F129)
	assert.Nil(t, struct2.F130)
	assert.NotNil(t, struct2.F131)
	assert.EqualValues(t, *struct2.F131, 9223372036854775807)
	assert.Nil(t, struct2.F132)
	assert.Nil(t, struct2.F133)
	assert.NotNil(t, struct2.F131)
	assert.EqualValues(t, *struct2.F134, uint64(0xFFFFFFFFFFFFFFFF))
	assert.Nil(t, struct2.F135)
	assert.Nil(t, struct2.F136)
	assert.NotNil(t, struct2.F137)
	assert.InEpsilon(t, *struct2.F137, 123.456, 0.0001)
	assert.Nil(t, struct2.F138)
	assert.Nil(t, struct2.F139)
	assert.NotNil(t, struct2.F140)
	assert.InEpsilon(t, *struct2.F140, -123.567e+123, 1e+123)
	assert.Nil(t, struct2.F141)
	assert.Nil(t, struct2.F142)
	assert.NotNil(t, struct2.F143)
	assert.EqualValues(t, *struct2.F143, fbe.DecimalFromString("123456.123456"))
	assert.Nil(t, struct2.F144)
	assert.Nil(t, struct2.F145)
	assert.NotNil(t, struct2.F146)
	assert.EqualValues(t, *struct2.F146, "Initial string!")
	assert.Nil(t, struct2.F147)
	assert.Nil(t, struct2.F148)
	assert.NotNil(t, struct2.F149)
	assert.True(t, struct2.F149.UnixNano() > fbe.TimestampFromDate(2018, 1, 1).UnixNano())
	assert.Nil(t, struct2.F150)
	assert.Nil(t, struct2.F151)
	assert.NotNil(t, struct2.F152)
	assert.EqualValues(t, *struct2.F152, fbe.UUIDFromString("123e4567-e89b-12d3-a456-426655440000"))
	assert.Nil(t, struct2.F153)
	assert.Nil(t, struct2.F154)
	assert.Nil(t, struct2.F155)
	assert.Nil(t, struct2.F156)
	assert.Nil(t, struct2.F157)
	assert.Nil(t, struct2.F158)
	assert.Nil(t, struct2.F159)
	assert.Nil(t, struct2.F160)
	assert.Nil(t, struct2.F161)
	assert.Nil(t, struct2.F162)
	assert.Nil(t, struct2.F163)
	assert.Nil(t, struct2.F164)
	assert.Nil(t, struct2.F165)

	assert.EqualValues(t, struct2.F1000, test.EnumSimple_ENUM_VALUE_0)
	assert.Nil(t, struct2.F1001)
	assert.EqualValues(t, struct2.F1002, test.EnumTyped_ENUM_VALUE_2)
	assert.Nil(t, struct2.F1003)
	assert.EqualValues(t, struct2.F1004, test.FlagsSimple_FLAG_VALUE_0)
	assert.Nil(t, struct2.F1005)
	assert.EqualValues(t, struct2.F1006, test.FlagsTyped_FLAG_VALUE_2|test.FlagsTyped_FLAG_VALUE_4|test.FlagsTyped_FLAG_VALUE_6)
	assert.Nil(t, struct2.F1007)
	assert.Nil(t, struct2.F1009)
	assert.Nil(t, struct2.F1011)

	assert.EqualValues(t, struct2.F1, struct1.F1)
	assert.EqualValues(t, struct2.F2, struct1.F2)
	assert.EqualValues(t, struct2.F3, struct1.F3)
	assert.EqualValues(t, struct2.F4, struct1.F4)
	assert.EqualValues(t, struct2.F5, struct1.F5)
	assert.EqualValues(t, struct2.F6, struct1.F6)
	assert.EqualValues(t, struct2.F7, struct1.F7)
	assert.EqualValues(t, struct2.F8, struct1.F8)
	assert.EqualValues(t, struct2.F9, struct1.F9)
	assert.EqualValues(t, struct2.F10, struct1.F10)
	assert.EqualValues(t, struct2.F11, struct1.F11)
	assert.EqualValues(t, struct2.F12, struct1.F12)
	assert.EqualValues(t, struct2.F13, struct1.F13)
	assert.EqualValues(t, struct2.F14, struct1.F14)
	assert.EqualValues(t, struct2.F15, struct1.F15)
	assert.EqualValues(t, struct2.F16, struct1.F16)
	assert.EqualValues(t, struct2.F17, struct1.F17)
	assert.EqualValues(t, struct2.F18, struct1.F18)
	assert.EqualValues(t, struct2.F19, struct1.F19)
	assert.EqualValues(t, struct2.F20, struct1.F20)
	assert.EqualValues(t, struct2.F21, struct1.F21)
	assert.EqualValues(t, struct2.F22, struct1.F22)
	assert.EqualValues(t, struct2.F23, struct1.F23)
	assert.EqualValues(t, struct2.F24, struct1.F24)
	assert.EqualValues(t, struct2.F25, struct1.F25)
	assert.InEpsilon(t, struct2.F26, struct1.F26, 0.0001)
	assert.EqualValues(t, struct2.F27, struct1.F27)
	assert.InEpsilon(t, struct2.F28, struct1.F28, 1e+123)
	assert.EqualValues(t, struct2.F29, struct1.F29)
	assert.EqualValues(t, struct2.F30, struct1.F30)
	assert.EqualValues(t, struct2.F31, struct1.F31)
	assert.EqualValues(t, struct2.F32, struct1.F32)
	assert.EqualValues(t, struct2.F33, struct1.F33)
	assert.EqualValues(t, struct2.F34, struct1.F34)
	assert.EqualValues(t, struct2.F35, struct1.F35)
	assert.EqualValues(t, struct2.F36, struct1.F36)
	assert.EqualValues(t, struct2.F37, struct1.F37)
	assert.EqualValues(t, struct2.F38, struct1.F38)
	assert.EqualValues(t, struct2.F39, struct1.F39)
	assert.EqualValues(t, struct2.F40, struct1.F40)

	assert.EqualValues(t, struct2.F100, struct1.F100)
	assert.EqualValues(t, struct2.F101, struct1.F101)
	assert.EqualValues(t, struct2.F102, struct1.F102)
	assert.EqualValues(t, struct2.F103, struct1.F103)
	assert.EqualValues(t, struct2.F104, struct1.F104)
	assert.EqualValues(t, struct2.F105, struct1.F105)
	assert.EqualValues(t, struct2.F106, struct1.F106)
	assert.EqualValues(t, struct2.F107, struct1.F107)
	assert.EqualValues(t, struct2.F108, struct1.F108)
	assert.EqualValues(t, struct2.F109, struct1.F109)
	assert.EqualValues(t, struct2.F110, struct1.F110)
	assert.EqualValues(t, struct2.F111, struct1.F111)
	assert.EqualValues(t, struct2.F112, struct1.F112)
	assert.EqualValues(t, struct2.F113, struct1.F113)
	assert.EqualValues(t, struct2.F114, struct1.F114)
	assert.EqualValues(t, struct2.F115, struct1.F115)
	assert.EqualValues(t, struct2.F116, struct1.F116)
	assert.EqualValues(t, struct2.F117, struct1.F117)
	assert.EqualValues(t, struct2.F118, struct1.F118)
	assert.EqualValues(t, struct2.F119, struct1.F119)
	assert.EqualValues(t, struct2.F120, struct1.F120)
	assert.EqualValues(t, struct2.F121, struct1.F121)
	assert.EqualValues(t, struct2.F122, struct1.F122)
	assert.EqualValues(t, struct2.F123, struct1.F123)
	assert.EqualValues(t, struct2.F124, struct1.F124)
	assert.EqualValues(t, struct2.F125, struct1.F125)
	assert.EqualValues(t, struct2.F126, struct1.F126)
	assert.EqualValues(t, struct2.F127, struct1.F127)
	assert.EqualValues(t, struct2.F128, struct1.F128)
	assert.EqualValues(t, struct2.F129, struct1.F129)
	assert.EqualValues(t, struct2.F130, struct1.F130)
	assert.EqualValues(t, struct2.F131, struct1.F131)
	assert.EqualValues(t, struct2.F132, struct1.F132)
	assert.EqualValues(t, struct2.F133, struct1.F133)
	assert.EqualValues(t, struct2.F134, struct1.F134)
	assert.EqualValues(t, struct2.F135, struct1.F135)
	assert.EqualValues(t, struct2.F136, struct1.F136)
	assert.InEpsilon(t, *struct2.F137, *struct1.F137, 0.0001)
	assert.EqualValues(t, struct2.F138, struct1.F138)
	assert.EqualValues(t, struct2.F139, struct1.F139)
	assert.InEpsilon(t, *struct2.F140, *struct1.F140, 1e+123)
	assert.EqualValues(t, struct2.F141, struct1.F141)
	assert.EqualValues(t, struct2.F142, struct1.F142)
	assert.EqualValues(t, struct2.F143, struct1.F143)
	assert.EqualValues(t, struct2.F144, struct1.F144)
	assert.EqualValues(t, struct2.F145, struct1.F145)
	assert.EqualValues(t, struct2.F146, struct1.F146)
	assert.EqualValues(t, struct2.F147, struct1.F147)
	assert.EqualValues(t, struct2.F148, struct1.F148)
	assert.EqualValues(t, struct2.F149, struct1.F149)
	assert.EqualValues(t, struct2.F150, struct1.F150)
	assert.EqualValues(t, struct2.F151, struct1.F151)
	assert.EqualValues(t, struct2.F152, struct1.F152)
	assert.EqualValues(t, struct2.F153, struct1.F153)
	assert.EqualValues(t, struct2.F154, struct1.F154)
	assert.EqualValues(t, struct2.F155, struct1.F155)
	assert.EqualValues(t, struct2.F156, struct1.F156)
	assert.EqualValues(t, struct2.F157, struct1.F157)

	assert.EqualValues(t, struct2.F1000, struct1.F1000)
	assert.EqualValues(t, struct2.F1001, struct1.F1001)
	assert.EqualValues(t, struct2.F1002, struct1.F1002)
	assert.EqualValues(t, struct2.F1003, struct1.F1003)
	assert.EqualValues(t, struct2.F1004, struct1.F1004)
	assert.EqualValues(t, struct2.F1005, struct1.F1005)
	assert.EqualValues(t, struct2.F1006, struct1.F1006)
	assert.EqualValues(t, struct2.F1007, struct1.F1007)
}

func TestSerializationStructBytes(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructBytes()
	bytesF1 := []byte("ABC")
	struct1.F1 = bytesF1
	bytesF2 := []byte("test")
	struct1.F2 = &bytesF2

	// Serialize the struct to the FBE stream
	writer := test.NewStructBytesModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 120)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 49)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructBytesModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 120)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 3)
	assert.EqualValues(t, struct2.F1[0], 65)
	assert.EqualValues(t, struct2.F1[1], 66)
	assert.EqualValues(t, struct2.F1[2], 67)
	assert.NotNil(t, struct2.F2)
	assert.EqualValues(t, len(*struct2.F2), 4)
	assert.EqualValues(t, (*struct2.F2)[0], 116)
	assert.EqualValues(t, (*struct2.F2)[1], 101)
	assert.EqualValues(t, (*struct2.F2)[2], 115)
	assert.EqualValues(t, (*struct2.F2)[3], 116)
	assert.Nil(t, struct2.F3)
}

func TestSerializationStructArray(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructArray()
	struct1.F1[0] = 48
	struct1.F1[1] = 65
	arrayF2 := byte(97)
	struct1.F2[0] = &arrayF2
	struct1.F2[1] = nil
	struct1.F3[0] = []byte("000")
	struct1.F3[1] = []byte("AAA")
	arrayF4 := []byte("aaa")
	struct1.F4[0] = &arrayF4
	struct1.F4[1] = nil
	struct1.F5[0] = test.EnumSimple_ENUM_VALUE_1
	struct1.F5[1] = test.EnumSimple_ENUM_VALUE_2
	arrayF6 := test.EnumSimple_ENUM_VALUE_1
	struct1.F6[0] = &arrayF6
	struct1.F6[1] = nil
	struct1.F7[0] = test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F7[1] = test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2 | test.FlagsSimple_FLAG_VALUE_3
	arrayF8 := test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F8[0] = &arrayF8
	struct1.F8[1] = nil
	struct1.F9[0] = *test.NewStructSimple()
	struct1.F9[1] = *test.NewStructSimple()
	struct1.F10[0] = test.NewStructSimple()
	struct1.F10[1] = nil

	// Serialize the struct to the FBE stream
	writer := test.NewStructArrayModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 125)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 1290)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructArrayModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 125)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 2)
	assert.EqualValues(t, struct2.F1[0], 48)
	assert.EqualValues(t, struct2.F1[1], 65)
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.EqualValues(t, *struct2.F2[0], 97)
	assert.Nil(t, struct2.F2[1])
	assert.EqualValues(t, len(struct2.F3), 2)
	assert.EqualValues(t, len(struct2.F3[0]), 3)
	assert.EqualValues(t, struct2.F3[0][0], 48)
	assert.EqualValues(t, struct2.F3[0][1], 48)
	assert.EqualValues(t, struct2.F3[0][2], 48)
	assert.EqualValues(t, len(struct2.F3[1]), 3)
	assert.EqualValues(t, struct2.F3[1][0], 65)
	assert.EqualValues(t, struct2.F3[1][1], 65)
	assert.EqualValues(t, struct2.F3[1][2], 65)
	assert.EqualValues(t, len(struct2.F4), 2)
	assert.NotNil(t, struct2.F4[0])
	assert.EqualValues(t, len(*struct2.F4[0]), 3)
	assert.EqualValues(t, (*struct2.F4[0])[0], 97)
	assert.EqualValues(t, (*struct2.F4[0])[1], 97)
	assert.EqualValues(t, (*struct2.F4[0])[2], 97)
	assert.Nil(t, struct2.F4[1])
	assert.EqualValues(t, len(struct2.F5), 2)
	assert.EqualValues(t, struct2.F5[0], test.EnumSimple_ENUM_VALUE_1)
	assert.EqualValues(t, struct2.F5[1], test.EnumSimple_ENUM_VALUE_2)
	assert.EqualValues(t, len(struct2.F6), 2)
	assert.EqualValues(t, *struct2.F6[0], test.EnumSimple_ENUM_VALUE_1)
	assert.Nil(t, struct2.F6[1])
	assert.EqualValues(t, len(struct2.F7), 2)
	assert.EqualValues(t, struct2.F7[0], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.EqualValues(t, struct2.F7[1], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	assert.EqualValues(t, len(struct2.F8), 2)
	assert.EqualValues(t, *struct2.F8[0], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.Nil(t, struct2.F8[1])
	assert.EqualValues(t, len(struct2.F9), 2)
	assert.EqualValues(t, struct2.F9[0].F2, true)
	assert.EqualValues(t, struct2.F9[0].F12, 0xFF)
	assert.EqualValues(t, struct2.F9[0].F32, "Initial string!")
	assert.EqualValues(t, struct2.F9[1].F2, true)
	assert.EqualValues(t, struct2.F9[1].F12, 0xFF)
	assert.EqualValues(t, struct2.F9[1].F32, "Initial string!")
	assert.EqualValues(t, len(struct2.F10), 2)
	assert.NotNil(t, struct2.F10[0])
	assert.EqualValues(t, struct2.F10[0].F2, true)
	assert.EqualValues(t, struct2.F10[0].F12, 0xFF)
	assert.EqualValues(t, struct2.F10[0].F32, "Initial string!")
	assert.Nil(t, struct2.F10[1])
}

func TestSerializationStructVector(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructVector()
	struct1.F1 = append(struct1.F1, 48)
	struct1.F1 = append(struct1.F1, 65)
	vectorF2 := byte(97)
	struct1.F2 = append(struct1.F2, &vectorF2)
	struct1.F2 = append(struct1.F2, nil)
	struct1.F3 = append(struct1.F3, []byte("000"))
	struct1.F3 = append(struct1.F3, []byte("AAA"))
	vectorF4 := []byte("aaa")
	struct1.F4 = append(struct1.F4, &vectorF4)
	struct1.F4 = append(struct1.F4, nil)
	struct1.F5 = append(struct1.F5, test.EnumSimple_ENUM_VALUE_1)
	struct1.F5 = append(struct1.F5, test.EnumSimple_ENUM_VALUE_2)
	vectorF6 := test.EnumSimple_ENUM_VALUE_1
	struct1.F6 = append(struct1.F6, &vectorF6)
	struct1.F6 = append(struct1.F6, nil)
	struct1.F7 = append(struct1.F7, test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	struct1.F7 = append(struct1.F7, test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	vectorF8 := test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F8 = append(struct1.F8, &vectorF8)
	struct1.F8 = append(struct1.F8, nil)
	struct1.F9 = append(struct1.F9, *test.NewStructSimple())
	struct1.F9 = append(struct1.F9, *test.NewStructSimple())
	struct1.F10 = append(struct1.F10, test.NewStructSimple())
	struct1.F10 = append(struct1.F10, nil)

	// Serialize the struct to the FBE stream
	writer := test.NewStructVectorModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 130)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 1370)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructVectorModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 130)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 2)
	assert.EqualValues(t, struct2.F1[0], 48)
	assert.EqualValues(t, struct2.F1[1], 65)
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.EqualValues(t, *struct2.F2[0], 97)
	assert.Nil(t, struct2.F2[1])
	assert.EqualValues(t, len(struct2.F3), 2)
	assert.EqualValues(t, len(struct2.F3[0]), 3)
	assert.EqualValues(t, struct2.F3[0][0], 48)
	assert.EqualValues(t, struct2.F3[0][1], 48)
	assert.EqualValues(t, struct2.F3[0][2], 48)
	assert.EqualValues(t, len(struct2.F3[1]), 3)
	assert.EqualValues(t, struct2.F3[1][0], 65)
	assert.EqualValues(t, struct2.F3[1][1], 65)
	assert.EqualValues(t, struct2.F3[1][2], 65)
	assert.EqualValues(t, len(struct2.F4), 2)
	assert.NotNil(t, struct2.F4[0])
	assert.EqualValues(t, len(*struct2.F4[0]), 3)
	assert.EqualValues(t, (*struct2.F4[0])[0], 97)
	assert.EqualValues(t, (*struct2.F4[0])[1], 97)
	assert.EqualValues(t, (*struct2.F4[0])[2], 97)
	assert.Nil(t, struct2.F4[1])
	assert.EqualValues(t, len(struct2.F5), 2)
	assert.EqualValues(t, struct2.F5[0], test.EnumSimple_ENUM_VALUE_1)
	assert.EqualValues(t, struct2.F5[1], test.EnumSimple_ENUM_VALUE_2)
	assert.EqualValues(t, len(struct2.F6), 2)
	assert.EqualValues(t, *struct2.F6[0], test.EnumSimple_ENUM_VALUE_1)
	assert.Nil(t, struct2.F6[1])
	assert.EqualValues(t, len(struct2.F7), 2)
	assert.EqualValues(t, struct2.F7[0], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.EqualValues(t, struct2.F7[1], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	assert.EqualValues(t, len(struct2.F8), 2)
	assert.EqualValues(t, *struct2.F8[0], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.Nil(t, struct2.F8[1])
	assert.EqualValues(t, len(struct2.F9), 2)
	assert.EqualValues(t, struct2.F9[0].F2, true)
	assert.EqualValues(t, struct2.F9[0].F12, 0xFF)
	assert.EqualValues(t, struct2.F9[0].F32, "Initial string!")
	assert.EqualValues(t, struct2.F9[1].F2, true)
	assert.EqualValues(t, struct2.F9[1].F12, 0xFF)
	assert.EqualValues(t, struct2.F9[1].F32, "Initial string!")
	assert.EqualValues(t, len(struct2.F10), 2)
	assert.NotNil(t, struct2.F10[0])
	assert.EqualValues(t, struct2.F10[0].F2, true)
	assert.EqualValues(t, struct2.F10[0].F12, 0xFF)
	assert.EqualValues(t, struct2.F10[0].F32, "Initial string!")
	assert.Nil(t, struct2.F10[1])
}

func TestSerializationStructList(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructList()
	struct1.F1 = append(struct1.F1, 48)
	struct1.F1 = append(struct1.F1, 65)
	vectorF2 := byte(97)
	struct1.F2 = append(struct1.F2, &vectorF2)
	struct1.F2 = append(struct1.F2, nil)
	struct1.F3 = append(struct1.F3, []byte("000"))
	struct1.F3 = append(struct1.F3, []byte("AAA"))
	vectorF4 := []byte("aaa")
	struct1.F4 = append(struct1.F4, &vectorF4)
	struct1.F4 = append(struct1.F4, nil)
	struct1.F5 = append(struct1.F5, test.EnumSimple_ENUM_VALUE_1)
	struct1.F5 = append(struct1.F5, test.EnumSimple_ENUM_VALUE_2)
	vectorF6 := test.EnumSimple_ENUM_VALUE_1
	struct1.F6 = append(struct1.F6, &vectorF6)
	struct1.F6 = append(struct1.F6, nil)
	struct1.F7 = append(struct1.F7, test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	struct1.F7 = append(struct1.F7, test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	vectorF8 := test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F8 = append(struct1.F8, &vectorF8)
	struct1.F8 = append(struct1.F8, nil)
	struct1.F9 = append(struct1.F9, *test.NewStructSimple())
	struct1.F9 = append(struct1.F9, *test.NewStructSimple())
	struct1.F10 = append(struct1.F10, test.NewStructSimple())
	struct1.F10 = append(struct1.F10, nil)

	// Serialize the struct to the FBE stream
	writer := test.NewStructListModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 131)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 1370)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructListModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 131)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 2)
	assert.EqualValues(t, struct2.F1[0], 48)
	assert.EqualValues(t, struct2.F1[1], 65)
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.EqualValues(t, *struct2.F2[0], 97)
	assert.Nil(t, struct2.F2[1])
	assert.EqualValues(t, len(struct2.F3), 2)
	assert.EqualValues(t, len(struct2.F3[0]), 3)
	assert.EqualValues(t, struct2.F3[0][0], 48)
	assert.EqualValues(t, struct2.F3[0][1], 48)
	assert.EqualValues(t, struct2.F3[0][2], 48)
	assert.EqualValues(t, len(struct2.F3[1]), 3)
	assert.EqualValues(t, struct2.F3[1][0], 65)
	assert.EqualValues(t, struct2.F3[1][1], 65)
	assert.EqualValues(t, struct2.F3[1][2], 65)
	assert.EqualValues(t, len(struct2.F4), 2)
	assert.NotNil(t, struct2.F4[0])
	assert.EqualValues(t, len(*struct2.F4[0]), 3)
	assert.EqualValues(t, (*struct2.F4[0])[0], 97)
	assert.EqualValues(t, (*struct2.F4[0])[1], 97)
	assert.EqualValues(t, (*struct2.F4[0])[2], 97)
	assert.Nil(t, struct2.F4[1])
	assert.EqualValues(t, len(struct2.F5), 2)
	assert.EqualValues(t, struct2.F5[0], test.EnumSimple_ENUM_VALUE_1)
	assert.EqualValues(t, struct2.F5[1], test.EnumSimple_ENUM_VALUE_2)
	assert.EqualValues(t, len(struct2.F6), 2)
	assert.EqualValues(t, *struct2.F6[0], test.EnumSimple_ENUM_VALUE_1)
	assert.Nil(t, struct2.F6[1])
	assert.EqualValues(t, len(struct2.F7), 2)
	assert.EqualValues(t, struct2.F7[0], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.EqualValues(t, struct2.F7[1], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	assert.EqualValues(t, len(struct2.F8), 2)
	assert.EqualValues(t, *struct2.F8[0], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.Nil(t, struct2.F8[1])
	assert.EqualValues(t, len(struct2.F9), 2)
	assert.EqualValues(t, struct2.F9[0].F2, true)
	assert.EqualValues(t, struct2.F9[0].F12, 0xFF)
	assert.EqualValues(t, struct2.F9[0].F32, "Initial string!")
	assert.EqualValues(t, struct2.F9[1].F2, true)
	assert.EqualValues(t, struct2.F9[1].F12, 0xFF)
	assert.EqualValues(t, struct2.F9[1].F32, "Initial string!")
	assert.EqualValues(t, len(struct2.F10), 2)
	assert.NotNil(t, struct2.F10[0])
	assert.EqualValues(t, struct2.F10[0].F2, true)
	assert.EqualValues(t, struct2.F10[0].F12, 0xFF)
	assert.EqualValues(t, struct2.F10[0].F32, "Initial string!")
	assert.Nil(t, struct2.F10[1])
}

func TestSerializationStructSet(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructSet()
	struct1.F1.Add(48)
	struct1.F1.Add(65)
	struct1.F1.Add(97)
	struct1.F2.Add(test.EnumSimple_ENUM_VALUE_1)
	struct1.F2.Add(test.EnumSimple_ENUM_VALUE_2)
	struct1.F3.Add(test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2)
	struct1.F3.Add(test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2 | test.FlagsSimple_FLAG_VALUE_3)
	s1 := test.NewStructSimple()
	s1.Id = 48
	struct1.F4.Add(*s1)
	s2 := test.NewStructSimple()
	s2.Id = 65
	struct1.F4.Add(*s2)

	// Serialize the struct to the FBE stream
	writer := test.NewStructSetModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 132)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 843)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructSetModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 132)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 3)
	assert.True(t, struct2.F1.Contains(48))
	assert.True(t, struct2.F1.Contains(65))
	assert.True(t, struct2.F1.Contains(97))
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.True(t, struct2.F2.Contains(test.EnumSimple_ENUM_VALUE_1))
	assert.True(t, struct2.F2.Contains(test.EnumSimple_ENUM_VALUE_2))
	assert.EqualValues(t, len(struct2.F3), 2)
	assert.True(t, struct2.F3.Contains(test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2))
	assert.True(t, struct2.F3.Contains(test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3))
	assert.EqualValues(t, len(struct2.F4), 2)
	s1.Id = 48
	assert.True(t, struct2.F4.Contains(*s1))
	s2.Id = 65
	assert.True(t, struct2.F4.Contains(*s2))
}

func TestSerializationStructMap(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructMap()
	struct1.F1[10] = 48
	struct1.F1[20] = 65
	mapF2 := byte(97)
	struct1.F2[10] = &mapF2
	struct1.F2[20] = nil
	struct1.F3[10] = []byte("000")
	struct1.F3[20] = []byte("AAA")
	mapF4 := []byte("aaa")
	struct1.F4[10] = &mapF4
	struct1.F4[20] = nil
	struct1.F5[10] = test.EnumSimple_ENUM_VALUE_1
	struct1.F5[20] = test.EnumSimple_ENUM_VALUE_2
	mapF6 := test.EnumSimple_ENUM_VALUE_1
	struct1.F6[10] = &mapF6
	struct1.F6[20] = nil
	struct1.F7[10] = test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F7[20] = test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2 | test.FlagsSimple_FLAG_VALUE_3
	mapF8 := test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F8[10] = &mapF8
	struct1.F8[20] = nil
	s1 := test.NewStructSimple()
	s1.Id = 48
	struct1.F9[10] = *s1
	s2 := test.NewStructSimple()
	s2.Id = 65
	struct1.F9[20] = *s2
	struct1.F10[10] = s1
	struct1.F10[20] = nil

	// Serialize the struct to the FBE stream
	writer := test.NewStructMapModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 140)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 1450)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructMapModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 140)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 2)
	assert.EqualValues(t, struct2.F1[10], 48)
	assert.EqualValues(t, struct2.F1[20], 65)
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.EqualValues(t, *struct2.F2[10], 97)
	assert.Nil(t, struct2.F2[20])
	assert.EqualValues(t, len(struct2.F3), 2)
	assert.EqualValues(t, len(struct2.F3[10]), 3)
	assert.EqualValues(t, len(struct2.F3[20]), 3)
	assert.EqualValues(t, len(struct2.F4), 2)
	assert.EqualValues(t, len(*struct2.F4[10]), 3)
	assert.Nil(t, struct2.F4[20])
	assert.EqualValues(t, len(struct2.F5), 2)
	assert.EqualValues(t, struct2.F5[10], test.EnumSimple_ENUM_VALUE_1)
	assert.EqualValues(t, struct2.F5[20], test.EnumSimple_ENUM_VALUE_2)
	assert.EqualValues(t, len(struct2.F6), 2)
	assert.EqualValues(t, *struct2.F6[10], test.EnumSimple_ENUM_VALUE_1)
	assert.Nil(t, struct2.F6[20])
	assert.EqualValues(t, len(struct2.F7), 2)
	assert.EqualValues(t, struct2.F7[10], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.EqualValues(t, struct2.F7[20], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	assert.EqualValues(t, len(struct2.F8), 2)
	assert.EqualValues(t, *struct2.F8[10], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.Nil(t, struct2.F8[20])
	assert.EqualValues(t, len(struct2.F9), 2)
	assert.EqualValues(t, struct2.F9[10].Id, 48)
	assert.EqualValues(t, struct2.F9[20].Id, 65)
	assert.EqualValues(t, len(struct2.F10), 2)
	assert.EqualValues(t, struct2.F10[10].Id, 48)
	assert.Nil(t, struct2.F10[20])
}

func TestSerializationStructHash(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructHash()
	struct1.F1["10"] = 48
	struct1.F1["20"] = 65
	hashF2 := byte(97)
	struct1.F2["10"] = &hashF2
	struct1.F2["20"] = nil
	struct1.F3["10"] = []byte("000")
	struct1.F3["20"] = []byte("AAA")
	hashF4 := []byte("aaa")
	struct1.F4["10"] = &hashF4
	struct1.F4["20"] = nil
	struct1.F5["10"] = test.EnumSimple_ENUM_VALUE_1
	struct1.F5["20"] = test.EnumSimple_ENUM_VALUE_2
	hashF6 := test.EnumSimple_ENUM_VALUE_1
	struct1.F6["10"] = &hashF6
	struct1.F6["20"] = nil
	struct1.F7["10"] = test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F7["20"] = test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2 | test.FlagsSimple_FLAG_VALUE_3
	hashF8 := test.FlagsSimple_FLAG_VALUE_1 | test.FlagsSimple_FLAG_VALUE_2
	struct1.F8["10"] = &hashF8
	struct1.F8["20"] = nil
	s1 := test.NewStructSimple()
	s1.Id = 48
	struct1.F9["10"] = *s1
	s2 := test.NewStructSimple()
	s2.Id = 65
	struct1.F9["20"] = *s2
	struct1.F10["10"] = s1
	struct1.F10["20"] = nil

	// Serialize the struct to the FBE stream
	writer := test.NewStructHashModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 141)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 1570)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructHashModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 141)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 2)
	assert.EqualValues(t, struct2.F1["10"], 48)
	assert.EqualValues(t, struct2.F1["20"], 65)
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.EqualValues(t, *struct2.F2["10"], 97)
	assert.Nil(t, struct2.F2["20"])
	assert.EqualValues(t, len(struct2.F3), 2)
	assert.EqualValues(t, len(struct2.F3["10"]), 3)
	assert.EqualValues(t, len(struct2.F3["20"]), 3)
	assert.EqualValues(t, len(struct2.F4), 2)
	assert.EqualValues(t, len(*struct2.F4["10"]), 3)
	assert.Nil(t, struct2.F4["20"])
	assert.EqualValues(t, len(struct2.F5), 2)
	assert.EqualValues(t, struct2.F5["10"], test.EnumSimple_ENUM_VALUE_1)
	assert.EqualValues(t, struct2.F5["20"], test.EnumSimple_ENUM_VALUE_2)
	assert.EqualValues(t, len(struct2.F6), 2)
	assert.EqualValues(t, *struct2.F6["10"], test.EnumSimple_ENUM_VALUE_1)
	assert.Nil(t, struct2.F6["20"])
	assert.EqualValues(t, len(struct2.F7), 2)
	assert.EqualValues(t, struct2.F7["10"], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.EqualValues(t, struct2.F7["20"], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2|test.FlagsSimple_FLAG_VALUE_3)
	assert.EqualValues(t, len(struct2.F8), 2)
	assert.EqualValues(t, *struct2.F8["10"], test.FlagsSimple_FLAG_VALUE_1|test.FlagsSimple_FLAG_VALUE_2)
	assert.Nil(t, struct2.F8["20"])
	assert.EqualValues(t, len(struct2.F9), 2)
	assert.EqualValues(t, struct2.F9["10"].Id, 48)
	assert.EqualValues(t, struct2.F9["20"].Id, 65)
	assert.EqualValues(t, len(struct2.F10), 2)
	assert.EqualValues(t, struct2.F10["10"].Id, 48)
	assert.Nil(t, struct2.F10["20"])
}

func TestSerializationStructHashEx(t *testing.T) {
	// Create a new struct
	struct1 := test.NewStructHashEx()
	s1 := test.NewStructSimple()
	s1.Id = 48
	struct1.F1[s1.Key()] = struct {
		Key   test.StructSimple
		Value test.StructNested
	}{*s1, *test.NewStructNested()}
	s2 := test.NewStructSimple()
	s2.Id = 65
	struct1.F1[s2.Key()] = struct {
		Key   test.StructSimple
		Value test.StructNested
	}{*s2, *test.NewStructNested()}
	struct1.F2[s1.Key()] = struct {
		Key   test.StructSimple
		Value *test.StructNested
	}{*s1, test.NewStructNested()}
	struct1.F2[s2.Key()] = struct {
		Key   test.StructSimple
		Value *test.StructNested
	}{*s2, nil}

	// Serialize the struct to the FBE stream
	writer := test.NewStructHashExModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 142)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Check the serialized FBE size
	assert.EqualValues(t, writer.Buffer().Size(), 7879)

	// Deserialize the struct from the FBE stream
	reader := test.NewStructHashExModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 142)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct2, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, len(struct2.F1), 2)
	assert.EqualValues(t, struct2.F1[s1.Key()].Value.F1002, test.EnumTyped_ENUM_VALUE_2)
	assert.EqualValues(t, struct2.F1[s2.Key()].Value.F1002, test.EnumTyped_ENUM_VALUE_2)
	assert.EqualValues(t, len(struct2.F2), 2)
	assert.EqualValues(t, struct2.F2[s1.Key()].Value.F1002, test.EnumTyped_ENUM_VALUE_2)
	assert.Nil(t, struct2.F2[s2.Key()].Value)
}

func TestSerializationPtrSelfReference(t *testing.T) {
	// Create a new struct
	struct1 := simple.NewSimple()
	struct1.Info = "simple1"
	struct1.Simple = simple.NewSimpleFromFieldValues("simple2", nil, 2, []*simple.Simple{}, []simple.Simple{}, map[int32]*simple.Simple{}, map[int32]simple.Simple{})
	struct1.Depth = 1
	struct3 := simple.NewSimple()
	struct3.Info = "simple3"
	struct3.Depth = 3
	struct1.Spv = append(struct1.Spv, struct3)

	struct4 := simple.NewSimple()
	struct4.Info = "simple4"
	struct4.Depth = 4

	struct1.Sv = append(struct1.Sv, *struct4)

	struct5 := simple.NewSimple()
	struct5.Info = "simple5"
	struct5.Depth = 5
	struct6 := simple.NewSimple()
	struct6.Info = "simple6"
	struct6.Depth = 6
	struct6.Simple = simple.NewSimpleFromFieldValues("simple7", nil, 7, []*simple.Simple{}, []simple.Simple{}, map[int32]*simple.Simple{}, map[int32]simple.Simple{})
	struct8 := simple.NewSimple()
	struct8.Info = "simple8"
	struct8.Depth = 8
	struct6.Spv = append(struct6.Spv, struct8)

	struct1.Spm[5] = struct5
	struct1.Sm[6] = *struct6

	assert.EqualValues(t, `Simple(info="simple1",simple=Simple(info="simple2",simple=null,depth=2,spv=[0][],sv=[0][],spm=[0]<{}>,sm=[0]<{}>),depth=1,spv=[1][Simple(info="simple3",simple=null,depth=3,spv=[0][],sv=[0][],spm=[0]<{}>,sm=[0]<{}>)],sv=[1][Simple(info="simple4",simple=null,depth=4,spv=[0][],sv=[0][],spm=[0]<{}>,sm=[0]<{}>)],spm=[1]<{5->Simple(info="simple5",simple=null,depth=5,spv=[0][],sv=[0][],spm=[0]<{}>,sm=[0]<{}>)}>,sm=[1]<{6->Simple(info="simple6",simple=Simple(info="simple7",simple=null,depth=7,spv=[0][],sv=[0][],spm=[0]<{}>,sm=[0]<{}>),depth=6,spv=[1][Simple(info="simple8",simple=null,depth=8,spv=[0][],sv=[0][],spm=[0]<{}>,sm=[0]<{}>)],sv=[0][],spm=[0]<{}>,sm=[0]<{}>)}>)`, struct1.String())

	// Serialize the struct to the FBE stream
	writer := simple.NewSimpleModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 1)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(struct1)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := simple.NewSimpleModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 1)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	struct1Copy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, struct1Copy.Info, "simple1")
	assert.EqualValues(t, struct1Copy.Depth, 1)
	assert.EqualValues(t, struct1Copy.Simple.Info, "simple2")
	assert.EqualValues(t, struct1Copy.Simple.Simple, (*simple.Simple)(nil))
	assert.EqualValues(t, struct1Copy.Simple.Depth, 2)
	assert.EqualValues(t, len(struct1Copy.Spv), 1)
	assert.EqualValues(t, struct1Copy.Spv[0].Info, "simple3")
	assert.EqualValues(t, struct1Copy.Spv[0].Depth, 3)
	assert.EqualValues(t, len(struct1Copy.Sv), 1)
	assert.EqualValues(t, struct1Copy.Sv[0].Info, "simple4")
	assert.EqualValues(t, struct1Copy.Sv[0].Depth, 4)
	assert.EqualValues(t, len(struct1Copy.Spm), 1)
	assert.EqualValues(t, struct1Copy.Spm[5].Info, "simple5")
	assert.EqualValues(t, struct1Copy.Spm[5].Depth, 5)
	assert.EqualValues(t, len(struct1Copy.Sm), 1)
	assert.EqualValues(t, struct1Copy.Sm[6].Info, "simple6")
	assert.EqualValues(t, struct1Copy.Sm[6].Depth, 6)
	assert.EqualValues(t, struct1Copy.Sm[6].Simple.Info, "simple7")
	assert.EqualValues(t, struct1Copy.Sm[6].Simple.Simple, (*simple.Simple)(nil))
	assert.EqualValues(t, struct1Copy.Sm[6].Simple.Depth, 7)
	assert.EqualValues(t, len(struct1Copy.Sm[6].Spv), 1)
	assert.EqualValues(t, struct1Copy.Sm[6].Spv[0].Info, "simple8")
	assert.EqualValues(t, struct1Copy.Sm[6].Spv[0].Depth, 8)
}

func TestSerializationImport(t *testing.T) {
	extra := osa.NewExtraFromFieldValues("extra", "detail", osa.Sex_male, osa.MyFLags_flag1)

	pkgInfo := pkg.NewInfoFromFieldValues("pkgInfo", osa.Sex_male, osa.MyFLags_flag1, *extra)

	// Serialize the struct to the FBE stream
	writer := pkg.NewInfoModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 1)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(pkgInfo)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := pkg.NewInfoModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 1)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	pkgInfoCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.EqualValues(t, pkgInfo, pkgInfoCopy)
	assert.EqualValues(t, "extra", pkgInfoCopy.Extra.Name)
}

// Test Expr
func TestSerializationVariantString(t *testing.T) {
	nestedStringVariant := variants_ptr.NewExprFromValue("nested variant string");
	value := variants_ptr.NewExprContainerFromFieldValues("test variant string", &nestedStringVariant, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewExprContainerModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 2)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewExprContainerModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 2)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())


	assert.Equal(t, "string",  reflect.TypeOf(valueCopy.E).String())
	v, ok := (valueCopy.E).(string)
	assert.True(t, ok)
	assert.Equal(t, "test variant string", v)
	
	assert.Equal(t, "string",  reflect.TypeOf(*valueCopy.Eo).String())
	assert.NotNil(t, valueCopy.Eo)
	nv, ok := (*valueCopy.Eo).(string)
	assert.True(t, ok)
	assert.Equal(t, "nested variant string", nv)
}

func TestSerializationVariantBool(t *testing.T) {
	nestedBoolVariant := variants_ptr.NewExprFromValue(true);
	value := variants_ptr.NewExprContainerFromFieldValues(true, &nestedBoolVariant, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewExprContainerModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 2)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewExprContainerModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 2)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "bool",  reflect.TypeOf(valueCopy.E).String())
	v, ok := (valueCopy.E).(bool)
	assert.True(t, ok)
	assert.Equal(t, true, v)

	assert.Equal(t, "bool",  reflect.TypeOf(*valueCopy.Eo).String())
	assert.NotNil(t, valueCopy.Eo)
	nv, ok := (*valueCopy.Eo).(bool)
	assert.True(t, ok)
	assert.Equal(t, true, nv)
}

func TestSerializationVariantInt32(t *testing.T) {
	nestedInt32Variant := variants_ptr.NewExprFromValue(int32(24));
	value := variants_ptr.NewExprContainerFromFieldValues(int32(42), &nestedInt32Variant, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewExprContainerModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 2)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewExprContainerModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 2)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "int32",  reflect.TypeOf(valueCopy.E).String())
	v, ok := (valueCopy.E).(int32)
	assert.True(t, ok)
	assert.Equal(t, int32(42), v)

	assert.Equal(t, "int32",  reflect.TypeOf(*valueCopy.Eo).String())
	assert.NotNil(t, valueCopy.Eo)
	nv, ok := (*valueCopy.Eo).(int32)
	assert.True(t, ok)
	assert.Equal(t, int32(24), nv)
}

func TestSerializationVariantByteSlices(t *testing.T) {
	var nestedByteSliceVariant variants_ptr.Expr = []byte("000");
	value := variants_ptr.NewExprContainerFromFieldValues([]byte("ABCD"), &nestedByteSliceVariant, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewExprContainerModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 2)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewExprContainerModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 2)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "[]uint8",  reflect.TypeOf(valueCopy.E).String())
	v, ok := (valueCopy.E).([]byte)
	assert.True(t, ok)
	assert.Equal(t, 4, len(v))
	assert.Equal(t, "ABCD", string(v))

	assert.Equal(t, "[]uint8",  reflect.TypeOf(*valueCopy.Eo).String())
	assert.NotNil(t, valueCopy.Eo)
	nv, ok := (*valueCopy.Eo).([]byte)
	assert.True(t, ok)
	assert.Equal(t, 3, len(nv))
	assert.Equal(t, "000", string(nv))
}

func TestSerializationVariantStruct(t *testing.T) {
	simple := variants_ptr.NewSimpleFromFieldValues("this is too simple")
	var f64 float64 = 3.14
	f64v := variants_ptr.NewVFromValue(f64);

	value := variants_ptr.NewValueFromFieldValues(*simple, &f64v, nil)


	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "variants_ptr.Simple",  reflect.TypeOf(valueCopy.V).String())
	s, ok := (valueCopy.V).(variants_ptr.Simple)
	assert.True(t, ok)
	assert.Equal(t, "this is too simple", s.Name)

	assert.Equal(t, "float64",  reflect.TypeOf(*valueCopy.Vo).String())
	assert.NotNil(t, valueCopy.Vo)
	f64copy, ok := (*valueCopy.Vo).(float64)
	assert.True(t, ok)
	assert.Equal(t, 3.14, f64copy)
}

func TestSerializationVariantVariant(t *testing.T) {
	var expr variants_ptr.Expr = variants_ptr.NewExprFromValue("simple is complex")

	value := variants_ptr.NewValueFromFieldValues(expr, nil, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "string",  reflect.TypeOf(valueCopy.V).String())
	s, ok := (valueCopy.V).(string)
	assert.True(t, ok)
	assert.NotNil(t, s)
	assert.Equal(t, "simple is complex", s)
}

func TestSerializationVariantPtrStruct(t *testing.T) {
	// Create a new struct
	simple := variants_ptr.NewSimpleFromFieldValues("simple is complex")

	value := variants_ptr.NewValueFromFieldValues(simple, nil, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "*variants_ptr.Simple",  reflect.TypeOf(valueCopy.V).String())
	s, ok := (valueCopy.V).(*variants_ptr.Simple)
	assert.True(t, ok)
	assert.NotNil(t, s)
	assert.Equal(t, "simple is complex", s.Name)
}

func TestSerializationVariantVectorStruct(t *testing.T) {
	// Create a new struct
	simpleVector := []variants_ptr.Simple{{Name: "simple"}, {Name: "complex"}}

	var stringVector variants_ptr.V = []string{"ABC", "123"}

	var simplePtrVector variants_ptr.V = []*variants_ptr.Simple{variants_ptr.NewSimpleFromFieldValues("123"), variants_ptr.NewSimpleFromFieldValues("ABC")}

	value := variants_ptr.NewValueFromFieldValues(simpleVector, &stringVector, &simplePtrVector)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "[]variants_ptr.Simple",  reflect.TypeOf(valueCopy.V).String())
	tt, ok := (valueCopy.V).([]variants_ptr.Simple)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tt))
	assert.Equal(t, "simple", tt[0].Name)
	assert.Equal(t, "complex", tt[1].Name)

	assert.Equal(t, "[]string",  reflect.TypeOf(*valueCopy.Vo).String())
	tto, ok := (*valueCopy.Vo).([]string)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tto))
	assert.Equal(t, "ABC", tto[0])
	assert.Equal(t, "123", tto[1])

	assert.Equal(t, "[]*variants_ptr.Simple",  reflect.TypeOf(*valueCopy.Vo2).String())
	tto2, ok := (*valueCopy.Vo2).([]*variants_ptr.Simple)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tto2))
	assert.Equal(t, "123", tto2[0].Name)
	assert.Equal(t, "ABC", tto2[1].Name)
}

func TestSerializationVariantVectorInt32(t *testing.T) {
	// Create a new struct
	int32Vector := []int32{24, 42}

	value := variants_ptr.NewValueFromFieldValues(int32Vector, nil, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "[]int32",  reflect.TypeOf(valueCopy.V).String())
	tt, ok := (valueCopy.V).([]int32)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tt))
	assert.Equal(t, int32(42), tt[1])
	assert.Equal(t, int32(24), tt[0])
}

func TestSerializationVariantMapStructInt(t *testing.T) {
	// Create a new struct
	simple2Int := make(map[int32]variants_ptr.Simple)
	simple2Int[24] = variants_ptr.Simple{Name: "simple"}
	simple2Int[42] = variants_ptr.Simple{Name: "complex"}

	var int322Bytes variants_ptr.V = map[int32][]byte{24: []byte("ABC"), 42: []byte("123")}
	var string2Bytes variants_ptr.V = map[string][]byte{"24": []byte("ABC"), "42": []byte("123")}

	value := variants_ptr.NewValueFromFieldValues(simple2Int, &int322Bytes, &string2Bytes)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "map[int32]variants_ptr.Simple",  reflect.TypeOf(valueCopy.V).String())
	tt, ok := (valueCopy.V).(map[int32]variants_ptr.Simple)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tt))
	assert.Equal(t, "simple", tt[24].Name)
	assert.Equal(t, "complex", tt[42].Name)

	assert.Equal(t, "map[int32][]uint8",  reflect.TypeOf(*valueCopy.Vo).String())
	tto, ok := (*valueCopy.Vo).(map[int32][]byte)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tto))
	assert.Equal(t, "ABC", string(tto[24]))
	assert.Equal(t, "123", string(tto[42]))

	assert.Equal(t, "map[string][]uint8",  reflect.TypeOf(*valueCopy.Vo2).String())
	tto2, ok := (*valueCopy.Vo2).(map[string][]byte)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tto2))
	assert.Equal(t, "ABC", string(tto2["24"]))
	assert.Equal(t, "123", string(tto2["42"]))
}

func TestSerializationVariantVectorBytes(t *testing.T) {
	// Create a new struct
	byteSlices := make([][]byte, 0)
	byteSlices = append(byteSlices, []byte("ABCDE"))
	byteSlices = append(byteSlices, []byte("12345"))

	value := variants_ptr.NewValueFromFieldValues(byteSlices, nil, nil)

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 3)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(value)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 3)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, "[][]uint8",  reflect.TypeOf(valueCopy.V).String())
	tt, ok := (valueCopy.V).([][]byte)
	assert.True(t, ok)
	assert.Equal(t, 2, len(tt))
	assert.Equal(t, "ABCDE", string(tt[0]))
	assert.Equal(t, "12345", string(tt[1]))
}

// Variant[] in struct
func TestSerializationContainerVariant(t *testing.T) {
	var v0 variants_ptr.V = int32(42)
	var v1 variants_ptr.V = variants_ptr.NewVFromValue("nested variant string")
	var v2 variants_ptr.V = variants_ptr.NewSimpleFromFieldValues("this is too simple")
	
	valueContainer := variants_ptr.NewValueContainerFromFieldValues([]variants_ptr.V{v0, v1, v2}, map[int32]variants_ptr.V{
		1: true,
		2: "this is a long story",
		3: variants_ptr.NewSimpleFromFieldValues("this is too complex"),
	})

	// Serialize the struct to the FBE stream
	writer := variants_ptr.NewValueContainerModel(fbe.NewEmptyBuffer())
	assert.EqualValues(t, writer.Model().FBEType(), 4)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4)
	serialized, err := writer.Serialize(valueContainer)
	assert.Nil(t, err)
	assert.EqualValues(t, serialized, writer.Buffer().Size())
	assert.True(t, writer.Verify())
	writer.Next(serialized)
	assert.EqualValues(t, writer.Model().FBEOffset(), 4+writer.Buffer().Size())

	// Deserialize the struct from the FBE stream
	reader := variants_ptr.NewValueContainerModel(writer.Buffer())
	assert.EqualValues(t, reader.Model().FBEType(), 4)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4)
	assert.True(t, reader.Verify())
	valueContainerCopy, deserialized, err := reader.Deserialize()
	assert.Nil(t, err)
	assert.EqualValues(t, deserialized, reader.Buffer().Size())
	reader.Next(deserialized)
	assert.EqualValues(t, reader.Model().FBEOffset(), 4+reader.Buffer().Size())

	assert.Equal(t, 3, len(valueContainerCopy.Vv))
	
	assert.Equal(t, "int32",  reflect.TypeOf(valueContainerCopy.Vv[0]).String())
	t0, ok := (valueContainerCopy.Vv[0]).(int32)
	assert.True(t, ok)
	assert.Equal(t, int32(42), t0)

	assert.Equal(t, "string",  reflect.TypeOf(valueContainerCopy.Vv[1]).String())
	t1, ok := (valueContainerCopy.Vv[1]).(string)
	assert.True(t, ok)
	assert.Equal(t, "nested variant string", t1)

	assert.Equal(t, "*variants_ptr.Simple",  reflect.TypeOf(valueContainerCopy.Vv[2]).String())
	t2, ok := (valueContainerCopy.Vv[2]).(*variants_ptr.Simple)
	assert.True(t, ok)
	assert.Equal(t, "this is too simple", t2.Name)

	assert.Equal(t, 3, len(valueContainerCopy.Vv))
	
	assert.Equal(t, "bool",  reflect.TypeOf(valueContainerCopy.Vm[1]).String())
	tt0, ok := (valueContainerCopy.Vm[1]).(bool)
	assert.True(t, ok)
	assert.Equal(t, true, tt0)

	assert.Equal(t, "string",  reflect.TypeOf(valueContainerCopy.Vm[2]).String())
	tt1, ok := (valueContainerCopy.Vm[2]).(string)
	assert.True(t, ok)
	assert.Equal(t, "this is a long story", tt1)

	assert.Equal(t, "*variants_ptr.Simple",  reflect.TypeOf(valueContainerCopy.Vm[3]).String())
	tt2, ok := (valueContainerCopy.Vm[3]).(*variants_ptr.Simple)
	assert.True(t, ok)
	assert.Equal(t, "this is too complex", tt2.Name)
}

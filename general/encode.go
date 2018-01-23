package general

import (
	"reflect"
	"github.com/thrift-iterator/go/spi"
	"github.com/thrift-iterator/go/protocol"
)

func ExportEncoders() map[reflect.Type]spi.ValEncoder{
	return map[reflect.Type]spi.ValEncoder {
		reflect.TypeOf(([]interface{})(nil)): &generalListEncoder{},
		reflect.TypeOf((map[interface{}]interface{})(nil)): &generalMapEncoder{},
		reflect.TypeOf((map[protocol.FieldId]interface{})(nil)): &generalStructEncoder{},
	}
}

func generalWriterOf(sample interface{}) (protocol.TType, func(val interface{}, stream spi.Stream)) {
	switch sample.(type) {
	case bool:
		return protocol.TypeBool, writeBool
	case int8:
		return protocol.TypeI08, writeInt8
	case uint8:
		return protocol.TypeI08, writeInt8
	case int16:
		return protocol.TypeI16, writeInt16
	case uint16:
		return protocol.TypeI16, writeUint16
	case int32:
		return protocol.TypeI32, writeInt32
	case uint32:
		return protocol.TypeI32, writeUint32
	case int64:
		return protocol.TypeI64, writeInt64
	case uint64:
		return protocol.TypeI64, writeUint64
	case float64:
		return protocol.TypeDouble, writeFloat64
	case string:
		return protocol.TypeString, writeString
	case []byte:
		return protocol.TypeString, writeBinary
	default:
		panic("unsupported type")
	}
}

func writeBool(val interface{}, stream spi.Stream) {
	stream.WriteBool(val.(bool))
}

func writeInt8(val interface{}, stream spi.Stream) {
	stream.WriteInt8(val.(int8))
}

func writeUint8(val interface{}, stream spi.Stream) {
	stream.WriteUint8(val.(uint8))
}

func writeInt16(val interface{}, stream spi.Stream) {
	stream.WriteInt16(val.(int16))
}

func writeUint16(val interface{}, stream spi.Stream) {
	stream.WriteUint16(val.(uint16))
}

func writeInt32(val interface{}, stream spi.Stream) {
	stream.WriteInt32(val.(int32))
}

func writeUint32(val interface{}, stream spi.Stream) {
	stream.WriteUint32(val.(uint32))
}

func writeInt64(val interface{}, stream spi.Stream) {
	stream.WriteInt64(val.(int64))
}

func writeUint64(val interface{}, stream spi.Stream) {
	stream.WriteUint64(val.(uint64))
}

func writeFloat64(val interface{}, stream spi.Stream) {
	stream.WriteFloat64(val.(float64))
}

func writeString(val interface{}, stream spi.Stream) {
	stream.WriteString(val.(string))
}

func writeBinary(val interface{}, stream spi.Stream) {
	stream.WriteBinary(val.([]byte))
}
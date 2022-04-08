package kotlin

import "strings"

func NameConv(in string) string {
	return strings.Title(in)
}

func TypeConv(in string) string {
	switch in {
	case "uint8":
		return "UByte"
	case "uint16":
		return "UShort"
	case "uint32":
		return "UInt"
	case "uint64":
		return "ULong"
	case "int8":
		return "Byte"
	case "int16":
		return "Short"
	case "int32":
		return "Int"
	case "int64":
		return "Long"
	case "float32":
		return "Float"
	case "float64":
		return "Double"
	case "bool":
		return "Boolean"
	case "string":
		return "String"
	case "bytes":
		return "ByteArray"
	default:
		return NameConv(in)
	}
}

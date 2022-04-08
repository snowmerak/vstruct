package kotlin

import "strings"

func NameConv(in string) string {
	return strings.Title(in)
}

func NumberConv(unsigned bool, bit int) string {
	switch unsigned {
	case true:
		switch bit {
		case 8:
			return "UByte"
		case 16:
			return "UShort"
		case 32:
			return "UInt"
		case 64:
			return "ULong"
		}
	case false:
		switch bit {
		case 8:
			return "Byte"
		case 16:
			return "Short"
		case 32:
			return "Int"
		case 64:
			return "Long"
		}
	}
	return "ULong"
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

package internal

import (
	"fmt"
	"strconv"
)

func TypeConvert(strValue string, dataType string, ptr bool) (value interface{}, err error) {
	// 去掉dataType前的*
	if len(dataType) > 1 && dataType[0] == '*' {
		ptr = true
		dataType = dataType[1:]
	}
	switch dataType {
	case "string":
		if ptr {
			return &strValue, nil
		} else {
			return strValue, nil
		}
	case "int":
		if ptr {
			return TypeConvertIntPtr(strValue)
		} else {
			return TypeConvertInt(strValue)
		}
	case "int8":
		if ptr {
			return TypeConvertInt8Ptr(strValue)
		} else {
			return TypeConvertInt8(strValue)
		}
	case "int16":
		if ptr {
			return TypeConvertInt16Ptr(strValue)
		} else {
			return TypeConvertInt16(strValue)
		}
	case "int32":
		if ptr {
			return TypeConvertInt32Ptr(strValue)
		} else {
			return TypeConvertInt32(strValue)
		}
	case "int64":
		if ptr {
			return TypeConvertInt64Ptr(strValue)
		} else {
			return TypeConvertInt64(strValue)
		}
	case "uint":
		if ptr {
			return TypeConvertUIntPtr(strValue)
		} else {
			return TypeConvertUInt(strValue)
		}
	case "uint8":
		if ptr {
			return TypeConvertUInt8Ptr(strValue)
		} else {
			return TypeConvertUInt8(strValue)
		}
	case "uint16":
		if ptr {
			return TypeConvertUInt16Ptr(strValue)
		} else {
			return TypeConvertUInt16(strValue)
		}
	case "uint32":
		if ptr {
			return TypeConvertUInt32Ptr(strValue)
		} else {
			return TypeConvertUInt32(strValue)
		}
	case "uint64":
		if ptr {
			return TypeConvertUInt64Ptr(strValue)
		} else {
			return TypeConvertUInt64(strValue)
		}
	case "float32":
		if ptr {
			return TypeConvertFloat32Ptr(strValue)
		} else {
			return TypeConvertFloat32(strValue)
		}
	case "float64":
		if ptr {
			return TypeConvertFloat64Ptr(strValue)
		} else {
			return TypeConvertFloat64(strValue)
		}
	case "bool":
		if ptr {
			return TypeConvertBoolPtr(strValue)
		} else {
			return TypeConvertBool(strValue)
		}
	default:
		return nil, fmt.Errorf("未知类型")
	}
}

func TypeConvertInt(strValue string) (value int, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	value = int(v2)
	return
}
func TypeConvertIntPtr(strValue string) (value *int, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	value = new(int)
	*value = int(v2)
	return
}

func TypeConvertInt8(strValue string) (value int8, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 8)
	if err != nil {
		return
	}
	value = int8(v2)
	return
}
func TypeConvertInt8Ptr(strValue string) (value *int8, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 8)
	if err != nil {
		return
	}
	*value = int8(v2)
	return
}
func TypeConvertInt16(strValue string) (value int16, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 16)
	if err != nil {
		return
	}
	value = int16(v2)
	return
}
func TypeConvertInt16Ptr(strValue string) (value *int16, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 16)
	if err != nil {
		return
	}
	*value = int16(v2)
	return
}
func TypeConvertInt32(strValue string) (value int32, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return
	}
	value = int32(v2)
	return
}
func TypeConvertInt32Ptr(strValue string) (value *int32, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return
	}
	*value = int32(v2)
	return
}
func TypeConvertInt64(strValue string) (value int64, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	value = v2
	return
}
func TypeConvertInt64Ptr(strValue string) (value *int64, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	*value = v2
	return
}
func TypeConvertFloat32(strValue string) (value float32, err error) {
	v2 := float64(0)
	v2, err = strconv.ParseFloat(strValue, 32)
	if err != nil {
		return
	}
	value = float32(v2)
	return
}
func TypeConvertFloat32Ptr(strValue string) (value *float32, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := float64(0)
	v2, err = strconv.ParseFloat(strValue, 32)
	if err != nil {
		return
	}
	*value = float32(v2)
	return
}
func TypeConvertFloat64(strValue string) (value float64, err error) {
	v2 := float64(0)
	v2, err = strconv.ParseFloat(strValue, 64)
	if err != nil {
		return
	}
	value = v2
	return
}
func TypeConvertFloat64Ptr(strValue string) (value *float64, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := float64(0)
	v2, err = strconv.ParseFloat(strValue, 64)
	if err != nil {
		return
	}
	*value = v2
	return
}
func TypeConvertBool(strValue string) (value bool, err error) {
	if strValue == "" {
		return false, nil
	}
	v2 := false
	v2, err = strconv.ParseBool(strValue)
	if err != nil {
		return
	}
	value = v2
	return
}
func TypeConvertBoolPtr(strValue string) (value *bool, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := false
	v2, err = strconv.ParseBool(strValue)
	if err != nil {
		return
	}
	*value = v2
	return
}

func TypeConvertUInt(strValue string) (value uint, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	value = uint(v2)
	return
}
func TypeConvertUIntPtr(strValue string) (value *uint, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	value = new(uint)
	*value = uint(v2)
	return
}

func TypeConvertUInt8(strValue string) (value uint8, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 8)
	if err != nil {
		return
	}
	value = uint8(v2)
	return
}
func TypeConvertUInt8Ptr(strValue string) (value *uint8, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 8)
	if err != nil {
		return
	}
	*value = uint8(v2)
	return
}
func TypeConvertUInt16(strValue string) (value uint16, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 16)
	if err != nil {
		return
	}
	value = uint16(v2)
	return
}
func TypeConvertUInt16Ptr(strValue string) (value *uint16, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 16)
	if err != nil {
		return
	}
	*value = uint16(v2)
	return
}
func TypeConvertUInt32(strValue string) (value uint32, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return
	}
	value = uint32(v2)
	return
}
func TypeConvertUInt32Ptr(strValue string) (value *uint32, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return
	}
	*value = uint32(v2)
	return
}
func TypeConvertUInt64(strValue string) (value uint64, err error) {
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	value = uint64(v2)
	return
}
func TypeConvertUInt64Ptr(strValue string) (value *uint64, err error) {
	if strValue == "" {
		return nil, nil
	}
	v2 := int64(0)
	v2, err = strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return
	}
	*value = uint64(v2)
	return
}

package iris_enhance

import (
	"context"
	"fmt"
	"github.com/celt237/iris-enhance/internal"
	"github.com/kataras/iris/v12"
	"reflect"
)

type ErrorWithCode interface {
	error
	Code() int
}

type ApiHandler interface {
	WrapContext(ctx iris.Context) context.Context
	Success(ctx iris.Context, produceType string, data interface{})
	CodeError(ctx iris.Context, produceType string, data interface{}, code int, message string)
	Error(ctx iris.Context, produceType string, data interface{}, message string)
	HandleCustomerAnnotation(ctx iris.Context, annotation string, opt ...string) error
}

func GetParamFromContext[T any](ctx iris.Context, paramName string, dataType string, paramType string, ptr bool, required bool) (value T, err error) {
	value, err = getDefaultValue[T]()
	var v1 interface{}
	if paramType == "path" {
		str := ctx.Params().Get(paramName)
		if str == "" {
			err = fmt.Errorf("参数%s不能为空", paramName)
			return
		}
		v1, err = internal.TypeConvert(str, dataType, ptr)
		if err != nil {
			return
		}
		if v1 != nil {
			value = v1.(T)
		}
	} else if paramType == "query" {
		str := ctx.URLParam(paramName)
		if str == "" && required {
			err = fmt.Errorf("参数%s不能为空", paramName)
			return
		}
		v1, err = internal.TypeConvert(str, dataType, ptr)
		if err != nil {
			return
		}
		if v1 != nil {
			value = v1.(T)
		}
	} else if paramType == "header" {
		str := ctx.GetHeader(paramName)
		if str == "" && required {
			err = fmt.Errorf("参数%s不能为空", paramName)
			return
		}
		v1, err = internal.TypeConvert(str, dataType, ptr)
		if err != nil {
			return
		}
		if v1 != nil {
			value = v1.(T)
		}
	} else if paramType == "body" {
		err = ctx.ReadJSON(value)
		if err != nil {
			return
		}
	}
	return
}

func getDefaultValue[T any]() (result T, err error) {
	var tempValue interface{}
	t := reflect.TypeOf((*T)(nil)).Elem()
	//v := reflect.New(t)
	switch t.Kind() {
	case reflect.Struct, reflect.Map:
		tempValue = reflect.Zero(t).Interface()
	case reflect.Slice, reflect.Array:
		tempValue = reflect.MakeSlice(t, 0, 0).Interface()
	case reflect.Ptr:
		subT := t.Elem()
		switch subT.Kind() {
		case reflect.Struct, reflect.Map:
			elem := reflect.New(t.Elem())
			elem.Elem().Set(reflect.Zero(t.Elem()))
			tempValue = elem.Interface()
		case reflect.Slice, reflect.Array:
			elem := reflect.New(t.Elem())
			elem.Elem().Set(reflect.MakeSlice(t.Elem(), 0, 0))
			tempValue = elem.Interface()
		default:
			break
		}
	default:
		break
	}
	//return v.Elem().Interface().(T), nil
	if tempValue != nil {
		result = tempValue.(T)
	}
	return result, nil
}

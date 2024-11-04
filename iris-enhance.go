package iris_enhance

import (
	"context"
	"fmt"
	"github.com/celt237/iris-enhance/internal"
	"github.com/kataras/iris/v12"
	"mime/multipart"
	"reflect"
)

type ErrorWithCode interface {
	error
	Code() int
}

type FileInfo interface {
	GetFileHeader() *multipart.FileHeader
	GetFile() multipart.File
}

type FileInfoImpl struct {
	fileHeader *multipart.FileHeader
	file       multipart.File
}

func (f *FileInfoImpl) GetFileHeader() *multipart.FileHeader {
	return f.fileHeader
}

func (f *FileInfoImpl) GetFile() multipart.File {
	return f.GetFile()
}

func NewFileInfo(file multipart.File, fileHeader *multipart.FileHeader) FileInfo {
	return &FileInfoImpl{file: file, fileHeader: fileHeader}
}

type ApiHandler interface {
	// WrapContext 从iris.Context中获取context.Context
	WrapContext(ctx iris.Context) context.Context

	// Success 成功返回
	// ctx iris.Context 上下文
	// produceType string 返回类型
	// data interface{} 返回数据
	Success(ctx iris.Context, produceType string, data interface{})

	// CodeError 失败返回
	// ctx iris.Context 上下文
	// produceType string 返回类型
	// data interface{} 返回数据
	// code int 错误码
	// err error 错误
	CodeError(ctx iris.Context, produceType string, data interface{}, code int, err error)

	// Error 失败返回
	// ctx iris.Context 上下文
	// produceType string 返回类型
	// data interface{} 返回数据
	// err error 错误
	Error(ctx iris.Context, produceType string, data interface{}, err error)

	// HandleCustomerAnnotation 处理自定义注解
	// ctx iris.Context 上下文
	// annotation string 注解名
	// opt ...string 参数
	HandleCustomerAnnotation(ctx iris.Context, annotation string, opt ...string) error
}

// GetParamFromContext 从iris.Context中获取参数
// ctx iris.Context 上下文
// paramName string 参数名
// dataType string 数据类型
// paramType string 参数类型
// ptr bool 是否指针
// required bool 是否必须
func GetParamFromContext[T any](ctx iris.Context, paramName string, dataType string, paramType string, ptr bool, required bool) (value T, err error) {
	value, err = getDefaultValue[T]()
	var v1 interface{}
	if paramType == internal.ParamTypePath {
		str := ctx.Params().Get(paramName)
		if str == "" {
			err = fmt.Errorf("param %s can not be empty", paramName)
			return value, err
		}
		v1, err = internal.TypeConvert(str, dataType, ptr)
		if err != nil {
			return value, err
		}
		if v1 != nil {
			value = v1.(T)
		}
	} else if paramType == internal.ParamTypeQuery {
		str := ctx.URLParam(paramName)
		if str == "" && required {
			err = fmt.Errorf("param %s can not be empty", paramName)
			return value, err
		}
		v1, err = internal.TypeConvert(str, dataType, ptr)
		if err != nil {
			return value, err
		}
		if v1 != nil {
			value = v1.(T)
		}
	} else if paramType == internal.ParamTypeHeader {
		str := ctx.GetHeader(paramName)
		if str == "" && required {
			err = fmt.Errorf("param %s can not be empty", paramName)
			return value, err
		}
		v1, err = internal.TypeConvert(str, dataType, ptr)
		if err != nil {
			return value, err
		}
		if v1 != nil {
			value = v1.(T)
		}
	} else if paramType == internal.ParamTypeBody {
		err = ctx.ReadJSON(value)
		if err != nil {
			return value, err
		}
	} else if paramType == internal.ParamFormData {
		var file multipart.File
		var fileHeader *multipart.FileHeader
		file, fileHeader, err = ctx.FormFile(paramName)
		if err != nil {
			return value, err
		}
		// 判断value的类型是FileInfo接口
		tType := reflect.TypeOf(value)
		if tType.Implements(reflect.TypeOf((*FileInfo)(nil)).Elem()) {
			fileInfo := NewFileInfo(file, fileHeader)
			value = fileInfo.(T)
		} else {
			err = fmt.Errorf("param %s type is not FileInfo", paramName)
			return value, err
		}

	}
	return value, err
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

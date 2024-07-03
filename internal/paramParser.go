package internal

import (
	"fmt"
	go_annotation "github.com/celt237/go-annotation"
	"strings"
)

type ParamParser struct {
	funcDecl *go_annotation.MethodDesc
}

func NewParamParser(funcDecl *go_annotation.MethodDesc) *ParamParser {
	return &ParamParser{funcDecl: funcDecl}
}

func (p *ParamParser) Parse() (params []*MethodParam, err error) {
	methodParams, err := p.getParamsFromMethod() // 获取方法参数
	if err != nil {
		return nil, err
	}
	commentParams, err := p.getParamsFromComment() // 解析方法注释中的参数信息
	if err != nil {
		return nil, err
	}
	params, err = p.handleParam(methodParams, commentParams) // 比较方法参数和注释中的参数信息
	return params, err
}

func (p *ParamParser) getParamsFromMethod() ([]*MethodParam, error) {
	params := make([]*MethodParam, 0)
	for _, param := range p.funcDecl.Params {
		if param.DataType == "context.Context" {
			continue
		}
		params = append(params, &MethodParam{Field: *param})
	}
	return params, nil
}

func (p *ParamParser) getParamsFromComment() (params map[string]*MethodParam, err error) {
	params = make(map[string]*MethodParam)
	for _, anno := range p.funcDecl.Annotations {
		if anno.Name == ZParamTag {
			for _, attribute := range anno.Attributes {
				if len(attribute) >= 5 {
					param := &MethodParam{
						Field: go_annotation.Field{
							Name: attribute["0"],
						},
						ParamDataType: attribute["2"],
						ParamType:     attribute["1"],
						Required:      strings.ToLower(attribute["3"]) == "true",
						Desc:          attribute["4"],
					}
					params[param.Name] = param
				}
			}

		}
	}
	return params, err
}

func (p *ParamParser) handleParam(methodParams []*MethodParam, commentParams map[string]*MethodParam) (params []*MethodParam, err error) {
	// 校验params 和method中的参数列表是否一致
	// 遍历methodParams
	//params = make([]*MethodParam, 0)
	// 遍历methodParams
	for _, param := range methodParams {
		if _, ok := commentParams[param.Name]; !ok {
			err = fmt.Errorf("param %s not found in comment", param.Name)
			return params, err
		}
		if !p.matchDataType(commentParams[param.Name].ParamType, param.RealDataType, commentParams[param.Name].ParamDataType) {
			err = fmt.Errorf("param %s type not match", param.Name)
			return params, err
		}
		param.ParamType = commentParams[param.Name].ParamType
		param.ParamDataType = commentParams[param.Name].ParamDataType
		param.Required = commentParams[param.Name].Required
		param.Desc = commentParams[param.Name].Desc
		//if param.RealDataType != commentParams[param.Name].RealDataType {
		//	err = fmt.Errorf("param %s type not match", param.Name)
		//	return params, err
		//}
		//params = append(params, commentParams[param.Name])
	}
	return methodParams, err
}

// int string number boolean array object
var typeMap = map[string][]string{
	"int":     {"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"},
	"string":  {"string"},
	"number":  {"float32", "float64"},
	"boolean": {"bool"},
	//"array":   {"[]"},
	//"object":  {"struct", "map"},
}

func (p *ParamParser) matchDataType(paramType string, goDataType string, commentDataType string) bool {
	// 匹配go数据类型和注释数据类型
	switch paramType {
	case ParamTypePath, ParamTypeQuery:
		for _, t := range typeMap[commentDataType] {
			if goDataType == t {
				return true
			}
		}
		return false
	case ParamTypeBody:
		// 判断是否切片
		if strings.HasPrefix(goDataType, "[]") {
			if commentDataType == DataTypeArray {
				return true
			} else {
				return false
			}
		} else if strings.HasPrefix(goDataType, "map") {
			if commentDataType == DataTypeObject {
				return true
			} else {
				return false
			}
		} else {
			if commentDataType == goDataType {
				return true
			} else {
				return false
			}
		}
	case ParamTypeHeader, ParamTypeCookie:
		// todo
		return true
	default:
		return false
	}
}

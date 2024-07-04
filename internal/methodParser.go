package internal

import (
	"fmt"
	go_annotation "github.com/celt237/go-annotation"
)

type MethodParser struct {
	funcDecl *go_annotation.MethodDesc
}

func NewMethodParser(funcDecl *go_annotation.MethodDesc) *MethodParser {
	return &MethodParser{funcDecl: funcDecl}
}

func (m *MethodParser) Parse(resultType string, errorCode string) (method *MethodDesc, err error) {
	method = &MethodDesc{
		MethodDesc:        *m.funcDecl,
		Params:            make([]*MethodParam, 0),
		CustomAnnotations: make(map[string]*go_annotation.Annotation),
		ApiResultType:     resultType,
		Produce:           "application/json",
		Accept:            "application/json",
		Description:       Ternary(m.funcDecl.Description != "", m.funcDecl.Description, m.funcDecl.Name),
		Summary:           Ternary(m.funcDecl.Description != "", m.funcDecl.Description, m.funcDecl.Name),
	}
	if errorCode != "" {
		method.ErrorCode = errorCode
	} else {
		method.ErrorCode = DefaultErrorCode
	}
	err = m.parseResult(method) // 解析方法返回值
	if err != nil {
		return nil, m.wrapperError(err)
	}
	method.ApiResultDataType = method.Result.RealDataType
	for _, annotation := range method.Annotations {
		commentParser := GetMethodCommentParser(annotation)
		err = commentParser.Parse(annotation, method)
		if err != nil {
			return nil, m.wrapperError(err)
		}
	}

	// 校验是否有缺失的字段
	if method.Path == "" || method.Method == "" {
		return nil, m.wrapperError(fmt.Errorf("no router tag found"))
	}
	if method.ApiResultType == "" {
		return nil, m.wrapperError(fmt.Errorf("no reply type found"))
	}
	paramParser := NewParamParser(m.funcDecl)
	method.Params, err = paramParser.Parse() // 解析方法参数
	if err != nil {
		err = m.wrapperError(fmt.Errorf("parse params error: %s", err.Error()))
		return nil, err
	}
	return method, err
}

func (m *MethodParser) wrapperError(err error) error {
	return fmt.Errorf("method:[%s] %s", m.funcDecl.Name, err.Error())
}

func (m *MethodParser) parseResult(method *MethodDesc) (err error) {
	if m.funcDecl.Results != nil {
		//fmt.Println("Results:")
		// 判断只能有2个返回值
		if len(m.funcDecl.Results) != 2 {
			err = m.wrapperError(fmt.Errorf("results length is not 2"))
			return err
		}
		for _, result := range m.funcDecl.Results {
			if result.DataType != "error" {
				method.Result = result
				return nil
			}
		}
	}
	return m.wrapperError(fmt.Errorf("no result type found"))
}

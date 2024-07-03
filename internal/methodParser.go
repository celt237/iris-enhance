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
		return nil, err
	}
	method.ApiResultDataType = method.Result.RealDataType
	for _, annotation := range method.Annotations {
		commentParser := GetMethodCommentParser(annotation)
		err = commentParser.Parse(annotation, method)
		if err != nil {
			return nil, err
		}
	}

	// 校验是否有缺失的字段
	if method.Path == "" || method.Method == "" {
		return nil, fmt.Errorf("no router tag found")
	}
	if method.ApiResultType == "" {
		return nil, fmt.Errorf("no reply type found")
	}
	paramParser := NewParamParser(m.funcDecl)
	method.Params, err = paramParser.Parse() // 解析方法参数
	if err != nil {
		return nil, err
	}
	//m.fillParamFormat(method) // 设置方法参数格式
	return method, err
}

func (m *MethodParser) parseResult(method *MethodDesc) (err error) {
	if m.funcDecl.Results != nil {
		//fmt.Println("Results:")
		// 判断只能有2个返回值
		if len(m.funcDecl.Results) != 2 {
			err = fmt.Errorf("results length is not 2")
			return err
		}
		for _, result := range m.funcDecl.Results {
			if result.DataType != "error" {
				method.Result = result
				return nil
			}
		}
	}
	return fmt.Errorf("no result type found")
}

//func (m *MethodParser) getAtComments(funcDecl *go_annotation.MethodDesc) (comments []string, err error) {
//	comments = make([]string, 0)
//	if funcDecl.Doc != nil {
//		//fmt.Println("Comments:")
//		for _, com := range funcDecl.Doc.List {
//			atIndex := strings.Index(com.Text, "@")
//			if atIndex != -1 {
//				commentText := com.Text[atIndex:]
//				comments = append(comments, commentText)
//			}
//		}
//	}
//	return comments, err
//}

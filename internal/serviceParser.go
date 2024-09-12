package internal

import (
	"fmt"
	go_annotation "github.com/celt237/go-annotation"
	"strings"
)

type ServiceParser struct {
	genDecl *go_annotation.StructDesc
}

func NewServiceParser(genDecl *go_annotation.StructDesc) *ServiceParser {
	return &ServiceParser{genDecl: genDecl}
}

func (s *ServiceParser) wrapperError(err error) error {
	return fmt.Errorf("service:[%s] %s", s.genDecl.Name, err.Error())
}

func (s *ServiceParser) Parse(resultType string, errorCode string, imports []string) (sDesc *ServiceDesc, err error) {
	importDict := make(map[string]*go_annotation.ImportDesc)
	importDict["context"] = &go_annotation.ImportDesc{Name: "context", HasAlias: false, Path: "context"}
	//importDict["fmt"] = &go_annotation.ImportDesc{Name: "fmt", HasAlias: false, Path: "fmt"}
	importDict["iris"] = &go_annotation.ImportDesc{Name: "iris", HasAlias: true, Path: "github.com/kataras/iris/v12"}
	importDict["iris-enhance"] = &go_annotation.ImportDesc{Name: "iris-enhance", HasAlias: false, Path: "github.com/celt237/iris-enhance"}
	if !strings.HasSuffix(s.genDecl.Name, "Service") || len(s.genDecl.Annotations) == 0 {
		return nil, nil
	}
	if _, ok := s.genDecl.Annotations[ZServiceTag]; !ok {
		return nil, nil
	}
	serviceImports := s.genDecl.Imports
	for name, imp := range importDict {
		if _, ok := serviceImports[name]; !ok {
			serviceImports[name] = imp
		}
	}
	extraImports := make([]string, 0)
	for _, imp := range imports {
		hasImport := false
		for name, imp2 := range serviceImports {
			if imp == imp2.Path {
				serviceImports[name] = imp2
				hasImport = true
				break
			}
		}
		if !hasImport {
			extraImports = append(extraImports, imp)
		}
	}
	sDesc = &ServiceDesc{
		Name:         strings.TrimSuffix(s.genDecl.Name, "Service"),
		PackageName:  "handler",
		Imports:      serviceImports,
		ExtraImports: extraImports,
		Comments:     s.genDecl.Comments,
		Annotations:  s.genDecl.Annotations,
		Description:  s.genDecl.Description,
		ReplyType:    resultType,
	}

	err = s.parseReplyType(resultType, sDesc)
	if err != nil {
		return nil, s.wrapperError(err)
	}

	for _, imp := range importDict {
		if _, ok := s.genDecl.Imports[imp.Name]; !ok {
			s.genDecl.Imports[imp.Name] = imp
		}
	}
	for _, methodData := range s.genDecl.Methods {
		methodParser := NewMethodParser(methodData)
		methodDesc, err := methodParser.Parse(sDesc.ReplyType, errorCode)
		if err != nil {
			return nil, s.wrapperError(err)
		}
		sDesc.Methods = append(sDesc.Methods, methodDesc)
	}
	return sDesc, nil
}

func (s *ServiceParser) parseReplyType(resultType string, sDesc *ServiceDesc) (err error) {
	// 判断是否有s.genDecl.Annotations[ZReplyTypeTag]
	if _, ok := s.genDecl.Annotations[ZReplyTypeTag]; ok {
		annotation := s.genDecl.Annotations[ZReplyTypeTag]
		if len(annotation.Attributes) > 1 {
			return s.wrapperError(fmt.Errorf(ZReplyTypeTag + " format is incorrect"))
		}
		attribute := annotation.Attributes[0]
		if len(attribute) >= 1 {
			sDesc.ReplyType = attribute["0"]
			return nil
		} else {
			return s.wrapperError(fmt.Errorf(ZReplyTypeTag + " format is incorrect"))
		}
	}
	sDesc.ReplyType = resultType
	return err
}

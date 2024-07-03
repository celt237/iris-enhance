package internal

import (
	"fmt"
	go_annotation "github.com/celt237/go-annotation"
	"strings"
)

type MethodCommentParser interface {
	Parse(annotation *go_annotation.Annotation, method *MethodDesc) error
}

func GetMethodCommentParser(annotation *go_annotation.Annotation) MethodCommentParser {
	switch annotation.Name {
	case ZRouterTag:
		return &RouterCommentParser{}
	case ZSummaryTag:
		return &SummaryCommentParser{}
	case ZDescriptionTag:
		return &DescriptionCommentParser{}
	case ZTagsTag:
		return &TagsCommentParser{}
	case ZAcceptTag:
		return &AcceptCommentParser{}
	case ZProduceTag:
		return &ProduceCommentParser{}
	//case ZParamTag:
	//	return &ParamCommentParser{}
	case ZReplyTypeTag:
		return &ReplyTypeCommentParser{}
	case ZReplyDataTag:
		return &ReplyDataCommentParser{}
	}
	// 如果commentSlice[0]以@x开头，那么就是自定义的tag
	if strings.HasPrefix(annotation.Name, "x") {
		return &CustomCommentParser{}
	}
	return &DefaultCommentParser{}
}

type RouterCommentParser struct{}

func (r *RouterCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZRouterTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 2 {
		method.Path = attribute["0"]
		method.Method = strings.TrimSuffix(strings.TrimPrefix(attribute["1"], "["), "]")
		return nil
	} else {
		return fmt.Errorf(ZRouterTag + " format is incorrect")
	}
}

type SummaryCommentParser struct{}

func (s *SummaryCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZSummaryTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.Summary = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZSummaryTag + " format is incorrect")
	}
}

type DescriptionCommentParser struct{}

func (d *DescriptionCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZDescriptionTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.Description = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZDescriptionTag + " format is incorrect")
	}
}

type TagsCommentParser struct{}

func (t *TagsCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZTagsTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.Tags = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZTagsTag + " format is incorrect")
	}
}

type AcceptCommentParser struct{}

func (a *AcceptCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZAcceptTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.Accept = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZAcceptTag + " format is incorrect")
	}
}

type ProduceCommentParser struct{}

func (p *ProduceCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZProduceTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.Produce = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZProduceTag + " format is incorrect")
	}
}

type ReplyTypeCommentParser struct{}

func (r *ReplyTypeCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZReplyTypeTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.ApiResultType = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZReplyTypeTag + " format is incorrect")
	}
}

type ReplyDataCommentParser struct{}

func (r *ReplyDataCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if len(annotation.Attributes) > 1 {
		return fmt.Errorf(ZReplyDataTag + " format is incorrect")
	}
	attribute := annotation.Attributes[0]
	if len(attribute) >= 1 {
		method.ApiResultDataType = attribute["0"]
		return nil
	} else {
		return fmt.Errorf(ZReplyDataTag + " format is incorrect")
	}
}

type CustomCommentParser struct{}

func (c *CustomCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	if method.CustomAnnotations == nil {
		method.CustomAnnotations = make(map[string]*go_annotation.Annotation)
	}
	method.CustomAnnotations[annotation.Name] = annotation
	return nil
}

type DefaultCommentParser struct{}

func (d *DefaultCommentParser) Parse(annotation *go_annotation.Annotation, method *MethodDesc) error {
	//fmt.Println("unknown comment tag:" + commentSlice[0])
	return nil
}

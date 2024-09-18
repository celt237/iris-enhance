package internal

import (
	go_annotation "github.com/celt237/go-annotation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMethodCommentParser(t *testing.T) {
	tests := []struct {
		name       string
		annotation *go_annotation.Annotation
		expected   MethodCommentParser
	}{
		{"RouterTag", &go_annotation.Annotation{Name: ZRouterTag}, &RouterCommentParser{}},
		{"SummaryTag", &go_annotation.Annotation{Name: ZSummaryTag}, &SummaryCommentParser{}},
		{"CustomTag", &go_annotation.Annotation{Name: "xcustom"}, &CustomCommentParser{}},
		{"UnknownTag", &go_annotation.Annotation{Name: "unknown"}, &DefaultCommentParser{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := GetMethodCommentParser(tt.annotation)
			assert.IsType(t, tt.expected, parser)
		})
	}
}

func TestRouterCommentParser_Parse(t *testing.T) {
	parser := &RouterCommentParser{}
	method := &MethodDesc{}

	tests := []struct {
		name           string
		annotation     *go_annotation.Annotation
		expectedErr    bool
		expectedPath   string
		expectedMethod string
	}{
		{
			name: "Valid Router",
			annotation: &go_annotation.Annotation{
				Name:       ZRouterTag,
				Attributes: []map[string]string{{"0": "/api/test", "1": "[GET]"}},
			},
			expectedErr:    false,
			expectedPath:   "/api/test",
			expectedMethod: "GET",
		},
		{
			name: "Invalid Router - Missing Method",
			annotation: &go_annotation.Annotation{
				Name:       ZRouterTag,
				Attributes: []map[string]string{{"0": "/api/test"}},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parser.Parse(tt.annotation, method)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPath, method.Path)
				assert.Equal(t, tt.expectedMethod, method.Method)
			}
		})
	}
}

// 可以为其他解析器添加类似的测试...

func TestCustomCommentParser_Parse(t *testing.T) {
	parser := &CustomCommentParser{}
	method := &MethodDesc{}

	annotation := &go_annotation.Annotation{
		Name:       "xcustom",
		Attributes: []map[string]string{{"0": "custom value"}},
	}

	err := parser.Parse(annotation, method)
	assert.NoError(t, err)
	assert.Contains(t, method.CustomAnnotations, "xcustom")
	assert.Equal(t, annotation, method.CustomAnnotations["xcustom"])
}

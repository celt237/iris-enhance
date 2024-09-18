package internal

import (
	go_annotation "github.com/celt237/go-annotation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethodParser_Parse(t *testing.T) {
	tests := []struct {
		name           string
		funcDecl       *go_annotation.MethodDesc
		resultType     string
		errorCode      string
		expectedMethod *MethodDesc
		expectedError  string
	}{
		{
			name: "Valid method",
			funcDecl: &go_annotation.MethodDesc{
				Name: "TestMethod",
				Annotations: map[string]*go_annotation.Annotation{
					ZRouterTag:    {Name: ZRouterTag, Attributes: []map[string]string{{"0": "/api/test", "1": "[GET]"}}},
					ZReplyTypeTag: {Name: ZReplyTypeTag, Attributes: []map[string]string{{"0": "TestReply"}}},
				},
				Results: []*go_annotation.Field{
					{
						Name:         "result",
						DataType:     "TestResult",
						PackageName:  "main",
						RealDataType: "TestResult",
						IsPtr:        false,
					},
					{
						Name:         "err",
						DataType:     "error",
						PackageName:  "",
						RealDataType: "error",
						IsPtr:        false,
					},
				},
			},
			resultType: "TestReply",
			errorCode:  "E1001",
			expectedMethod: &MethodDesc{
				MethodDesc: go_annotation.MethodDesc{
					Name: "TestMethod",
					Results: []*go_annotation.Field{
						{
							Name:         "result",
							DataType:     "TestResult",
							PackageName:  "main",
							RealDataType: "TestResult",
							IsPtr:        false,
						},
						{
							Name:         "err",
							DataType:     "error",
							PackageName:  "",
							RealDataType: "error",
							IsPtr:        false,
						},
					},
				},
				Path:              "/api/test",
				Method:            "GET",
				ApiResultType:     "TestReply",
				ApiResultDataType: "TestResult",
				ErrorCode:         "E1001",
				Produce:           "application/json",
				Accept:            "application/json",
				Result: &go_annotation.Field{
					Name:         "result",
					DataType:     "TestResult",
					PackageName:  "main",
					RealDataType: "TestResult",
					IsPtr:        false,
				},
			},
			expectedError: "",
		},
		{
			name: "Missing router tag",
			funcDecl: &go_annotation.MethodDesc{
				Name: "TestMethod",
				Annotations: map[string]*go_annotation.Annotation{
					ZReplyTypeTag: {Name: ZReplyTypeTag, Attributes: []map[string]string{{"0": "TestReply"}}},
				},
				Results: []*go_annotation.Field{
					{
						Name:         "result",
						DataType:     "TestResult",
						PackageName:  "main",
						RealDataType: "TestResult",
						IsPtr:        false,
					},
					{
						Name:         "err",
						DataType:     "error",
						PackageName:  "",
						RealDataType: "error",
						IsPtr:        false,
					},
				},
			},
			resultType:    "TestResultType",
			expectedError: "method:[TestMethod] no router tag found",
		},
		{
			name: "Invalid result count",
			funcDecl: &go_annotation.MethodDesc{
				Name: "TestMethod",
				Annotations: map[string]*go_annotation.Annotation{
					ZRouterTag: {Name: ZRouterTag, Attributes: []map[string]string{{"0": "/api/test", "1": "[GET]"}}},
				},
				Results: []*go_annotation.Field{
					{
						Name:         "result",
						DataType:     "TestResult",
						PackageName:  "main",
						RealDataType: "TestResult",
						IsPtr:        false,
					},
				},
			},
			resultType:    "TestResultType",
			expectedError: "method:[TestMethod] results length is not 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewMethodParser(tt.funcDecl)
			method, err := parser.Parse(tt.resultType, tt.errorCode)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedMethod.Path, method.Path)
				assert.Equal(t, tt.expectedMethod.Method, method.Method)
				assert.Equal(t, tt.expectedMethod.ApiResultType, method.ApiResultType)
				assert.Equal(t, tt.expectedMethod.ApiResultDataType, method.ApiResultDataType)
				assert.Equal(t, tt.expectedMethod.ErrorCode, method.ErrorCode)
				assert.Equal(t, tt.expectedMethod.Produce, method.Produce)
				assert.Equal(t, tt.expectedMethod.Accept, method.Accept)
			}
		})
	}
}

func TestMethodParser_parseResult(t *testing.T) {
	tests := []struct {
		name           string
		funcDecl       *go_annotation.MethodDesc
		expectedResult *go_annotation.Field
		expectedError  string
	}{
		{
			name: "Valid result",
			funcDecl: &go_annotation.MethodDesc{
				Name: "TestMethod",
				Results: []*go_annotation.Field{
					{
						Name:         "result",
						DataType:     "TestResult",
						PackageName:  "main",
						RealDataType: "TestResult",
						IsPtr:        false,
					},
					{
						Name:         "err",
						DataType:     "error",
						PackageName:  "",
						RealDataType: "error",
						IsPtr:        false,
					},
				},
			},
			expectedResult: &go_annotation.Field{
				Name:         "result",
				DataType:     "TestResult",
				PackageName:  "main",
				RealDataType: "TestResult",
				IsPtr:        false,
			},
			expectedError: "",
		},
		{
			name: "No result type found",
			funcDecl: &go_annotation.MethodDesc{
				Name: "TestMethod",
				Results: []*go_annotation.Field{
					{
						Name:         "err1",
						DataType:     "error",
						PackageName:  "",
						RealDataType: "error",
						IsPtr:        false,
					},
					{
						Name:         "err2",
						DataType:     "error",
						PackageName:  "",
						RealDataType: "error",
						IsPtr:        false,
					},
				},
			},
			expectedError: "no result type found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewMethodParser(tt.funcDecl)
			method := &MethodDesc{}
			err := parser.parseResult(method)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, method.Result)
			}
		})
	}
}

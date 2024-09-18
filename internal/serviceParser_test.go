package internal

import (
	go_annotation "github.com/celt237/go-annotation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceParser_Parse(t *testing.T) {
	tests := []struct {
		name            string
		genDecl         *go_annotation.StructDesc
		resultType      string
		errorCode       string
		imports         []string
		expectedService *ServiceDesc
		expectedError   string
	}{
		{
			name: "Valid service",
			genDecl: &go_annotation.StructDesc{
				Name: "TestService",
				Annotations: map[string]*go_annotation.Annotation{
					ZServiceTag: {},
					ZReplyTypeTag: {
						Attributes: []map[string]string{
							{"0": "CustomReply"},
						},
					},
				},
				Methods: []*go_annotation.MethodDesc{
					{
						Name: "TestMethod",
						Annotations: map[string]*go_annotation.Annotation{
							ZRouterTag: {
								Name: ZRouterTag,
								Attributes: []map[string]string{
									{"0": "/test", "1": "[GET]"},
								},
							},
						},
						Results: []*go_annotation.Field{
							{
								Name:         "result",
								DataType:     "CustomReply",
								PackageName:  "main",
								RealDataType: "CustomReply",
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
				},
				Imports: map[string]*go_annotation.ImportDesc{
					"custompackage": {
						Path:     "github.com/example/custompackage",
						HasAlias: false,
						Name:     "custompackage",
					},
				},
			},
			resultType: "DefaultReply",
			errorCode:  "E1001",
			imports:    []string{"github.com/example/custompackage"},
			expectedService: &ServiceDesc{
				Name:        "Test",
				PackageName: "handler",
				ReplyType:   "CustomReply",
				Methods: []*MethodDesc{
					{
						MethodDesc: go_annotation.MethodDesc{
							Name: "TestMethod",
						},
						Path:          "/test",
						Method:        "GET",
						ApiResultType: "CustomReply",
					},
				},
				Imports: map[string]*go_annotation.ImportDesc{
					"github.com/example/custompackage": {
						Path: "github.com/example/custompackage",
						Name: "custompackage",
					},
				},
				ExtraImports: []string{"github.com/example/custompackage"},
			},
			expectedError: "",
		},
		{
			name: "Invalid service name",
			genDecl: &go_annotation.StructDesc{
				Name: "TestInvalidName",
				Annotations: map[string]*go_annotation.Annotation{
					ZServiceTag: {},
				},
				Imports: map[string]*go_annotation.ImportDesc{},
			},
			expectedError: "",
		},
		{
			name: "Missing service annotation",
			genDecl: &go_annotation.StructDesc{
				Name:    "TestService",
				Imports: map[string]*go_annotation.ImportDesc{},
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewServiceParser(tt.genDecl)
			service, err := parser.Parse(tt.resultType, tt.errorCode, tt.imports)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else if tt.expectedService == nil {
				assert.Nil(t, service)
				assert.NoError(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedService.Name, service.Name)
				assert.Equal(t, tt.expectedService.PackageName, service.PackageName)
				assert.Equal(t, tt.expectedService.ReplyType, service.ReplyType)
				assert.Equal(t, len(tt.expectedService.Methods), len(service.Methods))
				for i, expectedMethod := range tt.expectedService.Methods {
					assert.Equal(t, expectedMethod.Name, service.Methods[i].Name)
					assert.Equal(t, expectedMethod.Path, service.Methods[i].Path)
					assert.Equal(t, expectedMethod.Method, service.Methods[i].Method)
					assert.Equal(t, expectedMethod.ApiResultType, service.Methods[i].ApiResultType)
				}
				assert.Contains(t, service.Imports, "custompackage")
			}
		})
	}
}

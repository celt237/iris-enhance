package internal

import (
	go_annotation "github.com/celt237/go-annotation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParamParser_Parse(t *testing.T) {
	tests := []struct {
		name           string
		funcDecl       *go_annotation.MethodDesc
		expectError    bool
		expectedLength int
	}{
		{
			name: "Valid function declaration",
			funcDecl: &go_annotation.MethodDesc{
				Params: []*go_annotation.Field{
					{
						Name:         "param1",
						DataType:     "string",
						PackageName:  "",
						RealDataType: "string",
						IsPtr:        false,
					},
					{
						Name:         "param2",
						DataType:     "int",
						PackageName:  "",
						RealDataType: "int",
						IsPtr:        false,
					},
				},
				Annotations: map[string]*go_annotation.Annotation{
					ZParamTag: {
						Name: ZParamTag,
						Attributes: []map[string]string{
							{"0": "param1", "1": "query", "2": "string", "3": "true", "4": "description1"},
							{"0": "param2", "1": "query", "2": "int", "3": "false", "4": "description2"},
						},
					},
				},
			},
			expectError:    false,
			expectedLength: 2,
		},
		{
			name: "Mismatch between method parameters and comment parameters",
			funcDecl: &go_annotation.MethodDesc{
				Params: []*go_annotation.Field{
					{
						Name:         "param1",
						DataType:     "string",
						PackageName:  "",
						RealDataType: "string",
						IsPtr:        false,
					},
					{
						Name:         "param2",
						DataType:     "int",
						PackageName:  "",
						RealDataType: "int",
						IsPtr:        false,
					},
				},
				Annotations: map[string]*go_annotation.Annotation{
					ZParamTag: {
						Name: ZParamTag,
						Attributes: []map[string]string{
							{"0": "param1", "1": "query", "2": "string", "3": "true", "4": "description1"},
						},
					},
				},
			},
			expectError:    true,
			expectedLength: 0,
		},
		{
			name: "Mismatch between method parameter type and comment parameter type",
			funcDecl: &go_annotation.MethodDesc{
				Params: []*go_annotation.Field{
					{
						Name:         "param1",
						DataType:     "string",
						PackageName:  "",
						RealDataType: "string",
						IsPtr:        false,
					},
				},
				Annotations: map[string]*go_annotation.Annotation{
					ZParamTag: {
						Name: ZParamTag,
						Attributes: []map[string]string{
							{"0": "param1", "1": "query", "2": "int", "3": "true", "4": "description1"},
						},
					},
				},
			},
			expectError:    true,
			expectedLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewParamParser(tt.funcDecl)
			params, err := parser.Parse()
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, params, tt.expectedLength)
			}
		})
	}
}

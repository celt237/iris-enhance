package internal

import (
	go_annotation "github.com/celt237/go-annotation"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 创建一个临时的配置文件
	tempFile, err := os.CreateTemp("", "config*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// 写入测试配置
	_, err = tempFile.WriteString(`
servicePath: ./services
handlePath: ./handlers
result: DefaultResult
errorCode: 1001
imports:
  - github.com/example/package1
  - github.com/example/package2
customAnnotations:
  - x-custom-annotation1
  - x-custom-annotation2
`)
	assert.NoError(t, err)
	tempFile.Close()

	tests := []struct {
		name        string
		configFile  string
		expectError bool
		expected    *Config
	}{
		{
			name:        "Valid config file",
			configFile:  tempFile.Name(),
			expectError: false,
			expected: &Config{
				ServicePath:       "./services",
				HandlePath:        "./handlers",
				ResultType:        "DefaultResult",
				ErrorCode:         1001,
				Imports:           []string{"github.com/example/package1", "github.com/example/package2"},
				CustomAnnotations: []string{"x-custom-annotation1", "x-custom-annotation2"},
			},
		},
		{
			name:        "Non-existent file",
			configFile:  "non_existent_file.yaml",
			expectError: true,
			expected:    nil,
		},
		{
			name:        "Empty file path",
			configFile:  "",
			expectError: true,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := LoadConfig(tt.configFile)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, config)
			}
		})
	}
}

func TestParseServiceDesc(t *testing.T) {
	tests := []struct {
		name           string
		fileDescList   []*go_annotation.FileDesc
		imports        []string
		resultType     string
		errorCode      string
		expectedLength int
		expectError    bool
	}{
		{
			name: "Valid service",
			fileDescList: []*go_annotation.FileDesc{
				{
					Structs: []*go_annotation.StructDesc{
						{
							Name: "TestService",
							Annotations: map[string]*go_annotation.Annotation{
								ZServiceTag: {},
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
											DataType:     "string",
											PackageName:  "",
											RealDataType: "string",
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
									Name:     "custompackage",
									HasAlias: false,
									Path:     "github.com/example/custompackage",
								},
							},
						},
					},
				},
			},
			imports:        []string{"github.com/example/package"},
			resultType:     "DefaultResult",
			errorCode:      "E1001",
			expectedLength: 1,
			expectError:    false,
		},
		{
			name:           "No valid services",
			fileDescList:   []*go_annotation.FileDesc{},
			imports:        []string{},
			resultType:     "DefaultResult",
			errorCode:      "E1001",
			expectedLength: 0,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceDescList, err := ParseServiceDesc(tt.fileDescList, tt.imports, tt.resultType, tt.errorCode)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, serviceDescList, tt.expectedLength)
			}
		})
	}
}

func TestGenerateCode(t *testing.T) {
	// 创建一个临时目录
	tempDir, err := os.MkdirTemp("", "testdir")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	serviceDesc := &ServiceDesc{
		Name:        "Test",
		PackageName: "handler",
		Methods:     []*MethodDesc{},
	}

	err = GenerateCode(tempDir, serviceDesc)
	assert.NoError(t, err)

	// 检查生成的文件是否存在
	_, err = os.Stat(tempDir + "/testHandler_gen.go")
	assert.NoError(t, err)
}

func TestTernary(t *testing.T) {
	assert.Equal(t, 1, Ternary(true, 1, 2))
	assert.Equal(t, 2, Ternary(false, 1, 2))
	assert.Equal(t, "yes", Ternary(true, "yes", "no"))
	assert.Equal(t, "no", Ternary(false, "yes", "no"))
}

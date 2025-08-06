package iris_enhance

import (
	"bytes"
	"github.com/celt237/iris-enhance/internal"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"mime/multipart"
	"testing"
)

func TestGetParamFromContext(t *testing.T) {
	// 创建测试用例
	tests := []struct {
		name      string
		setup     func() iris.Context
		paramName string
		dataType  string
		paramType string
		ptr       bool
		required  bool
		wantErr   bool
	}{
		{
			name: "测试路径参数",
			setup: func() iris.Context {
				app := iris.New()
				ctx := app.ContextPool.Acquire(httptest.NewRecorder(), httptest.NewRequest("GET", "/test/123", nil))
				ctx.Params().Set("id", "123")
				return ctx
			},
			paramName: "id",
			dataType:  "int",
			paramType: internal.ParamTypePath,
			ptr:       false,
			required:  true,
			wantErr:   false,
		},
		{
			name: "测试查询参数",
			setup: func() iris.Context {
				app := iris.New()
				ctx := app.ContextPool.Acquire(httptest.NewRecorder(), httptest.NewRequest("GET", "/test?name=test", nil))
				return ctx
			},
			paramName: "name",
			dataType:  "string",
			paramType: internal.ParamTypeQuery,
			ptr:       false,
			required:  true,
			wantErr:   false,
		},
		{
			name: "测试文件上传",
			setup: func() iris.Context {
				app := iris.New()
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				part, _ := writer.CreateFormFile("file", "test.txt")
				part.Write([]byte("test content"))
				writer.Close()

				req := httptest.NewRequest("POST", "/upload", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				ctx := app.ContextPool.Acquire(httptest.NewRecorder(), req)
				return ctx
			},
			paramName: "file",
			dataType:  "file",
			paramType: internal.ParamFormData,
			ptr:       false,
			required:  true,
			wantErr:   false,
		},
		{
			name: "测试普通表单字段",
			setup: func() iris.Context {
				app := iris.New()
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				writer.WriteField("age", "25")
				writer.Close()

				req := httptest.NewRequest("POST", "/form", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				ctx := app.ContextPool.Acquire(httptest.NewRecorder(), req)
				return ctx
			},
			paramName: "age",
			dataType:  "int",
			paramType: internal.ParamFormData,
			ptr:       false,
			required:  true,
			wantErr:   false,
		},
	}

	// 执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setup()

			switch tt.paramType {
			case internal.ParamTypePath:
				value, err := GetParamFromContext[int](ctx, tt.paramName, tt.dataType, tt.paramType, tt.ptr, tt.required)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetParamFromContext() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && value != 123 {
					t.Errorf("GetParamFromContext() = %v, want %v", value, 123)
				}

			case internal.ParamTypeQuery:
				value, err := GetParamFromContext[string](ctx, tt.paramName, tt.dataType, tt.paramType, tt.ptr, tt.required)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetParamFromContext() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err == nil && value != "test" {
					t.Errorf("GetParamFromContext() = %v, want %v", value, "test")
				}

			case internal.ParamFormData:
				if tt.paramName == "file" {
					value, err := GetParamFromContext[FileInfo](ctx, tt.paramName, tt.dataType, tt.paramType, tt.ptr, tt.required)
					if (err != nil) != tt.wantErr {
						t.Errorf("GetParamFromContext() error = %v, wantErr %v", err, tt.wantErr)
					}
					if err == nil && value == nil {
						t.Error("GetParamFromContext() returned nil FileInfo")
					}
				} else {
					value, err := GetParamFromContext[int](ctx, tt.paramName, tt.dataType, tt.paramType, tt.ptr, tt.required)
					if (err != nil) != tt.wantErr {
						t.Errorf("GetParamFromContext() error = %v, wantErr %v", err, tt.wantErr)
					}
					if err == nil && value != 25 {
						t.Errorf("GetParamFromContext() = %v, want %v", value, 25)
					}
				}
			}
		})
	}
}

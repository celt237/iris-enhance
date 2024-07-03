package internal

import go_annotation "github.com/celt237/go-annotation"

//type GenInfo struct {
//	Name        string                               // 名称
//	PackageName string                               // 包名
//	Imports     map[string]*go_annotation.ImportDesc // 导入包
//
//	Structs     *go_annotation.StructDesc            // 结构体
//}

type ServiceDesc struct {
	Name         string                               // 结构体名
	PackageName  string                               // 包名
	Imports      map[string]*go_annotation.ImportDesc // 导入信息
	ExtraImports []string                             // 额外导入包
	Comments     []string                             // 注释
	Annotations  map[string]*go_annotation.Annotation // 注解
	//Fields      []*Field               // 字段 暂不支持
	Methods     []*MethodDesc // 方法
	Description string        // 描述
	ReplyType   string        // 接口返回值类型
}

type MethodDesc struct {
	go_annotation.MethodDesc
	ErrorCode         string               // 错误编码
	Path              string               // url
	Method            string               // 全小写
	Params            []*MethodParam       // 参数
	Result            *go_annotation.Field // 返回值
	Summary           string
	Description       string
	Tags              string
	Accept            string
	Produce           string                               // 返回的数据格式 默认为json
	ApiResultType     string                               // 接口返回值类型 不含指针
	ApiResultDataType string                               // 接口返回值Data类型 不含指针
	CustomAnnotations map[string]*go_annotation.Annotation //自定义注解
}

type MethodParam struct {
	go_annotation.Field
	ParamType     string
	ParamDataType string
	Required      bool
	Desc          string
}

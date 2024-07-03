package internal

type Config struct {
	// service文件所在目录 必传
	ServicePath string `yaml:"servicePath"`

	// 要生成的控制器代码文件所在目录 必传
	HandlePath string `yaml:"handlePath"`

	// 接口返回结果的类型 必传
	ResultType string `yaml:"result"`

	// 默认的错误代码
	ErrorCode int `yaml:"errorCode"`

	// 需要额外导入的包
	Imports []string `yaml:"imports"`

	// 自定义属性，必须以x开头
	CustomAnnotations []string `yaml:"customAnnotations"`
}

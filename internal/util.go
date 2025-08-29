package internal

import (
	"bytes"
	"fmt"
	"github.com/celt237/go-annotation"
	"github.com/celt237/iris-enhance/internal/tmp"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
	"text/template"
	"unicode"
)

func LoadConfig(configFile string) (config *Config, err error) {
	// 加载配置文件
	if configFile != "" {
		yamlFile, err := os.ReadFile(configFile)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return nil, err
		}
		return config, nil
	} else {
		return nil, fmt.Errorf("config file is empty")
	}
}

func ParseServiceDesc(fileDescList []*go_annotation.FileDesc, imports []string, resultType string, errorCode string, basePath string) ([]*ServiceDesc, error) {
	serviceDescList := make([]*ServiceDesc, 0)
	for _, fileDesc := range fileDescList {
		for _, structData := range fileDesc.Structs {
			if structData == nil {
				continue
			}
			if !strings.HasSuffix(structData.Name, "Service") || len(structData.Annotations) == 0 {
				continue
			}
			if _, ok := structData.Annotations[ZServiceTag]; !ok {
				continue
			}
			serviceParser := NewServiceParser(structData)
			serviceDesc, err := serviceParser.Parse(resultType, errorCode, imports)
			if err != nil {
				return nil, err
			}
			serviceDesc.BasePath = basePath
			if serviceDesc != nil {
				serviceDescList = append(serviceDescList, serviceDesc)
			}
		}
	}
	return serviceDescList, nil
}

func GenerateCode(path string, serviceDesc *ServiceDesc) error {
	code := execute(serviceDesc)
	// 将生成的代码保存到文件
	// 判断是否存在handle文件夹
	if _, err := os.Stat(path); err != nil {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return err
		}
	}
	// 首字母小写
	fileName := strings.ToLower(serviceDesc.Name) + "Handler_gen.go"
	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		panic(err)
	}
	defer f.Close()
	// 清空文件内容
	err = f.Truncate(0)
	_, err = f.WriteString(code)
	fmt.Println("Generated code to", path+"/"+fileName)
	return nil
}

func execute(s *ServiceDesc) string {
	tmpText := tmp.Iris
	buf := new(bytes.Buffer)
	// 定义一个首字母大写的函数
	funcs := template.FuncMap{"capitalize": func(s string) string {
		if len(s) > 0 {
			return strings.ToUpper(s[:1]) + s[1:]
		}
		return ""
	}, "firstToLower": func(s string) string {
		if len(s) < 1 {
			return ""
		}
		r := []rune(s)
		r[0] = unicode.ToLower(r[0])
		return string(r)
	}}
	tmpl, err := template.New("iris").Funcs(funcs).Parse(strings.TrimSpace(tmpText))
	if err != nil {
		panic(err)
	}
	if err = tmpl.Execute(buf, s); err != nil {
		panic(err)
	}

	return strings.Trim(buf.String(), "\r\n")
}

// Ternary 三目运算
func Ternary[T any](a bool, b, c T) T {
	if a {
		return b
	}
	return c
}

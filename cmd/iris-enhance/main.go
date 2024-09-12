package main

import (
	"flag"
	"fmt"
	go_annotation "github.com/celt237/go-annotation"
	"github.com/celt237/iris-enhance/internal"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const version = "0.0.1"

func loadConfig() (config *internal.Config, err error) {
	// 配置文件
	configFile := flag.String("configFile", "", "The path to the configuration file")

	// service文件所在目录
	servicePath := flag.String("servicePath", "", "The path to the directory containing the service files")

	// handle文件所在目录
	handlePath := flag.String("handlePath", "", "The path to the directory containing the handle files")

	// 结果类型
	resultType := flag.String("result", "", "The type of the result")

	// 错误码 默认值为500
	errorCode := flag.String("errorCode", "", "The default error code")

	// 需要额外导入的包 多个以,分隔
	imports := flag.String("imports", "", "import packages")

	// 自定义注解 多个以,分隔
	customAnnotations := flag.String("customAnnotations", "", "custom annotations")

	versionFlag := flag.Bool("v", false, "print version number")
	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	config = &internal.Config{}
	// 加载配置文件
	if configFile != nil && *configFile != "" {
		config, err = internal.LoadConfig(*configFile)
		if err != nil {
			fmt.Println("Failed to load config file:", err)
			return nil, err
		}
	}
	if servicePath != nil && *servicePath != "" {
		config.ServicePath = *servicePath
	}
	if config.ServicePath == "" {
		return nil, fmt.Errorf("ServicePath is empty")
	}
	if handlePath != nil && *handlePath != "" {
		config.HandlePath = *handlePath
	}
	if config.HandlePath == "" {
		return nil, fmt.Errorf("HandlePath is empty")
	}
	if resultType != nil && *resultType != "" {
		config.ResultType = *resultType
	}
	if config.ResultType == "" {
		return nil, fmt.Errorf("ResultType is empty")
	}
	if errorCode != nil && *errorCode != "" {
		tempCode, err := strconv.Atoi(*errorCode)
		if err != nil {
			return nil, err
		}
		config.ErrorCode = tempCode
	}
	if config.ErrorCode == 0 {
		config.ErrorCode = 500
	}
	if imports != nil && *imports != "" {
		config.Imports = strings.Split(*imports, ",")
	}
	if customAnnotations != nil && *customAnnotations != "" {
		config.CustomAnnotations = strings.Split(*customAnnotations, ",")
	}
	return config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}
	// 遍历文件，找出以Service.go结尾的文件
	files, err := GetServiceFiles(config.ServicePath)
	if err != nil {
		fmt.Println("Failed to get service files:", err)
		return
	}
	if len(files) == 0 {
		fmt.Println("No service files found")
		return
	}
	fileDescList, err := go_annotation.GetFilesDescList(config.ServicePath, go_annotation.AnnotationModeArray)
	if err != nil {
		fmt.Println("Failed to get file desc list:", err)
		return
	}

	serviceDescList, err := internal.ParseServiceDesc(fileDescList, config.Imports, config.ResultType, strconv.Itoa(config.ErrorCode))
	if err != nil {
		fmt.Println("Failed to parse service desc:", err)
		return
	}
	for _, serviceDesc := range serviceDescList {

		err = internal.GenerateCode(config.HandlePath, serviceDesc)
		if err != nil {
			fmt.Println("Failed to generate code:", err)
			return
		}
	}
}

func GetServiceFiles(path string) (files []string, err error) {
	files = make([]string, 0)
	// 读取指定目录下的所有文件
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Failed to read directory:", err)
		return files, err
	}
	// 遍历文件，找出以.go结尾的文件
	for _, entry := range entries {
		// entry.Name() 以.go结尾
		fmt.Println("Found file:", entry.Name())

		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			files = append(files, filepath.Join(path, entry.Name()))
		}
	}
	return files, err
}

package iris_enhance

import (
	"bytes"
	"embed"
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	//go:embed static
	front   embed.FS
	docJson []byte
	s       service
)

type Config struct {
	RelativePath string
}

type service struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	SwaggerVersion string `json:"swaggerVersion"`
	Location       string `json:"location"`
}

func init() {
	var err error
	docJson, err = os.ReadFile("./docs/swagger.json")
	if err != nil {
		log.Println("no swagger.json found in ./docs")
	}
}

func readFileFromEmbedFS(fs embed.FS, filename string) (string, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// RegisterSwaggerDoc registers swagger documentation
func RegisterSwaggerDoc(app *iris.Application, path string) {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	path = strings.TrimSpace(path)
	if path == "" {
		path = "doc"
	}
	path = "/" + path
	app.Get(path+"/{any:path}", swagDocHandler(Config{RelativePath: path}))
}

// swagDocHandler is a handler for swagger documentation
func swagDocHandler(config Config) iris.Handler {
	docJsonPath := config.RelativePath + "/docJson"
	servicesPath := config.RelativePath + "/static/service"
	docPath := config.RelativePath + "/index"
	appjsPath := config.RelativePath + "/static/webjars/js/app.42aa019b.js"

	s.Url = "/docJson"
	s.Location = "/docJson"
	s.Name = "API Documentation"
	s.SwaggerVersion = "2.0"

	appjs, err := readFileFromEmbedFS(front, "static/webjars/js/app.42aa019b.js")
	if err != nil {
		log.Println(err)
	}
	appjs = strings.ReplaceAll(appjs, "{[(.RelativePath)]}", config.RelativePath)

	doc, err := readFileFromEmbedFS(front, "static/doc.html")
	if err != nil {
		log.Println(err)
	}
	doc = strings.ReplaceAll(doc, "{[(.RelativePath)]}", config.RelativePath)

	return func(ctx iris.Context) {
		if ctx.Request().Method != http.MethodGet {
			ctx.StatusCode(http.StatusMethodNotAllowed)
			ctx.StopExecution()
			return
		}
		switch ctx.Request().RequestURI {
		case appjsPath:
			ctx.WriteString(appjs)
		case servicesPath:
			ctx.JSON([]service{s})
		case docPath:
			ctx.HTML(doc)
		case docJsonPath:
			ctx.ContentType("application/json")
			ctx.Write(docJson)
		default:
			filePath := strings.TrimPrefix(ctx.Request().RequestURI, config.RelativePath)
			filePath = strings.TrimPrefix(filePath, "/")
			file, err := front.Open(filePath)
			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef("Error while opening file: %v", err)
				return
			}
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef("Error while reading file: %v", err)
				return
			}
			reader := bytes.NewReader(data)
			fileInfo, err := file.Stat()
			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef("Error while getting file info: %v", err)
				return
			}
			ctx.ServeContent(reader, fileInfo.Name(), fileInfo.ModTime())
		}
		ctx.Next()
	}

}

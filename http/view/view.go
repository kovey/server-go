package view

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/kovey/logger-go/logger"
	"github.com/kovey/server-go/util"
)

type View struct {
	path   string
	tpl    *template.Template
	data   map[string]interface{}
	writer http.ResponseWriter
}

func NewView(path string, writer http.ResponseWriter) *View {
	return &View{path: path, data: make(map[string]interface{}), writer: writer}
}

func (v *View) Set(field string, value interface{}) *View {
	v.data[field] = value
	return v
}

func (v *View) Load() error {
	logger.Debug("view load begin, path[%s]", v.path)
	if !util.IsFile(v.path) {
		return fmt.Errorf("path[%s] is not exists", v.path)
	}

	v.tpl = template.Must(template.New(path.Base(v.path)).Funcs(template.FuncMap{
		"add": Add, "sub": Sub, "mul": Mul, "div": Div, "map_get": MapGet,
	}).ParseFiles(v.path))

	logger.Debug("view load end, path[%s]", v.path)
	return nil
}

func (v *View) Render() error {
	return v.tpl.Execute(v.writer, v.data)
}

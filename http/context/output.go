package context

import (
	"encoding/json"
	"net/http"

	"github.com/kovey/logger-go/logger"
)

type Output struct {
	response http.ResponseWriter
	data     interface{}
	buf      []byte
}

func NewOutput(response http.ResponseWriter) *Output {
	return &Output{response: response}
}

func (o *Output) ToJson() {
	buf, err := json.Marshal(o.data)
	if err != nil {
		logger.Error("marsha data fail: %s", err)
		return
	}

	o.SetHeader("Content-Type", "application/json")
	o.buf = buf
}

func (o *Output) ToHtml() {
	o.ToString()
	o.SetHeader("Content-Type", "text/html")
}

func (o *Output) ToString() {
	data, ok := o.data.(string)
	if !ok {
		logger.Error("data is not string")
		return
	}

	o.buf = []byte(data)
}

func (o *Output) ToXml() {
	o.ToString()
	o.SetHeader("Content-Type", "text/xml")
}

func (o *Output) SetHeader(key string, value string) *Output {
	o.response.Header().Set(key, value)
	return o
}

func (o *Output) SetStatus(code int) *Output {
	o.response.WriteHeader(code)
	return o
}

func (o *Output) End() {
	_, err := o.response.Write(o.buf)
	if err != nil {
		logger.Error("response end fail: %s", err)
	}
}

func (o *Output) SetData(data interface{}) *Output {
	o.data = data
	return o
}

func (o *Output) PageNotFound() {
	o.data = "<p>kovey</><p>Page Not Found</p>"
	o.ToHtml()
	o.SetStatus(404)
}

func (o *Output) InternalError() {
	o.data = "<p>kovey</><p>Internal Error</p>"
	o.ToHtml()
	o.SetStatus(500)
}

func (o *Output) Redirect(url string) {
	o.SetHeader("Location", url)
	o.SetHeader("Content-Type", "text/html")
	o.SetStatus(302)
}

func (o *Output) Response() http.ResponseWriter {
	return o.response
}

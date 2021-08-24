package context

import (
	"encoding/json"
	"fmt"
	"github.com/kovey/logger-go/logger"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type Input struct {
	request *http.Request
	data map[string]interface{}
}

func NewInput(request *http.Request) *Input {
	input := &Input{request: request, data: make(map[string]interface{})}
	input.parseData()

	return input
}

func (i *Input) parseData() {
	if !i.isJson() {
		return
	}
	
	err := i.Json(&i.data)
	if err != nil {
		logger.Error("parse data to json fail, error: %s", err)
	}
}

func (i *Input) Method() string {
	return i.request.Method
}

func (i *Input) BaseUrl() string {
	return fmt.Sprintf("%s://%s", i.request.URL.Scheme, i.request.URL.Host)
}

func (i *Input) Path() string {
	if len(i.request.URL.Path) == 0 {
		return "/"
	}

	return i.request.URL.Path
}

func (i *Input) ClientIp() string {
	return i.request.RemoteAddr
}

func (i *Input) Uri() string {
	return i.request.RequestURI
}

func (i *Input) GetHeader(key string) string {
	return i.request.Header.Get(key)
}

func (i *Input) Query(key string) string {
	return i.request.FormValue(key)
}

func (i *Input) Post(key string) string {
	return i.request.FormValue(key)
}

func (i *Input) File(key string) (multipart.File, *multipart.FileHeader, error) {
	return i.request.FormFile(key)
}

func (i *Input) GetInt(key string) int {
	data, err := strconv.Atoi(i.Query(key))
	if err != nil {
		return 0
	}

	return data
}

func (i *Input) GetInt64(key string) int64 {
	data, err := strconv.ParseInt(i.Query(key), 10, 64)
	if err != nil {
		return 0
	}

	return data
}

func (i *Input) GetFloat64(key string) float64 {
	data, err := strconv.ParseFloat(i.Query(key), 64)
	if err != nil {
		return 0.0
	}

	return data
}

func (i *Input) GetBool(key string) bool {
	data, err := strconv.ParseBool(i.Query(key))
	if err != nil {
		return false
	}

	return data
}

func (i *Input) Is(method string) bool {
	return strings.ToLower(i.Method()) == strings.ToLower(method)
}

func (i *Input) QueryData() map[string][]string {
	i.request.ParseForm()

	return map[string][]string(i.request.Form)
}

func (i *Input) PostData() map[string][]string {
	if !i.IsMultipartForm() {
		i.request.ParseForm()
		return map[string][]string(i.request.PostForm)
	}

	i.request.ParseMultipartForm(2097152)
	return map[string][]string(i.request.MultipartForm.Value)
}

func (i *Input) isJson() bool {
	logger.Debug("content-type: %s", i.GetHeader("Content-Type"))
	return strings.ToLower(i.GetHeader("Content-Type")) == "application/json";
}

func (i *Input) Data() interface{} {
	if i.isJson() {
		return i.data
	}

	if i.Is("get") {
		return i.QueryData()
	}

	return i.PostData()
}

func (i *Input) IsMultipartForm() bool {
	return i.GetHeader("Content-Type") == "multipart/form-data"
}

func (i *Input) Headers() map[string][]string {
	return map[string][]string(i.request.Header)
}

func (i *Input) Json(data interface{}) error {
	return json.NewDecoder(i.request.Body).Decode(data)
}

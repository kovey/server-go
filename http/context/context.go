package context

import (
	"net/http"
)

type Context struct {
	intput       *Input
	output       *Output
	traceId      string
	spanId       string
	viewDisabled bool
}

func NewContext(request *http.Request, response http.ResponseWriter, traceId string, spanId string) *Context {
	return &Context{intput: NewInput(request), output: NewOutput(response), traceId: traceId, spanId: spanId, viewDisabled: false}
}

func (c *Context) Input() *Input {
	return c.intput
}

func (c *Context) Output() *Output {
	return c.output
}

func (c *Context) TraceId() string {
	return c.traceId
}

func (c *Context) SpanId() string {
	return c.spanId
}

func (c *Context) SetViewSwitch(viewDisabled bool) {
	c.viewDisabled = viewDisabled
}

func (c *Context) ViewDisabled() bool {
	return c.viewDisabled
}

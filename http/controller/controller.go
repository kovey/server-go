package controller

import (
	"github.com/kovey/logger-go/logger"
	"github.com/kovey/server-go/http/context"
	"github.com/kovey/server-go/http/view"
)

type ControllerInterface interface {
	SetCtx(*context.Context)
	Input() *context.Input
	Output() *context.Output
	TraceId() string
	SpanId() string
	Render() error
	Ctx() *context.Context
	View() *view.View
	Query(string) string
	Post(string) string
}

type Controller struct {
	ctx  *context.Context
	view *view.View
}

func NewController(ctx *context.Context, viewPath string) Controller {
	if len(viewPath) != 0 {
		return Controller{ctx: ctx, view: view.NewView(viewPath, ctx.Output().Response())}
	}

	return Controller{ctx: ctx}
}

func (c Controller) SetCtx(ctx *context.Context) {
	c.ctx = ctx
}

func (c Controller) Input() *context.Input {
	return c.ctx.Input()
}

func (c Controller) Output() *context.Output {
	return c.ctx.Output()
}

func (c Controller) TraceId() string {
	return c.ctx.TraceId()
}

func (c Controller) SpanId() string {
	return c.ctx.SpanId()
}

func (c Controller) Render() error {
	if c.ctx.ViewDisabled() {
		return nil
	}

	err := c.view.Load()
	if err != nil {
		logger.Error("view load failure, error: %s", err)
		return err
	}

	return c.view.Render()
}

func (c Controller) Ctx() *context.Context {
	return c.ctx
}

func (c Controller) View() *view.View {
	return c.view
}

func (c Controller) Query(field string) string {
	return c.Input().Query(field)
}

func (c Controller) Post(field string) string {
	return c.Input().Post(field)
}

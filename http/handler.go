package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kovey/logger-go/logger"
	"github.com/kovey/logger-go/monitor"
	"github.com/kovey/server-go/http/context"
	"github.com/kovey/server-go/http/router"
	"github.com/kovey/server-go/util"
)

type Handler struct {
	routers *router.Routers
}

func NewHandler() *Handler {
	return &Handler{routers: router.NewRouters()}
}

func (h *Handler) Routers() *router.Routers {
	return h.routers
}

func (h *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := context.NewContext(request, response, util.TraceId(), util.SpanId())
	monLog := util.GetMonitor(ctx.Input(), ctx.TraceId(), ctx.SpanId())

	defer func(monLog *monitor.Monitor, output *context.Output) {
		err := recover()
		monLog.Trace = util.GetTrace(err)
		monLog.End = time.Now().UnixNano() / 1e6
		monLog.Delay = float64(monLog.End-monLog.RequestTime) / 1e6

		if err == nil {
			monLog.Err = fmt.Sprintf("%s", err)
		} else {
			output.InternalError()
			output.End()
		}

		monitor.Write(*monLog)
	}(monLog, ctx.Output())

	r := h.routers.Router(ctx.Input().Path(), ctx.Input())
	if r == nil {
		ctx.Output().PageNotFound()
		ctx.Output().End()
		return
	}

	ctx.SetViewSwitch(r.ViewDisabled())

	err := r.Call(ctx)
	if err != nil {
		logger.Error("call error: %s", err)
		ctx.Output().End()
		return
	}

	if r.ViewDisabled() {
		ctx.Output().End()
	}
}

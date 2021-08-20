package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/kovey/logger-go/logger"
	"github.com/kovey/logger-go/monitor"
	"github.com/kovey/server-go/http/context"
)

func GetMonitor(input *context.Input, traceId string, spanId string) *monitor.Monitor {
	monLog := &monitor.Monitor{}
	monLog.Path = input.Path()
	monLog.Params = input.Data()
	monLog.RequestTime = time.Now().UnixNano() / 1e6
	monLog.ServiceType = "http"
	monLog.Args = input.Data()
	monLog.Ip = input.ClientIp()
	monLog.Time = time.Now().Unix()
	monLog.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	minute, err := strconv.ParseInt(time.Now().Format("200601021504"), 10, 64)
	if err != nil {
		monLog.Minute = minute
	}
	monLog.HttpCode = 200
	monLog.TraceId = traceId
	monLog.SpanId = spanId
	monLog.ParentId = ""
	monLog.ClientVersion = "1.0.0"
	monLog.ServerVersion = "1.0.0"
	monLog.ClientLang = "JS"
	monLog.ServerLang = "golang"
	monLog.From = ""

	return monLog
}

func GetTrace(err interface{}) string {
	if err == nil {
		logger.Debug("err is nil")
		return ""
	}

	logger.Error("panic error[%s]", err)

	traces := make([]string, 1)
	traces[0] = fmt.Sprintf("panic error[%s]", err)

	for i := 3; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		traces = append(traces, fmt.Sprintf("%s(%d)", file, line))
		logger.Error("%s(%d)", file, line)
	}

	return strings.Join(traces, "#")
}

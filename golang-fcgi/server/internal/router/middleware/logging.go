package middleware

import (
	"expvar"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gibriil/enterprise_portal_example/internal"
)

type responseLogger struct {
	http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func newLoggingResponseWriter(res http.ResponseWriter) *responseLogger {
	return &responseLogger{
		ResponseWriter: res,
		statusCode:     http.StatusTeapot,
	}
}

func (lw *responseLogger) WriteHeader(statusCode int) {
	lw.ResponseWriter.WriteHeader(statusCode)

	if !lw.headerWritten {
		lw.statusCode = statusCode
		lw.headerWritten = true
	}
}
func (lw *responseLogger) Write(b []byte) (int, error) {
	if !lw.headerWritten {
		lw.statusCode = http.StatusOK
	}
	lw.headerWritten = true
	return lw.ResponseWriter.Write(b)
}
func (lw *responseLogger) Unwrap() http.ResponseWriter { return lw.ResponseWriter }

func RequestLogging(app *internal.Application, next http.Handler) http.Handler {
	var (
		totalRequestsReceived           = expvar.NewInt("total_request_received")
		totalResponsesSent              = expvar.NewInt("total_responses_sent")
		totalProcessingTimeMicroseconds = expvar.NewInt("total_processing_time_μs")
		totalResponsesSentByStatus      = expvar.NewMap("total_responses_sent_by_status")
	)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()

		resLogger := newLoggingResponseWriter(res)

		totalRequestsReceived.Add(1)

		next.ServeHTTP(resLogger, req)

		totalResponsesSent.Add(1)

		statusCode := strconv.Itoa(resLogger.statusCode)

		totalResponsesSentByStatus.Add(statusCode, 1)

		app.Log.Info(fmt.Sprintf("%s Request Received", req.Method),
			slog.String("method", req.Method),
			slog.String("path", req.URL.Path),
			slog.String("protocol", req.Proto),
			slog.String("status", statusCode),
			slog.String("client_ip", req.RemoteAddr),
			slog.String("user_agent", req.UserAgent()),
		)

		duration := time.Since(start).Microseconds()
		totalProcessingTimeMicroseconds.Add(duration)
	})
}

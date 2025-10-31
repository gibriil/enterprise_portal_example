package helpers

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes a log entry at Error level (including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func ServerError(logger *slog.Logger, res http.ResponseWriter, req http.Request, err error) {
	var (
		method = req.Method
		uri    = req.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri), slog.String("stacktrace", trace))
	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user.
func ClientError(res http.ResponseWriter, status int) {
	http.Error(res, http.StatusText(status), status)
}

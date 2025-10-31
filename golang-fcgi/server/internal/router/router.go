package router

import (
	"expvar"
	"fmt"
	"net/http"
	"net/http/fcgi"

	"github.com/gibriil/enterprise_portal_example/internal"
	"github.com/gibriil/enterprise_portal_example/internal/router/middleware"
)

func CreateRouter(app *internal.Application) http.Handler {
	server := http.NewServeMux()

	// documentation := openapi.Server(app.Log)

	server.HandleFunc("GET /test.go", func(res http.ResponseWriter, req *http.Request) {
		_SERVER := fcgi.ProcessEnv(req)
		fmt.Println(_SERVER)
	})

	// server.HandleFunc("/healthz", RouteHandler.HealthCheck(app))
	server.Handle("/server-status", expvar.Handler())

	return createMiddlewareStack(app, server)
}

func createMiddlewareStack(app *internal.Application, server http.Handler) http.Handler {
	return middleware.PanicRecovery(app,
		middleware.RequestLogging(app,
			middleware.HandleBaseUrl(app,
				middleware.ApplyBaseHeaders(app,
					middleware.AuthenticatedUserContext(app, server)))))
}

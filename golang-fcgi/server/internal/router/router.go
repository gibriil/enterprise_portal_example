package router

import (
	"expvar"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/gibriil/enterprise_portal_example/internal"
	"github.com/gibriil/enterprise_portal_example/internal/helpers"
	"github.com/gibriil/enterprise_portal_example/internal/html"
	"github.com/gibriil/enterprise_portal_example/internal/models"
	"github.com/gibriil/enterprise_portal_example/internal/router/middleware"
)

func CreateRouter(app *internal.Application) http.Handler {
	server := http.NewServeMux()

	// documentation := openapi.Server(app.Log)

	server.HandleFunc("GET /header.go", func(res http.ResponseWriter, req *http.Request) {

		template, err := template.ParseFS(html.WrapperTemplates, "header.tmpl")
		if err != nil {
			helpers.ServerError(app.Log, res, *req, err)
		}

		user := app.CurrentReqContext.Value(internal.UserContext("user")).(models.User)

		template.Execute(io.MultiWriter(os.Stdout, res), user)
		// res.Write([]byte(`<br />`))
		// res.Write([]byte("<!--#include virtual=\"/_resources/include.inc\"-->"))
		// res.Write([]byte(`<br />`))
		// res.Write([]byte(fmt.Sprintf("Server Protocol: %s", user.Name)))
	})

	server.HandleFunc("GET /footer.go", func(res http.ResponseWriter, req *http.Request) {

		template, err := template.ParseFS(html.WrapperTemplates, "footer.tmpl")
		if err != nil {
			helpers.ServerError(app.Log, res, *req, err)
		}

		user := app.CurrentReqContext.Value(internal.UserContext("user")).(models.User)

		template.Execute(res, user)
	})

	server.HandleFunc("GET /user", func(res http.ResponseWriter, req *http.Request) {
		user := app.CurrentReqContext.Value(internal.UserContext("user")).(models.User)

		if authenticated := user.IsAuthenticated(); !authenticated {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte(user.Error()))
			return
		}

		res.Write(user.ToJson())
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

package middleware

import (
	"context"
	"net/http"

	"github.com/gibriil/enterprise_portal_example/internal"
	"github.com/gibriil/enterprise_portal_example/internal/helpers"
	"github.com/gibriil/enterprise_portal_example/internal/models"
)

func AuthenticatedUserContext(app *internal.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		var user models.User

		err := user.LoadClaims(req)
		if err != nil {
			helpers.ServerError(app.Log, res, *req, err)
		}

		contextKeys := make(map[internal.UserContext]any)

		if len(user.Roles) > 0 {
			contextKeys[internal.UserContext("user")] = user
		} else {
			contextKeys[internal.UserContext("user")] = nil
		}

		for key, value := range contextKeys {
			ctx = context.WithValue(ctx, key, value)
		}

		app.CurrentReqContext = ctx

		next.ServeHTTP(res, req.WithContext(ctx))

	})
}

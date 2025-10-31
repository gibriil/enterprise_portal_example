package helpers

import (
	"github.com/gibriil/enterprise_portal_example/internal"
	"github.com/gibriil/enterprise_portal_example/internal/models"
)

func GetAuthenticatedUser(app *internal.Application) (any, bool) {
	user, authenticated := app.CurrentReqContext.Value(internal.UserContext("user")).(models.User)

	// if !authenticated {
	// 	return models.User{}, errors.New("401: Unauthorized (Not Authenticated) - NO OIDC Information found")
	// }

	return user, authenticated
}

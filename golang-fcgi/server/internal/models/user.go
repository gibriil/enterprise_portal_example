package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/fcgi"
	"strings"
)

type User struct {
	Name              string            `json:"name"`
	GivenName         string            `json:"given_name"`
	FamilyName        string            `json:"family_name"`
	PreferredUsername string            `json:"preferred_username"`
	Id                string            `json:"id"`
	Email             string            `json:"email"`
	Roles             []string          `json:"roles"`
	Groups            []string          `json:"groups"`
	Claims            map[string]string `json:"claims"`
	Authenticated     bool              `json:"-"`
}

func (user *User) isAuthenticated() bool {
	return len(user.Claims) > 0
}

func (user *User) Error() string {
	return fmt.Sprintf("Error %d: %s", http.StatusUnauthorized, "portal user could not be created - check authentication claims/entitlements")
}

func (user *User) LoadClaims(req *http.Request) error {
	_SERVER := fcgi.ProcessEnv(req)

	user.Name = _SERVER["OIDC_CLAIM_name"]
	user.GivenName = _SERVER["OIDC_CLAIM_given_name"]
	user.FamilyName = _SERVER["OIDC_CLAIM_family_name"]
	user.PreferredUsername = _SERVER["OIDC_CLAIM_preferred_username"]
	user.Email = _SERVER["OIDC_CLAIM_email"]

	user.Claims = map[string]string{}

	for claim, value := range _SERVER {
		if strings.HasPrefix(claim, "OIDC_CLAIM_") {
			user.Claims[claim] = value
		}
	}

	user.Groups = strings.Split(_SERVER["OIDC_CLAIM_groups"], ",")

	user.Roles = strings.Split(_SERVER["OIDC_CLAIM_realm_roles"], ",")

	user.Roles = append(user.Roles, strings.Split(_SERVER["OIDC_CLAIM_client_roles"], ",")...)

	user.Authenticated = user.isAuthenticated()

	return nil
}

func (user *User) ToJson() []byte {
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return data
}

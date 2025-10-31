package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/fcgi"
)

type User struct {
	Name              string            `json:"name"`
	GivenName         string            `json:"given_name"`
	FamilyName        string            `json:"family_name"`
	PreferredUsername string            `json:"preferred_username"`
	Id                string            `json:"id"`
	Email             string            `json:"email"`
	Roles             []string          `json:"roles"`
	Claims            map[string]string `json:"claims"`
}

func (user *User) Error() string {
	return fmt.Sprintf("Error %d: %s", http.StatusUnauthorized, "myUVU user could not be created - check authentication claims/entitlements")
}

func (user *User) LoadClaims(req *http.Request) error {
	_SERVER := fcgi.ProcessEnv(req)

	fmt.Print(_SERVER)
	return nil
}

func (user *User) ToJson() []byte {
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return data
}

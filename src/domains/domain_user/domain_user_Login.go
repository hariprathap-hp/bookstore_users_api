package domain_user

import (
	"strings"

	"github.com/hariprathap-hp/bookstore_users_api/src/utils/errors"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *LoginUser) Validate() *errors.RestErr {
	if strings.TrimSpace(strings.ToLower(user.Email)) == "" {
		return errors.NewHTTPBadRequestError("email can not be empty")
	}
	if strings.TrimSpace(strings.ToLower(user.Password)) == "" {
		return errors.NewHTTPBadRequestError("password can not be empty")
	}
	return nil
}

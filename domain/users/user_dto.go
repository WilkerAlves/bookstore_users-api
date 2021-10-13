package users

import (
	"github.com/WilkerAlves/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id" binding:"gte=1"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadResquestError("invalid email address")
	}
	return nil
}

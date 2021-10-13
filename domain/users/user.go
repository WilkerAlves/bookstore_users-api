package users

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name" validate:"required,min=4,max=15"`
	LastName    string `json:"last_name" validate:"required,min=4,max=15"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"date_created"`
}

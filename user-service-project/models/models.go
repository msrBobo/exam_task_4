package models

type CreateUser struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Bio         string `json:"bio"`
	Website     string `json:"website"`
}

package entity

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Password string `json:"encrypted_password"`
}

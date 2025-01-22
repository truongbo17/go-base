package response

import "time"

type UserRegisterResponse struct {
	Token Token    `json:"token"`
	User  UserInfo `json:"user"`
}

type UserInfo struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

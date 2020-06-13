package presenter

import (
	"dast-api/internal/domain/model"
	"fmt"
)

type User struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type LogResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func RenderUser(user *model.User) *User {
	return &User{ID: user.ID, Name: user.Name, Email: user.Email}
}

func RenderSuccessLogIn(name string, token string) *LogResponse {
	return &LogResponse{Username: name, Message: fmt.Sprintf("%s successful logged in", name), Token: token}
}

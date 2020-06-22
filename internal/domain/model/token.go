package model

type LogIn struct {
	UserID string
	Token  string
}

func NewLogIn(ID string, token string) *LogIn {
	return &LogIn{UserID: ID, Token: token}
}

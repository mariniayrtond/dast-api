package model

type LogIn struct {
	ID    string
	Token string
}

func NewLogIn(ID string, token string) *LogIn {
	return &LogIn{ID: ID, Token: token}
}

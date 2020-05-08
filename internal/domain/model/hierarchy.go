package model

type Hierarchy struct {
	ID           string
	Name         string
	Description  string
	Owner        string
	Alternatives []string
	Criteria     []Criteria
}

func NewHierarchy(ID string, name string, description string, owner string, alternatives []string) *Hierarchy {
	return &Hierarchy{ID: ID, Name: name, Description: description, Owner: owner, Alternatives: alternatives, Criteria: []Criteria{}}
}

type Criteria struct {
	Level  int
	ID     string
	Name   string
	Parent string
	Score  Score
}

type Score struct {
	Local  float64
	Global float64
}

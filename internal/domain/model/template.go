package model

type CriteriaTemplate struct {
	ID          string
	Owner       string
	Description string
	Criteria    []Criteria
}

func NewCriteriaTemplate(ID string, description string, owner string, criteria []Criteria) *CriteriaTemplate {
	return &CriteriaTemplate{ID: ID, Description: description, Owner: owner, Criteria: criteria}
}

package model

type CriteriaTemplate struct {
	ID          string
	Owner       string
	Description string
	Public      bool
	Criteria    []Criteria
}

func NewCriteriaTemplate(ID string, description string, owner string, public bool, criteria []Criteria) *CriteriaTemplate {
	return &CriteriaTemplate{ID: ID, Description: description, Owner: owner, Public: public, Criteria: criteria}
}

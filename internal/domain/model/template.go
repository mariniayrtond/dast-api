package model

type CriteriaTemplate struct {
	ID       string
	Owner    string
	Public   bool
	Criteria []Criteria
}

func NewCriteriaTemplate(ID string, owner string, public bool, criteria []Criteria) *CriteriaTemplate {
	return &CriteriaTemplate{ID: ID, Owner: owner, Public: public, Criteria: criteria}
}

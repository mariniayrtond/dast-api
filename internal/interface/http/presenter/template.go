package presenter

import "dast-api/internal/domain/model"

type CriteriaTemplate struct {
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Criteria    []struct {
		Level       int    `json:"level"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Parent      string `json:"parent"`
	} `json:"criteria"`
}

func RenderCriteriaTemplates(templates []*model.CriteriaTemplate) []CriteriaTemplate {
	res := []CriteriaTemplate{}
	for _, template := range templates {
		res = append(res, RenderCriteriaTemplate(template))
	}
	return res
}

func RenderCriteriaTemplate(t *model.CriteriaTemplate) CriteriaTemplate {
	res := CriteriaTemplate{
		ID:          t.ID,
		Public:      t.Public,
		Owner:       t.Owner,
		Description: t.Description,
		Criteria: []struct {
			Level       int    `json:"level"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Parent      string `json:"parent"`
		}{},
	}

	for _, c := range t.Criteria {
		res.Criteria = append(res.Criteria, struct {
			Level       int    `json:"level"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Parent      string `json:"parent"`
		}{
			Level:       c.Level,
			ID:          c.ID,
			Description: c.Name,
			Parent:      c.Parent,
		})
	}

	return res
}

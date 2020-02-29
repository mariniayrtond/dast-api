package presenter

import "dast-api/internal/domain/model"

type hierarchyResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Owner        string   `json:"owner"`
	Objective    string   `json:"objective"`
	Alternatives []string `json:"alternatives"`
	Criteria     []struct {
		Level  int    `json:"level"`
		ID     string `json:"id"`
		Name   string `json:"name"`
		Parent string `json:"parent"`
		Score  struct {
			Local  float64 `json:"local"`
			Global float64 `json:"global"`
		} `json:"score"`
	} `json:"criteria"`
}

func RenderHierarchy(h *model.Hierarchy) hierarchyResponse {
	res := hierarchyResponse{
		ID:           h.ID,
		Name:         h.Name,
		Description:  h.Description,
		Owner:        h.Owner,
		Objective:    h.Objective,
		Alternatives: h.Alternatives,
		Criteria: []struct {
			Level  int    `json:"level"`
			ID     string `json:"id"`
			Name   string `json:"name"`
			Parent string `json:"parent"`
			Score  struct {
				Local  float64 `json:"local"`
				Global float64 `json:"global"`
			} `json:"score"`
		}{},
	}

	for _, c := range h.Criteria {
		res.Criteria = append(res.Criteria, struct {
			Level  int    `json:"level"`
			ID     string `json:"id"`
			Name   string `json:"name"`
			Parent string `json:"parent"`
			Score  struct {
				Local  float64 `json:"local"`
				Global float64 `json:"global"`
			} `json:"score"`
		}{
			Level:  c.Level,
			ID:     c.ID,
			Name:   c.Name,
			Parent: c.Parent,
			Score: struct {
				Local  float64 `json:"local"`
				Global float64 `json:"global"`
			}{
				Local:  c.Score.Local,
				Global: c.Score.Global,
			},
		})
	}

	return res
}

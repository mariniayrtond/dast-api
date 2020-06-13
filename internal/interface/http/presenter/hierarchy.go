package presenter

import "dast-api/internal/domain/model"

type HierarchyResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Owner        string   `json:"owner"`
	Alternatives []string `json:"alternatives"`
	Criteria     []struct {
		Level       int    `json:"level"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Parent      string `json:"parent"`
		Score       struct {
			Local  float64 `json:"local"`
			Global float64 `json:"global"`
		} `json:"score"`
	} `json:"criteria"`
}

func RenderHierarchy(h *model.Hierarchy) HierarchyResponse {
	res := HierarchyResponse{
		ID:           h.ID,
		Name:         h.Name,
		Description:  h.Description,
		Owner:        h.Owner,
		Alternatives: h.Alternatives,
		Criteria: []struct {
			Level       int    `json:"level"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Parent      string `json:"parent"`
			Score       struct {
				Local  float64 `json:"local"`
				Global float64 `json:"global"`
			} `json:"score"`
		}{},
	}

	for _, c := range h.Criteria {
		res.Criteria = append(res.Criteria, struct {
			Level       int    `json:"level"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Parent      string `json:"parent"`
			Score       struct {
				Local  float64 `json:"local"`
				Global float64 `json:"global"`
			} `json:"score"`
		}{
			Level:       c.Level,
			ID:          c.ID,
			Description: c.Name,
			Parent:      c.Parent,
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

func RenderHierarchies(h []*model.Hierarchy) []HierarchyResponse {
	toRet := []HierarchyResponse{}
	for _, hierarchy := range h {
		toRet = append(toRet, RenderHierarchy(hierarchy))
	}

	return toRet
}

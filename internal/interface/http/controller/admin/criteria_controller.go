package admin

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type criteriaRequest []criteria

type criteria struct {
	Level    int    `json:"level"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

func (m *criteriaRequest) UnmarshalJSON(b []byte) error {
	var body []criteria
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}

	if len(body) == 0 {
		return errors.New("criteria cannot be empty")
	}

	for i, c := range body {
		if c.Name == "" {
			return errors.New(fmt.Sprintf("criteria on index:%d. name must be != zero value", i+1))
		}
		if c.ID == "" {
			return errors.New(fmt.Sprintf("criteria on index:%d. id must be != zero value", i+1))
		}
	}

	for i, c := range body {
		n := c.Name
		for k := i+1; k<len(body); k++ {
			if strings.ToUpper(n) == strings.ToUpper(body[k].Name) {
				return fmt.Errorf("%s is duplicated", n)
			}
		}
	}

	*m = body
	return nil
}

func (m criteriaRequest) ToCriteriaModel() []model.Criteria {
	ret := []model.Criteria{}
	for _, c := range m {
		ret = append(ret, model.Criteria{
			Level:  c.Level,
			ID:     c.ID,
			Name:   c.Name,
			Parent: c.ParentID,
			Score:  model.Score{
				Local:  0,
				Global: 0,
			},
		})
	}
	return ret
}

func NewCriteriaAdminController(uc usecase.HierarchyCRUD) *criteriaAdminController {
	return &criteriaAdminController{useCase: uc}
}

type criteriaAdminController struct {
	useCase usecase.HierarchyCRUD
}

func (cac criteriaAdminController) Fill(c *gin.Context) {
	var input criteriaRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_criteria_request", err))
		return
	}

	h, err := cac.useCase.SetCriteria(c.Param("id"), input.ToCriteriaModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_adding_criteria", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderHierarchy(h))
}

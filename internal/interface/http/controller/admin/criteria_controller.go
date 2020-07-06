package admin

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type criteriaRequest []criteria

type criteria struct {
	ID          string `json:"id"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}

type template struct {
	Owner       string          `json:"owner" binding:"required"`
	Description string          `json:"description" binding:"required"`
	Criteria    criteriaRequest `json:"criteria" binding:"required"`
}

func (m *criteriaRequest) UnmarshalJSON(b []byte) error {
	var body []criteria
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}

	if len(body) == 0 {
		return errors.New("criteria cannot be empty")
	}

	if len(body) == 1 {
		return errors.New("criteria must be two or more")
	}

	for i, c := range body {
		if c.Description == "" {
			return errors.New(fmt.Sprintf("criteria on index:%d. name must be != zero value", i+1))
		}
		if c.ID == "" {
			return errors.New(fmt.Sprintf("criteria on index:%d. name must be != zero value", i+1))
		}
		if c.Level < 0 {
			return errors.New(fmt.Sprintf("criteria on index:%d. level must be > 0", i+1))
		}
		if c.Parent == "" && c.Level > 0 {
			return errors.New(fmt.Sprintf("criteria on index:%d has level > 0 but parent empty", i+1))
		}
	}

	for i, c := range body {
		for k := i + 1; k < len(body); k++ {
			if strings.ToUpper(c.Description) == strings.ToUpper(body[k].Description) {
				return fmt.Errorf("%s is duplicated", c.Description)
			}
			if strings.TrimSpace(strings.ToUpper(c.ID)) == strings.TrimSpace(strings.ToUpper(body[k].ID)) {
				return fmt.Errorf("%s is duplicated", c.ID)
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
			Name:   c.Description,
			Parent: c.Parent,
			Score: model.Score{
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
		logger.Errorf("error filling hierarchy", err)
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_criteria_request", err))
		return
	}

	context, _ := c.Get(c.Param("id"))
	h, err := cac.useCase.SetCriteria(context.(*model.Hierarchy), input.ToCriteriaModel())
	if err != nil {
		logger.Errorf("error getting hierarchy:%s", err, c.Param("id"))
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_adding_criteria", err))
		return
	}

	logger.Infof("hierarchy:%s filled successfully", c.Param("id"))
	c.JSON(http.StatusOK, presenter.RenderHierarchy(h))
}

func (cac criteriaAdminController) SaveTemplate(c *gin.Context) {
	var input template
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_template", err))
		return
	}

	template, err := cac.useCase.SaveCriteriaTemplate(input.Owner, input.Description, input.Criteria.ToCriteriaModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_saving_template", err))
		return
	}

	c.JSON(http.StatusCreated, presenter.RenderCriteriaTemplate(template))
}

func (cac criteriaAdminController) SearchPublicTemplates(c *gin.Context) {
	templates, err := cac.useCase.SearchPublicTemplates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_searching_templates", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderCriteriaTemplates(templates))
}

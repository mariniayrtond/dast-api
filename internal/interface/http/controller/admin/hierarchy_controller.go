package admin

import (
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type hierarchyRequest struct {
	Name         string   `json:"name" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	Owner        string   `json:"owner" binding:"required"`
	Objective    string   `json:"objective" binding:"required"`
	Alternatives []string `json:"alternatives" binding:"required"`
}

func NewHierarchyAdminController(uc usecase.HierarchyCRUD) *hierarchyAdminController {
	return &hierarchyAdminController{useCase: uc}
}

type hierarchyAdminController struct {
	useCase usecase.HierarchyCRUD
}

func (hac hierarchyAdminController) Create(c *gin.Context) {
	var input hierarchyRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_hierarchy_request", err))
		return
	}

	if len(input.Alternatives) < 2 {
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_alternatives", errors.New("the size of alternatives must be > 1")))
		return
	}

	res, err := hac.useCase.RegisterHierarchy(input.Name, input.Description, input.Owner, input.Alternatives, input.Objective)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_saving_hierarchy", err))
		return
	}

	c.JSON(http.StatusCreated, presenter.RenderHierarchy(res))
}

func (hac hierarchyAdminController) Get(c *gin.Context) {
	id := c.Param("id")
	res, err := hac.useCase.GetHierarchy(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_getting_hierarchy", err))
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, presenter.NewNotFound(id))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderHierarchy(res))
}

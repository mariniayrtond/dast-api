package pairwise

import (
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewPairwiseController(uc usecase.PairwiseComparison) *pairwiseController {
	return &pairwiseController{useCase: uc}
}

type pairwiseController struct {
	useCase usecase.PairwiseComparison
}

func (p pairwiseController) GenerateCriteriaMatrices(c *gin.Context) {
	judgements, err := p.useCase.GenerateMatrices(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_generating_criteria_matrices", err))
	}
	c.JSON(http.StatusCreated, presenter.RenderCriteriaJudgements(judgements))
}

func (p pairwiseController) SetCriteriaScore(c *gin.Context) {
	judgements, err := p.useCase.GenerateMatrices(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_generating_criteria_matrices", err))
	}
	c.JSON(http.StatusCreated, presenter.RenderCriteriaJudgements(judgements))
}

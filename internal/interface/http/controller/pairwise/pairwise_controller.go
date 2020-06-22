package pairwise

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type judgementsRequest struct {
	CriteriaComparison    []pairwiseComparison `json:"criteria_comparison" binding:"required"`
	AlternativeComparison []matrixContext      `json:"alternative_comparison" binding:"required"`
}

func (j judgementsRequest) ToJudgementsModel() *model.CriteriaJudgements {
	ret := model.CriteriaJudgements{
		CriteriaComparison:    []model.CriteriaPairwiseComparison{},
		AlternativeComparison: []model.MatrixContext{},
	}

	for _, alternativeMatrix := range j.AlternativeComparison {
		ret.AlternativeComparison = append(ret.AlternativeComparison, model.MatrixContext{
			ComparedTo: alternativeMatrix.ComparedTo,
			Elements:   alternativeMatrix.Elements,
			Judgements: alternativeMatrix.Judgements,
		})
	}

	for _, criteriaComparison := range j.CriteriaComparison {
		ret.CriteriaComparison = append(ret.CriteriaComparison, model.CriteriaPairwiseComparison{
			Level:         criteriaComparison.Level,
			MatrixContext: model.MatrixContext{
				ComparedTo: criteriaComparison.MatrixContext.ComparedTo,
				Elements:   criteriaComparison.MatrixContext.Elements,
				Judgements: criteriaComparison.MatrixContext.Judgements,
			},
		})
	}

	return &ret
}

type pairwiseComparison struct {
	Level         int           `json:"level"`
	MatrixContext matrixContext `json:"matrix_context"`
}

type matrixContext struct {
	ComparedTo string      `json:"compared_to"`
	Elements   []string    `json:"elements"`
	Judgements [][]float64 `json:"judgements"`
}

func NewPairwiseController(uc usecase.PairwiseComparison) *pairwiseController {
	return &pairwiseController{useCase: uc}
}

type pairwiseController struct {
	useCase usecase.PairwiseComparison
}

func (p pairwiseController) GenerateCriteriaMatrices(c *gin.Context) {
	context, _ := c.Get(c.Param("id"))
	judgements, err := p.useCase.GenerateMatrices(context.(*model.Hierarchy))
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_generating_criteria_matrices", err))
	}
	c.JSON(http.StatusCreated, presenter.RenderCriteriaJudgements(judgements))
}

func (p pairwiseController) GetJudgements(c *gin.Context) {
	j, err := p.useCase.GetJudgements(c.Param("id"), c.Param("judgements_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_getting_judgements", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderCriteriaJudgements(j))
}

func (p pairwiseController) SetJudgements(c *gin.Context) {
	var input judgementsRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_judgements", err))
		return
	}

	j, err := p.useCase.UpdateJudgements(c.Param("id"), c.Param("judgements_id"), input.ToJudgementsModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_updating_judgements", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderCriteriaJudgements(j))
}

func (p pairwiseController) GenerateResults(c *gin.Context) {
	context, _ := c.Get(c.Param("id"))
	judgements, err := p.useCase.GenerateResults(context.(*model.Hierarchy), c.Param("judgements_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_generating_criteria_matrices", err))
	}
	c.JSON(http.StatusOK, presenter.RenderCriteriaJudgements(judgements))
}

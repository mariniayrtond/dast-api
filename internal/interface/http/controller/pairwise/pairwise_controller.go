package pairwise

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/interface/http/presenter"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
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
			Level: criteriaComparison.Level,
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
		logger.Errorf("error at matrix generation for hierarchy:%s", err, c.Param("id"))
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_generating_criteria_matrices", err))
		return
	}

	logger.Infof("matrices generated for hierarchy:%s", c.Param("id"))
	c.JSON(http.StatusCreated, presenter.RenderCriteriaJudgements(judgements))
}

func (p pairwiseController) GetJudgements(c *gin.Context) {
	j, err := p.useCase.GetJudgements(c.Param("id"), c.Param("judgements_id"))
	if err != nil {
		logger.Errorf("error getting judgements", err, c.Param("id"))
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_getting_judgements", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderCriteriaJudgements(j))
}

func (p pairwiseController) GetAllJudgementsByHierarchyID(c *gin.Context) {
	jj, err := p.useCase.GetJudgementsByHierarchyId(c.Param("id"))
	if err != nil {
		logger.Errorf("error getting judgements by hierarchy id", err, c.Param("id"))
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_getting_judgements", err))
		return
	}

	c.JSON(http.StatusOK, presenter.RenderSomeCriteriaJudgements(jj))
}

func (p pairwiseController) SetJudgements(c *gin.Context) {
	var input judgementsRequest
	if err := c.BindJSON(&input); err != nil {
		logger.Error("error setting judgements", err)
		c.JSON(http.StatusBadRequest, presenter.NewBadRequest("error_parsing_judgements", err))
		return
	}

	j, err := p.useCase.UpdateJudgements(c.Param("id"), c.Param("judgements_id"), input.ToJudgementsModel())
	if err != nil {
		logger.Error("error setting judgements", err)
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_updating_judgements", err))
		return
	}

	logger.Infof("judgements updated hierarchy:%s - judgements:%s", c.Param("id"), c.Param("judgements_id"))
	c.JSON(http.StatusOK, presenter.RenderCriteriaJudgements(j))
}

func (p pairwiseController) GenerateResults(c *gin.Context) {
	context, _ := c.Get(c.Param("id"))
	judgements, err := p.useCase.GenerateResults(context.(*model.Hierarchy), c.Param("judgements_id"))
	if err != nil {
		logger.Error("error generating judgements", err)
		c.JSON(http.StatusInternalServerError, presenter.NewInternalServerError("error_generating_criteria_matrices", err))
		return
	}

	logger.Infof("results generated. hierarchy:%s - judgements:%s", c.Param("id"), c.Param("judgements_id"))
	c.JSON(http.StatusOK, presenter.RenderCriteriaJudgements(judgements))
}

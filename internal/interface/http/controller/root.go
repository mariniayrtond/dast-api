package controller

import (
	"dast-api/internal/interface/http/controller/admin"
	"dast-api/internal/interface/http/controller/pairwise"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterAdminControllers(e *gin.Engine, uc usecase.HierarchyCRUD) {
	hc := admin.NewHierarchyAdminController(uc)
	e.POST("/hierarchy", hc.Create)
	e.GET("/hierarchy/:id", hc.Get)

	c := admin.NewCriteriaAdminController(uc)
	e.PUT("/hierarchy/:id/criteria", c.Fill)
}

func RegisterPairwiseControllers(e *gin.Engine, uc usecase.PairwiseComparison) {
	pwise := pairwise.NewPairwiseController(uc)
	e.POST("/pairwise/:id/generate", pwise.GenerateCriteriaMatrices)
}

package controller

import (
	"dast-api/internal/interface/http/controller/admin"
	"dast-api/internal/interface/http/controller/pairwise"
	"dast-api/internal/interface/http/controller/user"
	"dast-api/internal/interface/http/middleware"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPingController(e *gin.Engine) {
	e.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ping")
		return
	})
}

func RegisterUserControllers(e *gin.Engine, uc usecase.UserUseCase) {
	controller := user.NewUserController(uc)
	e.POST("dast/user/login", controller.LogIn)
	e.POST("dast/user/create", controller.Create)
	e.POST("dast/user/validate", controller.ValidateToken)
	e.GET("dast/user/:id", controller.Get)
}

func RegisterAdminControllers(e *gin.Engine, uc usecase.HierarchyCRUD, userUseCase usecase.UserUseCase, auth *middleware.AuthHandler) {
	hc := admin.NewHierarchyAdminController(uc, userUseCase)
	e.GET("dast/hierarchies/:username/search", auth.ValidateUsername(), hc.SearchByUsername)
	e.POST("dast/hierarchy", hc.Create)
	e.GET("dast/hierarchy/:id", hc.Get)

	c := admin.NewCriteriaAdminController(uc)
	e.PUT("dast/hierarchy/:id/criteria", auth.ValidateToken(), c.Fill)
	e.POST("dast/criteria/template", auth.IsAdmin(), c.SaveTemplate)
	e.GET("dast/criteria/template/search", c.SearchPublicTemplates)
}

func RegisterPairwiseControllers(e *gin.Engine, uc usecase.PairwiseComparison, auth *middleware.AuthHandler) {
	pwise := pairwise.NewPairwiseController(uc)
	e.POST("dast/pairwise/:id/generate", auth.ValidateToken(), pwise.GenerateCriteriaMatrices)
	e.GET("dast/pairwise/:id/search/judgements", auth.ValidateToken(), pwise.GetAllJudgementsByHierarchyID)
	e.GET("dast/pairwise/:id/judgements/:judgements_id", pwise.GetJudgements)
	e.PUT("dast/pairwise/:id/judgements/:judgements_id", auth.ValidateToken(), pwise.SetJudgements)
	e.POST("dast/pairwise/:id/judgements/:judgements_id/resolve", auth.ValidateToken(), pwise.GenerateResults)
}

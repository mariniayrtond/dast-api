package test

import (
	"dast-api/internal/domain/service"
	"dast-api/internal/interface/http/controller/admin"
	"dast-api/internal/interface/http/controller/pairwise"
	"dast-api/internal/interface/persistance/memory"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
)

var controllers *gin.Engine


func SetupNetworkControllers() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.TestMode)
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = false

	hRepo := memory.NewHierarchyRepository()
	pRepo := memory.NewCriteriaJudgementsRepository()

	hCrudUC := usecase.NewHierarchyCRUD(hRepo, service.NewCriteriaService(hRepo))
	hc := admin.NewHierarchyAdminController(hCrudUC)
	router.POST("/hierarchy", hc.Create)
	router.GET("/hierarchy/:id", hc.Get)

	c := admin.NewCriteriaAdminController(hCrudUC)
	router.PUT("/hierarchy/:id/criteria", c.Fill)

	pwise := pairwise.NewPairwiseController(usecase.NewPairwiseComparisonUC(hRepo, pRepo, service.NewPairwiseService()))
	router.POST("/pairwise/:id/generate", pwise.GenerateCriteriaMatrices)
	router.POST("/pairwise/:id/judgements/:judgements_id", pwise.SetCriteria)
	router.POST("/pairwise/:id/judgements/:judgements_id/resolve", pwise.GenerateResults)

	return router
}

//PerformRequest used for make a request in tests
var performRequest = func(method, target, body string, engine *gin.Engine) *httptest.ResponseRecorder {
	payload := strings.NewReader(body)
	req := httptest.NewRequest(method, target, payload)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	return res
}

func init() {
	controllers = SetupNetworkControllers()
}

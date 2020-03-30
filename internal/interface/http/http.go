package http

import (
	"dast-api/internal/app"
	"dast-api/internal/interface/http/controller"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *app.Container) {
	controller.RegisterPingController(server)
	controller.RegisterAdminControllers(server, ctn.Resolve("hierarchy-usecase").(usecase.HierarchyCRUD))
	controller.RegisterPairwiseControllers(server, ctn.Resolve("pwise-usecase").(usecase.PairwiseComparison))
}

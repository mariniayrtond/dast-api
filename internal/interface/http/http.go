package http

import (
	"dast-api/internal/app"
	"dast-api/internal/interface/http/controller"
	"dast-api/internal/interface/http/middleware"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Apply(server *gin.Engine, ctn *app.Container) {
	server.Use()
	uc := ctn.Resolve("user-usecase").(usecase.UserUseCase)
	hc := ctn.Resolve("hierarchy-usecase").(usecase.HierarchyCRUD)
	pwisec := ctn.Resolve("pwise-usecase").(usecase.PairwiseComparison)
	auth := middleware.NewAuthHandler(hc, uc)

	controller.RegisterPingController(server)
	controller.RegisterAdminControllers(server, hc, uc, auth)
	controller.RegisterPairwiseControllers(server, pwisec, auth)
	controller.RegisterUserControllers(server, uc)

}

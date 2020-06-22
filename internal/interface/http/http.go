package http

import (
	"dast-api/internal/app"
	"dast-api/internal/interface/http/controller"
	"dast-api/internal/interface/http/middleware"
	"dast-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func Apply(server *gin.Engine, ctn *app.Container) {
	server.Use()
	uc := ctn.Resolve("user-usecase").(usecase.UserUseCase)
	if uc == nil {
		log.Fatal("no se pudo resolver el caso de uso User")
	}
	hc := ctn.Resolve("hierarchy-usecase").(usecase.HierarchyCRUD)
	if hc == nil {
		log.Fatal("no se pudo resolver el caso de uso Hierarchy")
	}
	pwisec := ctn.Resolve("pwise-usecase").(usecase.PairwiseComparison)
	if pwisec == nil {
		log.Fatal("no se pudo resolver el caso de uso Pairwise")
	}
	auth := middleware.NewAuthHandler(hc, uc)

	controller.RegisterPingController(server)
	controller.RegisterAdminControllers(server, hc, uc, auth)
	controller.RegisterPairwiseControllers(server, pwisec, auth)
	controller.RegisterUserControllers(server, uc)
}

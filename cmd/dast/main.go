package main

import (
	"dast-api/internal/app"
	"dast-api/internal/interface/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	ctn, err := app.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
	server := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AddAllowHeaders("X-Auth-Token")
	server.Use(cors.New(corsCfg))
	http.Apply(server, ctn)
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("failed at running server: %v", err)
	}
}

package main

import (
	"dast-api/internal/app"
	"dast-api/internal/interface/http"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	logger.SetLevel(logger.InfoLevel)
	logger.SetFormatter(&logger.JSONFormatter{})
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
	if err := server.Run(GetPort()); err != nil {
		log.Fatalf("failed at running server: %v", err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

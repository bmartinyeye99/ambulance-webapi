package main

import (
	"log"
	"os"
	"strings"

	"github.com/bmartinyeye99/ambulance-webapi/api"
	"github.com/bmartinyeye99/ambulance-webapi/internal/ambulance_wl"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())

	handleFunctions := &ambulance_wl.ApiHandleFunctions{
		AmbulanceConditionsAPI:  ambulance_wl.NewAmbulanceConditionsApi(),
		AmbulanceWaitingListAPI: ambulance_wl.NewAmbulanceWaitingListApi(),
	}
	ambulance_wl.NewRouterWithGinEngine(engine, *handleFunctions)

	// request routings
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)

}

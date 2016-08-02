package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"ymir/settings"
	"ymir/web/api"
)

func Run() {
	var (
		router *gin.Engine
	)

	if settings.IsDevelop() {
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	api.Bluesprint(router)

	// Start http server
	fmt.Printf("listen http server %s %s on %s\n", settings.Name, settings.Version, settings.HttpListen)

	if err := router.Run(settings.HttpListen); err != nil {
		log.Fatalf("error in Engine.Run: %s", err)
	}
}

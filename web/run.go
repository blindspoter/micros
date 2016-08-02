package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gin/settings"
	"gin/web/api"
)

func Run() {
	settings.Parse()
	var (
		router *gin.Engine
	)

	if settings.IsDevelop() {
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	api.Strap(router)

	// Start http server
	fmt.Printf("listen http server %s %s on %s\n", settings.Name, settings.Version, settings.HttpListen)

	if err := router.Run(settings.HttpListen); err != nil {
		log.Fatalf("error in Engine.Run: %s", err)
	}
}

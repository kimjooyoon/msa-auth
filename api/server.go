package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var Server *gin.Engine

func RunServer(mode bool, allowOrigin string) {
	if mode {
		gin.SetMode(gin.ReleaseMode)
	}

	if Server != nil {
		return
	}

	Server = gin.Default()
	Server.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{allowOrigin},
			AllowHeaders: []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
			AllowMethods: []string{"*"},
		}))

	HandlerSetup()

	err := Server.Run()
	if err != nil {
		log.Panicf("RunServer Error\nerr=%v", err)
	}

	return
}

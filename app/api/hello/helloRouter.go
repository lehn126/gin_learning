package hello

import (
	"gin_app/app/api/hello/controller"
	"gin_app/app/middleware"

	"github.com/gin-gonic/gin"
)

/*
	 func RegHelloHandlers() {
		http.HandleFunc("/hello", controller.SayHello)
	}
*/

func RegHelloHandlersGin(engine *gin.Engine) {
	// use router group
	helloGroup := engine.Group("/hello")
	helloGroup.Use(middleware.ResponseTimeMW)
	helloGroup.GET("/", controller.SayHelloGin) // path "/hello"
}

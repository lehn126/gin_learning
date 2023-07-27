package api

import (
	"gin_app/app/api/alarm"
	"gin_app/app/api/hello"
	"gin_app/app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register all handlers in this app
/* func RegisterHandlers() {
	hello.RegHelloHandlers()
} */

// 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// Access-Control-Allow-Credentials=true 和 Access-Control-Allow-Origin="*" 有冲突
		// 如果客户端开启了withCredentials, 仍然不能正确访问。需要把Access-Control-Allow-Origin的值设置为客户端传过来的值
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func loadHtml(engine *gin.Engine) {
	// don't use LoadHTMLFiles and LoadHTMLGlob at same time to avoid template undefined error
	//engine.LoadHTMLFiles("web/common/index.html")
	engine.LoadHTMLGlob("web/**/*")
	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", nil) // use definition name in index.html
	})
	engine.NoRoute(middleware.Error404MW) // 404 error
	engine.Use(middleware.Error500MW)     // 500 error
}

func RegisterHandlersGin(engine *gin.Engine) {
	engine.Use(Cors())
	loadHtml(engine)
	hello.RegHelloHandlersGin(engine)
	alarm.RegAlarmHandlersGin(engine)
}

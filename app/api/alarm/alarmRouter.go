package alarm

import (
	"gin_app/app/api/alarm/controller"

	"github.com/gin-gonic/gin"
)

func RegAlarmHandlersGin(engine *gin.Engine) {
	engine.GET("/alarm/:id", controller.TestGetAlarmGin)
}

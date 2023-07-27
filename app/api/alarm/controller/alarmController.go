package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAlarmForm struct {
	ID int64 `uri:"id" binding:"required"`
}

func TestGetAlarmGin(ctx *gin.Context) {
	form := new(GetAlarmForm)
	//id := ctx.Param("id")
	e := ctx.ShouldBindUri(form)
	if e != nil {
		panic(e.Error())
	}

	id := form.ID
	msg := fmt.Sprintf("get alarm %v", id)
	ctx.String(http.StatusOK, msg)
}

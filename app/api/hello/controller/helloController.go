package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		name = "go lang http"
	}
	msg := []byte(fmt.Sprintf("hello %v", name))
	writer.Write(msg)
} */

type HelloForm struct {
	Name string `form:"name"`
}

func SayHelloGin(ctx *gin.Context) {
	form := new(HelloForm)
	//name := ctx.Request.URL.Query().Get("name")
	//name := ctx.Query("name")
	e := ctx.Bind(form)
	if e != nil {
		return
	}

	name := form.Name
	if name == "" {
		name = "go lang http"
	}

	msg := fmt.Sprintf("hello %v", name)
	ctx.String(http.StatusOK, msg)
}

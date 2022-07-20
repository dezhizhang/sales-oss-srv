package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (this *BaseController) Success(c *gin.Context) {
	fmt.Println("success")
}

func (this *BaseController) Fail(c *gin.Context) {
	fmt.Println("fail")
}

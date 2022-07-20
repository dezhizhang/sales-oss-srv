package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UploadController struct {
	BaseController
}

var Upload *UploadController

func (this *UploadController) List(c *gin.Context) {
	fmt.Println("hello list")
}

func init() {
	Upload = new(UploadController)
}

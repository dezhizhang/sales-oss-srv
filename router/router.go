package router

import (
	"github.com/gin-gonic/gin"
	"sales-oss/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/upload/list", controller.Upload.List)
	}
	return r
}

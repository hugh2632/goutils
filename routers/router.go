package routers

import (
	"github.com/gin-gonic/gin"
	"goutils/pkg/setting"
	"goutils/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		//apiv1.GET("/", v1.GetPdfBytes)
		//apiv1.GET("/GetPdfDoc", v1.GetPdfDoc)
		//apiv1.POST("/GetArray", v1.GetArray)
		apiv1.GET("/search", v1.Search)
		apiv1.POST("/getdata", v1.GetData)
	}

	return r
}
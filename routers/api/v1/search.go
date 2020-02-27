package v1

import (
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	_ = c.Request.URL.RawQuery
	//从es加载
}
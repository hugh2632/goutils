package v1

import (
	"github.com/gin-gonic/gin"
	"goutils/pkg/crawler"
	"net/http"
	"time"
)

func GetPdfBytes(c *gin.Context) {
	url := c.Request.URL.RawQuery
	bytes, err := crawler.GetPdfBytes(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"iserror": true,
			"msg":     "获取PDF失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"iserror": false,
			"msg":     "获取PDF成功",
			"data":    bytes,
		})
	}
}

func GetPdfDoc(c *gin.Context) {
	url := c.Request.URL.RawQuery
	bytes, err := crawler.GetPdfBytes(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"iserror": true,
			"msg":     "获取PDF失败",
		})
	} else {
		downloadName := time.Now().Format("20060102150405") + ".pdf"
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= " + downloadName}
		if err != nil {
			c.String(http.StatusOK, "%v", err)
			return
		}
		c.Writer.Write(bytes)
	}
}

func GetArray(c *gin.Context) {
	var arr, ok = c.GetPostFormArray("lead_id")
	if !ok{
		c.Writer.Write([]byte("参数错误"))
	}else{
		c.Writer.Write([]byte("结果是:" + arr[0] + "," + arr[1]))
	}
}

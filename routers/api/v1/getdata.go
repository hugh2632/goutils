package v1

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goutils/model"
	"goutils/pkg/crawler"
	mysqlUtil "goutils/util"
	"log"
	"net/http"
)

func GetData(c *gin.Context) {
	var d, _ = c.GetPostForm("data")
	mysqlUtil.NewMysql("30wishuser:Wish30!@#@tcp(192.168.30.156:3306)/sic", func(db *sql.DB) {
		rows, err := db.Query(`select * from TaskInfo where Name='` + d + "'")
		if err != nil{
			log.Fatal(err)
		}
		if rows.Next(){
			var t model.TaskInfo
			err = rows.Scan(&t.Name,&t.Url, &t.Rules)
			res, _ := json.Marshal(t)
			log.Println(string(res))
			err = crawler.GetData(t)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"iserror": true,
					"msg":     "发起扫描失败!" + err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"iserror": false,
					"msg":     "扫描成功",
				})
			}
		}
	})
}

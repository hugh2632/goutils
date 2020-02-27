package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goutils/pkg/setting"
	"goutils/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	router.LoadHTMLGlob("web/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "123.html",gin.H{
			"title": "Posts",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}


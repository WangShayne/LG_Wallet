package server

import (
	"LG_wallet/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Server()  *gin.Engine{

	//写入日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.Static("/img/", "./static/images/")
	r.POST("/", router.Routers)
	return r
}
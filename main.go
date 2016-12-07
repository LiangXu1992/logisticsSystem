package main

import (
	"github.com/gin-gonic/gin"
    "route"
)

//执行一个监听器
func main() {
    //初始化httpServer
	router := gin.Default()
    //初始化路由表
    route.InitRoute(router)
    //监听端口,建立http server
    router.Run("localhost:9999")
}
package main

import (
	"fugin/routers"
	"github.com/gin-gonic/gin"
)


func main() {
	// 创建路由
	r := gin.Default()
	// 配置路由
	routers.New(r)
	// 运行服务
	r.Run()
	

}




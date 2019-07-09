package routers

import (
	"fugin/controllers"
	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine)  {
	r.GET("/", controllers.HandleFunc)
	r.POST("/login", controllers.HandlePost)
}


package controllers

import (
	"fugin/conf"
	"fugin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"golang":"hello world"})
	db := conf.New()

	user := models.Default{Name:"go", Age:24}
	db.Create(&user)

	defer db.Close()
}

func HandlePost(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status_code":200})
}

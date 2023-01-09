package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.New()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
	return router
}
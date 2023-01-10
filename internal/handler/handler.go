package handler

import (
	"github.com/gin-gonic/gin"
	//"net/http"
  "fmt"
  "github.com/nailus/workout/internal/repository"
)

func InitRouters(r *repository.Repository) *gin.Engine {
	router := gin.New()
  fmt.Println(r.GetUserAll())
  users, _ := r.GetUserAll()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, users)
  })
	return router
}
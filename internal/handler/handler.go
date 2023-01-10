package handler

import (
	"github.com/gin-gonic/gin"
  "github.com/nailus/workout/internal/repository"
)

func InitRouters(r *repository.Repository) *gin.Engine {
	router := gin.New()

  // auth := router.Group("/auth") {
  //   auth.POST("/sign-up", h.signUp)
  //   auth.POST("/sign-in", h.signIn)
  // }

  api := router.Group("/api", h.userIdentity)


	return router
}
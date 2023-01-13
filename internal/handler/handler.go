package handler

import (
	"github.com/gin-gonic/gin"
  "github.com/nailus/workout/internal/service"
  "fmt"
  "net/http"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
  return &Handler{service: s}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

  api := router.Group("/api")
  {
    exercises := api.Group("exercises")
    exercises.POST("/", h.createExercise)
    exercises.GET("/", h.getAllExercises)
    exercises.GET("/:id", h.getExerciseById)
    exercises.PUT("/:id", h.updateExercise)
    exercises.DELETE("/:id", h.deleteExercise)
  }
	return router
}

func (h *Handler) createExercise(c *gin.Context) {
  fmt.Println("createExercise handler")
}

func (h *Handler) getAllExercises(c *gin.Context) {
  exercises, _ := h.service.GetAllExercises()
  
  c.JSON(http.StatusOK, exercises)
}

func (h *Handler) getExerciseById(c *gin.Context) {
  fmt.Println("getExerciseById handler")
}

func (h *Handler) updateExercise(c *gin.Context) {
  fmt.Println("updateExercise handler")
}

func (h *Handler) deleteExercise(c *gin.Context) {
  fmt.Println("deleteExercise handler")
}
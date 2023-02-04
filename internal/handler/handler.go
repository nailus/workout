package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nailus/workout/internal/service"
	"github.com/nailus/workout/internal/entity"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func ResponseError(c *gin.Context, status_code int, message string) {
	c.AbortWithStatusJSON(status_code, gin.H{"message": message})
}

func ResponseForbiddenError(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.authorizedUser)
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

func (h *Handler) signUp(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		ResponseError(c, http.StatusBadRequest, "input params are invalid")
		return
	}

	h.service.CreateUser(&user)
}

func (h *Handler) signIn(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		ResponseError(c, http.StatusBadRequest, "input params are invalid")
		return
	}
	
	token, err := h.service.GenerateAuthToken(&user)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "user didnt found")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}
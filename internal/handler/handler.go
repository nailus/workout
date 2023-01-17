package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nailus/workout/internal/service"
	"github.com/nailus/workout/internal/entity"

	"log"

	//"log"
)

type Handler struct {
	service *service.Service
}

type UserAuthCred struct {
	email string `db:"email"`
	password string `db:"password"`
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

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

func (h *Handler) signUp(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		//newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	h.service.CreateUser(&user)
}

func (h *Handler) signIn(c *gin.Context) {
	// TODO: Разобраться с ошибками!!
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		log.Fatal(err)
		//c.JSON(http.StatusBadRequest, gin.H{"error": "TEST!!!"})
		return
	}
	
	fundedUser, err := h.service.GetUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, fundedUser)
}
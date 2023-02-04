package handler

import (
	"strings"
	"log"
	"github.com/gin-gonic/gin"
)

// CHANGE NAME!!!
func (h *Handler) authorizedUser(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		log.Fatal("Authorization not found")
		ResponseForbiddenError(c)
		return
	}

	bearerHeader := strings.Split(authorizationHeader, " ")
	if len(bearerHeader) != 2 || bearerHeader[0] != "Bearer" {
		log.Fatal("Authorization Bearer not found")
		ResponseForbiddenError(c)
		return
	}	

	userId, err := h.service.ParseAuthToken(bearerHeader[1])

	if err != nil {
		log.Fatal(err)
		ResponseForbiddenError(c)
		return
	}

	c.Set("userId", userId)
}
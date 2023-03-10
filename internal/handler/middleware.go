package handler

import (
	"strings"
	"github.com/gin-gonic/gin"
)

// CHANGE NAME!!!
func (h *Handler) authorizedUser(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		ResponseForbiddenError(c)
		return
	}

	bearerHeader := strings.Split(authorizationHeader, " ")
	if len(bearerHeader) != 2 || bearerHeader[0] != "Bearer" {
		ResponseForbiddenError(c)
		return
	}	

	userId, err := h.service.ParseAuthToken(bearerHeader[1])

	if err != nil {
		ResponseForbiddenError(c)
		return
	}

	c.Set("userId", userId)
}
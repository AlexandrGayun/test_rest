package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type User struct {
	Username string `form:"username"`
}

func (h *Handler) GetProfiles(c *gin.Context) {
	var (
		user     User
		username *string
	)
	if err := c.ShouldBind(&user); err == nil {
		if strings.TrimSpace(user.Username) != "" {
			username = &user.Username
		}
	}
	if username != nil {
		res, err := h.service.GetProfileByUsername(c.Request.Context(), *username)
		if err != nil {
			h.logger.Error("/profiles with username error", zap.String("username", user.Username), zap.Error(err))
			internalServerErrorResponse(c)
			return
		}
		if res == nil {
			notFoundErrorResponse(c)
			return
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res, err := h.service.GetProfiles(c.Request.Context())
	if err != nil {
		h.logger.Error("/profiles error", zap.Error(err))
		internalServerErrorResponse(c)
	}
	if res == nil {
		notFoundErrorResponse(c)
		return
	}
	c.JSON(http.StatusOK, res)
}

func internalServerErrorResponse(c *gin.Context) {
	res := gin.H{"error": "internal error, try again later"}
	c.JSON(http.StatusInternalServerError, res)
}
func notFoundErrorResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "no results found"})
}

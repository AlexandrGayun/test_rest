package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"test_task/business/service"
)

type AuthChecker struct {
	logger  *zap.Logger
	service service.Implementor
}

func NewAuthChecker(logger *zap.Logger, service service.Implementor) *AuthChecker {
	return &AuthChecker{logger: logger, service: service}
}

func (m *AuthChecker) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := strings.TrimSpace(c.GetHeader("Api-key"))
		if len(key) != 0 {
			if err := m.service.CheckAuth(key); err != nil {
				m.logger.Error("authorization failed", zap.String("api-key", key), zap.Error(err))
				res := gin.H{"error": "access forbidden"}
				c.JSON(http.StatusForbidden, res)
				return
			}
			c.Next()
		} else {
			res := gin.H{"error": "no api key provided"}
			c.JSON(http.StatusForbidden, res)
		}
	}
}

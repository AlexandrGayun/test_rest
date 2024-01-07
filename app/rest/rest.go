package rest

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"test_task/app/rest/handler"
	"test_task/app/rest/middleware"
	"test_task/business/service"
	"test_task/utils/config"
	"time"
)

type Server struct {
	logger  *zap.Logger
	service service.Implementor
	config  config.Config
}

func New(logger *zap.Logger, service service.Implementor, cfg config.Config) *Server {
	return &Server{
		logger:  logger,
		service: service,
		config:  cfg,
	}
}

func (s *Server) Start() {
	r := s.setupRouter()
	s.logger.Info("REST API server running", zap.String("port", s.config.HttpPort))
	if err := r.Run(fmt.Sprintf(":%s", s.config.HttpPort)); err != nil {
		s.logger.Fatal("failed to start http server", zap.Error(err))
	}
}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.New()
	m := middleware.NewAuthChecker(s.logger, s.service)
	r.Use(ginzap.Ginzap(s.logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(s.logger, true))
	r.Use(m.Auth())
	h := handler.New(s.logger, s.service)
	r.GET("/profile", h.GetProfiles)
	return r
}

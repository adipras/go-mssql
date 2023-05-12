package server

import (
	"context"
	"fmt"
	"go-mssql/internal/infra/server/handler/area"
	"go-mssql/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr    string
	engine      *gin.Engine
	areaService services.AreaService
}

func New(ctx context.Context, host string, port uint, areaService services.AreaService) Server {
	server := Server{
		engine:      gin.Default(),
		httpAddr:    fmt.Sprintf("%s:%d", host, port),
		areaService: areaService,
	}

	server.routes()
	return server
}

func (s *Server) routes() {
	s.engine.GET("/area/list", area.ListHandler(s.areaService))
}

func (s *Server) Run() error {
	err := s.engine.Run(s.httpAddr)
	if err != nil {
		return err
	}
	log.Println("Server running on :", s.httpAddr)

	return nil
}

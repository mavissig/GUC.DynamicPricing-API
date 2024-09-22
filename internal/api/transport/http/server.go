package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport"
	"log"
)

type Http struct {
	router *gin.Engine
	data   DataUC
	cfg    *transport.Config
}

func New(
	cfg *transport.Config,
	clientUC DataUC,
) *Http {
	return &Http{
		cfg:    cfg,
		router: gin.Default(),
		data:   clientUC,
	}
}

func (s *Http) Run() {
	s.registerHandlers()
	if err := s.router.Run(s.cfg.HTTP.Address); err != nil {
		log.Println("error run http server: ", err)
	}
}

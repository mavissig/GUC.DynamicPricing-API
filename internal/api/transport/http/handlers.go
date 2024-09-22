package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mavissig/GUC.DynamicPricing-API/docs"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Http) registerHandlers() {
	apiGroup := s.router.Group("/api/v1")
	s.registerHandlersData(apiGroup)
	s.registerHandlersDocs(apiGroup)
}

func (s *Http) registerHandlersDocs(g *gin.RouterGroup) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("/ping", s.docsEntrypointPing)
}

func (s *Http) registerHandlersData(g *gin.RouterGroup) {
	clientGroup := g.Group("/data")
	clientGroup.POST("/", s.dataEntrypointAdd)
}

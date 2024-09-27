package http_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// docsEntrypointPing пинг сервера
// @Summary Пропинговать API
// @Description Пропинговать API
// @Tags docs
// @Accept json
// @Produce json
// @Success 200 "Успешный пинг API"
// @Failure 400 "Ошибка при пинге API"
// @Router /ping [get]
func (s *Http) docsEntrypointPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"server time": time.Now().String(),
	})
}

func (s *Http) docsEntrypointPingData(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

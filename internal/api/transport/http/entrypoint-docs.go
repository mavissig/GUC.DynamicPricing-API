package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// clientEntrypointAdd добавляет нового клиента
// @Summary Отправить данные на расчет
// @Description Отправляет данные в сервис для рассвета
// @Tags data
// @Accept json
// @Produce json
// @Param   data body domain.Client true "входные данные для расчета"
// @Success 200 "Сообщение об успешной отправке на расчет"
// @Failure 400 "Ошибка при отправке данных"
// @Router /data [post]
func (s *Http) docsEntrypointPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"server time": time.Now().String(),
	})
}

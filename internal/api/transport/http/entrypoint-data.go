package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/domain"
	"net/http"
)

type DataUC interface {
	DataAdd(data *domain.Data) error
}

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
func (s *Http) dataEntrypointAdd(ctx *gin.Context) {
	var jsonData domain.Data
	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[API][DATA][ERROR]: %v", err),
		})
		return
	}
	err := s.data.DataAdd(&jsonData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[API][DATA][ERROR]: %v", err),
		})
	}
}

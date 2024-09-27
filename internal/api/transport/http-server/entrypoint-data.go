package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/domain"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport/common"
	"net/http"
)

type DataUC interface {
	DataAdd(data *domain.Data) (uuid.UUID, error)
	DataGetByKey(key uuid.UUID) (*domain.Data, error)
}

// dataEntrypointAdd отправляет данные в сервис для расчета
// @Summary Отправить данные на расчет
// @Description Отправляет данные в сервис для рассвета
// @Tags data
// @Accept json
// @Producer json
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

	id, err := s.data.DataAdd(&jsonData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("[API][DATA][ERROR]: %v", err),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// dataEntrypointGetByKey получает данные из сервиса по ключу
// @Summary Получить данные по ключу
// @Description Получает данные из сервиса по ключу
// @Tags data
// @Accept json
// @Producer json
// @Param   key query string true "ключ для поиска данных uuid"
// @Success 200 "Данные по ключу"
// @Failure 404 "Данные не найдены"
// @Router /data [get]
func (s *Http) dataEntrypointGetByKey(ctx *gin.Context) {
	key := ctx.Query("id")

	id, err := uuid.Parse(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	res, err := s.data.DataGetByKey(id)
	if err != nil {
		ctx.JSON(common.ParseErrToHttpStatus(err.Error()), gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

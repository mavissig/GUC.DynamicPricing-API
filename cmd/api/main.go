// @title           DynamicPricing-API
// @version         1.0
// @description     API для работы с динамическим ценообразованием

// @contact.name   API Support(Егор)
// @contact.url    https://t.me/senior_stepik

// @license.name  MIT
// @license.url   http://opensource.org/licenses/MIT

// @host      127.0.0.1:8080
// @BasePath  /api/v1

// https://github.com/mavissig/GUC.DynamicPricing-API

package main

import (
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/domain"

	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport"
	thttp "github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport/http"
)

func main() {
	transportCfg := transport.LoadConfig()

	useCase := domain.New()

	serverHttp := thttp.New(transportCfg, useCase)
	serverHttp.Run()
}

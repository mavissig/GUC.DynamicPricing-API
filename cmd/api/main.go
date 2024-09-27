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
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/repository"
	"log"

	repositoryRedis "github.com/mavissig/GUC.DynamicPricing-API/internal/api/repository/redis"
	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport"
	httpSrv "github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport/http-server"
	kafkaClient "github.com/mavissig/GUC.DynamicPricing-API/internal/api/transport/kafka-client"
)

func main() {
	transportCfg := transport.LoadConfig()
	repositoryCfg := repository.LoadConfig()

	storageRedis := repositoryRedis.New(repositoryCfg)
	cKafka, err := kafkaClient.New(transportCfg)
	if err != nil {
		log.Fatal("[API][MAIN][KAFKA-CLIENT][NEW][ERROR]:")
	}

	useCase := domain.New(cKafka, storageRedis)

	serverHttp := httpSrv.New(transportCfg, useCase)
	serverHttp.Run()
}

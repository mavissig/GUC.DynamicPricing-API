package transport

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	HTTP *HttpConfig
	API  *ApiConfig
}

type ApiConfig struct {
	DefaultPageSize int `envconfig:"API_DEFAULT_PAGE_SIZE" required:"false"`
}

type HttpConfig struct {
	Address string `envconfig:"HTTP_ROUTER_ADDRESS" required:"true"`
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env.transport"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[TRANSPORT][CONFIG] ERROR: ", err)
		}
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg)

	return cfg
}

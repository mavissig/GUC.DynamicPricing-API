package repository

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	PGRedis *RedisConfig
}

type RedisConfig struct {
	Address  string `envconfig:"DB_REDIS_ADDRESS" required:"true"`
	Password string `envconfig:"DB_REDIS_PASSWORD" required:"true"`
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env.storage"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[API][TRANSPORT][CONFIG] WARN ", err)
		}
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[IMS]REDIS CONFIG: \n")
	fmt.Printf("Address:       %s\n", cfg.PGRedis.Address)

	return cfg
}

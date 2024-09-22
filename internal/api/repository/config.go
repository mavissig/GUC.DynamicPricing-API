package repository

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	PG *PgConfig
}

type PgConfig struct {
	Name     string `envconfig:"DB_POSTGRES_NAME" required:"true"`
	Host     string `envconfig:"DB_POSTGRES_HOST" required:"true"`
	User     string `envconfig:"DB_POSTGRES_USER" required:"true"`
	Password string `envconfig:"DB_POSTGRES_PASSWORD" required:"true"`
	Port     string `envconfig:"DB_POSTGRES_PORT" required:"true"`

	InitTables bool `envconfig:"DB_POSTGRES_INIT_TABLE" required:"true"`
	ClearStart bool `envconfig:"DB_POSTGRES_CLEAR_START" required:"true"`
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env.storage"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[STORAGE][POSTGRES][CONFIG] ERROR: ", err)
		}
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg)

	return cfg
}

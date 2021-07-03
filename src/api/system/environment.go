package system

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type environment struct {
	PostgreSQLURI string `env:"POSTGRESQL_URI" envDefault:"postgres://ticketier:ticketier@localhost:5432/ticketier"`
	MinioEndpoint string `env:"MINIO_ENDPOINT" envDefault:"localhost:9000"`
	MinioUsername string `env:"MINIO_USERNAME" envDefault:"ticketier"`
	MinioPassword string `env:"MINIO_PASSWORD" envDefault:"ticketier"`
	MinioUseSSL   bool   `env:"MINIO_USE_SSL" envDefault:"false"`
}

func NewEnvironment() (*environment, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	_env := environment{}
	err = env.Parse(&_env)
	if err != nil {
		return nil, err
	}

	return &_env, nil
}

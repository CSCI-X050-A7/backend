package config

import (
	"github.com/caarlos0/env/v7"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type conf struct {
	// server config
	Debug bool   `env:"DEBUG"               envDefault:"false"`
	Host  string `env:"HOST"                envDefault:"127.0.0.1"`
	Port  int    `env:"PORT"                envDefault:"5000"`
	Https bool   `env:"HTTPS"               envDefault:"false"`

	// sqlite config
	DBFilename string `env:"DB_FILENAME" envDefault:"sqlite3.db"`
	DBEcho     bool   `env:"DB_ECHO"     envDefault:"true"`

	// jwt config
	JWTSecret        string `env:"JWT_SECRET"                 envDefault:"secret"`
	JWTExpireSeconds int64  `env:"JWT_EXPIRE_SECONDS"         envDefault:"1209600"` // 14 days, in seconds
}

var Conf *conf

func init() {
	cfg := conf{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("failed to parse config: %+v", err)
	}
	logrus.Debugf("config object: %+v", cfg)
	Conf = &cfg
}

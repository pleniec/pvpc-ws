package config

import (
	"log"
)

var (
	config Config
)


type Config struct {
	AMQP struct {
		Username string
		Password string
		Host string
		Port int
	}
	Redis struct {
		Host string
		Port int
	}
}

func init() {
	log.Fatal("sraka")
}

func Get() Config {
	return config
}

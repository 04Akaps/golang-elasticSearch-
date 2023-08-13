package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Elastic struct {
		Uri      string
		User     string
		Password string
	}
	Server struct {
		Port string
	}
}

func NewConfig(file string) *Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}

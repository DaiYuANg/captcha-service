package config

import (
	"captcha-service/internal/constant"
	"strconv"
)

type Config struct {
	Http   Http   `koanf:"http"`
	Store  Store  `koanf:"store"`
	Logger Logger `koanf:"logger"`
}

type Logger struct {
	Level string `koanf:"level"`
}

type Http struct {
	Port      int  `koanf:"port"`
	Prefork   bool `koanf:"prefork"`
	LowMemory bool `koanf:"low_memory"`
}

func (h Http) GetPort() string {
	return strconv.Itoa(h.Port)
}

type Store struct {
	Type constant.StoreType `koanf:"type"`
	Url  string             `koanf:"url"`
}

func defaultConfig() Config {
	return Config{
		Http: Http{
			Port:    80,
			Prefork: true,
		},
		Store: Store{Type: constant.Memory},
	}
}

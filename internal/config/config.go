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
	Type     constant.StoreType `koanf:"type"`
	Url      string             `koanf:"url"`
	Port     int                `koanf:"port"`
	Username string             `koanf:"username"`
	Password string             `koanf:"password"`
	Database string             `koanf:"database"`
}

func defaultConfig() Config {
	return Config{
		Http: Http{
			Port:    18080,
			Prefork: false,
		},
		Store: Store{Type: constant.Memory},
	}
}

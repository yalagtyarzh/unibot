package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

const StructDateFormat = "2006-01-02"
const SecretJWT = "?"

type Config struct {
	IsDebug       *bool `yaml:"is_debug" env:"UNIBOT-IsDebug" env-default:"false" env-required:"true"`
	IsDevelopment *bool `yaml:"is_development" env:"UNIBOT-IsDevelopment" env-default:"false" env-required:"true"`
	Listen        struct {
		Type   string `yaml:"type" env:"UNIBOT-ListenType" env-default:"port"`
		BindIP string `yaml:"bind_ip" env:"UNIBOT-BindIP" env-default:"localhost"`
		Port   string `yaml:"port" env:"UNIBOT-Port" env-default:"8080"`
	} `yaml:"listen" env-required:"true"`
	Telegram struct {
		Token string `yaml:"token" env:"UNIBOT-TelegramToken" env-required:"true"`
	}
	AppConfig AppConfig `yaml:"app" env-required:"true"`
}

type AppConfig struct {
	LogLevel string `yaml:"log_level" env:"UNIBOT-LogLevel" env-default:"error" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		log.Printf("read application config in path %s", path)

		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}

// package config

// import (
// 	"flag"
// 	"log"
// 	"os"

// 	"github.com/ilyakaznacheev/cleanenv"
// )

// type HTTPServer struct {
// 	Addr string `yaml:"address" env-required:"true"`
// }

// // env-default:"production"

// type Config struct {
// 	Env         string `yaml:"env" env:"ENV" env-required:"true"`
// 	StoragePath string `yaml:"storage_path" env-required:"true"`
// 	HTTPServer  `yaml:"http_server"`
// }

// func MustLoad() *Config {
// 	var configPath string

// 	configPath = os.Getenv("CONFIG_PATH")

// 	if configPath == "" {
// 		flags := flag.String("config", "", "path to the configuration file")
// 		flag.Parse()

// 		configPath = *flags

// 		if configPath == "" {
// 			log.Fatal("Config path is not set")
// 		}
// 	}

// 	if _, err := os.Stat(configPath); os.IsNotExist(err) {
// 		log.Fatalf("config file does not exist: %s", configPath)
// 	}

// 	var cfg Config

// 	err := cleanenv.ReadConfig(configPath, &cfg)
// 	if err != nil {
// 		log.Fatalf("can not read config file: %s", err.Error())
// 	}

// 	return &cfg
// }

package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `env:"PORT" env-required:"true"` // Render sets PORT automatically
}

type Config struct {
	Env         string `env:"ENV" env-required:"true"`
	StoragePath string `env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  HTTPServer
}

func MustLoad() *Config {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("failed to load config from env: %s", err.Error())
	}

	// Convert PORT to address format if needed
	cfg.HTTPServer.Addr = ":" + cfg.HTTPServer.Addr

	return &cfg
}

package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name        string
		Description string
		Version     string
	}
	Server struct {
		Ip   string
		Port string
	}
	Database struct {
		Host     string
		Port     uint
		User     string
		Password string
		Dbname   string
	}
}

func Properties() Config {

	log.Printf("Loading Environment")

	cfg := Config{}

	viper := viper.GetViper()
	viper.SetConfigName("application")
	viper.AddConfigPath("resources")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[VIPER] Error reading configuration files, %s", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return cfg
}

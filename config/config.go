package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// InitConfig - Init Config Files
func InitConfig() {
	config = viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := config.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}
func GetConfig() *viper.Viper {
	return config
}

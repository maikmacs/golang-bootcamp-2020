package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Port string
	}

	Data struct {
		Source string
		Output string
	}
}

// Structure - Export Config Structure
var Structure config

// ReadConfig - Init Config Files
func ReadConfig() {
	Config := &Structure

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
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

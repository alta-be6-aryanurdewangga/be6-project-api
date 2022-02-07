package configs

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Adress   string `yaml:"adress"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Pasword  string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()
	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "crud_go_test"
	defaultConfig.Database.Adress = "localhost"
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = "root"
	defaultConfig.Database.Pasword = "root"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	// log.Info(viper.ReadInConfig())
	if err := viper.ReadInConfig(); err != nil {
		// fmt.Println("kok mlaku")
		log.Info("error in open file")
		return &defaultConfig
	}

	var finalConfig AppConfig

	if err := viper.Unmarshal(&finalConfig); err != nil {
		log.Info("error in extract external config, must use default config")
		return &defaultConfig
	}
	return &finalConfig
}

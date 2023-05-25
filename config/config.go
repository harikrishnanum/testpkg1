package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Key1 string `mapstructure:"key1"`
	Key2 int    `mapstructure:"key2"`
}

var Conf Config

func init() {
	Conf = Config{}

	//viper.AddConfigPath("pkg/config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Error reading config file, %s", err)
		return
	}
	viper.WatchConfig()
	err = viper.Unmarshal(Conf)

	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return
	}
}

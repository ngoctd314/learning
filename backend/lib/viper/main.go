package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

/*
 */
func main() {
	c := viper.New()
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	c.AddConfigPath(".")
	// viper.AddConfigPath("$HOME/.appname")
	err := c.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, ignore error if desired
			panic("config file not found")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Sprintf("config file found but, err: %v", err))
		}
	}

	// env
	c.SetEnvPrefix("")
	c.AutomaticEnv()
	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	log.Println(c.Get("author"))
	log.Println(c.Get("my.env"))
}

func setConfigFile() {

}

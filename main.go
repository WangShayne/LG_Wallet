package main

import (
	"LG_wallet/db"
	. "LG_wallet/server"
	"fmt"
	"github.com/spf13/viper"
)

var port string

func  initViper()  {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func init(){
	initViper()
}

func main() {
	port := viper.GetString("server.port")
	db.ConnecToDB()
	s:=Server()
	s.Run(":"+port)
}

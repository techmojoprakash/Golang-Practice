// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/spf13/viper"
// )

// type Config struct {
// 	Port             string `mapstructure:"port"` // mapping the json val "port" to this struct Port
// 	ConnectionString string `mapstructure:"connection_string"`
// }

// var AppConfig *Config // by using this variable we can access port and ConnectionString from anywhere in project

// func LoadAppConfig() {
// 	log.Println("Loading Server Configurations...")
// 	viper.AddConfigPath("./configdata") // config folder path
// 	viper.SetConfigName("advConfig")    // config file name
// 	viper.SetConfigType("json")         // config data type
// 	err := viper.ReadInConfig()         // reading the config data
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = viper.Unmarshal(&AppConfig) // convert struct to JSON
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// func main() {
// 	LoadAppConfig()
// 	fmt.Println("Hello")
// 	fmt.Println(AppConfig.Port) // accessing the data from AppConfig variable of Type Config
// 	fmt.Println(AppConfig.ConnectionString)

// }

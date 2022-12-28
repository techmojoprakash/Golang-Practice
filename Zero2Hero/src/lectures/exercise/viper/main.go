// // viper basic

// package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// var Vp = viper.New()

// func main() {
// 	// Vp := viper.New()                // create object for viper ; without creating the object also possible please check advConfig
// 	Vp.AddConfigPath("./configdata") // add config path
// 	Vp.SetConfigName("config")       // set config file name
// 	Vp.SetConfigType("json")         // set config data type

// 	err := Vp.ReadInConfig() // read the data from object
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(Vp.GetString("appname")) // then get the actual data by the key
// 	fmt.Println(Vp.GetString("server.host"))
// 	fmt.Println(Vp.GetString("server.port"))
// 	fmt.Println(Vp.GetString("database.name"))

// }

// // --------------------------------------------

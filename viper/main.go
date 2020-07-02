package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	value := viper.Get("GOMAXPROCS")
	fmt.Println(value)
	viper.Set("GOMAXPROCS", 13)
	value = viper.Get("GOMAXPROCS")
	fmt.Println(value)

	viper.BindEnv("NEW_VARIABLE")
	valueNew := viper.Get("NEW_VARIABLE")
	if valueNew == nil {
		fmt.Println("NEW_VARIABLE is not defined")
		return
	}
	fmt.Println(valueNew)

}

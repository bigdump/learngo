package main

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	conf *viper.Viper
	lock *sync.RWMutex
}

func main() {
	fmt.Println("read file test")
	vp := viper.New()
	vp.AddConfigPath("./global.toml")
	vp.ReadInConfig()

	fmt.Println(vp.GetString("title"))

}

package main

import (
	"fmt"
	config "github.com/ilholbea/gym-rat/config"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic("unable to load config")
	}

	fmt.Printf(conf.DbConfig.Host)
}

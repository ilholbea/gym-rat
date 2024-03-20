package main

import (
	"fmt"
	config "github.com/ilholbea/gym-rat/config"
)

func main() {
	conf := config.NewConfig()

	fmt.Printf(conf.Database.Host)
}

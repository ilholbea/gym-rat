package main

import (
	"github.com/ilholbea/gym-rat/config"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("unable to find env file")
	}

	conf, err := config.NewConfig()
	if err != nil {
		panic("unable to load config")
	}

}

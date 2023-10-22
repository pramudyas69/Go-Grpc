package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("unable to load env %s", err)
	}

	return &Config{
		Mongo: Mongo{
			Uri: os.Getenv("MONGO_URI"),
		},
		Server: Server{
			Port: os.Getenv("PORT"),
		},
		Token: Token{
			Access_Token: os.Getenv("SECRET_KEY"),
		},
	}
}

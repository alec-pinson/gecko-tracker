package main

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	Database struct {
		Url      string
		Username string
		Password string
	}
	Sources []string
}

var config Config

func LoadConfiguration() {
	config.Database.Url = os.Getenv("DATABASE_URL")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")

	if os.Getenv("SALE_SOURCES") != "" {
		config.Sources = strings.Split(os.Getenv("SALE_SOURCES"), ",")
	} else {
		// defaults
		config.Sources = []string{"Preloved", "Facebook"}
	}
	log.Println("Configured Sale Sources: " + strings.Join(config.Sources, ", "))
}

package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Database struct {
		Url      string
		Username string
		Password string
		Name     string
	}
	Sources   []string
	HatchTime time.Duration
}

var config Config

func LoadConfiguration() {
	config.Database.Url = os.Getenv("DATABASE_URL")
	config.Database.Name = os.Getenv("DATABASE_NAME")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")

	if os.Getenv("SALE_SOURCES") != "" {
		config.Sources = strings.Split(os.Getenv("SALE_SOURCES"), ",")
	} else {
		// defaults
		config.Sources = []string{"Preloved", "Facebook"}
	}
	log.Println("Configured Sale Sources: " + strings.Join(config.Sources, ", "))

	if os.Getenv("HATCH_DAYS") != "" {
		hatchTime, err := time.ParseDuration(DaysToHours(os.Getenv("HATCH_DAYS"))) // hatch eta 60 days, will automatically generate from average of eggs later
		if err != nil {
			log.Println(err)
		}
		config.HatchTime = hatchTime
	} else {
		config.HatchTime, _ = time.ParseDuration("1440h") // hatch eta 60 days, will automatically generate from average of eggs later
	}
	log.Println("Configured hatch time is " + config.HatchTime.String())
}

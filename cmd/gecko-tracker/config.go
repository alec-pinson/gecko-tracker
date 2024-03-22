package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Database struct {
		Url        string
		Username   string
		Password   string
		Name       string
		BackupName string
	}
	Sources   []string
	LayTime   time.Duration
	HatchTime time.Duration
}

var config Config

func LoadConfiguration() {
	config.Database.Url = os.Getenv("DATABASE_URL")
	config.Database.Name = os.Getenv("DATABASE_NAME")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database.BackupName = os.Getenv("DATABASE_BACKUP_NAME")

	if config.Database.Url == "" || config.Database.Name == "" || config.Database.Username == "" || config.Database.Password == "" {
		log.Println("Please set the following environment variables")
		log.Println("DATABASE_URL")
		log.Println("DATABASE_NAME")
		log.Println("DATABASE_USERNAME")
		log.Println("DATABASE_PASSWORD")
		os.Exit(1)
	}

	if os.Getenv("SALE_SOURCES") != "" {
		config.Sources = strings.Split(os.Getenv("SALE_SOURCES"), ",")
	} else {
		// defaults
		config.Sources = []string{"Preloved", "Facebook"}
	}
	log.Println("Configured Sale Sources: " + strings.Join(config.Sources, ", "))

	if os.Getenv("LAY_DAYS") != "" {
		layTime, err := time.ParseDuration(DaysToHours(os.Getenv("LAY_DAYS"))) // lay eta 14 days, will automatically generate from average of eggs later
		if err != nil {
			log.Println(err)
		}
		config.LayTime = layTime
	} else {
		config.LayTime, _ = time.ParseDuration("336h") // lay eta 14 days, will automatically generate from average of eggs later
	}
	log.Println("Configured lay time is " + config.LayTime.String())

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

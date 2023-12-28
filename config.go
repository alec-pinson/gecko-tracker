package main

import "os"

type Config struct {
	Database struct {
		Url      string
		Username string
		Password string
	}
}

var config Config

func LoadConfiguration() {
	config.Database.Url = os.Getenv("DATABASE_URL")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
}

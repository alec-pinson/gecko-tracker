package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Gecko struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var LayTime, _ = time.ParseDuration("336h") // lay eta 14 days, will automatically generate from average of eggs later

func AddGecko(id int, description string) (*Gecko, error) {
	// Check if a gecko with the same ID already exists
	for _, existingGecko := range geckos {
		if existingGecko.ID == id {
			return nil, errors.New("Gecko with the same ID already exists")
		}
	}

	var gecko Gecko
	gecko.ID = id
	gecko.Description = description
	geckos = append(geckos, gecko)

	log.Println("Added gecko '" + description + "' (id: " + strconv.Itoa(id) + ")")

	WriteToDB("gecko", gecko, Incubator{}, Egg{}, Sale{})

	return &gecko, nil
}

func LoadGecko(id int, description string) (*Gecko, error) {
	var gecko Gecko
	gecko.ID = id
	gecko.Description = description
	geckos = append(geckos, gecko)

	log.Println("Loaded gecko '" + description + "' (id: " + strconv.Itoa(id) + ")")

	return &gecko, nil
}

func GetGecko(id int) (Gecko, error) {
	for _, gecko := range geckos {
		if gecko.ID == id {
			return gecko, nil
		}
	}
	log.Println("Gecko not found: " + strconv.Itoa(id))
	return Gecko{}, fmt.Errorf("Gecko not found with ID: %d", id)
}

func (gecko Gecko) GetLastLayDate() time.Time {
	eggsLaid := gecko.GetLayHistory()
	var lastLayDate time.Time
	for _, egg := range eggsLaid {
		if egg.LayDate.After(lastLayDate) {
			lastLayDate = egg.LayDate
		}
	}

	return lastLayDate
}

func (gecko Gecko) GetLayETA() time.Time {
	return gecko.GetLastLayDate().Add(LayTime)
}

func (gecko Gecko) GetLayHistory() []Egg {
	var eggsLaid []Egg
	for _, egg := range eggs {
		if egg.GeckoID == gecko.ID {
			eggsLaid = append(eggsLaid, egg)
		}
	}

	return eggsLaid
}

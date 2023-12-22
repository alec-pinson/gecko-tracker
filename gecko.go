package main

import (
	"log"
	"strconv"
	"time"
)

type Gecko struct {
	Id          int
	Description string
}

var LayTime, _ = time.ParseDuration("336h") // lay eta 14 days, will automatically generate from average of eggs later

func AddGecko(id int, description string) *Gecko {
	var gecko Gecko
	gecko.Id = id
	gecko.Description = description
	geckos = append(geckos, gecko)

	log.Println("Added gecko '" + description + "' (id: " + strconv.Itoa(id) + ")")

	return &gecko
}

func GetGecko(id int) Gecko {
	for _, gecko := range geckos {
		if gecko.Id == id {
			return gecko
		}
	}
	log.Println("Gecko not found: " + strconv.Itoa(id))
	var geckoNotFound Gecko
	return geckoNotFound
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
		if egg.GeckoID == gecko.Id {
			eggsLaid = append(eggsLaid, egg)
		}
	}

	return eggsLaid
}

package main

import (
	"fmt"
	"log"
	"time"
)

type Egg struct {
	EggID                 string
	GeckoID               int       `json:"geckoId"`
	Count                 int       `json:"count"`
	LayDate               time.Time `json:"layDate"`
	FormattedLayDate      string
	HasHatched            bool      `json:"hasHatched"`
	HatchDate             time.Time `json:"hatchDate"`
	FormattedHatchDateETA string
}

func generateUniqueID() string {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	uniqueID := fmt.Sprintf("%d", currentTime)

	return uniqueID
}

var HatchTime, _ = time.ParseDuration("1440h") // hatch eta 60 days, will automatically generate from average of eggs later

func AddEgg(geckoId int, eggCount int, layDate string, hatchDate string) *Egg {
	var egg Egg
	egg.EggID = generateUniqueID()
	egg.GeckoID = geckoId
	egg.Count = eggCount
	LayDate, err := time.Parse("02/01/2006", layDate)
	if err != nil {
		log.Println(err)
	}
	egg.LayDate = LayDate
	if hatchDate != "" {
		HatchDate, err := time.Parse("02/01/2006", hatchDate)
		if err != nil {
			log.Println(err)
		}
		egg.HatchDate = HatchDate
	}
	eggs = append(eggs, egg)

	// AddEggToDB(egg)

	return &egg
}

func (egg *Egg) GetLayDateString() string {
	return egg.LayDate.Format("02-01-2006")
}

func (egg *Egg) Hatched() {
	egg.HasHatched = true
	egg.HatchDate = time.Now()
}

func (egg *Egg) GetHatchETA() time.Time {
	return egg.LayDate.Add(HatchTime)
}

func GetHatchETAString(egg *Egg) string {
	return egg.GetHatchETA().Format("02-01-2006")
}

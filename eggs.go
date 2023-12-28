package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Egg struct {
	ID          string
	IncubatorID int      `json:"incubatorId"`
	Incubator   struct { // slot position in the incubator
		Row    int `json:"incubatorRow"`
		Column int `json:"incubatorColumn"`
	}
	GeckoID               int       `json:"geckoId"`
	Count                 int       `json:"count"`
	LayDate               time.Time `json:"layDate"`
	FormattedLayDate      string    `json:"formattedLayDate"`
	HasHatched            bool      `json:"hasHatched"`
	HatchDate             time.Time `json:"hatchDate"`
	FormattedHatchDateETA string    `json:"formattedHatchDateETA"`
}

func generateUniqueID() string {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	uniqueID := fmt.Sprintf("%d", currentTime)

	return uniqueID
}

var HatchTime, _ = time.ParseDuration("1440h") // hatch eta 60 days, will automatically generate from average of eggs later

func AddEgg(incubatorId int, row int, column, geckoId int, eggCount int, layDate string, hatchDate string) *Egg {
	var egg Egg
	egg.ID = generateUniqueID()
	egg.IncubatorID = incubatorId
	egg.Incubator.Row = row
	egg.Incubator.Column = column
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
	egg.FormattedLayDate = egg.LayDate.Format("02-01-2006")
	egg.FormattedHatchDateETA = egg.GetHatchETAString() // Assuming HatchDate is the ETA
	eggs = append(eggs, egg)

	log.Println("Added new egg to incubator " + strconv.Itoa(egg.IncubatorID) + " slot " + strconv.Itoa(egg.Incubator.Row) + "," + strconv.Itoa(egg.Incubator.Column))

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

func (egg *Egg) GetHatchETAString() string {
	return egg.GetHatchETA().Format("02-01-2006")
}

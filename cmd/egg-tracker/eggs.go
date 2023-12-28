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
		Row    int `json:"row"`
		Column int `json:"column"`
	} `json:"incubator"`
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
	egg.FormattedHatchDateETA = egg.GetHatchETAString()
	eggs = append(eggs, egg)

	log.Println("Added new egg to incubator " + strconv.Itoa(egg.IncubatorID) + " slot " + strconv.Itoa(egg.Incubator.Row) + "," + strconv.Itoa(egg.Incubator.Column))

	WriteToDB("egg", Gecko{}, Incubator{}, egg, Sale{})

	return &egg
}

func LoadEgg(id string, incubatorId int, row int, column, geckoId int, eggCount int, layDate time.Time, hatchDate time.Time, formattedLayDate string, formattedHatchDate string, hasHatched bool) *Egg {
	var egg Egg
	egg.ID = id
	egg.IncubatorID = incubatorId
	egg.Incubator.Row = row
	egg.Incubator.Column = column
	egg.GeckoID = geckoId
	egg.Count = eggCount
	egg.LayDate = layDate
	egg.HatchDate = hatchDate
	egg.FormattedLayDate = formattedLayDate
	egg.FormattedHatchDateETA = formattedHatchDate
	egg.HasHatched = hasHatched
	eggs = append(eggs, egg)

	log.Println("Loaded egg, incubator " + strconv.Itoa(egg.IncubatorID) + " slot " + strconv.Itoa(egg.Incubator.Row) + "," + strconv.Itoa(egg.Incubator.Column))

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
	return egg.LayDate.Add(config.HatchTime)
}

func (egg *Egg) GetHatchETAString() string {
	return egg.GetHatchETA().Format("02-01-2006")
}

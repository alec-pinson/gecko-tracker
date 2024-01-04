package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Egg struct {
	ID          string   `json:"id"`
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
	Colour                string    // used for rag status
}

func generateUniqueID() string {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	uniqueID := fmt.Sprintf("%d", currentTime)

	return uniqueID
}

func GetEgg(id string) *Egg {
	for _, egg := range eggs {
		if egg.ID == id {
			return egg
		}
	}
	return &Egg{}
}

func AddEgg(incubatorId int, row int, column, geckoId int, eggCount int, layDate string) *Egg {
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
	egg.HatchDate = egg.GetHatchETA()
	egg.FormattedLayDate = egg.LayDate.Format("02-01-2006")
	egg.FormattedHatchDateETA = egg.GetHatchETAString()
	eggs = append(eggs, &egg)

	log.Println("Added new egg to incubator " + strconv.Itoa(egg.IncubatorID) + " slot " + strconv.Itoa(egg.Incubator.Row) + "," + strconv.Itoa(egg.Incubator.Column))

	WriteToDB("egg", Gecko{}, Incubator{}, egg, Sale{}, Tank{})

	return &egg
}

func LoadEgg(id string, incubatorId int, row int, column, geckoId int, eggCount int, formattedLayDate string, formattedHatchDate string, hasHatched bool) *Egg {
	var egg Egg
	egg.ID = id
	egg.IncubatorID = incubatorId
	egg.Incubator.Row = row
	egg.Incubator.Column = column
	egg.GeckoID = geckoId
	egg.Count = eggCount
	layDate, err := time.Parse("02-01-2006", formattedLayDate)
	if err != nil {
		log.Println(err)
	}
	egg.LayDate = layDate
	hatchDate, err := time.Parse("02-01-2006", formattedHatchDate)
	if err != nil {
		log.Println(err)
	}
	egg.HatchDate = hatchDate
	egg.FormattedLayDate = formattedLayDate
	egg.FormattedHatchDateETA = formattedHatchDate
	egg.HasHatched = hasHatched
	eggs = append(eggs, &egg)

	log.Println("Loaded egg, incubator " + strconv.Itoa(egg.IncubatorID) + " slot " + strconv.Itoa(egg.Incubator.Row) + "," + strconv.Itoa(egg.Incubator.Column))

	return &egg
}

func (egg *Egg) GetLayDateString() string {
	return egg.LayDate.Format("02-01-2006")
}

func (egg *Egg) Hatched() {
	egg.HasHatched = true
	egg.HatchDate = time.Now()
	egg.IncubatorID = 0
	egg.Incubator.Row = 0
	egg.Incubator.Column = 0
	log.Print("Marked egg as hatched - " + egg.ID)
	UpdateDB("egg", Gecko{}, Incubator{}, *egg, Sale{}, Tank{})
}

func (egg *Egg) GetHatchETA() time.Time {
	return egg.LayDate.Add(GetAverageHatchTimeDuration())
}

func (egg *Egg) GetHatchETAString() string {
	return egg.GetHatchETA().Format("02-01-2006")
}

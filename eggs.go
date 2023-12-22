package main

import (
	"log"
	"strconv"
	"time"
)

type Egg struct {
	SlotId     int       `json:"slotId"`
	GeckoID    int       `json:"geckoId"`
	Count      int       `json:"count"`
	LayDate    time.Time `json:"layDate"`
	HasHatched bool      `json:"hasHatched"`
	HatchDate  time.Time `json:"hatchDate"`
}

var HatchTime, _ = time.ParseDuration("1440h") // hatch eta 60 days, will automatically generate from average of eggs later

func AddEgg(slotId int, geckoId int, eggCount int, layDate string, hatchDate string) *Egg {
	var egg Egg
	egg.SlotId = slotId
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

	log.Println("Added new egg to slot " + strconv.Itoa(egg.SlotId))

	AddEggToDB(egg)

	return &egg
}

func GetEgg(slotId int) *Egg {
	for _, egg := range eggs {
		if egg.SlotId == slotId {
			return &egg
		}
	}
	var eggNotFound Egg
	return &eggNotFound
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

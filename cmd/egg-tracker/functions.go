package main

import (
	"fmt"
	"strconv"
	"time"
)

func GetNextLayDateInfo() string {
	var LayETA time.Time = time.Now().Add(time.Hour * 99999)
	var geckoId int = 0
	for _, gecko := range geckos {
		if gecko.GetLayETA().Before(LayETA) {
			LayETA = gecko.GetLayETA()
			geckoId = gecko.ID
		}
	}

	return LayETA.Format("02-01-2006") + " (gecko " + strconv.Itoa(geckoId) + ")"
}

func GetNextHatchDateInfo() string {
	var HatchETA time.Time = time.Now().Add(time.Hour * 99999)
	for _, egg := range eggs {
		if egg.GetHatchETA().Before(HatchETA) {
			HatchETA = egg.GetHatchETA()
		}

	}

	return HatchETA.Format("02-01-2006")
}

func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

func sortEggsIntoGrid(eggs []Egg) map[string]Egg {
	eggMap := make(map[string]Egg)

	for _, egg := range eggs {
		eggMap[strconv.Itoa(egg.IncubatorID)+","+strconv.Itoa(egg.Incubator.Row)+","+strconv.Itoa(egg.Incubator.Column)] = egg
	}

	return eggMap
}

func toSlotID(incubatorId int, row int, column int) string {
	return fmt.Sprintf("%s,%s,%s", strconv.Itoa(incubatorId), strconv.Itoa(row), strconv.Itoa(column))
}

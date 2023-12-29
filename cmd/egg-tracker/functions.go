package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetNextLayDateInfo() string {
	var LayETA time.Time = time.Now().Add(time.Hour * 999999)
	var geckoId int = 0
	for _, gecko := range geckos {
		if gecko.GetLayETA().Before(LayETA) {
			LayETA = gecko.GetLayETA()
			geckoId = gecko.ID
		}
	}

	if len(eggs) == 0 {
		return "No gecko eggs have been recorded"
	}

	return LayETA.Format("02-01-2006") + " (gecko " + strconv.Itoa(geckoId) + ")"
}

func GetNextHatchDateInfo() string {
	var HatchETA time.Time = time.Now().Add(time.Hour * 999999)
	var unhatchedEggsFound = false
	for _, egg := range eggs {
		if egg.GetHatchETA().Before(HatchETA) && !egg.HasHatched {
			HatchETA = egg.GetHatchETA()
		}
		if !egg.HasHatched {
			unhatchedEggsFound = true
		}
	}

	if len(eggs) == 0 {
		return "No gecko eggs have been recorded"
	}

	// check if any eggs waiting to hatch
	if !unhatchedEggsFound {
		return "No eggs waiting to hatch"
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

func DaysToHours(days string) string {
	iDays, err := strconv.Atoi(strings.TrimSuffix(days, "d"))
	if err != nil {
		log.Println(err)
		return ""
	}
	iHours := iDays * 24

	return strconv.Itoa(iHours) + "h"
}

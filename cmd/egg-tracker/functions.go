package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type NextLayDate struct {
	Value  string
	Colour string
}

func GetNextLayDateInfo() NextLayDate {
	var LayETA time.Time = time.Now().Add(time.Hour * 999999)
	var nextLayDate NextLayDate
	nextLayDate.Colour = "000000" // black
	var geckoId int = 0
	for _, gecko := range geckos {
		if gecko.GetLayETA().Before(LayETA) {
			LayETA = gecko.GetLayETA()
			geckoId = gecko.ID
		}
	}

	if len(eggs) == 0 {
		nextLayDate.Value = "No gecko eggs have been recorded"
		return nextLayDate
	}

	if time.Now().After(LayETA) {
		// if lay date is after current date
		// make it red
		nextLayDate.Colour = "#FF0000"
	} else if time.Now().Add(time.Hour * (2 * 24)).After(LayETA) {
		// if lay date is in the next 2 days
		// make it amber
		nextLayDate.Colour = "#FF9B00"
	} else {
		// normal colour
		nextLayDate.Colour = "000000" // black
	}

	nextLayDate.Value = LayETA.Format("02-01-2006") + " (gecko " + strconv.Itoa(geckoId) + ")"

	return nextLayDate
}

type NextHatchDate struct {
	Value  string
	Colour string
}

func GetNextHatchDateInfo() NextHatchDate {
	var HatchETA time.Time = time.Now().Add(time.Hour * 999999)
	var nextHatchDate NextHatchDate
	nextHatchDate.Colour = "000000" // black
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
		nextHatchDate.Value = "No gecko eggs have been recorded"
		return nextHatchDate
	}

	// check if any eggs waiting to hatch
	if !unhatchedEggsFound {
		nextHatchDate.Value = "No eggs waiting to hatch"
		return nextHatchDate
	}

	if time.Now().After(HatchETA) {
		// if hatch date is after current date
		// make it red
		nextHatchDate.Colour = "#FF0000"
	} else if time.Now().Add(time.Hour * (2 * 24)).After(HatchETA) {
		// if hatch date is in the next 2 days
		// make it amber
		nextHatchDate.Colour = "#FF9B00"
	} else {
		// normal colour
		nextHatchDate.Colour = "000000" // black
	}

	nextHatchDate.Value = HatchETA.Format("02-01-2006")

	return nextHatchDate
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

func GetEggRAG(egg Egg) string {
	if time.Now().After(egg.GetHatchETA()) {
		// if hatch date is after current date
		// make it red
		return "#FF0000"
	} else if time.Now().Add(time.Hour * (2 * 24)).After(egg.GetHatchETA()) {
		// if hatch date is in the next 2 days
		// make it amber
		return "#FF9B00"
	} else {
		// normal colour
		return "#008721" // green
	}
}

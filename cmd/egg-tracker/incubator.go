package main

import (
	"log"
	"strconv"
)

type Incubator struct {
	ID      int `json:"id"`
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

func AddIncubator(rows int, columns int) *Incubator {
	var incubator Incubator
	incubator.ID = len(incubators) + 1
	incubator.Rows = rows
	incubator.Columns = columns
	incubators = append(incubators, &incubator)

	log.Println("Added new incubator, " + strconv.Itoa(incubator.ID) + ", size " + strconv.Itoa(incubator.Columns) + " x " + strconv.Itoa(incubator.Rows))

	WriteToDB("incubator", Gecko{}, incubator, Egg{}, Sale{})

	return &incubator
}

func LoadIncubator(id int, rows int, columns int) *Incubator {
	var incubator Incubator
	incubator.ID = id
	incubator.Rows = rows
	incubator.Columns = columns
	incubators = append(incubators, &incubator)

	log.Println("Loaded incubator, " + strconv.Itoa(incubator.ID) + ", size " + strconv.Itoa(incubator.Columns) + " x " + strconv.Itoa(incubator.Rows))

	return &incubator
}

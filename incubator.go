package main

import (
	"log"
	"strconv"
)

type Incubator struct {
	ID      int
	Rows    int
	Columns int
}

func AddIncubator(rows int, columns int) *Incubator {
	var incubator Incubator
	incubator.ID = len(incubators) + 1
	incubator.Rows = rows
	incubator.Columns = columns
	incubators = append(incubators, incubator)

	log.Println("Added new incubator, " + strconv.Itoa(incubator.ID) + ", size " + strconv.Itoa(incubator.Columns) + " x " + strconv.Itoa(incubator.Rows))

	return &incubator
}
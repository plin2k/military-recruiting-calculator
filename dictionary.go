package main

import (
	"github.com/pkg/errors"
	"strings"
)

const (
	groupA = iota // А
	groupB        // Б
	groupC        // В
	groupD        // Г
	groupE        // Д

	specialityMedicine = iota // Медицина
	specialityIT              // IT
	specialityArmy            // Army
	specialityFactory         // factory

	defaultAirport = "SVO" // Sheremetyevo
)

var airports map[string]int

func (c *calculator) loadAirports(filePath string) error {

	records, err := loadCSVFile(filePath)
	if err != nil {
		return errors.Wrap(err, "can't load airports")
	}

	airports = make(map[string]int, len(records))

	for i, row := range records {
		airports[row[0]] = i
	}

	return nil
}

func getGroupFloat64(group string) float64 {
	switch strings.ToUpper(group) {
	case "А":
		return groupA
	case "Б":
		return groupB
	case "В":
		return groupC
	case "Г":
		return groupD
	default:
		return groupE
	}
}

func getAirportFloat64(airport string) float64 {
	if v, ok := airports[airport]; ok {
		return float64(v)
	}

	return float64(airports[defaultAirport])
}

func getSpecialityFloat64(speciality string) float64 {
	switch strings.ToUpper(speciality) {
	case "MED":
		return specialityMedicine
	case "FACTORY":
		return specialityFactory
	case "ARMY":
		return specialityArmy
	default:
		return specialityIT
	}
}

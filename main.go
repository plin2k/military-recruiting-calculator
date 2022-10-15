package main

import (
	"flag"
	"log"
)

func main() {

	var (
		flagServed, flagDepartmental, flagInvited                 bool
		flagCategory                                              int
		flagGroup, flagSpeciality, flagDeparture, flagDestination string
	)

	flag.BoolVar(&flagServed, "served", false, "Служил?")
	flag.BoolVar(&flagDepartmental, "departmental", false, "Военная кафедра?")
	flag.BoolVar(&flagInvited, "invited", false, "Пришла повестка?")
	flag.IntVar(&flagCategory, "category", 1, "Категория запаса (11 страница)")
	flag.StringVar(&flagGroup, "group", "А", "Группа учета (11 страница) (А/Б/В/Г/Д)")
	flag.StringVar(&flagSpeciality, "speciality", "IT", "Специальность (IT, MED, Factory, Army)")
	flag.StringVar(&flagDeparture, "departure", "SVO", "IATA (Код) Аэропорта вылета")
	flag.StringVar(&flagDestination, "destination", "SVO", "IATA (Код) Аэропорта прилета")

	flag.Parse()

	calc, err := NewNN("dataset.csv")
	if err != nil {
		panic(err)
	}

	log.Println(calc.Forward(
		flagServed,
		flagDepartmental,
		flagInvited,
		flagCategory,
		flagGroup,
		flagSpeciality,
		flagDeparture,
		flagDestination,
	))

}

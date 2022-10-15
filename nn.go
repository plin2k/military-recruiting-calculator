package main

import (
	"fmt"
	"github.com/fxsjy/gonn/gonn"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

type calculator struct {
	nn     *gonn.NeuralNetwork
	input  [][]float64
	target [][]float64
}

func New(path string) (*calculator, error) {
	var c = &calculator{
		nn:     gonn.DefaultNetwork(8, 16, 2, false),
		input:  [][]float64{},
		target: [][]float64{},
	}

	if !strings.Contains(strings.ToLower(path), ".csv") {
		return nil, errors.New("only CSV files")
	}

	if err := c.LoadCSVDataset(path); err != nil {
		return nil, errors.Wrap(err, "can't load CSV dataset'")
	}

	return c, nil
}

func (c *calculator) Load(path string) {
	c.nn = gonn.LoadNN(path)
}

func (c *calculator) Forward(served bool, departmental bool, invited bool, category int, group string, speciality, departure, destination string) string {
	result := c.nn.Forward(
		constructCase(served, departmental, invited, category, group, speciality, departure, destination))

	var max float64 = -99999
	pos := -1
	// Ищем позицию нейрона с самым большим весом.
	for i, value := range result {
		if value > max {
			max = value
			pos = i
		}
	}

	// Теперь, в зависимости от позиции, возвращаем решение.
	switch pos {
	case 0:
		return "Не призовут"
	default:
		return "Призовут"
	}
}

// AddCase
// Служил?
// Военная кафедра?
// Приходила повестка?
// Категория запаса (11 страница) (1/2/3)
// Группа учета (11 страница) (А/Б/В/Г/Д)
// Специальность (IT, MED, Factory, Army)
// Аэропорт вылета (LED)
// Аэропорт прилета (MRV)
// Результат прохождения
func (c *calculator) AddCase(served bool, departmental bool, invited bool, category int, group string, speciality, departure, destination string, result bool) {

	c.input = append(c.input,
		constructCase(served, departmental, invited, category, group, speciality, departure, destination))

	c.addTarget(result)
}

func (c *calculator) addTarget(result bool) {
	switch result {
	case true:
		c.target = append(c.target, []float64{1, 0})
	case false:
		c.target = append(c.target, []float64{0, 1})
	}
}

func constructCase(served bool, departmental bool, invited bool, category int, group string, speciality, departure, destination string) []float64 {
	return []float64{
		boolToFloat64(served),
		boolToFloat64(departmental),
		boolToFloat64(invited),
		float64(category),
		getGroupFloat64(group),
		getSpecialityFloat64(speciality),
		getAirportFloat64(departure),
		getAirportFloat64(destination),
	}
}

func (c *calculator) Apply() error {
	fmt.Println(c.input)
	c.nn.Train(c.input, c.target, 100000)

	if err := os.MkdirAll("./dump", os.ModePerm); err != nil {
		return errors.New("can't create dump directory")
	}

	gonn.DumpNN(fmt.Sprintf("./dump/%s", time.Now().Format("2006-01-02T15:04:05")), c.nn)

	return nil
}

func boolToFloat64(in bool) float64 {
	if in {
		return 1
	}
	return 0
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
	switch strings.ToUpper(airport) {
	case "KUF":
		return ruKUF
	case "SVO":
		return ruSVO
	case "LED":
		return ruLED
	case "VKO":
		return ruVKO
	case "KRR":
		return ruKRR
	case "GZP":
		return turGZP
	case "ESB":
		return turESB
	case "AYT":
		return turAYT
	case "BJV":
		return turBJV
	case "DLM":
		return turDLM
	case "ADB":
		return turADB
	case "IST":
		return turIST
	default:
		return ruDME
	}
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
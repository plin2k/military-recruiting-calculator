package main

import (
	"encoding/csv"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func (c *calculator) LoadCSVDataset(filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return errors.Wrapf(err, "unable to read input file %s", filePath)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return errors.Wrapf(err, "unable to parse file as CSV for %s", filePath)
	}

	if err = c.loadCSVDatasetStrings(records); err != nil {
		return errors.Wrap(err, "LoadCSVDataset")
	}

	if err = c.Apply(); err != nil {
		return errors.Wrap(err, "unable to apply")
	}

	return nil
}

func (c *calculator) loadCSVDatasetStrings(records [][]string) error {
	var served, departmental, invited, result bool
	var category int64
	var err error

	for _, record := range records[1:] {

		if served, err = strconv.ParseBool(record[0]); err != nil {
			return errors.New("can't parse boolean of served")
		}

		if departmental, err = strconv.ParseBool(record[1]); err != nil {
			return errors.New("can't parse boolean of departmental")
		}

		if invited, err = strconv.ParseBool(record[2]); err != nil {
			return errors.New("can't parse boolean of invited")
		}

		if category, err = strconv.ParseInt(record[3], 10, 32); err != nil {
			return errors.New("can't parse int of category")
		}

		if result, err = strconv.ParseBool(record[8]); err != nil {
			return errors.New("can't parse boolean of result")
		}

		c.AddCase(
			served,
			departmental,
			invited,
			int(category),
			record[4],
			record[5],
			record[6],
			record[7],
			result,
		)
	}

	return nil
}

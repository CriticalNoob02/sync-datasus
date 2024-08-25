package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type DataFilterStruct struct {
	MonthPosition [2]int
	YearPosition  [2]int
}

type FilterType string

const (
	start      FilterType = "start"
	end        FilterType = "end"
	contain    FilterType = "contain"
	containAny FilterType = "containAny"
)

func FilterString(stringFilter string, fileName string, filterType FilterType) bool {
	switch filterType {
	case start:
		if strings.HasPrefix(fileName, stringFilter) {
			return true
		}
	case end:
		if strings.HasSuffix(fileName, stringFilter) {
			return true
		}
	case contain:
		if strings.Contains(fileName, stringFilter) {
			return true
		}
	case containAny:
		if strings.ContainsAny(fileName, stringFilter) {
			return true
		}
	}
	return false
}

func FilterDate(dateFilter string, fileName string, model DataFilterStruct) (bool, error) {
	layout := "06/01"
	baseFilter, err := time.Parse(layout, dateFilter)
	if err != nil {
		return false, errors.New(err.Error())
	}

	filterDate := fmt.Sprintf(
		"%s/%s",
		fileName[model.MonthPosition[0]:model.MonthPosition[1]],
		fileName[model.YearPosition[0]:model.YearPosition[1]],
	)
	listDate, err := time.Parse(layout, filterDate)
	if err != nil {
		return false, errors.New(err.Error())
	}

	if listDate.After(baseFilter) {
		return true, nil
	}

	return false, nil
}

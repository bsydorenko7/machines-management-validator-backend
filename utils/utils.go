package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const SPACE string = " "

const (
	HUMAN_NANOSECONDS  string = "nanoseconds"
	HUMAN_MICROSECONDS string = "microseconds"
	HUMAN_MILISECONDS  string = "milliseconds"
	HUMAN_SECONDS      string = "seconds"
	HUMAN_MINUTES      string = "minutes"
	HUMAN_HOURS        string = "hours"
	HUMAN_DAYS         string = "days"
	HUMAN_WEEKS        string = "weeks"
	HUMAN_MONTHS       string = "months"
	HUMAN_YEARS        string = "years"
)

const (
	NANOSECONDS_UNIT_KEY   string = "ns"
	MISCROSECONDS_UNIT_KEY string = "us"
	MILISECONDS_UNIT_KEY   string = "ms"
	SECONDS_UNIT_KEY       string = "s"
	MINUTES_UNIT_KEY       string = "m"
	HOURS_UNIT_KEY         string = "h"
)

const (
	DAYS_TO_HOURS_KEF   float64 = 24
	WEEKS_TO_HOURS_KEF  float64 = 168  //24*7=168
	MONTHS_TO_HOURS_KEF float64 = 720  //24*30=720
	YEARS_TO_HOURS_KEF  float64 = 8760 //24*365=8760
)

func GetAgeInNanoseconds(age string) (float64, error) {
	var err error
	numberAndTimeUnit := strings.SplitN(age, SPACE, 2)

	switch numberAndTimeUnit[1] {
	case HUMAN_NANOSECONDS:
		numberAndTimeUnit[1] = NANOSECONDS_UNIT_KEY
	case HUMAN_MICROSECONDS:
		numberAndTimeUnit[1] = MISCROSECONDS_UNIT_KEY
	case HUMAN_MILISECONDS:
		numberAndTimeUnit[1] = MILISECONDS_UNIT_KEY
	case HUMAN_SECONDS:
		numberAndTimeUnit[1] = SECONDS_UNIT_KEY
	case HUMAN_MINUTES:
		numberAndTimeUnit[1] = MINUTES_UNIT_KEY
	case HUMAN_HOURS:
		numberAndTimeUnit[1] = HOURS_UNIT_KEY
	case HUMAN_DAYS:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], DAYS_TO_HOURS_KEF)
		numberAndTimeUnit[1] = HOURS_UNIT_KEY
	case HUMAN_WEEKS:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], WEEKS_TO_HOURS_KEF)
		numberAndTimeUnit[1] = HOURS_UNIT_KEY
	case HUMAN_MONTHS:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], MONTHS_TO_HOURS_KEF)
		numberAndTimeUnit[1] = HOURS_UNIT_KEY
	case HUMAN_YEARS:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], YEARS_TO_HOURS_KEF)
		numberAndTimeUnit[1] = HOURS_UNIT_KEY
	default:
		err = errors.New("can't parse time unit")
	}

	duration, err := time.ParseDuration(numberAndTimeUnit[0] + numberAndTimeUnit[1])

	return float64(duration.Nanoseconds()), err
}

func toDays(number string, kef float64) (string, error) {
	daysInFloat64, err := strconv.ParseFloat(strings.TrimSpace(number), 64)
	return fmt.Sprintf("%v", daysInFloat64*kef), err
}

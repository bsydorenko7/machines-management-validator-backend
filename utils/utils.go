package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	space string = " "

	humanNanoseconds  string = "nanoseconds"
	humanMicroseconds string = "microseconds"
	humanMilliseconds string = "milliseconds"
	humanSeconds      string = "seconds"
	humanMinutes      string = "minutes"
	humanHours        string = "hours"
	humanDays         string = "days"
	humanWeeks        string = "weeks"
	humanMonths       string = "months"
	humanYears        string = "years"

	nanosecondsUnitKey  string = "ns"
	microsecondsUnitKey string = "us"
	millisecondsUnitKey string = "ms"
	secondsUnitKey      string = "s"
	minutesUnitKey      string = "m"
	hoursUnitKey        string = "h"

	daysToHoursKef   float64 = 24
	weeksToHoursKef  float64 = 168  //24*7=168
	monthsToHoursKef float64 = 720  //24*30=720
	yearsToHoursKef  float64 = 8760 //24*365=8760
)

func GetAgeInNanoseconds(age string) (float64, error) {
	var err error
	numberAndTimeUnit := strings.SplitN(age, space, 2)

	switch numberAndTimeUnit[1] {
	case humanNanoseconds:
		numberAndTimeUnit[1] = nanosecondsUnitKey
	case humanMicroseconds:
		numberAndTimeUnit[1] = microsecondsUnitKey
	case humanMilliseconds:
		numberAndTimeUnit[1] = millisecondsUnitKey
	case humanSeconds:
		numberAndTimeUnit[1] = secondsUnitKey
	case humanMinutes:
		numberAndTimeUnit[1] = minutesUnitKey
	case humanHours:
		numberAndTimeUnit[1] = hoursUnitKey
	case humanDays:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], daysToHoursKef)
		numberAndTimeUnit[1] = hoursUnitKey
	case humanWeeks:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], weeksToHoursKef)
		numberAndTimeUnit[1] = hoursUnitKey
	case humanMonths:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], monthsToHoursKef)
		numberAndTimeUnit[1] = hoursUnitKey
	case humanYears:
		numberAndTimeUnit[0], err = toDays(numberAndTimeUnit[0], yearsToHoursKef)
		numberAndTimeUnit[1] = hoursUnitKey
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

package parsers

import "strconv"

func ParseTemperatureFloat(temperature string) (float64, error) {
	return strconv.ParseFloat(temperature, 64)
}

func ParseTemperatureInt(temperature string) (int64, error) {
	return strconv.ParseInt(temperature[:len(temperature)-2]+temperature[len(temperature)-1:], 10, 64)
}

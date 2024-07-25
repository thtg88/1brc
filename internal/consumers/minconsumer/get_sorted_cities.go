package minconsumer

import "slices"

func (mtc *MinTempConsumer) GetSortedCities() []string {
	cities := mtc.GetCities()

	slices.Sort(cities)

	return cities
}

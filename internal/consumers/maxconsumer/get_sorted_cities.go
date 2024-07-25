package maxconsumer

import "slices"

func (mtc *MaxTempConsumer) GetSortedCities() []string {
	cities := mtc.GetCities()

	slices.Sort(cities)

	return cities
}

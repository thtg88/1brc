package avgconsumer

import "slices"

func (atc *AvgTempConsumer) GetSortedCities() []string {
	cities := atc.GetCities()

	slices.Sort(cities)

	return cities
}

package sequentialconsumer

import "slices"

func (sc *SequentialConsumer) GetSortedCities() []string {
	cities := sc.GetCities()

	slices.Sort(cities)

	return cities
}

package bufferedsequentialconsumer

import "slices"

func (bsc *BufferedSequentialConsumer) GetSortedCities() []string {
	cities := bsc.GetCities()

	slices.Sort(cities)

	return cities
}

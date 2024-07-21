package sequentialconsumer

func (sc *SequentialConsumer) GetCities() []string {
	stats := sc.Stats

	cities := make([]string, 0, len(stats))
	for k := range stats {
		cities = append(cities, k)
	}

	return cities
}

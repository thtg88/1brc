package sequentialconsumer

func (sc *SequentialConsumer) GetCities() []string {
	// No need to call GetStats here as we only need the key and are not interested in the average temp
	stats := sc.Stats

	cities := make([]string, 0, len(stats))
	for k := range stats {
		cities = append(cities, k)
	}

	return cities
}

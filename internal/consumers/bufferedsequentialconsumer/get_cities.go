package bufferedsequentialconsumer

func (bsc *BufferedSequentialConsumer) GetCities() []string {
	stats := bsc.Stats

	cities := make([]string, 0, len(stats))
	for k := range stats {
		cities = append(cities, k)
	}

	return cities
}

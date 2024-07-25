package bufferedsequentialconsumer

func (bsc *BufferedSequentialConsumer) GetCities() []string {
	// No need to call GetStats here as we only need the key and are not interested in the average temp
	stats := bsc.Stats

	cities := make([]string, len(stats))
	idx := 0
	for k := range stats {
		cities[idx] = k
		idx++
	}

	return cities
}

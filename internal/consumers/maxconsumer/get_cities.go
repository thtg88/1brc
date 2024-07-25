package maxconsumer

func (mtc *MaxTempConsumer) GetCities() []string {
	stats := mtc.Stats

	cities := make([]string, len(stats))
	idx := 0
	for k := range stats {
		cities[idx] = k
		idx++
	}

	return cities
}

package minconsumer

func (mtc *MinTempConsumer) GetCities() []string {
	stats := mtc.Stats

	cities := make([]string, len(stats))
	idx := 0
	for k := range stats {
		cities[idx] = k
		idx++
	}

	return cities
}

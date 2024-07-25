package avgconsumer

func (atc *AvgTempConsumer) GetRecordsConsumed() uint64 {
	return atc.RecordsConsumed
}

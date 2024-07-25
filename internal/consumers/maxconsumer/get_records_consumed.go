package maxconsumer

func (mtc *MaxTempConsumer) GetRecordsConsumed() uint64 {
	return mtc.RecordsConsumed
}

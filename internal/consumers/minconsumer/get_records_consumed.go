package minconsumer

func (mtc *MinTempConsumer) GetRecordsConsumed() uint64 {
	return mtc.RecordsConsumed
}

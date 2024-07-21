package sequentialconsumer

func (sc *SequentialConsumer) GetRecordsConsumed() uint64 {
	return sc.RecordsConsumed
}

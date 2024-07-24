package bufferedsequentialconsumer

func (bsc *BufferedSequentialConsumer) GetRecordsConsumed() uint64 {
	return bsc.RecordsConsumed
}

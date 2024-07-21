package nilconsumer

func (nc *NilConsumer) GetRecordsConsumed() uint64 {
	return nc.RecordsConsumed
}

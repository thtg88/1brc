package bufferedproducer

func (mp *BufferedProducer) GetRecordsProduced() uint64 {
	return mp.RecordsProduced
}

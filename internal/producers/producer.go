package producers

type Producer interface {
	GetRecordsProduced() uint64
	Start()
}

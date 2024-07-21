package csvreadermock

type CSVReaderMock struct{}

func NewCSVReaderMock() *CSVReaderMock {
	return &CSVReaderMock{}
}

func (crm *CSVReaderMock) Read() (record []string, err error) {
	return []string{}, nil
}

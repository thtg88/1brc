package resultwritermock

import "github.com/thtg88/1brc/internal/models"

type WriterMock struct {
	rows uint64
}

func NewWriterMock() *WriterMock {
	return &WriterMock{}
}

func (wm *WriterMock) Write(sortedStats []models.CityStats) error {
	wm.rows++
	return nil
}

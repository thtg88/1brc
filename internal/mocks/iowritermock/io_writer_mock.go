package iowritermock

import "sync"

type IOWriterMock struct {
	sync.Mutex
	Bytes []byte
}

func NewIOWriterMock() *IOWriterMock {
	return &IOWriterMock{}
}

func (w *IOWriterMock) Write(p []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	w.Bytes = append(w.Bytes, p...)

	return len(p), nil
}

func (w *IOWriterMock) GetBytes() []byte {
	return w.Bytes
}

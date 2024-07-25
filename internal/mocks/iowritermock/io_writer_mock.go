package iowritermock

import "sync"

type IOWriterMock struct {
	sync.Mutex
	bytes []byte
}

func NewIOWriterMock() *IOWriterMock {
	return &IOWriterMock{}
}

func (w *IOWriterMock) Write(p []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	w.bytes = append(w.bytes, p...)

	return len(p), nil
}

func (w *IOWriterMock) GetBytes() []byte {
	return w.bytes
}

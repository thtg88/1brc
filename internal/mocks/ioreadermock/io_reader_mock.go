package ioreadermock

import (
	"errors"
	"io"
)

type IOReaderMock struct {
	bytes      []byte
	forceEOF   bool
	forceError bool
	position   uint64
}

func NewIOReaderMock(bytes []byte, position uint64, forceEOF bool, forceError bool) *IOReaderMock {
	return &IOReaderMock{
		bytes:      bytes,
		forceEOF:   forceEOF,
		forceError: forceError,
		position:   position,
	}
}

func (irm *IOReaderMock) Read(p []byte) (n int, err error) {
	if irm.forceEOF {
		return 0, io.EOF
	}
	if irm.forceError {
		return 0, errors.New("forced error on IO reader")
	}
	remainingBytesCount := len(irm.bytes) - int(irm.position)
	if len(p) == 0 {
		n = remainingBytesCount
	} else {
		n = min(len(p), remainingBytesCount)
	}

	endPosition := n + int(irm.position)

	copy(p, irm.bytes[irm.position:endPosition])

	irm.position = uint64(endPosition)

	return n, nil
}

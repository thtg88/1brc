package ioreadermock

import (
	"errors"
	"fmt"
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

func (irm *IOReaderMock) Seek(offset int64, whence int) (int64, error) {
	if irm.forceError {
		return 0, errors.New("forced error")
	}

	var newPosition int64

	length := int64(len(irm.bytes))

	switch whence {
	case io.SeekStart:
		newPosition = offset
	case io.SeekCurrent:
		newPosition = int64(irm.position) + offset
	case io.SeekEnd:
		newPosition = length + offset
	default:
		return 0, fmt.Errorf("invalid whence %d", whence)
	}

	if newPosition >= length {
		return 0, fmt.Errorf("newPosition %d greater than length %d", newPosition, length)
	}

	if newPosition < 0 {
		return 0, fmt.Errorf("newPosition %d less than 0", newPosition)
	}

	irm.position = uint64(newPosition)

	return 0, nil
}

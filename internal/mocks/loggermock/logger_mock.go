package loggermock

import "sync"

type LoggerMock struct {
	printfLock  sync.RWMutex
	printfCalls uint64

	printlnLock  sync.RWMutex
	printlnCalls uint64

	fatalLock  sync.RWMutex
	fatalCalls uint64

	fatalfLock  sync.RWMutex
	fatalfCalls uint64
}

func NewLoggerMock() *LoggerMock {
	return &LoggerMock{}
}

func (ml *LoggerMock) GetPrintfCalls() uint64 {
	ml.printfLock.RLock()
	defer ml.printfLock.RUnlock()

	return ml.printfCalls
}

func (ml *LoggerMock) GetPrintlnCalls() uint64 {
	ml.printlnLock.RLock()
	defer ml.printlnLock.RUnlock()

	return ml.printlnCalls
}

func (ml *LoggerMock) GetFatalCalls() uint64 {
	ml.fatalLock.RLock()
	defer ml.fatalLock.RUnlock()

	return ml.fatalCalls
}

func (ml *LoggerMock) GetFatalfCalls() uint64 {
	ml.fatalfLock.RLock()
	defer ml.fatalfLock.RUnlock()

	return ml.fatalfCalls
}

func (ml *LoggerMock) Printf(format string, v ...any) {
	ml.printfLock.Lock()
	defer ml.printfLock.Unlock()

	ml.printfCalls++
}

func (ml *LoggerMock) Println(v ...any) {
	ml.printlnLock.Lock()
	defer ml.printlnLock.Unlock()

	ml.printlnCalls++
}

func (ml *LoggerMock) Fatalf(format string, v ...any) {
	ml.fatalfLock.Lock()
	defer ml.fatalfLock.Unlock()

	ml.fatalfCalls++
}

func (ml *LoggerMock) Fatal(v ...any) {
	ml.fatalLock.Lock()
	defer ml.fatalLock.Unlock()

	ml.fatalCalls++
}

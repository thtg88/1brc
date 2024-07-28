package sequentialconsumer

import "time"

func (sc *SequentialConsumer) Start() {
	for {
		select {
		case reading, more := <-sc.DataChannel:
			if !more {
				sc.Logger.Println("consumer channel closed")

				return
			}

			sc.ProcessReading(reading)
		default:
			if sc.Config.Debug {
				sc.Logger.Println("waiting for readings to be produced")
			}

			time.Sleep(time.Duration(sc.Config.WaitingRecordsSleepDurationMs) * time.Millisecond)
		}
	}
}

package bufferedsequentialconsumer

import (
	"time"
)

func (bsc *BufferedSequentialConsumer) Start() {
	for {
		select {
		case readings, more := <-bsc.DataChannel:
			if !more {
				bsc.Logger.Println("consumer channel closed")

				return
			}

			for _, reading := range readings {
				bsc.ProcessReading(reading)
			}
		default:
			if bsc.Config.Debug {
				bsc.Logger.Println("waiting for readings to be produced")
			}

			time.Sleep(time.Duration(bsc.Config.WaitingRecordsSleepDurationMs) * time.Millisecond)
		}
	}
}

package sequentialconsumer

import "time"

func (sc *SequentialConsumer) Start() {
	for {
		select {
		case reading, more := <-sc.DataChannel:
			if !more {
				sc.Config.Progress.Enabled = false

				sc.Logger.Println("consumer channel closed")

				return
			}

			sc.ProcessReading(reading)
		case <-sc.DoneChannel:
			sc.Config.Progress.Enabled = false

			sc.Logger.Println("done channel closed")

			return
		default:
			if sc.Config.Debug {
				sc.Logger.Println("waiting for readings to be produced")
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

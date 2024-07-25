package minconsumer

func (mtc *MinTempConsumer) Start() {
	for {
		readings, more := <-mtc.DataChannel
		if !more {
			mtc.Logger.Println("consumer channel closed")

			break
		}

		for _, reading := range readings {
			mtc.ProcessReading(reading)
		}
	}
}

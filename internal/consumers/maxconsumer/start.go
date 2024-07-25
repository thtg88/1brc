package maxconsumer

func (mtc *MaxTempConsumer) Start() {
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

	mtc.Config.Progress.Enabled = false
	mtc.Logger.Println("done channel closed")
}

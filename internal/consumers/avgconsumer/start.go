package avgconsumer

func (atc *AvgTempConsumer) Start() {
	for {
		readings, more := <-atc.DataChannel
		if !more {
			atc.Logger.Println("consumer channel closed")

			break
		}

		for _, reading := range readings {
			atc.ProcessReading(reading)
		}
	}

	atc.Config.Progress.Enabled = false
	atc.Logger.Println("done channel closed")
}

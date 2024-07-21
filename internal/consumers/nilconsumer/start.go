package nilconsumer

func (nc *NilConsumer) Start() {
	for {
		reading, more := <-nc.DataChannel
		if !more {
			nc.Logger.Println("consumer channel closed")

			break
		}

		nc.ProcessReading(reading)
	}
}

package nilconsumer

import "github.com/thtg88/1brc/internal/models"

func (nc *NilConsumer) ProcessReading(reading *models.TemperatureReading) {
	nc.RecordsConsumed++
}

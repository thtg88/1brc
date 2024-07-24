package builders

type TemperatureReadingCSVRowBuilder struct {
	csvRow []string
}

const (
	CSVRowLength                                    = 2
	TemperatureReadingCSVRowBuilder_TestTemperature = "12.3"
)

func NewTemperatureReadingCSVRowBuilder() *TemperatureReadingCSVRowBuilder {
	return &TemperatureReadingCSVRowBuilder{
		csvRow: make([]string, CSVRowLength),
	}
}

func (trcrb *TemperatureReadingCSVRowBuilder) Build() []string {
	return trcrb.csvRow
}

func (trcrb *TemperatureReadingCSVRowBuilder) WithTestValues() *TemperatureReadingCSVRowBuilder {
	trcrb.csvRow[0] = TemperatureReadingBuilder_TestCity
	trcrb.csvRow[1] = TemperatureReadingCSVRowBuilder_TestTemperature
	return trcrb
}

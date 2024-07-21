package configs

type DataGeneratorConfig struct {
	DestinationFilePath string
	SourceFilePath      string
}

func NewDefaultDataGeneratorConfig() *DataGeneratorConfig {
	return &DataGeneratorConfig{
		DestinationFilePath: DefaultTemperaturesFilePath,
		SourceFilePath:      DefaultWeatherStationsFilePath,
	}
}

package configs

const (
	DefaultBufferedChannelSize            = 50_000
	DefaultCalculateAverageForEachReading = false
	DefaultDebug                          = false
	DefaultLimit                          = 999_999_000

	DefaultStatsFilePath           = "./data/stats.csv"
	DefaultTemperaturesFilePath    = "./data/temperatures.csv"
	DefaultWeatherStationsFilePath = "./data/weather_stations.csv"

	// Profiling
	DefaultProfileEnabled           = false
	DefaultProfileCPUFilePath       = "./profiles/cpuprofile"
	DefaultProfileExecutionFilePath = "./profiles/exeprofile"
	DefaultProfileMemoryFilePath    = "./profiles/memprofile"

	// Progress
	DefaultProgressEnabled         = true
	DefaultProgressSleepDurationMs = 10_000
)

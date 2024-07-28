package configs

type FilePositioningConfig struct {
	Enabled               bool
	ReadUntilFilePosition int64
}

type ProgressSolverConfig struct {
	SleepDurationMs uint64
	Enabled         bool
}

type ProfileSolverConfig struct {
	Enabled           bool
	CPUFilePath       string
	ExecutionFilePath string
	MemoryFilePath    string
}

type SolverConfig struct {
	BufferedChannelSize            uint64
	CalculateAverageForEachReading bool
	Debug                          bool
	DestinationFilePath            string
	FilePositioning                *FilePositioningConfig
	Limit                          uint64
	Profile                        *ProfileSolverConfig
	Progress                       *ProgressSolverConfig
	SourceFilePath                 string
	WaitingRecordsSleepDurationMs  uint64
}

func NewDefaultSolverConfig() *SolverConfig {
	return &SolverConfig{
		BufferedChannelSize:            DefaultBufferedChannelSize,
		CalculateAverageForEachReading: DefaultCalculateAverageForEachReading,
		Debug:                          DefaultDebug,
		DestinationFilePath:            DefaultStatsFilePath,
		Limit:                          DefaultLimit,
		SourceFilePath:                 DefaultTemperaturesFilePath,
		WaitingRecordsSleepDurationMs:  DefaultWaitingRecordsSleepDurationMs,
		FilePositioning: &FilePositioningConfig{
			Enabled: DefaultFilePositioningEnabled,
		},
		Profile: &ProfileSolverConfig{
			Enabled:           DefaultProfileEnabled,
			CPUFilePath:       DefaultProfileCPUFilePath,
			ExecutionFilePath: DefaultProfileExecutionFilePath,
			MemoryFilePath:    DefaultProfileMemoryFilePath,
		},
		Progress: &ProgressSolverConfig{
			Enabled:         DefaultProgressEnabled,
			SleepDurationMs: DefaultProgressSleepDurationMs,
		},
	}
}

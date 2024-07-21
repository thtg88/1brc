package configs

type ProgressSolverConfig struct {
	SleepDurationMs uint64
	ShouldReport    bool
}

type ProfileSolverConfig struct {
	Enabled           bool
	CPUFilePath       string
	ExecutionFilePath string
	MemoryFilePath    string
}

type SolverConfig struct {
	BufferedChannelSize uint64
	Debug               bool
	DestinationFilePath string
	Limit               uint64
	Profile             *ProfileSolverConfig
	Progress            *ProgressSolverConfig
	SourceFilePath      string
}

func NewDefaultSolverConfig() *SolverConfig {
	return &SolverConfig{
		BufferedChannelSize: DefaultBufferedChannelSize,
		Debug:               DefaultDebug,
		DestinationFilePath: DefaultStatsFilePath,
		Limit:               DefaultLimit,
		SourceFilePath:      DefaultTemperaturesFilePath,
		Profile: &ProfileSolverConfig{
			Enabled:           DefaultProfileEnabled,
			CPUFilePath:       DefaultProfileCPUFilePath,
			ExecutionFilePath: DefaultProfileExecutionFilePath,
			MemoryFilePath:    DefaultProfileMemoryFilePath,
		},
		Progress: &ProgressSolverConfig{
			SleepDurationMs: DefaultProgressSleepDurationMs,
			ShouldReport:    DefaultShouldReportProgress,
		},
	}
}

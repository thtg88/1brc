package configs_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/configs"
)

func TestNewDefaultSolverConfig(t *testing.T) {
	t.Parallel()

	wantConfig := &configs.SolverConfig{
		BufferedChannelSize:            configs.DefaultBufferedChannelSize,
		CalculateAverageForEachReading: configs.DefaultCalculateAverageForEachReading,
		Debug:                          configs.DefaultDebug,
		DestinationFilePath:            configs.DefaultStatsFilePath,
		Limit:                          configs.DefaultLimit,
		SourceFilePath:                 configs.DefaultTemperaturesFilePath,
		WaitingRecordsSleepDurationMs:  configs.DefaultWaitingRecordsSleepDurationMs,
		FilePositioning: &configs.FilePositioningConfig{
			Enabled: configs.DefaultFilePositioningEnabled,
		},
		Profile: &configs.ProfileSolverConfig{
			Enabled:           configs.DefaultProfileEnabled,
			CPUFilePath:       configs.DefaultProfileCPUFilePath,
			ExecutionFilePath: configs.DefaultProfileExecutionFilePath,
			MemoryFilePath:    configs.DefaultProfileMemoryFilePath,
		},
		Progress: &configs.ProgressSolverConfig{
			SleepDurationMs: configs.DefaultProgressSleepDurationMs,
			Enabled:         configs.DefaultProgressEnabled,
		},
	}

	actualConfig := configs.NewDefaultSolverConfig()

	require.Equal(t, wantConfig, actualConfig)
}

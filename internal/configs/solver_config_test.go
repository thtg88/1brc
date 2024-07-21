package configs_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/configs"
)

func TestNewDefaultSolverConfig(t *testing.T) {
	t.Parallel()

	wantConfig := &configs.SolverConfig{
		BufferedChannelSize: configs.DefaultBufferedChannelSize,
		Debug:               configs.DefaultDebug,
		DestinationFilePath: configs.DefaultStatsFilePath,
		Limit:               configs.DefaultLimit,
		SourceFilePath:      configs.DefaultTemperaturesFilePath,
		Profile: &configs.ProfileSolverConfig{
			Enabled:           true,
			CPUFilePath:       configs.DefaultProfileCPUFilePath,
			ExecutionFilePath: configs.DefaultProfileExecutionFilePath,
			MemoryFilePath:    configs.DefaultProfileMemoryFilePath,
		},
		Progress: &configs.ProgressSolverConfig{
			SleepDurationMs: configs.DefaultProgressSleepDurationMs,
			ShouldReport:    configs.DefaultShouldReportProgress,
		},
	}

	actualConfig := configs.NewDefaultSolverConfig()

	require.Equal(t, wantConfig, actualConfig)
}
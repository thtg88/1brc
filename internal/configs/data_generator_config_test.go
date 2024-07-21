package configs_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thtg88/1brc/internal/configs"
)

func TestNewDefaultDataGeneratorConfig(t *testing.T) {
	t.Parallel()

	expectedConfig := &configs.DataGeneratorConfig{
		DestinationFilePath: configs.DefaultTemperaturesFilePath,
		SourceFilePath:      configs.DefaultWeatherStationsFilePath,
	}

	actualConfig := configs.NewDefaultDataGeneratorConfig()

	require.Equal(t, expectedConfig, actualConfig)
}

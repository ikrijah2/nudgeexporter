package nudgeexporter

import (

	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for file exporter.
type Config struct {
	config.ExporterSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
}

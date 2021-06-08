// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package obsreport

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/ikrijah2/nudge/config/configtelemetry"
)

const (
	// ProcessorKey is the key used to identify processors in metrics and traces.
	ProcessorKey = "processor"

	// DroppedSpansKey is the key used to identify spans dropped by the Collector.
	DroppedSpansKey = "dropped_spans"

	// DroppedMetricPointsKey is the key used to identify metric points dropped by the Collector.
	DroppedMetricPointsKey = "dropped_metric_points"

	// DroppedLogRecordsKey is the key used to identify log records dropped by the Collector.
	DroppedLogRecordsKey = "dropped_log_records"
)

var (
	tagKeyProcessor, _ = tag.NewKey(ProcessorKey)

	processorPrefix = ProcessorKey + nameSep

	// Processor metrics. Any count of data items below is in the internal format
	// of the collector since processors only deal with internal format.
	mProcessorAcceptedSpans = stats.Int64(
		processorPrefix+AcceptedSpansKey,
		"Number of spans successfully pushed into the next component in the pipeline.",
		stats.UnitDimensionless)
	mProcessorRefusedSpans = stats.Int64(
		processorPrefix+RefusedSpansKey,
		"Number of spans that were rejected by the next component in the pipeline.",
		stats.UnitDimensionless)
	mProcessorDroppedSpans = stats.Int64(
		processorPrefix+DroppedSpansKey,
		"Number of spans that were dropped.",
		stats.UnitDimensionless)
	mProcessorAcceptedMetricPoints = stats.Int64(
		processorPrefix+AcceptedMetricPointsKey,
		"Number of metric points successfully pushed into the next component in the pipeline.",
		stats.UnitDimensionless)
	mProcessorRefusedMetricPoints = stats.Int64(
		processorPrefix+RefusedMetricPointsKey,
		"Number of metric points that were rejected by the next component in the pipeline.",
		stats.UnitDimensionless)
	mProcessorDroppedMetricPoints = stats.Int64(
		processorPrefix+DroppedMetricPointsKey,
		"Number of metric points that were dropped.",
		stats.UnitDimensionless)
	mProcessorAcceptedLogRecords = stats.Int64(
		processorPrefix+AcceptedLogRecordsKey,
		"Number of log records successfully pushed into the next component in the pipeline.",
		stats.UnitDimensionless)
	mProcessorRefusedLogRecords = stats.Int64(
		processorPrefix+RefusedLogRecordsKey,
		"Number of log records that were rejected by the next component in the pipeline.",
		stats.UnitDimensionless)
	mProcessorDroppedLogRecords = stats.Int64(
		processorPrefix+DroppedLogRecordsKey,
		"Number of log records that were dropped.",
		stats.UnitDimensionless)
)

// BuildProcessorCustomMetricName is used to be build a metric name following
// the standards used in the Collector. The configType should be the same
// value used to identify the type on the config.
func BuildProcessorCustomMetricName(configType, metric string) string {
	return buildComponentPrefix(processorPrefix, configType) + metric
}

// ProcessorMetricViews builds the metric views for custom metrics of processors.
func ProcessorMetricViews(configType string, legacyViews []*view.View) []*view.View {
	var allViews []*view.View
	if gLevel != configtelemetry.LevelNone {
		for _, legacyView := range legacyViews {
			// Ignore any nil entry and views without measure or aggregation.
			// These can't be registered but some code registering legacy views may
			// ignore the errors.
			if legacyView == nil || legacyView.Measure == nil || legacyView.Aggregation == nil {
				continue
			}
			newView := *legacyView
			viewName := legacyView.Name
			if viewName == "" {
				viewName = legacyView.Measure.Name()
			}
			newView.Name = BuildProcessorCustomMetricName(configType, viewName)
			allViews = append(allViews, &newView)
		}
	}

	return allViews
}

var gProcessor = &Processor{level: configtelemetry.LevelNone}

// Processor is a helper to add observability to a component.Processor.
type Processor struct {
	level    configtelemetry.Level
	mutators []tag.Mutator
}

// ProcessorSettings are settings for creating a Processor.
type ProcessorSettings struct {
	Level         configtelemetry.Level
	ProcessorName string
}

// NewProcessor creates a new Processor.
func NewProcessor(cfg ProcessorSettings) *Processor {
	return &Processor{
		level:    cfg.Level,
		mutators: []tag.Mutator{tag.Upsert(tagKeyProcessor, cfg.ProcessorName, tag.WithTTL(tag.TTLNoPropagation))},
	}
}

// TracesAccepted reports that the trace data was accepted.
func (por *Processor) TracesAccepted(ctx context.Context, numSpans int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedSpans.M(int64(numSpans)),
			mProcessorRefusedSpans.M(0),
			mProcessorDroppedSpans.M(0),
		)
	}
}

// TracesRefused reports that the trace data was refused.
func (por *Processor) TracesRefused(ctx context.Context, numSpans int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedSpans.M(0),
			mProcessorRefusedSpans.M(int64(numSpans)),
			mProcessorDroppedSpans.M(0),
		)
	}
}

// TracesDropped reports that the trace data was dropped.
func (por *Processor) TracesDropped(ctx context.Context, numSpans int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedSpans.M(0),
			mProcessorRefusedSpans.M(0),
			mProcessorDroppedSpans.M(int64(numSpans)),
		)
	}
}

// MetricsAccepted reports that the metrics were accepted.
func (por *Processor) MetricsAccepted(ctx context.Context, numPoints int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedMetricPoints.M(int64(numPoints)),
			mProcessorRefusedMetricPoints.M(0),
			mProcessorDroppedMetricPoints.M(0),
		)
	}
}

// MetricsRefused reports that the metrics were refused.
func (por *Processor) MetricsRefused(ctx context.Context, numPoints int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedMetricPoints.M(0),
			mProcessorRefusedMetricPoints.M(int64(numPoints)),
			mProcessorDroppedMetricPoints.M(0),
		)
	}
}

// MetricsDropped reports that the metrics were dropped.
func (por *Processor) MetricsDropped(ctx context.Context, numPoints int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedMetricPoints.M(0),
			mProcessorRefusedMetricPoints.M(0),
			mProcessorDroppedMetricPoints.M(int64(numPoints)),
		)
	}
}

// LogsAccepted reports that the logs were accepted.
func (por *Processor) LogsAccepted(ctx context.Context, numRecords int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedLogRecords.M(int64(numRecords)),
			mProcessorRefusedLogRecords.M(0),
			mProcessorDroppedLogRecords.M(0),
		)
	}
}

// LogsRefused reports that the logs were refused.
func (por *Processor) LogsRefused(ctx context.Context, numRecords int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedLogRecords.M(0),
			mProcessorRefusedLogRecords.M(int64(numRecords)),
			mProcessorDroppedMetricPoints.M(0),
		)
	}
}

// LogsDropped reports that the logs were dropped.
func (por *Processor) LogsDropped(ctx context.Context, numRecords int) {
	if por.level != configtelemetry.LevelNone {
		stats.RecordWithTags(
			ctx,
			por.mutators,
			mProcessorAcceptedLogRecords.M(0),
			mProcessorRefusedLogRecords.M(0),
			mProcessorDroppedLogRecords.M(int64(numRecords)),
		)
	}
}

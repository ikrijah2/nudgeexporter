package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/service/defaultcomponents"

	//nudgeexporter "github.com/ikrijah2/nudge/nudgeexporter"
)


func components() (
	component.Factories,
	error,
) {
	var errs []error

	factories, err := defaultcomponents.Components()
	if err != nil {
		errs = append(errs, err)
	}

	exporters := []component.ExporterFactory{
	//	nudgeexporter.NewFactory(),
	}

	for _, pr := range factories.Exporters {
		exporters = append(exporters, pr)
	}

	factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)
	if err != nil {
		errs = append(errs, err)
	}

	return factories, consumererror.Combine(errs)
}
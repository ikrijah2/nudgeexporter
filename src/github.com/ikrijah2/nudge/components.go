package main

import (
	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/consumer/consumererror"
	"github.com/ikrijah2/nudge/service/defaultcomponents"

//"github.com/ikrijah2/nudge/exporter/fileexporter"

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
		//fileexporter.NewFactory(),
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
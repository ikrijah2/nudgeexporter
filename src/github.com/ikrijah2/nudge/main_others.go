
package main

import "go.opentelemetry.io/collector/service"

func run(params service.Parameters) error {
	return runInteractive(params)
}

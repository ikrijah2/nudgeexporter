package nudgeexporter

import (

	
)


// Se comporte comme le file exporter et écrit les données de télémétrie dans un
// fichier au format Protobuf-JSON
type fileExporter struct {
	path  string
	file  io.WriteCloser
	mutex sync.Mutex
}
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

package nudgeexporter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/ikrijah2/nudge/rawdata"

	"github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/internal"
)

var marshaler = &jsonpb.Marshaler{}

// Marshaler configuration used for marhsaling Protobuf to JSON. Use default config.
var spanCount int

// fileExporter is the implementation of file exporter that writes telemetry data to a file
// in Protobuf-JSON format.
type fileExporter struct {
	file  io.WriteCloser
	mutex sync.Mutex
}

func (e *fileExporter) ConsumeTraces(_ context.Context, td pdata.Traces) error {
	spanCount = td.SpanCount()
	//fmt.Printf("nombres de spans %v", td.SpanCount())

	return exportMessageAsLine(e, internal.TracesToOtlp(td.InternalRep()))
}

func (e *fileExporter) ConsumeMetrics(_ context.Context, md pdata.Metrics) error {
	return exportMessageAsLine(e, internal.MetricsToOtlp(md.InternalRep()))
}

func (e *fileExporter) ConsumeLogs(_ context.Context, ld pdata.Logs) error {
	request := internal.LogsToOtlp(ld.InternalRep())
	return exportMessageAsLine(e, request)
}

func exportMessageAsLine(e *fileExporter, message proto.Message) error {
	// Ensure only one write operation happens at a time.
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// e.file correspond au nom donn?? dans config.yml lors du lancement d'otel collector.
	// Serialisation du message protobuf en JSON. (??criture d'une ligne JSON dans le file)
	if err := marshaler.Marshal(e.file, message); err != nil {
		return err
	}

	var traces AutoGenerated

	var buf bytes.Buffer

	// Marshal serializes a protobuf message as JSON into buf.
	if err := marshaler.Marshal(&buf, message); err != nil {
		return err
	}

	var str = buf.String()

	// Unmarshal unmarshals a JSON object from []byte(str) into &traces.
	if err := json.Unmarshal([]byte(str), &traces); err != nil {
		return err
	}

	// Utilisation du rawdata.pb.go
	var rd = rawdata.RawData{}

	// Population des donn??es ?? partir de la structure Autogenerated
	entryPoint(traces, &rd)

	// Marshal serializes a protobuf message as JSON into buf.
	if err := marshaler.Marshal(&buf, message); err != nil {
		return err
	}

	// Encode and write rawdata to file.
	// WriteAddRawdata(&rd)

	// // Read the existing rawdata file.
	// in, err := ioutil.ReadFile("rawdata.dat")
	// if err != nil {
	// 	log.Fatalln("Error reading file:", err)
	// }

	// // Pour voir si les traces sont bien export??es
	// rawdataUnmarshal := rawdata.RawData{}
	// if err := googleproto.Unmarshal(in, &rawdataUnmarshal); err != nil {
	// 	log.Fatalln("Failed to parse rawdata file:", err)
	// }
	//var m = int(*rawdataUnmarshal.Id)
	//fmt.Printf("id lecture unmarshal  %v \n AgentID : %v \n", m, *rawdataUnmarshal.AgentId)

	// Envoi
	SendRawdataToNudge()

	if _, err := io.WriteString(e.file, "\n"); err != nil {
		return err
	}
	return nil
}

func (e *fileExporter) Start(context.Context, component.Host) error {
	return nil
}

// Shutdown stops the exporter and is invoked during shutdown.
func (e *fileExporter) Shutdown(context.Context) error {
	return e.file.Close()
}

/* fonction qui envoie les rawdata au format binaire dans le serveur nudge sur l'adresse
   http://localhost:8080/nudge-war/collect/rawdata/<uuid>
   o?? <uuid> correspond ?? l'id de l'app
*/

func SendRawdataToNudge() {

	file, _ := os.Open("rawdata.dat")

	defer file.Close()

	uuid, err := ioutil.ReadFile("uuid")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	uuidStr := string(uuid)

	url := "http://nudge-server:8080/collect/rawdata/" + uuidStr

	r, err := http.NewRequest("PUT", url, file)
	if err != nil {
		fmt.Printf("Erreur http : %v \n", err)
	}

	r.Header.Set("Content-Type", "text/pain")

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Print("Erreur r??ponse : %v \n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erreur body r??ponse : %v \n", err)
	}

	fmt.Println("voici le body de la r??ponse : %v \n", string(body))

}

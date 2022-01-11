package nudgeexporter

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ikrijah2/nudge/rawdata"
)

type Span struct {
	TraceID           string `json:"traceId"`
	SpanID            string `json:"spanId"`
	ParentSpanID      string `json:"parentSpanId"`
	Name              string `json:"name"`
	Kind              string `json:"kind"`
	StartTimeUnixNano string `json:"startTimeUnixNano"`
	EndTimeUnixNano   string `json:"endTimeUnixNano"`
	Attributes        []struct {
		Key   string `json:"key"`
		Value struct {
			IntValue    string `json:"intValue,omitempty"`
			StringValue string `json:"stringValue,omitempty"`
		} `json:"value"`
	} `json:"attributes"`
	Status struct {
	} `json:"status,omitempty"`
}

func sqlToRawdata(span Span, rd *rawdata.RawData) {

	uuid, err := ioutil.ReadFile("uuid")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	uuidStr := string(uuid)
	uuidPtr := &uuidStr

	var Transactions []*rawdata.Transaction
	var Layers []*rawdata.Layer
	var Calls []*rawdata.LayerDetail

	var Transaction rawdata.Transaction
	var Layer rawdata.Layer
	var Call rawdata.LayerDetail

	var extCodesTable []*rawdata.KeyValue
	var extCodes rawdata.KeyValue
	var dbname string

	Layer.LayerName = PtrCreateAndAssign("SQL")
	Layer.Count = PtrInt64CreateAndAssign(int64(1))

	// faire la différence endtime - starttime
	Layer.Time = PtrInt64CreateAndAssign(int64(1))

	var StartTimePointer *int64 = TimeToInt64(span.StartTimeUnixNano)
	var EndTimePointer *int64 = TimeToInt64(span.EndTimeUnixNano)

	// Transaction.Type = TransactionType
	Transaction.StartTime = StartTimePointer
	Transaction.EndTime = EndTimePointer
	Transaction.Uuid = uuidPtr

	// TODO MAPPING
	//Transaction.TxType = PtrCreateAndAssign(traces.ResourceSpans[0].Resource.Attributes[0].Value.StringValue)

	Call.Timestamp = TimeToFloat64(span.StartTimeUnixNano)

	Call.Count = PtrInt64CreateAndAssign(int64(1))
	Call.Errors = PtrInt64CreateAndAssign(int64(0))

	// faire la différence endtime - starttime
	Call.Time = PtrInt64CreateAndAssign(int64(1))

	var UrlPointer *string = PtrCreateAndAssign(url)
	fmt.Printf("voici l'url dans le sql ! %s", url)
	Transaction.Url = UrlPointer

	for _, attribute := range span.Attributes {

		if attribute.Key == "sql-query" {
			Call.Code = PtrCreateAndAssign(attribute.Value.StringValue)
		}

		if attribute.Key == "db.name" {
			dbname = attribute.Value.StringValue
		}
	}

	extCodes.Key = PtrCreateAndAssign("jdbc.url")
	extCodes.Value = PtrCreateAndAssign(dbname)
	extCodesPtr := &extCodes
	extCodesTable = append(extCodesTable, extCodesPtr)
	Call.ExtCodes = extCodesTable

	var CallPtr *rawdata.LayerDetail
	CallPtr = &Call
	Calls = append(Calls, CallPtr)

	Layer.Calls = Calls

	var LayerPtr = &Layer
	Layers = append(Layers, LayerPtr)

	Transaction.Layers = Layers

	// On créé le pointeur qui va pointer sur Transaction
	var TransactionPtr *rawdata.Transaction

	TransactionPtr = &Transaction

	// On ajoute la Transaction dans le tableau des Transactions
	Transactions = append(Transactions, TransactionPtr)

	rd.Transactions = Transactions

	WriteAddRawdata(rd)

}

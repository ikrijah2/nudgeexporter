package nudgeexporter

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/ikrijah2/nudge/rawdata"
)

func httpToRawdata(span Span, rd *rawdata.RawData) {

	uuid, err := ioutil.ReadFile("uuid")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	uuidStr := string(uuid)
	uuidPtr := &uuidStr

	var Transactions []*rawdata.Transaction
	var ExtendedStatusCodes []*rawdata.KeyValue
	var Errors []*rawdata.Error
	var Transaction rawdata.Transaction

	var ExtendedStatusCodePtr *rawdata.KeyValue
	var ExtendedStatusCode rawdata.KeyValue

	var Error rawdata.Error
	var StartTimePointer *int64 = TimeToInt64(span.StartTimeUnixNano)
	var EndTimePointer *int64 = TimeToInt64(span.EndTimeUnixNano)

	Transaction.StartTime = StartTimePointer
	Transaction.EndTime = EndTimePointer
	Transaction.Uuid = uuidPtr
	//Transaction.TxType = PtrCreateAndAssign(traces.ResourceSpans[0].Resource.Attributes[0].Value.StringValue)

	for _, attribute := range span.Attributes {
		if attribute.Key == "http.url" {
			var UrlPointer *string = PtrCreateAndAssign(attribute.Value.StringValue)
			Transaction.Url = UrlPointer
			url = attribute.Value.StringValue
		}

		if attribute.Key == "http.status_code" {
			ExtendedStatusCode.Key = PtrCreateAndAssign(attribute.Key)

			ExtendedStatusCode.Value = PtrCreateAndAssign(attribute.Value.IntValue)

			value, _ := strconv.Atoi(attribute.Value.IntValue)

			if !(value >= 200 && value < 400) {
				msg := "Http status :" + attribute.Value.IntValue
				msgPtr := &msg
				Error.Code = msgPtr
				Error.StartTime = StartTimePointer

				var status rawdata.Transaction_Status = 1
				statusPointer := &status

				Transaction.Status = statusPointer

				var ErrorPtr *rawdata.Error = &Error
				Errors = append(Errors, ErrorPtr)
			}
		}
	}

	ExtendedStatusCodePtr = &ExtendedStatusCode
	ExtendedStatusCodes = append(ExtendedStatusCodes, ExtendedStatusCodePtr)

	Transaction.Errors = Errors
	Transaction.ExtendedCodes = ExtendedStatusCodes

	// On créé le pointeur qui va pointer sur Transaction
	var TransactionPtr *rawdata.Transaction
	TransactionPtr = &Transaction

	// On ajoute la Transaction dans le tableau des Transactions
	Transactions = append(Transactions, TransactionPtr)

	rd.Transactions = Transactions

	WriteAddRawdata(rd)

}

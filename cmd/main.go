package main

import (
	"fmt"

	"github.com/AlexCorn999/metrics-service/internal/transport/apiserver"
	"github.com/AlexCorn999/metrics-service/internal/transport/billing"
	"github.com/AlexCorn999/metrics-service/internal/transport/email"
	"github.com/AlexCorn999/metrics-service/internal/transport/incidents"
	"github.com/AlexCorn999/metrics-service/internal/transport/mms"
	"github.com/AlexCorn999/metrics-service/internal/transport/sms"
	"github.com/AlexCorn999/metrics-service/internal/transport/support"
	voicecall "github.com/AlexCorn999/metrics-service/internal/transport/voiceCall"
)

func main() {
	server := apiserver.NewAPIServer()
	// if err := server.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	server.SMS = sms.NewSms("./sms.data")
	server.MMS = mms.NewMMS()
	server.Support = support.NewSupport()
	server.Incident = incidents.NewIncident()
	server.VoiceCall = voicecall.NewVoiceCall("./voice.data")
	server.Email = email.NewEmail("./email.data")
	server.Billing = billing.NewBilling("./billing.data")
	dataSMS, err := server.SMS.CheckSMSSystem()
	fmt.Println(dataSMS, err)
	fmt.Println("------------------------------------------------")
	dataMMS, err := server.MMS.CheckMMSSystem()
	fmt.Println(dataMMS, err)
	fmt.Println("------------------------------------------------")
	dataVoiceCall, err := server.VoiceCall.CheckVoiceCallSystem()
	fmt.Println(dataVoiceCall, err)
	fmt.Println("------------------------------------------------")
	dataEmail, err := server.Email.CheckEmailSystem()
	fmt.Println(dataEmail, err)
	fmt.Println("------------------------------------------------")
	dataBilling, _ := server.Billing.CheckBillingSystem()
	fmt.Printf("%+v\n", dataBilling)
	fmt.Println("------------------------------------------------")
	dataSupport, _ := server.Support.CheckSupportData()
	fmt.Printf("%+v\n", dataSupport)
	fmt.Println("------------------------------------------------")
	dataIncident, _ := server.Incident.CheckIncidentData()
	fmt.Printf("%+v\n", dataIncident)
	//res, _ := dataresult.GetResultData()
	//fmt.Println(*res)

}

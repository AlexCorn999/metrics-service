package main

import (
	"fmt"

	"github.com/AlexCorn999/metrics-service/internal/transport/apiserver"
	"github.com/AlexCorn999/metrics-service/internal/transport/mms"
	"github.com/AlexCorn999/metrics-service/internal/transport/sms"
	voicecall "github.com/AlexCorn999/metrics-service/internal/transport/voiceCall"
)

func main() {
	server := apiserver.NewAPIServer()
	// if err := server.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	server.SMS = sms.NewSms("./sms.data")
	server.MMS = mms.NewMMS()
	server.VoiceCall = voicecall.NewVoiceCall("./voice.data")
	dataSMS, err := server.SMS.CheckSMSSystem()
	fmt.Println(dataSMS, err)
	fmt.Println("------------------------------------------------")
	dataMMS, err := server.MMS.CheckMMSSystem()
	fmt.Println(dataMMS, err)
	fmt.Println("------------------------------------------------")
	dataVoiceCall, err := server.VoiceCall.CheckVoiceCallSystem()
	fmt.Println(dataVoiceCall, err)
	//res, _ := dataresult.GetResultData()
	//fmt.Println(*res)

}

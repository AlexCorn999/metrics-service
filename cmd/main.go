package main

import (
	"fmt"

	"github.com/AlexCorn999/metrics-service/internal/support"
)

func main() {
	//server := apiserver.NewAPIServer()
	//if err := server.Start(); err != nil {
	//	log.Fatal(err)
	//}

	// result, _ := sms.CheckSMSSystem("./sms.data")
	// for _, value := range result {
	// 	fmt.Println(value)
	// }

	// result, err := mms.CheckMMSSystem()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// result, _ := voicecall.CheckVoiceCall("./voice.data")
	// for _, value := range result {
	// 	fmt.Println(value)
	// }

	// result, _ := email.CheckEmails("./email.data")
	// for _, value := range result {
	// 	fmt.Println(value)
	// }

	// result, _ := billing.CheckBilling("./billing.data")
	// fmt.Println(result)

	result, _ := support.CheckSupportData()
	fmt.Printf("%+v", result)
}

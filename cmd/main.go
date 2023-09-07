package main

import (
	"fmt"

	sms "github.com/AlexCorn999/metrics-service/internal/smsSystemData"
)

func main() {
	//server := apiserver.NewAPIServer()
	//if err := server.Start(); err != nil {
	//	log.Fatal(err)
	//}

	result, _ := sms.CheckSMSSystem("./sms.data")
	for _, value := range result {
		fmt.Println(value)
	}

}

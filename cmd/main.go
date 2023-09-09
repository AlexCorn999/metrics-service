package main

import (
	"fmt"

	dataresult "github.com/AlexCorn999/metrics-service/internal/dataResult"
)

func main() {
	// server := apiserver.NewAPIServer()
	// if err := server.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	res, _ := dataresult.GetResultData()
	fmt.Println(*res)

}

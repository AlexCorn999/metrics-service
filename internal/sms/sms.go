package sms

import (
	"os"
	"sort"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func CheckSMSSystem(path string) ([]SMSData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(data), "\n")

	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str

	}

	var smsData []SMSData
	for _, entry := range result {

		values := strings.Split(entry, ";")
		if len(values) != 4 {
			continue
		}

		sms := SMSData{
			Country:      values[0],
			Bandwidth:    values[1],
			ResponseTime: values[2],
			Provider:     values[3],
		}
		smsData = append(smsData, sms)
	}

	newData, err := CheckCountries(smsData)
	if err != nil {
		return nil, err
	}

	res, err := CheckProviders(newData)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func CheckCountries(sms []SMSData) ([]SMSData, error) {
	var filteredSms []SMSData
	for _, value := range sms {
		if _, ok := data.Countries[value.Country]; ok {
			filteredSms = append(filteredSms, value)
		}
	}
	return filteredSms, nil
}

func CheckProviders(sms []SMSData) ([]SMSData, error) {
	var filteredSms []SMSData
	for _, value := range sms {
		if _, ok := data.ProvidersSMSMMS[value.Provider]; ok {
			filteredSms = append(filteredSms, value)
		}
	}
	return filteredSms, nil
}

func ResultSMSSystem(sms *[]SMSData) *[][]SMSData {
	var resultData [][]SMSData

	for i := 0; i < len(*sms); i++ {
		country := data.Countries[(*sms)[i].Country]
		(*sms)[i].Country = country
	}

	sms2 := make([]SMSData, len(*sms))
	copy(sms2, *sms)

	sortByProvider(sms)
	sortByCountry(&sms2)

	resultData = [][]SMSData{*sms, sms2}
	return &resultData
}

func sortByProvider(sms *[]SMSData) {
	sort.Slice(*sms, func(i, j int) bool {
		return (*sms)[i].Provider < (*sms)[j].Provider
	})
}

func sortByCountry(sms *[]SMSData) {
	sort.Slice(*sms, func(i, j int) bool {
		return (*sms)[i].Country < (*sms)[j].Country
	})
}

package sms

import (
	"io/ioutil"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

// CheckSMSSystem проверяет файл sms.
func CheckSMSSystem(path string) ([]SMSData, error) {
	data, err := ioutil.ReadFile(path)
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

// CheckCountries проверяет на корректность страны.
func CheckCountries(sms []SMSData) ([]SMSData, error) {
	var filteredSms []SMSData
	for _, value := range sms {
		if _, ok := data.Countries[value.Country]; ok {
			filteredSms = append(filteredSms, value)
		}
	}
	return filteredSms, nil
}

// CheckProviders проверяет на корректность провайдера.
func CheckProviders(sms []SMSData) ([]SMSData, error) {
	var filteredSms []SMSData
	for _, value := range sms {
		if _, ok := data.ProvidersSMSMMS[value.Provider]; ok {
			filteredSms = append(filteredSms, value)
		}
	}
	return filteredSms, nil
}

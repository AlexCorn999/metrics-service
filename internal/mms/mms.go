package mms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// CheckMMSSystem получает данные о состоянии системы через GET запрос.
func CheckMMSSystem() ([]MMSData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mms []MMSData
	if err := json.Unmarshal(data, &mms); err != nil {
		return nil, err
	}

	newData, err := CheckCountries(mms)
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
func CheckCountries(mms []MMSData) ([]MMSData, error) {
	var filteredMMS []MMSData
	for _, value := range mms {
		if _, ok := data.Countries[value.Country]; ok {
			filteredMMS = append(filteredMMS, value)
		}
	}
	return filteredMMS, nil
}

// CheckProviders проверяет на корректность провайдера.
func CheckProviders(mms []MMSData) ([]MMSData, error) {
	var filteredMMS []MMSData
	for _, value := range mms {
		if _, ok := data.ProvidersSMSMMS[value.Provider]; ok {
			filteredMMS = append(filteredMMS, value)
		}
	}
	return filteredMMS, nil
}

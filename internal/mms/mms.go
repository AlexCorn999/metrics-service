package mms

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func CheckMMSSystem() ([]MMSData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	data, err := io.ReadAll(resp.Body)
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

func CheckCountries(mms []MMSData) ([]MMSData, error) {
	var filteredMMS []MMSData
	for _, value := range mms {
		if _, ok := data.Countries[value.Country]; ok {
			filteredMMS = append(filteredMMS, value)
		}
	}
	return filteredMMS, nil
}

func CheckProviders(mms []MMSData) ([]MMSData, error) {
	var filteredMMS []MMSData
	for _, value := range mms {
		if _, ok := data.ProvidersSMSMMS[value.Provider]; ok {
			filteredMMS = append(filteredMMS, value)
		}
	}
	return filteredMMS, nil
}

func ResultMMSSystem(mms *[]MMSData) *[][]MMSData {
	var resultData [][]MMSData

	for i := 0; i < len(*mms); i++ {
		country := data.Countries[(*mms)[i].Country]
		(*mms)[i].Country = country
	}

	mms2 := make([]MMSData, len(*mms))
	copy(mms2, *mms)

	sortByProvider(mms)
	sortByCountry(&mms2)

	resultData = [][]MMSData{*mms, mms2}
	return &resultData
}

func sortByProvider(mms *[]MMSData) {
	sort.Slice(*mms, func(i, j int) bool {
		return (*mms)[i].Provider < (*mms)[j].Provider
	})
}

func sortByCountry(mms *[]MMSData) {
	sort.Slice(*mms, func(i, j int) bool {
		return (*mms)[i].Country < (*mms)[j].Country
	})
}

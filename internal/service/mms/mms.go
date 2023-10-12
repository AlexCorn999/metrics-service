package mmsservice

import (
	"encoding/json"
	"sort"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type MMSService struct {
}

// ValidateMMSData производит проверку данных и переводит данные в структуры MMS.
func (m *MMSService) ValidateMMSData(data []byte) ([]domain.MMSData, error) {
	var mms []domain.MMSData
	if err := json.Unmarshal(data, &mms); err != nil {
		return []domain.MMSData{}, nil
	}

	return mms, nil
}

// CheckCountries проверяет на валидность стран и не допускает данные не прошедшие проверку на страны.
func (m *MMSService) CheckCountries(mmsData *[]domain.MMSData) {
	var filteredMMSData []domain.MMSData
	for _, mms := range *mmsData {
		if _, ok := domain.Countries[mms.Country]; ok {
			filteredMMSData = append(filteredMMSData, mms)
		}
	}
	*mmsData = filteredMMSData
}

// CheckProviders проверяет на валидность провайдеров и не допускает данные не прошедшие проверку на провайдера.
func (m *MMSService) CheckProviders(mmsData *[]domain.MMSData) {
	var filteredMMSData []domain.MMSData
	for _, mms := range *mmsData {
		if _, ok := domain.ProvidersSMSMMS[mms.Provider]; ok {
			filteredMMSData = append(filteredMMSData, mms)
		}
	}
	*mmsData = filteredMMSData
}

// SortByProvider сортирует mms данные по полю провайдер от A до Z.
func (m *MMSService) SortByProvider(mms *[]domain.MMSData) {
	sort.Slice(*mms, func(i, j int) bool {
		return (*mms)[i].Provider < (*mms)[j].Provider
	})
}

// SortByCountry сортирует mms данные по полю страны от A до Z.
func (m *MMSService) SortByCountry(mms *[]domain.MMSData) {
	sort.Slice(*mms, func(i, j int) bool {
		return (*mms)[i].Country < (*mms)[j].Country
	})
}

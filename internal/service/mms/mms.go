package mmsservice

import (
	"encoding/json"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type MMSService struct {
}

// ValidateSMSData производит проверку данных и переводит данные в структуры SMS.
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

package smsservice

import (
	"sort"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type SMSService struct {
}

// ValidateSMSData производит проверку данных и переводит данные в структуры SMS.
func (s *SMSService) ValidateSMSData(data []byte) []domain.SMSData {
	result := strings.Split(string(data), "\n")

	// удаление пробелов из строк
	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str
	}

	var smsData []domain.SMSData

	for _, data := range result {
		values := strings.Split(data, ";")

		if len(values) != 4 {
			continue
		}

		sms := domain.SMSData{
			Country:      values[0],
			Bandwidth:    values[1],
			ResponseTime: values[2],
			Provider:     values[3],
		}
		smsData = append(smsData, sms)
	}

	return smsData
}

// CheckCountries проверяет на валидность стран и не допускает данные не прошедшие проверку на страны.
func (s *SMSService) CheckCountries(smsData *[]domain.SMSData) {
	var filteredSmsData []domain.SMSData
	for _, sms := range *smsData {
		if _, ok := domain.Countries[sms.Country]; ok {
			filteredSmsData = append(filteredSmsData, sms)
		}
	}
	*smsData = filteredSmsData
}

// CheckProviders проверяет на валидность провайдеров и не допускает данные не прошедшие проверку на провайдера.
func (s *SMSService) CheckProviders(smsData *[]domain.SMSData) {
	var filteredSmsData []domain.SMSData
	for _, sms := range *smsData {
		if _, ok := domain.ProvidersSMSMMS[sms.Provider]; ok {
			filteredSmsData = append(filteredSmsData, sms)
		}
	}
	*smsData = filteredSmsData
}

// SortByProvider сортирует sms данные по полю провайдер от A до Z.
func (s *SMSService) SortByProvider(sms *[]domain.SMSData) {
	sort.Slice(*sms, func(i, j int) bool {
		return (*sms)[i].Provider < (*sms)[j].Provider
	})
}

// SortByCountry сортирует sms данные по полю страны от A до Z.
func (s *SMSService) SortByCountry(sms *[]domain.SMSData) {
	sort.Slice(*sms, func(i, j int) bool {
		return (*sms)[i].Country < (*sms)[j].Country
	})
}

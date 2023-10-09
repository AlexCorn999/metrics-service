package sms

import (
	"os"
	"sort"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	smsservice "github.com/AlexCorn999/metrics-service/internal/service/sms"
)

type SMSSystem interface {
	ValidateSMSData(data []byte) []domain.SMSData
	CheckCountries(smsData *[]domain.SMSData)
	CheckProviders(smsData *[]domain.SMSData)
}

type SMS struct {
	smsSystem SMSSystem
	filePath  string
}

func NewSms(filePath string) *SMS {
	return &SMS{
		filePath:  filePath,
		smsSystem: &smsservice.SMSService{},
	}
}

// CheckSMSSystem собирает данные из SMS системы.
func (s *SMS) CheckSMSSystem() ([]domain.SMSData, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	smsData := s.smsSystem.ValidateSMSData(data)
	s.smsSystem.CheckCountries(&smsData)
	s.smsSystem.CheckProviders(&smsData)
	return smsData, nil

}

func ResultSMSSystem(sms *[]domain.SMSData) *[][]domain.SMSData {
	var resultData [][]domain.SMSData

	for i := 0; i < len(*sms); i++ {
		country := domain.Countries[(*sms)[i].Country]
		(*sms)[i].Country = country
	}

	sms2 := make([]domain.SMSData, len(*sms))
	copy(sms2, *sms)

	sortByProvider(sms)
	sortByCountry(&sms2)

	resultData = [][]domain.SMSData{*sms, sms2}
	return &resultData
}

func sortByProvider(sms *[]domain.SMSData) {
	sort.Slice(*sms, func(i, j int) bool {
		return (*sms)[i].Provider < (*sms)[j].Provider
	})
}

func sortByCountry(sms *[]domain.SMSData) {
	sort.Slice(*sms, func(i, j int) bool {
		return (*sms)[i].Country < (*sms)[j].Country
	})
}

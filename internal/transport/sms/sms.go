package sms

import (
	"os"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	smsservice "github.com/AlexCorn999/metrics-service/internal/service/sms"
)

type SMSSystem interface {
	ValidateSMSData(data []byte) []domain.SMSData
	CheckCountries(smsData *[]domain.SMSData)
	CheckProviders(smsData *[]domain.SMSData)
	SortByProvider(sms *[]domain.SMSData)
	SortByCountry(sms *[]domain.SMSData)
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

// ResultSMSSystem сортирует данные и заполняет срез результатов по sms системе.
func (s *SMS) ResultSMSSystem(sms *[]domain.SMSData) *[][]domain.SMSData {
	var resultSMS [][]domain.SMSData

	// замена кода страны на полное название
	for i := 0; i < len(*sms); i++ {
		country := domain.Countries[(*sms)[i].Country]
		(*sms)[i].Country = country
	}

	sms2 := make([]domain.SMSData, len(*sms))
	copy(sms2, *sms)

	s.smsSystem.SortByProvider(sms)
	s.smsSystem.SortByCountry(&sms2)

	resultSMS = [][]domain.SMSData{*sms, sms2}
	return &resultSMS
}

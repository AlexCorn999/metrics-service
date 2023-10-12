package email

import (
	"os"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	emailservice "github.com/AlexCorn999/metrics-service/internal/service/email"
)

type EmailSystem interface {
	ValidateEmailData(data []byte) []domain.EmailData
	CheckCountries(emailData *[]domain.EmailData)
	CheckProviders(emailData *[]domain.EmailData)
	SortByDeliveryTimeUp(email *[]domain.EmailData)
	SortByDeliveryTimeDown(email *[]domain.EmailData)
}

type Email struct {
	emailSystem EmailSystem
	filePath    string
}

func NewEmail(filePath string) *Email {
	return &Email{
		filePath:    filePath,
		emailSystem: &emailservice.EmailService{},
	}
}

// CheckEmailSystem собирает данные из Email системы.
func (s *Email) CheckEmailSystem() ([]domain.EmailData, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	emailData := s.emailSystem.ValidateEmailData(data)
	s.emailSystem.CheckCountries(&emailData)
	s.emailSystem.CheckProviders(&emailData)
	return emailData, nil

}

// ResultEmailSystem сортирует данные и заполняет срез результатов по email системе.
func (s *Email) ResultEmailSystem(emails *[]domain.EmailData) *[][]domain.EmailData {
	var resultEmail [][]domain.EmailData

	emails2 := make([]domain.EmailData, len(*emails))
	copy(emails2, *emails)

	s.emailSystem.SortByDeliveryTimeUp(emails)
	s.emailSystem.SortByDeliveryTimeDown(&emails2)

	resultEmail = [][]domain.EmailData{(*emails)[:3], emails2[:3]}
	return &resultEmail
}

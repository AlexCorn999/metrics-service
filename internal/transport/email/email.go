package email

import (
	"os"
	"sort"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	emailservice "github.com/AlexCorn999/metrics-service/internal/service/email"
)

type EmailSystem interface {
	ValidateEmailData(data []byte) []domain.EmailData
	CheckCountries(emailData *[]domain.EmailData)
	CheckProviders(emailData *[]domain.EmailData)
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

/*
func ResultEmailSystem(emails *[]EmailData) *map[string][][]EmailData {
	resultData := make(map[string][][]EmailData)
	sortByCountry(emails)

	countries := make(map[string][]EmailData)
	for _, email := range *emails {
		countries[email.Country] = append(countries[email.Country], email)
	}

	for country, value := range countries {
		fastestProviders := make([]EmailData, 3)
		slowestProviders := make([]EmailData, 3)

		for _, email := range value {
			// Sort providers by delivery time in ascending order
			sort.SliceStable(value, func(i, j int) bool {
				return email.DeliveryTime < email.DeliveryTime
			})

			// Add fastest providers
			if len(email) >= 3 {
				fastestProviders = append(fastestProviders, email.Provider...)
				fastestProviders = fastestProviders[:3]
			} else {
				fastestProviders = append(fastestProviders, email.Provider...)
			}

			// Add slowest providers
			if len(email.Provider) >= 3 {
				slowestProviders = append(slowestProviders, email.Provider...)
				slowestProviders = slowestProviders[:3]
			} else {
				slowestProviders = append(slowestProviders, email.Provider...)
			}
		}

	}

	fmt.Println("_____________")
	fmt.Println(countries)
	fmt.Println("_____________")

	return &resultData
}*/

func sortByCountry(email *[]domain.EmailData) {
	sort.SliceStable(*email, func(i, j int) bool {
		return (*email)[i].Country < (*email)[j].Country
	})
}

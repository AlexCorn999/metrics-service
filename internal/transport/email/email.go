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

func (s *Email) ResultEmailSystem(emails *[]domain.EmailData) *map[string][][]domain.EmailData {
	resultEmail := make(map[string][][]domain.EmailData)
	fastestProviders := make([]domain.EmailData, 0)
	slowestProviders := make([]domain.EmailData, 0)

	sortByCountryDeliveryTimeUp(emails)

	var currentCountry string
	var currentProvider string
	if len(*emails) > 0 {
		currentCountry = (*emails)[0].Country
		currentProvider = ""
	} else {
		return nil
	}

	if len(*emails) <= 1 {
		slowestProviders = append(slowestProviders, (*emails)[0])
		fastestProviders = append(fastestProviders, (*emails)[0])
		resultEmail[(*emails)[0].Country] = [][]domain.EmailData{fastestProviders, slowestProviders}
		return &resultEmail
	}

	for el, email := range *emails {
		// инициализация ключа в виде кода страны
		if _, ok := resultEmail[email.Country]; !ok {
			resultEmail[email.Country] = [][]domain.EmailData{}
		}

		if email.Country == currentCountry {
			if email.Provider != currentProvider && len(fastestProviders) < 3 {
				fastestProviders = append(fastestProviders, email)
			}
			currentProvider = email.Provider

			if el == len(*emails)-1 {
				resultEmail[currentCountry] = [][]domain.EmailData{fastestProviders}
			}

		} else {
			resultEmail[currentCountry] = [][]domain.EmailData{fastestProviders}
			fastestProviders = make([]domain.EmailData, 0)
			currentCountry = email.Country
			fastestProviders = append(fastestProviders, email)
		}
	}

	sortByCountryDeliveryTimeDown(emails)
	currentCountry = (*emails)[0].Country
	currentProvider = ""

	for el, email := range *emails {

		if email.Country == currentCountry {

			if email.Provider != currentProvider && len(slowestProviders) < 3 {
				slowestProviders = append(slowestProviders, email)
			}
			currentProvider = email.Provider

			if el == len(*emails)-1 {
				resultEmail[currentCountry] = append(resultEmail[currentCountry], slowestProviders)
			}

		} else {
			resultEmail[currentCountry] = append(resultEmail[currentCountry], slowestProviders)
			slowestProviders = make([]domain.EmailData, 0)
			currentCountry = email.Country
			slowestProviders = append(slowestProviders, email)
		}
	}

	return &resultEmail
}

func sortByCountryDeliveryTimeUp(email *[]domain.EmailData) {
	sort.SliceStable(*email, func(i, j int) bool {
		if (*email)[i].Country != (*email)[j].Country {
			return (*email)[i].Country < (*email)[j].Country
		} else {
			return (*email)[i].DeliveryTime < (*email)[j].DeliveryTime
		}
	})
}

func sortByCountryDeliveryTimeDown(email *[]domain.EmailData) {
	sort.SliceStable(*email, func(i, j int) bool {
		if (*email)[i].Country != (*email)[j].Country {
			return (*email)[i].Country < (*email)[j].Country
		} else {
			return (*email)[i].DeliveryTime > (*email)[j].DeliveryTime
		}
	})
}

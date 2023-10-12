package emailservice

import (
	"sort"
	"strconv"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type EmailService struct {
}

// ValidateEmailData производит проверку данных и переводит данные в структуры Email.
func (s *EmailService) ValidateEmailData(data []byte) []domain.EmailData {
	result := strings.Split(string(data), "\n")

	// удаление пробелов из строк
	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str
	}

	var emailData []domain.EmailData

	for _, data := range result {
		values := strings.Split(data, ";")

		if len(values) != 3 {
			continue
		}

		deliveryTime, err := strconv.Atoi(values[2])
		if err != nil {
			continue
		}

		email := domain.EmailData{
			Country:      values[0],
			Provider:     values[1],
			DeliveryTime: deliveryTime,
		}
		emailData = append(emailData, email)
	}

	return emailData
}

// CheckCountries проверяет на валидность стран и не допускает данные не прошедшие проверку на страны.
func (s *EmailService) CheckCountries(emailData *[]domain.EmailData) {
	var filteredEmailData []domain.EmailData
	for _, email := range *emailData {
		if _, ok := domain.Countries[email.Country]; ok {
			filteredEmailData = append(filteredEmailData, email)
		}
	}
	*emailData = filteredEmailData
}

// CheckProviders проверяет на валидность провайдеров и не допускает данные не прошедшие проверку на провайдера.
func (s *EmailService) CheckProviders(emailData *[]domain.EmailData) {
	var filteredEmailData []domain.EmailData
	for _, email := range *emailData {
		if _, ok := domain.ProvidersEmails[email.Provider]; ok {
			filteredEmailData = append(filteredEmailData, email)
		}
	}
	*emailData = filteredEmailData
}

// SortByDeliveryTimeUp сортирует данные email по полю deliveryTime.
func (s *EmailService) SortByDeliveryTimeUp(email *[]domain.EmailData) {
	sort.SliceStable(*email, func(i, j int) bool {
		return (*email)[i].DeliveryTime < (*email)[j].DeliveryTime
	})
}

// SortByDeliveryTimeDown сортирует данные email по полю deliveryTime.
func (s *EmailService) SortByDeliveryTimeDown(email *[]domain.EmailData) {
	sort.SliceStable(*email, func(i, j int) bool {
		return (*email)[i].DeliveryTime > (*email)[j].DeliveryTime
	})
}

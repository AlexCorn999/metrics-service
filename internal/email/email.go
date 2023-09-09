package email

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func CheckEmails(path string) ([]EmailData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(data), "\n")

	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str

	}

	var emailData []EmailData
	for _, entry := range result {
		values := strings.Split(entry, ";")

		if len(values) != 3 {
			continue
		}

		deliveryTimeInt, err := strconv.Atoi(values[2])
		if err != nil {
			continue
		}

		email := EmailData{
			Country:      values[0],
			Provider:     values[1],
			DeliveryTime: deliveryTimeInt,
		}
		emailData = append(emailData, email)
	}

	newData, err := CheckCountries(emailData)
	if err != nil {
		return nil, err
	}

	res, err := CheckProviders(newData)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func CheckCountries(emails []EmailData) ([]EmailData, error) {
	var filteredEmails []EmailData
	for _, value := range emails {
		if _, ok := data.Countries[value.Country]; ok {
			filteredEmails = append(filteredEmails, value)
		}
	}
	return filteredEmails, nil
}

func CheckProviders(emails []EmailData) ([]EmailData, error) {
	var filteredEmails []EmailData
	for _, value := range emails {
		if _, ok := data.ProvidersEmails[value.Provider]; ok {
			filteredEmails = append(filteredEmails, value)
		}
	}
	return filteredEmails, nil
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

func sortByCountry(email *[]EmailData) {
	sort.SliceStable(*email, func(i, j int) bool {
		return (*email)[i].Country < (*email)[j].Country
	})
}

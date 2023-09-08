package email

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

// CheckEmails проверяет файл email.data.
func CheckEmails(path string) ([]EmailData, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// получаем массив строк, которые необходимо разделить ;
	result := strings.Split(string(data), "\n")

	// убираем пробелы.
	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str

	}

	var emailData []EmailData
	for _, entry := range result {
		// делим строку на данные.
		values := strings.Split(entry, ";")
		// должно быть 3 типа данных. Если нету, то пропускаем.
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

// CheckCountries проверяет на корректность страны.
func CheckCountries(emails []EmailData) ([]EmailData, error) {
	var filteredEmails []EmailData
	for _, value := range emails {
		if _, ok := data.Countries[value.Country]; ok {
			filteredEmails = append(filteredEmails, value)
		}
	}
	return filteredEmails, nil
}

// CheckProviders проверяет на корректность провайдера.
func CheckProviders(emails []EmailData) ([]EmailData, error) {
	var filteredEmails []EmailData
	for _, value := range emails {
		if _, ok := data.ProvidersEmails[value.Provider]; ok {
			filteredEmails = append(filteredEmails, value)
		}
	}
	return filteredEmails, nil
}

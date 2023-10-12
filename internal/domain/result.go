package domain

import "errors"

type ResultT struct {
	// true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Status bool `json:"status"`
	// заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Data ResultSetT `json:"data"`
	// пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)
	Error string `json:"error"`
}

var (
	ErrEmptyField = errors.New("the field of structure is empty")
)

type ResultSetT struct {
	SMS       [][]SMSData     `json:"sms"`
	MMS       [][]MMSData     `json:"mms"`
	VoiceCall []VoiceCallData `json:"voice_call"`
	Email     [][]EmailData   `json:"email"`
	Billing   BillingData     `json:"billing"`
	Support   []int           `json:"support"`
	Incidents []IncidentData  `json:"incident"`
}

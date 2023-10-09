package dataresult

import (
	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type ResultT struct {
	// true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Status bool `json:"status"`
	// заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Data ResultSetT `json:"data"`
	// пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)
	Error string `json:"error"`
}

type ResultSetT struct {
	SMS       [][]domain.SMSData              `json:"sms"`
	MMS       [][]domain.MMSData              `json:"mms"`
	VoiceCall []domain.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]domain.EmailData `json:"email"`
	Billing   domain.BillingData              `json:"billing"`
	Support   []int                           `json:"support"`
	Incidents []domain.IncidentData           `json:"incident"`
}

func GetResultData() (*ResultSetT, error) {
	var result ResultSetT

	// SMS system
	//resultSMS, err := sms.CheckSMSSystem("./sms.data")
	//if err != nil {
	//	return nil, err
	//}
	//result.SMS = *sms.ResultSMSSystem(&resultSMS)

	// MMS system
	//resultMMS, err := mms.CheckMMSSystem()
	//if err != nil {
	//	return nil, err
	//}
	//result.MMS = *mms.ResultMMSSystem(&resultMMS)

	// VoiceCall data
	//resultVoiceCall, err := voicecall.CheckVoiceCall("./voice.data")
	//if err != nil {
	//	return nil, err
	//}
	//result.VoiceCall = resultVoiceCall

	// // Emails data
	// resultEmail, err := email.CheckEmails("./email.data")
	// if err != nil {
	// 	return nil, err
	// }
	// result.Email = *email.ResultEmailSystem(&resultEmail)

	// Billing data
	//resultBilling, err := billing.CheckBilling("./billing.data")
	//if err != nil {
	//	return nil, err
	//}
	//result.Billing = *resultBilling

	// // Support data
	// resultSupport, err := support.CheckSupportData()
	// if err != nil {
	// 	return nil, err
	// }
	// result.Support = resultSupport

	// Incidents data
	//resultIncidents, err := accendent.CheckAccendentData()
	//if err != nil {
	//	return nil, err
	//}
	//accendent.ResultAccendentSystem(&resultIncidents)
	//result.Incidents = resultIncidents

	return &result, nil
}

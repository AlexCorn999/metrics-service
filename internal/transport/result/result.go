package dataresult

import (
	"github.com/AlexCorn999/metrics-service/internal/domain"
	"github.com/AlexCorn999/metrics-service/internal/transport/billing"
	"github.com/AlexCorn999/metrics-service/internal/transport/email"
	"github.com/AlexCorn999/metrics-service/internal/transport/incidents"
	"github.com/AlexCorn999/metrics-service/internal/transport/mms"
	"github.com/AlexCorn999/metrics-service/internal/transport/sms"
	"github.com/AlexCorn999/metrics-service/internal/transport/support"
	voicecall "github.com/AlexCorn999/metrics-service/internal/transport/voiceCall"
)

type Result struct {
	SMS       *sms.SMS
	MMS       *mms.MMS
	VoiceCall *voicecall.VoiceCall
	Email     *email.Email
	Billing   *billing.Billing
	Support   *support.Support
	Incident  *incidents.Incident
	Result    *Result
}

func NewResult() *Result {
	return &Result{
		SMS: sms.NewSms("./sms.data"),
		MMS: mms.NewMMS(),
	}
}

func (r *Result) GetResultData() (domain.ResultSetT, error) {
	var result domain.ResultSetT

	// SMS system
	smsSystem, err := r.SMS.CheckSMSSystem()
	if err != nil {
		return domain.ResultSetT{}, err
	}
	resultSMS := r.SMS.ResultSMSSystem(&smsSystem)
	result.SMS = *resultSMS

	// MMS system
	mmsSystem, err := r.MMS.CheckMMSSystem()
	if err != nil {
		return domain.ResultSetT{}, err
	}
	resultMMS := r.MMS.ResultMMSSystem(&mmsSystem)
	result.MMS = *resultMMS

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

	return result, nil
}

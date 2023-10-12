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
		SMS:       sms.NewSms(domain.SMSFilename),
		MMS:       mms.NewMMS(),
		VoiceCall: voicecall.NewVoiceCall(domain.VoiceFilename),
		Email:     email.NewEmail(domain.EmailFilename),
		Billing:   billing.NewBilling(domain.BillingFilename),
		Support:   support.NewSupport(),
		Incident:  incidents.NewIncident(),
	}
}

func (r *Result) GetResultData() (domain.ResultSetT, error) {
	var result domain.ResultSetT

	// SMS system
	smsSystem, err := r.SMS.CheckSMSSystem()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}

	resultSMS := r.SMS.ResultSMSSystem(&smsSystem)
	result.SMS = *resultSMS

	// MMS system
	mmsSystem, err := r.MMS.CheckMMSSystem()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}
	resultMMS := r.MMS.ResultMMSSystem(&mmsSystem)
	result.MMS = *resultMMS

	// VoiceCall system
	voiceCallSystem, err := r.VoiceCall.CheckVoiceCallSystem()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}
	result.VoiceCall = voiceCallSystem

	// Emails system
	emailSystem, err := r.Email.CheckEmailSystem()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}
	resultEmail := r.Email.ResultEmailSystem(&emailSystem)
	result.Email = *resultEmail

	// Billing system
	billingSystem, err := r.Billing.CheckBillingSystem()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}
	result.Billing = billingSystem

	// Support system
	supportSystem, err := r.Support.CheckSupportData()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}
	resultSupport := r.Support.ResultSupportSystem(&supportSystem)
	result.Support = *resultSupport

	// Incident system
	incidentSystem, err := r.Incident.CheckIncidentData()
	if err != nil {
		return domain.ResultSetT{}, domain.ErrEmptyField
	}
	r.Incident.ResultIncidentSystem(&incidentSystem)
	result.Incidents = incidentSystem

	return result, nil
}

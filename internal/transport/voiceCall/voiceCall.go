package voicecall

import (
	"os"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	voicecallservice "github.com/AlexCorn999/metrics-service/internal/service/voicecall"
)

type VoiceCallSystem interface {
	ValidateVoiceCallData(data []byte) []domain.VoiceCallData
	CheckCountries(voice *[]domain.VoiceCallData)
	CheckProviders(voice *[]domain.VoiceCallData)
}

type VoiceCall struct {
	voiceCallSystem VoiceCallSystem
	filePath        string
}

func NewVoiceCall(filePath string) *VoiceCall {
	return &VoiceCall{
		filePath:        filePath,
		voiceCallSystem: &voicecallservice.VoiceCallService{},
	}
}

// CheckVoiceCallSystem собирает данные из VoiceCall системы.
func (s *VoiceCall) CheckVoiceCallSystem() ([]domain.VoiceCallData, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	voiceCallData := s.voiceCallSystem.ValidateVoiceCallData(data)
	s.voiceCallSystem.CheckCountries(&voiceCallData)
	s.voiceCallSystem.CheckProviders(&voiceCallData)
	return voiceCallData, nil
}

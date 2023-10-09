package voicecallservice

import (
	"strconv"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type VoiceCallService struct {
}

// ValidateVoiceCallData производит проверку данных и переводит данные в структуры VoiceCall.
func (v *VoiceCallService) ValidateVoiceCallData(data []byte) []domain.VoiceCallData {
	result := strings.Split(string(data), "\n")

	// удаление пробелов из строк
	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str
	}

	var voiceCallData []domain.VoiceCallData

	for _, data := range result {
		values := strings.Split(data, ";")

		if len(values) != 8 {
			continue
		}

		connection, err := strconv.ParseFloat(values[4], 32)
		if err != nil {
			continue
		}

		ttfb, err := strconv.Atoi(values[5])
		if err != nil {
			continue
		}

		voicePurity, err := strconv.Atoi(values[6])
		if err != nil {
			continue
		}

		medianOfCallsTime, err := strconv.Atoi(values[7])
		if err != nil {
			continue
		}

		voiceCall := domain.VoiceCallData{
			Country:             values[0],
			Bandwidth:           values[1],
			ResponseTime:        values[2],
			Provider:            values[3],
			ConnectionStability: float32(connection),
			TTFB:                ttfb,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   medianOfCallsTime,
		}
		voiceCallData = append(voiceCallData, voiceCall)
	}

	return voiceCallData
}

// CheckCountries проверяет на валидность стран и не допускает данные не прошедшие проверку на страны.
func (v *VoiceCallService) CheckCountries(voiceCallData *[]domain.VoiceCallData) {
	var filteredVoiceCallData []domain.VoiceCallData
	for _, voiceCall := range *voiceCallData {
		if _, ok := domain.Countries[voiceCall.Country]; ok {
			filteredVoiceCallData = append(filteredVoiceCallData, voiceCall)
		}
	}
	*voiceCallData = filteredVoiceCallData
}

// CheckProviders проверяет на валидность провайдеров и не допускает данные не прошедшие проверку на провайдера.
func (v *VoiceCallService) CheckProviders(voiceCallData *[]domain.VoiceCallData) {
	var filteredVoiceCallData []domain.VoiceCallData
	for _, voiceCall := range *voiceCallData {
		if _, ok := domain.ProvidersVoiceCall[voiceCall.Provider]; ok {
			filteredVoiceCallData = append(filteredVoiceCallData, voiceCall)
		}
	}
	*voiceCallData = filteredVoiceCallData
}

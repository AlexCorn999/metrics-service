package voicecall

import (
	"os"
	"strconv"
	"strings"

	"github.com/AlexCorn999/metrics-service/internal/data"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func CheckVoiceCall(path string) ([]VoiceCallData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(data), "\n")

	for i := 0; i < len(result); i++ {
		str := strings.ReplaceAll(result[i], " ", "")
		result[i] = str

	}

	var voiceData []VoiceCallData
	for _, entry := range result {
		values := strings.Split(entry, ";")

		if len(values) != 8 {
			continue
		}

		connectionFloat, err := strconv.ParseFloat(values[4], 32)
		if err != nil {
			continue
		}

		TTFBInt, err := strconv.Atoi(values[5])
		if err != nil {
			continue
		}

		voicePurityInt, err := strconv.Atoi(values[6])
		if err != nil {
			continue
		}

		medianOfCallsTimeInt, err := strconv.Atoi(values[7])
		if err != nil {
			continue
		}

		voice := VoiceCallData{
			Country:             values[0],
			Bandwidth:           values[1],
			ResponseTime:        values[2],
			Provider:            values[3],
			ConnectionStability: float32(connectionFloat),
			TTFB:                TTFBInt,
			VoicePurity:         voicePurityInt,
			MedianOfCallsTime:   medianOfCallsTimeInt,
		}
		voiceData = append(voiceData, voice)
	}

	newData, err := CheckCountries(voiceData)
	if err != nil {
		return nil, err
	}

	res, err := CheckProviders(newData)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func CheckCountries(voice []VoiceCallData) ([]VoiceCallData, error) {
	var filteredVoiceCall []VoiceCallData
	for _, value := range voice {
		if _, ok := data.Countries[value.Country]; ok {
			filteredVoiceCall = append(filteredVoiceCall, value)
		}
	}
	return filteredVoiceCall, nil
}

func CheckProviders(voice []VoiceCallData) ([]VoiceCallData, error) {
	var filteredVoiceCall []VoiceCallData
	for _, value := range voice {
		if _, ok := data.ProvidersVoiceCall[value.Provider]; ok {
			filteredVoiceCall = append(filteredVoiceCall, value)
		}
	}
	return filteredVoiceCall, nil
}

package mms

import (
	"io"
	"net/http"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	mmsservice "github.com/AlexCorn999/metrics-service/internal/service/mms"
)

type MMSSystem interface {
	ValidateMMSData(data []byte) ([]domain.MMSData, error)
	CheckCountries(mmsData *[]domain.MMSData)
	CheckProviders(mmsData *[]domain.MMSData)
	SortByProvider(mms *[]domain.MMSData)
	SortByCountry(mms *[]domain.MMSData)
}

type MMS struct {
	mmsSystem MMSSystem
}

func NewMMS() *MMS {
	return &MMS{
		mmsSystem: &mmsservice.MMSService{},
	}
}

// CheckMMSSystem собирает данные из MMS системы.
func (m *MMS) CheckMMSSystem() ([]domain.MMSData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return []domain.MMSData{}, domain.ErrEmptyField
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	mmsData, err := m.mmsSystem.ValidateMMSData(data)
	if err != nil {
		return nil, err
	}

	m.mmsSystem.CheckCountries(&mmsData)
	m.mmsSystem.CheckProviders(&mmsData)

	return mmsData, nil
}

// ResultMMSSystem сортирует данные и заполняет срез результатов по mms системе.
func (m *MMS) ResultMMSSystem(mms *[]domain.MMSData) *[][]domain.MMSData {
	var resultMMS [][]domain.MMSData

	// замена кода страны на полное название
	for i := 0; i < len(*mms); i++ {
		country := domain.Countries[(*mms)[i].Country]
		(*mms)[i].Country = country
	}

	mms2 := make([]domain.MMSData, len(*mms))
	copy(mms2, *mms)

	m.mmsSystem.SortByProvider(mms)
	m.mmsSystem.SortByCountry(&mms2)

	resultMMS = [][]domain.MMSData{*mms, mms2}
	return &resultMMS
}

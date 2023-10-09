package mms

import (
	"io"
	"net/http"
	"sort"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	mmsservice "github.com/AlexCorn999/metrics-service/internal/service/mms"
)

type MMSSystem interface {
	ValidateMMSData(data []byte) ([]domain.MMSData, error)
	CheckCountries(mmsData *[]domain.MMSData)
	CheckProviders(mmsData *[]domain.MMSData)
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
		return []domain.MMSData{}, nil
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

func ResultMMSSystem(mms *[]domain.MMSData) *[][]domain.MMSData {
	var resultData [][]domain.MMSData

	for i := 0; i < len(*mms); i++ {
		country := domain.Countries[(*mms)[i].Country]
		(*mms)[i].Country = country
	}

	mms2 := make([]domain.MMSData, len(*mms))
	copy(mms2, *mms)

	sortByProvider(mms)
	sortByCountry(&mms2)

	resultData = [][]domain.MMSData{*mms, mms2}
	return &resultData
}

func sortByProvider(mms *[]domain.MMSData) {
	sort.Slice(*mms, func(i, j int) bool {
		return (*mms)[i].Provider < (*mms)[j].Provider
	})
}

func sortByCountry(mms *[]domain.MMSData) {
	sort.Slice(*mms, func(i, j int) bool {
		return (*mms)[i].Country < (*mms)[j].Country
	})
}

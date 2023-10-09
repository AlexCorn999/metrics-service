package support

import (
	"io"
	"net/http"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	supportservice "github.com/AlexCorn999/metrics-service/internal/service/support"
)

type SupportSystem interface {
	ValidateSupportData(data []byte) ([]domain.SupportData, error)
}

type Support struct {
	supportSystem SupportSystem
}

func NewSupport() *Support {
	return &Support{
		supportSystem: &supportservice.SupportService{},
	}
}

// CheckSupportData собирает данные из Support системы.
func (s *Support) CheckSupportData() ([]domain.SupportData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return []domain.SupportData{}, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	supportData, err := s.supportSystem.ValidateSupportData(data)
	if err != nil {
		return nil, err
	}

	return supportData, nil
}

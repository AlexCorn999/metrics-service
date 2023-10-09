package supportservice

import (
	"encoding/json"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type SupportService struct {
}

// ValidateSupportData производит проверку данных и переводит данные в структуры Support.
func (s *SupportService) ValidateSupportData(data []byte) ([]domain.SupportData, error) {
	var support []domain.SupportData
	if err := json.Unmarshal(data, &support); err != nil {
		return []domain.SupportData{}, nil
	}

	return support, nil
}

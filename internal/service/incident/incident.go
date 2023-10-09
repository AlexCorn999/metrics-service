package incidentservice

import (
	"encoding/json"

	"github.com/AlexCorn999/metrics-service/internal/domain"
)

type IncidentService struct {
}

// ValidateIncidentData производит проверку данных и переводит данные в структуры Incident.
func (s *IncidentService) ValidateIncidentData(data []byte) ([]domain.IncidentData, error) {
	var incident []domain.IncidentData
	if err := json.Unmarshal(data, &incident); err != nil {
		return []domain.IncidentData{}, nil
	}

	return incident, nil
}

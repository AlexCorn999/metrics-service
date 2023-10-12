package incidents

import (
	"io"
	"net/http"
	"sort"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	incidentservice "github.com/AlexCorn999/metrics-service/internal/service/incident"
)

type IncidentSystem interface {
	ValidateIncidentData(data []byte) ([]domain.IncidentData, error)
}

type Incident struct {
	incidentSystem IncidentSystem
}

func NewIncident() *Incident {
	return &Incident{
		incidentSystem: &incidentservice.IncidentService{},
	}
}

// CheckIncidentData собирает данные из Incident системы.
func (i *Incident) CheckIncidentData() ([]domain.IncidentData, error) {
	resp, err := http.Get(domain.AccendentListFilename)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return []domain.IncidentData{}, domain.ErrEmptyField
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	incidentData, err := i.incidentSystem.ValidateIncidentData(data)
	if err != nil {
		return nil, err
	}

	return incidentData, nil
}

// ResultIncidentSystem сортирует данные о системе incidents. Статусы active сверху, а close снизу.
func (i *Incident) ResultIncidentSystem(incident *[]domain.IncidentData) {
	sort.Slice(*incident, func(i, j int) bool {
		return (*incident)[i].Status == "active" && (*incident)[j].Status != "active"
	})
}

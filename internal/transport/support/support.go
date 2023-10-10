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
		return []domain.SupportData{}, domain.ErrEmptyField
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

// ResultSupportSystem заполняет срез результатов по support системе.
func (s *Support) ResultSupportSystem(support *[]domain.SupportData) *[]int {
	var resultSupport []int
	var needTime float32
	var ticketTime float32 = 60 / 18
	var status int
	var statusResult int

	for _, data := range *support {
		needTime += float32(data.ActiveTickets) * ticketTime
		status += data.ActiveTickets
	}

	status = status / len(*support)

	if status < 9 {
		statusResult = 1
	} else if status < 16 {
		statusResult = 2
	} else {
		statusResult = 3
	}

	resultSupport = append(resultSupport, statusResult)
	resultSupport = append(resultSupport, int(needTime))
	return &resultSupport
}

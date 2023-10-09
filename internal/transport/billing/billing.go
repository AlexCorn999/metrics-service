package billing

import (
	"os"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	billingservice "github.com/AlexCorn999/metrics-service/internal/service/billing"
)

type BillingSystem interface {
	ValidateBillingData(data []byte) (domain.BillingData, error)
}

type Billing struct {
	billingSystem BillingSystem
	filePath      string
}

func NewBilling(filePath string) *Billing {
	return &Billing{
		filePath:      filePath,
		billingSystem: &billingservice.BillingService{},
	}
}

// CheckBillingSystem собирает данные из Billing системы.
func (s *Billing) CheckBillingSystem() (domain.BillingData, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return domain.BillingData{}, err
	}

	billingData, err := s.billingSystem.ValidateBillingData(data)
	if err != nil {
		return domain.BillingData{}, err
	}

	return billingData, nil
}

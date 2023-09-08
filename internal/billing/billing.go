package billing

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

// CheckBilling проверяет файл billing.data.
func CheckBilling(path string) (*BillingData, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// убираем пробелы.
	str := strings.ReplaceAll(string(data), " ", "")

	// Интерпретируем битовую маску и сохраняем сумму степеней каждого бита.
	sum := uint8(0)
	for i := len(str) - 1; i >= 0; i-- {
		bit, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования:%s", err)
		}

		if bit == 1 {
			sum += 1 << (len(str) - 1 - i)
		}
	}

	// Проверяем каждый бит на соответствие 1 и сохраняем результаты в структуру.
	billingData := BillingData{
		CreateCustomer: sum&1 > 0,
		Purchase:       sum&2 > 0,
		Payout:         sum&4 > 0,
		Recurring:      sum&8 > 0,
		FraudControl:   sum&16 > 0,
		CheckoutPage:   sum&32 > 0,
	}

	return &billingData, nil
}

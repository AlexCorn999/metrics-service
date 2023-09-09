package support

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func CheckSupportData() ([]SupportData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var support []SupportData
	if err := json.Unmarshal(data, &support); err != nil {
		return nil, err
	}

	return support, nil
}

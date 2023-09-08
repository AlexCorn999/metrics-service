package accendent

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

// CheckAccendentData получает данные о состоянии системы через GET запрос.
func CheckAccendentData() ([]IncidentData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var accendents []IncidentData
	if err := json.Unmarshal(data, &accendents); err != nil {
		return nil, err
	}

	return accendents, nil
}

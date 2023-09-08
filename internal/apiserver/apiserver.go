package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AlexCorn999/metrics-service/internal/accendent"
	"github.com/AlexCorn999/metrics-service/internal/billing"
	"github.com/AlexCorn999/metrics-service/internal/email"
	"github.com/AlexCorn999/metrics-service/internal/mms"
	"github.com/AlexCorn999/metrics-service/internal/sms"
	voicecall "github.com/AlexCorn999/metrics-service/internal/voiceCall"
	"github.com/gorilla/mux"
)

type ResultT struct {
	Status bool       `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Data   ResultSetT `json:"data"`   // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)
}

type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voicecall.VoiceCallData      `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []accendent.IncidentData       `json:"incident"`
}

type APIServer struct {
	router *mux.Router
}

func NewAPIServer() *APIServer {
	return &APIServer{
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureRouter(); err != nil {
		return err
	}
	log.Println("starting server ...")
	return http.ListenAndServe(":8080", s.router)
}

func (s *APIServer) ConfigureRouter() error {
	s.router.HandleFunc("/", s.handleConnection)
	return nil
}

func (s *APIServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	var result ResultT

	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("%+v", string(data))
	w.Write(data)
}

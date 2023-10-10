package apiserver

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/AlexCorn999/metrics-service/internal/domain"
	dataresult "github.com/AlexCorn999/metrics-service/internal/transport/result"
	"github.com/gorilla/mux"
)

type APIServer struct {
	router *mux.Router
	Result *dataresult.Result
	//SMS       *sms.SMS
	//MMS       *mms.MMS
	//VoiceCall *voicecall.VoiceCall
	//Email    *email.Email
	//Billing  *billing.Billing
	//Support  *support.Support
	//Incident *incidents.Incident
}

func NewAPIServer() *APIServer {
	return &APIServer{
		router: mux.NewRouter(),
		Result: dataresult.NewResult(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureRouter(); err != nil {
		return err
	}

	//s.SMS = sms.NewSms("./sms.data")
	//s.MMS = mms.NewMMS()
	//s.VoiceCall = voicecall.NewVoiceCall("./voice.data")
	//s.Email = email.NewEmail("./email.data")
	//s.Billing = billing.NewBilling("./billing.data")
	//s.Support = support.NewSupport()
	//s.Incident = incidents.NewIncident()

	log.Println("starting server ...")
	return http.ListenAndServe(":8080", s.router)
}

func (s *APIServer) ConfigureRouter() error {
	s.router.HandleFunc("/", s.handleConnection)
	return nil
}

func (s *APIServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	var result domain.ResultT
	result.Status = true

	res, err := s.Result.GetResultData()
	if err != nil {
		if errors.Is(err, domain.ErrEmptyField) {
			result.Status = false
		} else {
			result.Error = err.Error()
		}
	}

	if err == nil {
		result.Data = res
	}

	data, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

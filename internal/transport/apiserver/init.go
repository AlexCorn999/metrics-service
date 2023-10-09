package apiserver

import (
	"log"
	"net/http"

	"github.com/AlexCorn999/metrics-service/internal/transport/mms"
	"github.com/AlexCorn999/metrics-service/internal/transport/sms"
	voicecall "github.com/AlexCorn999/metrics-service/internal/transport/voiceCall"
	"github.com/gorilla/mux"
)

type APIServer struct {
	router    *mux.Router
	SMS       *sms.SMS
	MMS       *mms.MMS
	VoiceCall *voicecall.VoiceCall
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

	s.SMS = sms.NewSms("./sms.data")
	s.MMS = mms.NewMMS()
	s.VoiceCall = voicecall.NewVoiceCall("./voice.data")

	log.Println("starting server ...")
	return http.ListenAndServe(":8080", s.router)
}

func (s *APIServer) ConfigureRouter() error {
	s.router.HandleFunc("/", s.handleConnection)
	return nil
}

func (s *APIServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	//var result ResultT

	//data, err := json.Marshal(result)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	w.WriteHeader(http.StatusOK)
	//fmt.Println("%+v", string(data))
	//w.Write(data)
}

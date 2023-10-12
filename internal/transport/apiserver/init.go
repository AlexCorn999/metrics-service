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

	log.Println("starting server ...")
	return http.ListenAndServe(":8080", s.router)
}

func (s *APIServer) ConfigureRouter() error {
	s.router.HandleFunc("/test", s.handleConnection).Methods("GET", "OPTIONS")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(data)
}

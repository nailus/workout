package apiserver

import (
	//"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"io"
	"github.com/nailus/workout/internal/database"
	"fmt"
)

type ApiServer struct {
	config *Config
	//logger *logger.Logrus
	router *mux.Router
	database *database.Database
}

func New(config *Config) *ApiServer {
	return &ApiServer {
		config: config,
		//logger: logrus.New(),
		router: mux.NewRouter(),
		database: database.NewDatabase(config.Database),
	}
}

func (s  *ApiServer) Start() error {
	
	s.configureRouter()
	
	
	if err := s.configureDatabase(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *ApiServer) configureDatabase() error {
	db := database.NewDatabase(s.config.Database)
	if err := db.Open(); err != nil {
		return err
	}

	

	s.database = db

	return nil
}
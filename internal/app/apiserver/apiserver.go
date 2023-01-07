package apiserver

import (
	//"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"io"
)

type ApiServer struct {
	config *Config
	//logger *logger.Logrus
	router *mux.Router
}

func New(config *Config) *ApiServer {
	return &ApiServer {
		config: config,
		//logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s  *ApiServer) Start() error {
	s.configureRouter()

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
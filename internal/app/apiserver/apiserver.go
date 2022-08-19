package apiserver

import (
	"context"
	"example.com/prj/store"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewServer(config *Config) *ApiServer {
	return &ApiServer{config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start(ctx context.Context) error {
	if err := s.initLogger(); err != nil {
		return err
	}
	s.initRouter()
	if err := s.initStore(ctx); err != nil {
		return err
	}
	s.logger.Info("Starting API server...")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ApiServer) initLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) initRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.Handle("/metrics", promhttp.Handler())
}

func (s *ApiServer) initStore(ctx context.Context) error {
	st := store.New(s.config.Store)
	if err := st.Open(ctx); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "my_app_ops_total",
		Help: "Total number of /hello calls",
	})
	return func(w http.ResponseWriter, req *http.Request) {
		s.logger.Info("A ", req.Method, " with path = ", req.RequestURI)
		io.WriteString(w, "Hey:)")
		opsProcessed.Inc()
	}
}

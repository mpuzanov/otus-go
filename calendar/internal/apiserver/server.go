package apiserver

import (
	"github.com/mpuzanov/otus-go/calendar/internal/storage"
	"github.com/mpuzanov/otus-go/calendar/pkg/logger"
	"go.uber.org/zap"
)

type server struct {
	//router       *mux.Router
	logger *zap.Logger
	store  storage.Storage
}

// Start Запуск сервера
func Start() error {
	err := storage.NewStorage(storage.Memory)
	if err != nil {
		return err
	}
	_ = newServer(storage.DB)

	storage.ISample()

	return nil
}

func newServer(store storage.Storage) *server {
	s := &server{
		//router:       mux.NewRouter(),
		logger: logger.InitLogger(),
		store:  store,
	}
	return s
}

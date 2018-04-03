package browser

import (
	"errors"
	"net/http"
	"reflect"
	"time"

	"github.com/RenegadeTech/MysticalTutor/interfaces"
	"github.com/RenegadeTech/MysticalTutor/logger"
	"github.com/gorilla/mux"
)

const (
	url = "localhost:8080"
)

type display struct {
	engine prototype.Engine
	server *mux.Router
}

func New() prototype.Display {
	return &display{
		server: mux.NewRouter(),
	}
}

func (d *display) Connect(engine prototype.Engine) error {
	if engine == nil {
		return errors.New("Engine is nil")
	}
	if reflect.TypeOf(engine).Kind() != reflect.Ptr {
		return errors.New("Engine needs to be a reference")
	}
	d.engine = engine
	return nil
}

func (d *display) Initialise() prototype.Display {
	return d
}

func (d *display) Update() {

}

func (d *display) Run() error {
	logger.GetInstance().Log(logger.Entry{Level: logger.Trace,
		Data: "Starting display",
	})
	if err := open("http://" + url); err != nil {
		return err
	}
	srv := &http.Server{
		Handler:      d.server,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}

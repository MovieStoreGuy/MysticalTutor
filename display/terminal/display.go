package terminal

import (
	"errors"
	"reflect"

	"github.com/RenegadeTech/MysticalTutor/interfaces"
	"github.com/RenegadeTech/MysticalTutor/logger"
	ui "github.com/gizak/termui"
)

func New() prototype.Display {
	return &display{
		eng:    nil,
		points: []ui.Bufferer{},
	}
}

type display struct {
	eng    prototype.Engine
	points []ui.Bufferer
}

func (d *display) Connect(engine prototype.Engine) error {
	if engine == nil {
		return errors.New("engine was nil")
	}
	if reflect.TypeOf(engine).Kind() != reflect.Ptr {
		return errors.New("engine needs to be a pointer not a copy")
	}
	d.eng = engine
	return nil
}

func (d *display) Initialise() prototype.Display {
	if err := ui.Init(); err != nil {
		logger.GetInstance().Log(logger.Entry{Level: logger.Fatal,
			Data: "Unable to start UI due to: " + err.Error(),
		})
		panic(err)
	}
	return d
}

func (d *display) Update() {
	ui.Render(d.points...)
}

func (d *display) Run() error {
	logger.GetInstance().Log(logger.Entry{Level: logger.Debug,
		Data: "Starting to run the terminal gui",
	})
	// Ensure that the console is stopped once the program has stopped
	defer ui.Close()
	for {
		d.Update()
	}
	return nil
}

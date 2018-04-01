package terminal

import (
	"errors"
	"reflect"

	"github.com/RenegadeTech/MysticalTutor/interfaces"
)

func New() prototype.Display {
	return &display{}
}

type display struct {
	eng prototype.Engine
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
	return d
}

func (d *display) Update() {

}

func (d *display) Run() error {
	return nil
}

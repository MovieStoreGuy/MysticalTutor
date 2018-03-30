package engine

import (
	"errors"
	"reflect"

	"github.com/RenegadeTech/MysticalTutor/interfaces"
)

type driver struct {
	disp prototype.Display
}

func (d *driver) Connect(disp prototype.Display) error {
	if disp == nil {
		return errors.New("nil was passed in for the display")
	}
	// As we want to ensure that we are interfacing with the same type
	// That isn't a copy, we need to ensure that we get what we want.
	if reflect.TypeOf(disp).Kind() != reflect.Ptr {
		return errors.New("Display needs to be a pointer to an object")
	}
	d.disp = disp
	return nil
}

func (d *driver) Initialise() prototype.Engine {
	return d
}

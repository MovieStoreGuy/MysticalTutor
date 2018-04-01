package prototype

import "errors"

func Connect(e Engine, d Display) error {
	if e == nil || d == nil {
		return errors.New("Unable to connect due to nil object")
	}
	if err := e.Connect(d); err != nil {
		return err
	}
	return d.Connect(e)
}

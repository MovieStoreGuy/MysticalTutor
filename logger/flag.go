package logger

import (
	"errors"
	"strings"
)

type Flag struct {
	level Level
}

func (f *Flag) String() string {
	return f.level.String()
}

func (f *Flag) Set(value string) error {
	switch strings.ToLower(value) {
	case "info":
		f.level = Info
	case "fatal":
		f.level = Fatal
	case "trace":
		f.level = Trace
	case "debug":
		f.level = Debug
	default:
		return errors.New("Unknown type being passed")
	}
	return nil
}

func (f *Flag) GetLevel() Level {
	return f.level
}

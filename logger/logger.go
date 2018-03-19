package logger

import (
	"fmt"
	"io"
	"sync"
)

type Level int8

type Entry struct {
	Level Level
	Data  string
}

type instance struct {
	entries chan Entry
	done    chan bool
	output  io.Writer
	level   Level
	running bool
}

var (
	printer map[Level]string = map[Level]string{
		Fatal: "Fatal",
		Info:  "Info",
		Debug: "Debug",
		Trace: "Trace",
	}
)

func (l Level) String() string {
	if data, exist := printer[l]; exist {
		return data
	}
	return "Unknown"
}

const (
	Fatal Level = iota
	Info
	Debug
	Trace
)

var (
	isnt *instance
	once sync.Once
)

func GetInstance() *instance {
	once.Do(func() {
		isnt = &instance{
			level:   Fatal,
			entries: make(chan Entry, 100),
			done:    make(chan bool),
		}
	})
	return isnt
}

func (i *instance) Log(e Entry) {
	i.entries <- e
}

func (i *instance) Set(w io.Writer, level Level) *instance {
	if i.output == nil {
		i.output = w
	}
	i.level = level
	return i
}

func (i *instance) Start() {
	if i.running {
		return
	}
	i.running = true
	go func() {
		for data := range i.entries {
			if data.Level <= i.level {
				fmt.Fprintf(i.output, "[%s] %s\n", i.level, data.Data)
			}
		}
		i.done <- true
	}()
}

func (i *instance) Stop() {
	for {
		if len(i.entries) == 0 {
			break
		}
		// Wait for the buffer to empty
	}
	close(i.entries)
	<-i.done
}

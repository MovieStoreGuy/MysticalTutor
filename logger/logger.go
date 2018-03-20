package logger

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"sync"
)

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
	isnt       *instance
	once       sync.Once
	BufferSize int
)

func init() {
	// BufferSize defines how many log entries can store before locking the application
	BufferSize = 100
}

func GetInstance() Log {
	once.Do(func() {
		isnt = &instance{
			level:   Fatal,
			entries: make(chan Entry, BufferSize),
			done:    make(chan bool),
		}
	})
	return isnt
}

func (i *instance) Log(e Entry) Log {
	if !i.running {
		return i
	}
	_, fn, line, _ := runtime.Caller(1)
	e.Data = fmt.Sprintf("[%s:%d] %s", path.Base(fn), line, e.Data)
	i.entries <- e
	return i
}

func (i *instance) Set(w io.Writer, level Level) Log {
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
				fmt.Fprintf(i.output, "[%s] %s\n", data.Level, data.Data)
			}
		}
		i.done <- true
	}()
}

func (i *instance) Stop() {
	i.running = false
	for {
		if len(i.entries) == 0 {
			break
		}
		// Wait for the buffer to empty
	}
	close(i.entries)
	<-i.done
	close(i.done)
}

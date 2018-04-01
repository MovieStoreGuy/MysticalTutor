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
	isnt *instance
	once sync.Once
	// BufferSize defines how many log entries can store before locking the application
	BufferSize int
)

func init() {
	BufferSize = 100
}

// GetInstance will returns the logger singleton
// ready for use.
func GetInstance() Log {
	once.Do(func() {
		isnt = &instance{
			level: Fatal,
		}
	})
	return isnt
}

func (i *instance) Log(e Entry) Log {
	if !i.running {
		return i
	}
	_, fn, line, _ := runtime.Caller(1)
	e.Data = fmt.Sprintf("[%s:%d]\t%s", path.Base(fn), line, e.Data)
	i.entries <- e
	if i.level == Debug {
		var m runtime.MemStats
		convert := func(val uint64) uint64 {
			return val / 1024 / 1024
		}
		runtime.ReadMemStats(&m)
		i.entries <- Entry{
			Level: Debug,
			Data:  fmt.Sprintf("[Current Usage] %v MiB, [GC Count] %v", convert(m.Alloc), m.NumGC),
		}
	}
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
	i.entries = make(chan Entry, BufferSize)
	i.done = make(chan bool)
	go func() {
		for data := range i.entries {
			if data.Level <= i.level {
				fmt.Fprintf(i.output, "[%s]\t%s\n", data.Level, data.Data)
			}
		}
		i.done <- true
	}()
}

func (i *instance) Stop() {
	if !i.running {
		return
	}
	i.Log(Entry{
		Level: Info,
		Data:  "Logger is being stopped",
	})
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

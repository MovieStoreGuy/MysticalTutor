// MysticalTutor is an application that allows for the user
// to create the most optimal deck given their needs.
package main

import (
	"flag"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/RenegadeTech/MysticalTutor/engine"
	"github.com/RenegadeTech/MysticalTutor/interfaces"
	"github.com/RenegadeTech/MysticalTutor/logger"
)

var (
	logLevel  *logger.Flag = &logger.Flag{}
	logWriter io.Writer    = os.Stderr
	log                    = logger.GetInstance()
)

func init() {
	flag.Var(logLevel, "log-level", "The ability to set the log level to something more invovled")
}

func main() {
	flag.Parse()
	log.Set(logWriter, logLevel.GetLevel()).
		Start()
	log.Log(logger.Entry{Level: logger.Info,
		Data: "Started running application " + path.Base(os.Args[0]),
	})
	log.Log(logger.Entry{Level: logger.Info,
		Data: "Log level is set to " + logLevel.String(),
	})
	log.Log(logger.Entry{Level: logger.Info,
		Data: "Golang Version: " + runtime.Version(),
	})
	var display prototype.Display
	var engine prototype.Engine = engine.New()
	if err := prototype.Connect(engine, display); err != nil {
		log.Log(logger.Entry{Level: logger.Fatal,
			Data: err.Error(),
		})
		goto cleanup
	}
	if err := display.Run(); err != nil {
		log.Log(logger.Entry{Level: logger.Fatal,
			Data: "Application stopped due to " + err.Error(),
		})
		goto cleanup
	}

cleanup:
	// Ensures that we flush any logs waiting to be written
	log.Stop()
}

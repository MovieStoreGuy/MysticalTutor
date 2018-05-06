// MysticalTutor is an application that allows for the user
// to create the most optimal deck given their needs.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/RenegadeTech/MysticalTutor/display"
	"github.com/RenegadeTech/MysticalTutor/engine"
	"github.com/RenegadeTech/MysticalTutor/interfaces"
	"github.com/RenegadeTech/MysticalTutor/logger"
)

var (
	logLevel  *logger.Flag = &logger.Flag{}
	logWriter io.Writer    = os.Stderr
	log                    = logger.GetInstance()

	enableBrowser bool
)

const (
	message string = `%s -- A Magic the Gathering™ deck building app

This app will help you gauge and decide what cards should be paired together.
Using varying methods to help choose what would be best, it is every brewers delight.

You can use the following command line flags are mainly used to help debug the app
or change the output.

Flags:
`
)

func init() {
	flag.Var(logLevel, "log-level", "The ability to set the log level to something more invovled")
	flag.BoolVar(&enableBrowser, "enable-browser", true, "Allows the user to view everything view their browser")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, message, os.Args[0])
		flag.PrintDefaults()
	}
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
	var display prototype.Display = util.GetDisplay(enableBrowser)
	var engine prototype.Engine = engine.New()
	if err := prototype.Connect(engine, display); err != nil {
		log.Log(logger.Entry{Level: logger.Fatal,
			Data: err.Error(),
		})
		goto cleanup
	}
	engine.Initialise()
	if err := display.Initialise().Run(); err != nil {
		log.Log(logger.Entry{Level: logger.Fatal,
			Data: "Application stopped due to " + err.Error(),
		})
		goto cleanup
	}

cleanup:
	// Ensures that we flush any logs waiting to be written
	log.Stop()
}

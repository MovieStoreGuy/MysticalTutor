package main

import (
	"io"
	"os"

	"github.com/RenegadeTech/MysticalTutor/logger"
)

var (
	logLevel  logger.Level = logger.Info
	logWriter io.Writer    = os.Stderr
)

func main() {
	log := logger.GetInstance()
	log.Set(logWriter, logLevel).
		Start()
	log.Log(logger.Entry{Level: logger.Info,
		Data: "Started running application" + os.Args[0],
	})

	log.Stop()
}

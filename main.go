package main

import (
	"os"

	"github.com/RenegadeTech/MysticalTutor/logger"
)

var (
	logLevel logger.Level = logger.Fatal
)

func main() {
	log := logger.GetInstance()
	log.Set(os.Stderr, logLevel).
		Start()
	log.Log(logger.Entry{Level: logger.Info,
		Data: "Started running application" + os.Args[0],
	})

	log.Stop()
}

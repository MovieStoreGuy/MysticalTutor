package browser

import (
	"fmt"
	"os/exec"

	"github.com/RenegadeTech/MysticalTutor/logger"
)

func open(url string) error {
	// Preference: macOS, Linux, and Windows
	for _, browser := range []string{"open", "xdg-open", "start"} {
		logger.GetInstance().Log(logger.Entry{Level: logger.Trace,
			Data: "Trying to start: " + browser,
		})
		err := exec.Command(browser, url).Start()
		if err == nil {
			return nil
		}
		logger.GetInstance().Log(logger.Entry{Level: logger.Trace,
			Data: "Issue trying to start the browser: " + err.Error(),
		})
	}
	return fmt.Errorf("Unable to open default browser")
}

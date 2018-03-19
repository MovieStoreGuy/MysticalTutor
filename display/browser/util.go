package browser

import (
	"fmt"
	"os/exec"
)

func open(url string) error {
	// Preference: macOS, Linux, and Windows
	for _, browser := range []string{"open", "xdg-open", "start"} {
		err := exec.Command(browser, url).Start()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("Unable to open default browser")
}

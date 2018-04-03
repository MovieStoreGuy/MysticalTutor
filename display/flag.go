package util

import (
	"github.com/RenegadeTech/MysticalTutor/display/browser"
	"github.com/RenegadeTech/MysticalTutor/display/terminal"
	"github.com/RenegadeTech/MysticalTutor/interfaces"
)

func GetDisplay(allowBrowser bool) prototype.Display {
	if allowBrowser {
		return browser.New()
	}
	return terminal.New()
}

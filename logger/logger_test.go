package logger_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/RenegadeTech/MysticalTutor/logger"
)

func TestCreateLogger(t *testing.T) {
	i := logger.GetInstance()
	if i == nil {
		t.Fatal("Created a nil object")
	}
	if i != logger.GetInstance() {
		t.Fatal("Should create a singleton")
	}
}

func TestLogger(t *testing.T) {
	i := logger.GetInstance()
	if i == nil {
		t.Fatal("Created a nil object")
	}
	b := &bytes.Buffer{}
	i.Set(b, logger.Trace)
	i.Log(logger.Entry{
		Level: logger.Info,
		Data:  "What is it good for",
	})
	i.Start()
	// testing starting the logger twice
	i.Start()
	i.Log(logger.Entry{
		Level: logger.Info,
		Data:  "What is it good for",
	})
	i.Stop()
	expected := fmt.Sprintf("^[Trace] .* %s$", "What is it good for")
	if regexp.MustCompile(expected).MatchString(string(b.Bytes())) {
		t.Log("Expected:", expected)
		t.Log("Given:", string(b.Bytes()))
		t.Fatal("Incorrect details logged")
	}
	i.Stop()
}

func TestLogLevel(t *testing.T) {
	var level logger.Level = -1
	if level.String() != "Unknown" {
		t.Fatal("Should report an unknown level")
	}
}

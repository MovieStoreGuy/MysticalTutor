package logger_test

import (
	"bytes"
	"fmt"
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
	i.Stop()
	expected := fmt.Sprintf("[Trace]%s\n", "What is it good for")
	if string(b.Bytes()) != expected {
		t.Log("Expected:", expected)
		t.Log("Given:", string(b.Bytes()))
		t.Fatal("Incorrect details logged")
	}
}

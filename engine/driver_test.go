package engine_test

import (
	"os"
	"testing"

	"github.com/RenegadeTech/MysticalTutor/engine"
)

func TestInitialise(t *testing.T) {
	defer os.RemoveAll(engine.CollectionPath)
	d := engine.New()
	collection := d.Initialise().GetEntireCollection()
	if len(collection) == 0 {
		t.Fatal("Failed to load cards from storage")
	}
	t.Log("Number of cards loaded", len(collection))
}

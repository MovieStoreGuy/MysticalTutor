package engine

import (
	"errors"
	"reflect"

	"github.com/RenegadeTech/MysticalTutor/interfaces"
	"github.com/RenegadeTech/MysticalTutor/types"
)

const (
	// JsonCollectionURL is the URL used as the reference point for all cards available
	JsonCollectionURL = "https://mtgjson.com/json/AllCards.json.zip"
)

var (
	// CollectionPath is the local file directory as to where the "storage bank"
	// of cards to use is kept. Can be changed incase the user wants to use a different sellection
	// of cards
	CollectionPath = "resources/AllCards.json"
)

type driver struct {
	disp  prototype.Display
	store types.Collection
}

func (d *driver) Connect(disp prototype.Display) error {
	if disp == nil {
		return errors.New("nil was passed in for the display")
	}
	// As we want to ensure that we are interfacing with the same type
	// That isn't a copy, we need to ensure that we get what we want.
	if reflect.TypeOf(disp).Kind() != reflect.Ptr {
		return errors.New("Display needs to be a pointer to an object")
	}
	d.disp = disp
	return nil
}

func (d *driver) Initialise() prototype.Engine {
	return d
}

func (d *driver) AddProcessor(p prototype.Processor) error {
	return nil
}

func (d *driver) RemoveProcessor(p prototype.Processor) error {
	return nil
}

func (d *driver) AddCollection(c types.Collection) error {
	return nil
}

func (d *driver) RemoveCollection(c types.Collection) error {
	return nil
}

func (d *driver) GetProcessors() []prototype.Processor {
	return nil
}

func (d *driver) GetCollections() []types.Collection {
	return nil
}

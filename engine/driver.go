package engine

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/RenegadeTech/MysticalTutor/interfaces"
	"github.com/RenegadeTech/MysticalTutor/logger"
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
	disp        prototype.Display
	store       types.Collection
	collections []types.Collection
	processors  []prototype.Processor
}

func New() prototype.Engine {
	return &driver{
		store:       types.Collection{},
		collections: []types.Collection{},
		processors:  []prototype.Processor{},
	}
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
	logger.GetInstance().Log(logger.Entry{
		Level: logger.Info,
		Data:  "Initialising the engine",
	})
	d.store = types.Collection{}
	_, err := os.Stat(CollectionPath)
	switch {
	case os.IsNotExist(err):
		if err := d.downloadCardCollection(); err != nil {
			logger.GetInstance().Log(logger.Entry{
				Level: logger.Info,
				Data:  "Recieved error: " + err.Error(),
			})
		}
	default:
		if err := d.loadCollectionFromDisk(); err != nil {
			// Failed to load what we need
			logger.GetInstance().Log(logger.Entry{
				Level: logger.Fatal,
				Data:  "Recieved error: " + err.Error(),
			})
		}
	}
	d.collections, d.processors = []types.Collection{}, []prototype.Processor{}
	return d
}

func (d *driver) AddProcessor(p prototype.Processor) error {
	if p == nil {
		return errors.New("Process was nil")
	}
	d.processors = append(d.processors, p)
	return nil
}

func (d *driver) RemoveProcessor(p prototype.Processor) error {
	for i, proc := range d.processors {
		if proc == p {
			d.processors = append(d.processors[:i], d.processors[i+1:]...)
			break
		}
	}
	return nil
}

func (d *driver) AddCollection(c types.Collection) error {
	if c == nil {
		return errors.New("Collection was nil")
	}
	d.collections = append(d.collections, c)
	return nil
}

func (d *driver) RemoveCollection(c types.Collection) error {
	for i, collection := range d.collections {
		if reflect.DeepEqual(c, collection) {
			d.collections = append(d.collections[:i], d.collections[i+1:]...)
			break
		}
	}
	return nil
}

func (d *driver) GetProcessors() []prototype.Processor {
	return d.processors
}

func (d *driver) GetCollections() []types.Collection {
	return d.collections
}

func (d *driver) GetEntireCollection() types.Collection {
	return d.store
}

func (d *driver) ProcessCollectionID(id int) {
	// TODO(Sean Marciniak): Consider how we are going to process a given collection
}

func (d *driver) downloadCardCollection() error {
	files, err := downloadZip(JsonCollectionURL)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.FileInfo().IsDir() {
			continue
		}
		f, err := file.Open()
		if err != nil {
			return err
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		dto := map[string]*types.Card{}
		err = json.Unmarshal(b, &dto)
		for _, card := range dto {
			d.store = append(d.store, card)
		}
	}
	return err
}

func (d *driver) loadCollectionFromDisk() error {
	if _, err := os.Stat(CollectionPath); os.IsNotExist(err) {
		return errors.New("File does not exist")
	}
	buff, err := ioutil.ReadFile(CollectionPath)
	if err != nil {
		return err
	}
	dto := map[string]*types.Card{}
	if err := json.Unmarshal(buff, &dto); err != nil {
		return err
	}
	for _, card := range dto {
		d.store = append(d.store, card)
	}
	return nil
}

package prototype

import "github.com/RenegadeTech/MysticalTutor/types"

// Engine allows for Factory like interactions
// and well defined model processing and management
type Engine interface {
	Connect(display Display) error

	Initialise() Engine

	AddProcessor(p Processor) error

	AddCollection(c types.Collection) error

	RemoveProcessor(p Processor) error

	RemoveCollection(c types.Collection) error

	GetProcessors() []Processor

	GetCollections() []types.Collection

	GetEntireCollection() types.Collection
}

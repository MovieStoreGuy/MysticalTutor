package prototype

import "github.com/RenegadeTech/MysticalTutor/types"

// Engine ...
type Engine interface {
	Connect(display Display) error

	AddProcessor(p Processor) error

	AddCollection(c types.Collection) error

	Initialise() Engine
}

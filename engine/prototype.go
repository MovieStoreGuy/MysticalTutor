package engine

// Engine ...
type Engine interface {
	Connect(display interface{}) error
}

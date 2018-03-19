package prototype

// Display allows to support various types of actual displays
// and keep our code clean as possible.
type Display interface {
	Connect(engine interface{}) error

	Initialise() Display

	Update()

	Run() error
}

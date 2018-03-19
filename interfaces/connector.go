package prototype

func Connect(e Engine, d Display) error {
	if err := e.Connect(d); err != nil {
		return err
	}
	return d.Connect(e)
}

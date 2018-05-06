package logger

type Level int8

var (
	printer map[Level]string = map[Level]string{
		Fatal: "Fatal",
		Info:  "Info",
		Debug: "Debug",
		Trace: "Trace",
	}
)

func (l Level) String() string {
	if data, exist := printer[l]; exist {
		return data
	}
	return "Unknown"
}

const (
	Fatal Level = iota
	Info
	Trace
	Debug
)

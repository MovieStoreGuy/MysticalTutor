package logger

import "io"

// Log is an abstraction for minimal, thread safe logger.
type Log interface {
	// Log will push the entry into logger
	// ready to be printed when possible
	// This method can block if the logger has a full buffer.
	// Any calls to Log require the logger to be started
	Log(e Entry) Log

	// Set will update the current logger to use the
	// given writen and log level.
	// Note that you can only set the write once in your application,
	// Further calls to update the writer will be ignored.
	// The log level can be updated at any point.
	Set(w io.Writer, level Level) Log

	// Start will run the logging tool in the background
	// and be ready to process any entries in the buffer.
	Start()

	// Stop will shutdown the logger but will flush what is currently
	// in the buffer.
	Stop()
}

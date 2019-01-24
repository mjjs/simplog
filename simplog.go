// Package simplog provides a general, platform agnostic logging library.
package simplog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Severity levels for the logger
const (
	EMERGENCY   string = "emerg"
	ALERT       string = "alert"
	CRITICAL    string = "crit"
	ERROR       string = "err"
	WARNING     string = "warning"
	NOTICE      string = "notice"
	INFORMATION string = "info"
	DEBUG       string = "debug"
)

// Logger represents a single logging instance. Each instance holds their own logfile file handler and the severity level.
type Logger struct {
	logFile     *os.File
	LogFileName string
	Severity    string
}

// New constructs a Logger instance, and opens a new file handle for the logfile.
// Returns a non-nil error if the filename doesn't parse into a correct filepath
// or if the file cannot be opened.
func New(filename string, severity string) (*Logger, error) {
	path, err := filepath.Abs(filename)

	if err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, err
	}

	logger := &Logger{
		LogFileName: filename,
		Severity:    severity,
		logFile:     logFile,
	}

	return logger, nil
}

// Close releases the logfile handle of the Logger instance.
func (l *Logger) Close() {
	l.logFile.Close()
}

// Write writes a log entry to the logfile. Uses the severity defined in the Logger instance.
// Returns an error if the entry cannot be written.
func (l *Logger) Write(message string) error {
	return l.writeLog(message, l.Severity)
}

// Alert writes an ALERT level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Alert(message string) error {
	return l.writeLog(message, ALERT)
}

// Crit writes an CRITICAL level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Crit(message string) error {
	return l.writeLog(message, CRITICAL)
}

// Debug writes an DEBUG level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Debug(message string) error {
	return l.writeLog(message, DEBUG)
}

// Emerg writes an EMERGENCY level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Emerg(message string) error {
	return l.writeLog(message, EMERGENCY)
}

// Err writes an ERROR level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Err(message string) error {
	return l.writeLog(message, ERROR)
}

// Info writes an INFO level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Info(message string) error {
	return l.writeLog(message, INFORMATION)
}

// Notice writes an NOTICE level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Notice(message string) error {
	return l.writeLog(message, NOTICE)
}

// Warning writes an WARNING level log entry to the logfile. Returns an error if the entry cannot be written.
func (l *Logger) Warning(message string) error {
	return l.writeLog(message, WARNING)
}

// getLogTimestamp returns a timestamp of the moment of time it was called on.
// The timestamp is in the format Dec 24 12.
func getLogTimestamp() string {
	now := time.Now()
	day := now.Day()
	month := now.Month().String()[0:3]
	year := now.Year()
	return fmt.Sprintf("%s %d %d", month, day, year)
}

// writeLog assembles the line written into the log file and writes it.
// Returns an error if the line could not be written.
func (l *Logger) writeLog(message string, severity string) error {
	line := fmt.Sprintf(
		"%s - %s - %s\n",
		getLogTimestamp(),
		severity,
		message,
	)

	_, err := l.logFile.WriteString(line)

	return err
}

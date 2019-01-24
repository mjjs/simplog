package simplog

import (
	"os"
	"regexp"
	"testing"
)

func TestNewLogger(t *testing.T) {
	const filename string = "testfile_simplog.log"
	os.Remove(filename)
	l, _ := New(filename, INFORMATION)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("The New function did not create a new logfile")
	}

	l.Close()
	os.Remove(filename)
}

func TestGetLogTimestamp(t *testing.T) {
	re := regexp.MustCompile(`(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s\d+\s\d+`)
	timestamp := getLogTimestamp()

	if !re.MatchString(timestamp) {
		t.Errorf("Timestamp did not match required regex")
	}
}

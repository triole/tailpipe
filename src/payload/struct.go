package payload

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/hpcloud/tail"
)

var (
	dateLayout = "Mon, 02 Jan 2006 03:04:05.000 MST"
)

// Payload holds the information that are transmitted
type Payload struct {
	Date         string
	LogFile      string
	Host         string
	Text         string
	TailError    error
	TailErrorStr string
}

// NewPayload returns a new payload object
func NewPayload(line *tail.Line) (p Payload) {
	return Payload{
		Date:      line.Time.Format(dateLayout),
		Host:      getHostName(),
		TailError: line.Err,
		Text:      line.Text,
	}
}

// NewTestPayload returns a test payload to send using the mail parameter
func NewTestPayload(text string) (p Payload) {
	return Payload{
		Date: time.Now().Format(dateLayout),
		Host: getHostName(),
		TailError: errors.New(
			"This is a test error. Keep calm, no real error occured.",
		),
		Text: text,
	}
}

func getHostName() (s string) {
	s, err := os.Hostname()
	if err != nil {
		fmt.Printf("Can not get hostname: %q\n", err)
	}
	return
}

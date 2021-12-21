package payload

import (
	"fmt"
	"os"

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

func getHostName() (s string) {
	s, err := os.Hostname()
	if err != nil {
		fmt.Printf("Can not get hostname: %q\n", err)
	}
	return
}

package mail

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func (m *Mail) addAttachments() {
	for _, filename := range m.AttachmentFiles {
		fileBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error adding attachment %q\n", err)
		} else {
			fileMIMEType := http.DetectContentType(fileBytes)
			fileData := base64.StdEncoding.EncodeToString(fileBytes)
			m.AttachFileName = cleanseFilename(filename)
			m.AttachMime = fileMIMEType
			m.addTemplateToBody("attachment_header.tpl")
			m.addStringToBody("\n")
			// add file data
			lineMaxLength := 64
			nbrLines := len(fileData) / lineMaxLength
			for i := 0; i < nbrLines; i++ {
				m.Body.WriteString(fileData[i*lineMaxLength:(i+1)*lineMaxLength] + "\n")
			}
			m.addStringToBody("\n")
		}
	}
}

func cleanseFilename(filename string) (s string) {
	rx := regexp.MustCompile(`[^a-zA-Z0-9-_\.]`)
	s = rx.ReplaceAllString(filename, "")
	s = strings.Trim(s, ".")
	return
}

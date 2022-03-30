package mail

import (
	"os"
	"regexp"
	"strings"
)

func (m *Mail) addAttachments() {
	for _, filename := range m.AttachmentFiles {
		m.addAttachment(filename)
	}
}

func (m *Mail) addAttachment(filename string) {
	fileBytes, err := readFile(filename)
	if err == nil {
		m.AttachFileName = cleanseFilename(filename)
		m.AttachMime = "text/plain"
		if strings.Split(m.AttachFileName, ".")[1] == "json" {
			m.AttachMime = "application/json"
		}

		// fileData := base64.StdEncoding.EncodeToString(fileBytes)
		m.addTemplateToBody("attachment_header.tpl")
		m.addStringToBody("\n")
		m.addBase64WrapToBody([]byte(fileBytes))
		// m.addStringToBody("\n")
	}
}

func cleanseFilename(filename string) (s string) {
	s = filename
	if strings.Contains(filename, string(os.PathSeparator)) {
		arr := strings.Split(filename, string(os.PathSeparator))
		s = arr[len(arr)-1]
	}
	rx := regexp.MustCompile(`[^a-zA-Z0-9-_\.]`)
	s = rx.ReplaceAllString(s, "")
	s = strings.Trim(s, ".")
	return
}

package mail

import (
	"fmt"
	"net/smtp"
	"strconv"
	"tailpipe/src/payload"
)

// SendMail sends an email via smtp
func SendMail(payload payload.Payload, mail Mail) {
	mail.Payload = payload
	mail.Payload.Host = "Just Testing"
	if mail.Payload.TailError != nil {
		mail.Payload.TailErrorStr = fmt.Sprintf("%s", mail.Payload.TailError)
	}

	auth := smtp.PlainAuth("", mail.AddrFrom, mail.Pass, mail.Host)
	if mail.Encryption == "none" {
		auth = unencryptedAuth{
			smtp.PlainAuth(
				"", mail.AddrFrom, mail.Pass, mail.Host,
			),
		}
	}

	if len(mail.AttachmentFiles) == 0 {
		// plain text mail
		mail.addTemplateToBody("header_plain.tpl")
		mail.addStringToBody("\n")
		mail.addStringToBody(payload.Text)
	} else {
		// multipart mail
		mail.Boundary = randStr(48)
		mail.addTemplateToBody("header_multipart.tpl")
		mail.addTemplateToBody("ct_header_text.tpl")
		mail.addStringToBody("\n")
		mail.addStringToBody(payload.Text)
		mail.addStringToBody("\n\n\n\n")
		mail.addAttachments()
		mail.addStringToBody("--" + mail.Boundary + "--")
	}

	if mail.Print == true {
		mail.print()
	} else {
		err := smtp.SendMail(
			mail.Host+":"+strconv.Itoa(mail.Port),
			auth, mail.AddrFrom, mail.AddrTo, mail.Body.Bytes(),
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Email sent to via %s to %s\n", mail.Host, mail.AddrTo)
	}
}

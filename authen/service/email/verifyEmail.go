package email

import (
	"crypto/tls"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/config"
	gomail "gopkg.in/mail.v2"
)

func SendVerifyUrl(email string,token string)(err error){
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Get().Email)

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Please verify you Email")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Please click on the link below to verify your Email\n" + config.Get().VerifyUrl+token)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, config.Get().Email, config.Get().EmailPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


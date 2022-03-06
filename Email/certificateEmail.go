package email

import (
	"bytes"
	"log"
	"text/template"

	mail "github.com/xhit/go-simple-mail"
)

func SendCertificateMail(email string, name string) {

	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Username = "ife.borngreat@gmail.com"
	server.Password = "uvszqidwzopjeexv"
	server.Port = 465
	server.KeepAlive = false
	server.Encryption = mail.EncryptionSSL

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	emailSent := mail.NewMSG()
	emailSent.SetFrom("Octamile@octimile.com")
	emailSent.AddTo(email)
	emailSent.SetSubject("Certificate")

	t, _ := template.ParseFiles("./email-templates/certificate.html")
	var body bytes.Buffer

	t.Execute(&body, struct {
		Email string

		Name string
	}{
		Email: email,

		Name: name,
	})

	emailSent.SetBody(mail.TextHTML, body.String())

	// Send email
	err = emailSent.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}

}

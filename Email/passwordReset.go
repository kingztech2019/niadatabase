package email

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	mail "github.com/xhit/go-simple-mail"
)

func SendPasswordToken(email string, token string)  {
	fmt.Println(email, token)
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Username="ife.borngreat@gmail.com"
	server.Password="uvszqidwzopjeexv"
	 server.Port =  465
	server.KeepAlive=false 
	server.Encryption = mail.EncryptionSSL

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	emailSent := mail.NewMSG()
	emailSent.SetFrom("Octamile@octimile.com")
	emailSent.AddTo(email)
	emailSent.SetSubject("Password Reset Token")

	t,_:= template.ParseFiles("./email-templates/reset.html")
	var body bytes.Buffer

	t.Execute(&body, struct {
		Email string
		Token string
	 }{
	   Email:email,
	   Token: token,
	 })

	 emailSent.SetBody(mail.TextHTML,  body.String())

	// Send email
	err = emailSent.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}

}

 
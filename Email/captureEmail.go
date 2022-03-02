package email

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	mail "github.com/xhit/go-simple-mail"
)

func SendCaptureUrl(email string, name string,id string)  {
	fmt.Println(email, id)
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
	emailSent.SetSubject("Capture Verification")

	t,_:= template.ParseFiles("./email-templates/imagecapture.html")
	var body bytes.Buffer

	t.Execute(&body, struct {
		Email string
		Id string
		Name string
	 }{
	   Email:email,
	   Id: id,
	   Name:name,
	 })

	 emailSent.SetBody(mail.TextHTML,  body.String())

	// Send email
	err = emailSent.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}

}

// func SendCaptureUrl(email string, id string)  {
// 	// Sender data.
// 	from := "ife.borngreat@gmail.com"
// 	password := "uvszqidwzopjeexv"
  
// 	// Receiver email address.
// 	to := []string{
// 	  email,
// 	}
  
// 	// smtp server configuration.
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"
	

  
// 	// Authentication.
// 	auth := smtp.PlainAuth("", from, password, smtpHost)
  
// 	t, _ := template.ParseFiles("./email-templates/capture.html")
  
// 	var body bytes.Buffer
  
// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))
  
// 	t.Execute(&body, struct {
// 	   Email string
// 	   ID string
// 	}{
// 	  Email:email,
// 	  ID: id,
// 	})
//   log.Println(email,id)
// 	// Sending email.
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
// 	if err != nil {
// 	  fmt.Println(err)
// 	  return
// 	}
// 	fmt.Println("Email Sent!")
	
// }
package main

import (

	"crypto/tls"	
  
  	"fmt"

	"log"

	"net/smtp"

)

func main() {

	fmt.Print(SendMail())

}


func SendMail() string {

	password := "your app password"

	from := "your email"

	to := "any email whatsoever"

	host := "smtp.mail.yahoo.com"

	port := "465"

	subject := "Hey, I'm Just Checking On You."

	body := "Hope you're doing okay! How are you doing today. "

	headers := make(map[string]string)

	headers["From"] = from

	headers["To"] = to

	headers["Subject"] = subject

	message := ""

	for k, v := range headers {

		message += fmt.Sprintf("%s: %s\r", k, v)

	}

	message += "\r" + body

	serverAddress := host + ":" + port

	authenticate := smtp.PlainAuth("", from, password, host)

	// TLS config

	tlsConfig := &tls.Config{

		InsecureSkipVerify: true,

		ServerName:         host,

	}

	conn, err := tls.Dial("tcp", serverAddress, tlsConfig)

	if err != nil {

		log.Panic(err)

	}

	c, err := smtp.NewClient(conn, host)

	if err != nil {

		log.Panic(err)

	}

	// Auth

	if err = c.Auth(authenticate); err != nil {

		log.Panic(err)

	}

	// To && From

	if err = c.Mail(from); err != nil {

		log.Panic(err)

	}

	if err = c.Rcpt(headers["To"]); err != nil {

		log.Panic(err)

	}

	// Data

	writer, err := c.Data()

	if err != nil {

		log.Panic(err)

	}

	_, err = writer.Write([]byte(message))

	if err != nil {

		log.Panic(err)

	}

	err = writer.Close()

	if err != nil {

		log.Panic(err)

	}

	err = c.Quit()

	if err != nil {

		return ""

	}

	return "Successful, the mail was sent!"

}

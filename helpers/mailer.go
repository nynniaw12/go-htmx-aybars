package helpers 

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/nynniaw12/go-htmx-aybars/creds"
)


func SendMail(body, sender, subject string) (string, string, error) {
	mg := mailgun.NewMailgun(creds.MailgunDomain, creds.MailgunAPIKey)
	
	recipient := "aybars.ari26@gmail.com"

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
        return resp, id, err
	}

    fmt.Printf("ID: %s Resp: %s\n", id, resp)
        return resp, id, nil 
}

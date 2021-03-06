package email

import (
	"context"
	pb_email "github.com/onezerobinary/email-box/proto"
	"errors"
	"time"
)

type EmailServiceServer struct {

}

func (s *EmailServiceServer) SendEmail(ctx context.Context, recipient *pb_email.Recipient) (*pb_email.EmailResponse, error) {

	emailResponse := pb_email.EmailResponse{}

	ok := SendConfirmRegistrationEmail(recipient.Email, recipient.Token)

	time.Sleep(2 * time.Second)

	if !ok {
		emailResponse.Code = 400
		return &emailResponse, errors.New("It was not possible to send the email")
	}

	emailResponse.Code = 200

	return &emailResponse, nil
}
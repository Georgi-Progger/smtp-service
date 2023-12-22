package repository

import (
	"cmd/main.go/internal/model/otp"
	"context"
	"log"
)

func (r *repo) SendOtpToMail(ctx context.Context, otp otp.OtpBody) error {
	query := `
			INSERT INTO otp(Recipient,Thema_message, Body_message, Date_of_sending) 
			VALUES($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, query, otp.Recipient, otp.ThemaMessage, otp.BodyMessage, otp.DateOfSending)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

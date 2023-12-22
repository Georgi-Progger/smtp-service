package service

import (
	"cmd/main.go/internal/model/otp"
	"cmd/main.go/internal/repository"
	"cmd/main.go/pkg/otpmailstorage"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (s *Service) SendingMessageOtp(c echo.Context) error {
	var otpBody otp.OtpBody
	if err := c.Bind(&otpBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	otpRepo := repository.NewRepo(s.Db)
	if err := otpRepo.SendOtpToMail(c.Request().Context(), otpBody); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert otp"})
	}

	emailSettings := otpmailstorage.NewSmtpSettings(
		"smtp.gmail.com", "agievrazet@gmail.com", "xeoosvtkndoslszq", "587",
	)
	emailService := otpmailstorage.NewEmailService(*emailSettings)
	err := emailService.SendEmail(otpBody.Recipient, otpBody.ThemaMessage, otpBody.BodyMessage)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Otp successfully sended"})

}

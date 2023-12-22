package otp

type OtpBody struct {
	Id            uint
	Recipient     string
	BodyMessage   string
	ThemaMessage  string
	DateOfSending string
}

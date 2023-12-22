package otpmailstorage

type smtpSettings struct {
	Server   string
	Port     string
	User     string
	Password string
}

func NewSmtpSettings(server, user, password, port string) *smtpSettings {
	return &smtpSettings{
		Server:   server,
		Port:     port,
		User:     user,
		Password: password,
	}
}

type emailService struct {
	smtpSettings
}

func NewEmailService(settings smtpSettings) *emailService {
	return &emailService{smtpSettings: settings}
}

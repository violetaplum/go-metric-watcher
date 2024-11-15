package notifier

type EmailNotifier struct {
	from     string
	password string
	host     string
	port     string
}

func NewEmailNotifier(from, password, host, port string) *EmailNotifier {
	return &EmailNotifier{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}

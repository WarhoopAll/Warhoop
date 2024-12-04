package email

import (
	"bytes"
	"grimoire/app/config"
	"html/template"
	"net/smtp"
)

var cfg = config.Get()

func GetSubject(notificationType, locale string) string {
	subjects := map[string]map[string]string{
		"verify": {
			"en": "Please verify your email address.",
			"ru": "Пожалуйста, подтвердите ваш адрес электронной почты.",
		},
		"changePassword": {
			"en": "You have successfully changed your password.",
			"ru": "Вы успешно сменили пароль.",
		},
		"changeEmail": {
			"en": "You have successfully changed your Email.",
			"ru": "Вы успешно сменили Email.",
		},
		"password": {
			"en": "Please reset your password.",
			"ru": "Пожалуйста, сбросьте ваш пароль.",
		},
		"login": {
			"en": "Sign-in from new location.",
			"ru": "Вход с нового местоположения.",
		},
	}
	return subjects[notificationType][locale]
}

type Setup struct {
	identity string
	username string
	password string
	host     string
	port     string
}

var setup *Setup

type Request struct {
	from    string
	to      []string
	subject string
	body    string
	From    string
	CC      string
	BCC     string
}

func init() {
	setup = &Setup{}
}

func LoadConfig(cfg *Setup) {
	setup = &Setup{
		identity: cfg.identity,
		username: cfg.username,
		password: cfg.password,
		host:     cfg.host,
		port:     cfg.port,
	}
}

func NewRequest(from string, to []string, subject, body string) *Request {
	return &Request{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) Send() error {
	auth := smtp.PlainAuth(setup.identity, setup.username, setup.password, setup.host)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	head := "Subject: " + r.subject + "\n"
	if len(r.From) > 0 {
		head = head + "From: " + r.From + "<" + r.from + ">\n"
	}
	if len(r.CC) > 0 {
		head = head + "CC: " + r.CC + "\n"
	}
	if len(r.BCC) > 0 {
		head = head + "BCC: " + r.BCC + "\n"
	}
	mesg := []byte(head + mime + "\n" + r.body)
	addr := setup.host + ":" + setup.port

	if err := smtp.SendMail(addr, auth, r.from, r.to, mesg); err != nil {
		return err
	}

	return nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

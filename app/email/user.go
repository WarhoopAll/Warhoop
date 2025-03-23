package email

import (
	"fmt"
	"path/filepath"
	"time"
	"warhoop/app/config"
	"warhoop/app/log"
	"warhoop/app/model/auth"
)

func SendFunction(acc *auth.Account, token, subject, templateName, urlPath string, extraData map[string]string) error {
	cfg := config.Get()

	url := ""
	if token != "" && urlPath != "" {
		url = fmt.Sprintf("%v%v?token=%v", config.Get().AppUrl, urlPath, token)
	}

	config := &Setup{
		identity: "",
		username: config.Get().MailUser,
		password: config.Get().MailPassword,
		host:     config.Get().MailServer,
		port:     config.Get().MailPort,
	}

	LoadConfig(config)

	request := NewRequest(cfg.MailUser, []string{acc.Email}, subject, "")
	request.From = cfg.MailSanderName
	request.BCC = acc.Email

	data := map[string]interface{}{
		"Name":  acc.Username,
		"Email": acc.Email,
		"URL":   url,
	}

	for k, v := range extraData {
		data[k] = v
	}

	template := filepath.Join(cfg.MailFolderTemplates, "ru", templateName)

	err := request.ParseTemplate(template, data)
	if err != nil {
		log.Get().Error("email.request.ParseTemplate",
			log.String("err", err.Error()),
			log.String("template", template),
			log.Object("data", data),
		)
		return err

	}

	log.Get().Debug("email.sent",
		log.String("recipient", template),
		log.Object("subject", data),
	)

	err = request.Send()
	if err != nil {
		log.Get().Error("email.request.Send",
			log.String("err", err.Error()),
			log.String("recipient", acc.Email),
			log.String("subject", subject),
		)
		return err
	}

	log.Get().Debug("email.sent",
		log.String("recipient", acc.Email),
		log.String("subject", subject),
	)

	return nil
}

func NotifyLogin(acc *auth.Account, ips []string, date time.Time) error {
	cfg := config.Get()
	subject := cfg.MailSanderName + " " + GetSubject("login", "ru")
	extraData := map[string]string{
		"Name": acc.Username,
		"IP":   fmt.Sprintf("%v", ips),
		"Date": date.Format("2006-01-02 15:04:05"),
	}
	return SendFunction(acc, "", subject, cfg.MailTemplateLogin, "", extraData)
}

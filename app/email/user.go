package email

import (
	"fmt"
	"warhoop/app/log"
	"warhoop/app/model"
	"path/filepath"
	"time"
)

func SendFunction(acc *model.Account, token, subject, templateName, urlPath string, extraData map[string]string) error {
	url := ""
	if token != "" && urlPath != "" {
		url = fmt.Sprintf("%v%v?token=%v", cfg.Service.AppUrl, urlPath, token)
	}

	config := &Setup{
		identity: "",
		username: cfg.Mail.User,
		password: cfg.Mail.Password,
		host:     cfg.Mail.Server,
		port:     cfg.Mail.Port,
	}

	LoadConfig(config)

	request := NewRequest(cfg.Mail.User, []string{acc.Email}, subject, "")
	request.From = cfg.Mail.SanderName
	request.BCC = acc.Email

	data := map[string]interface{}{
		"Name":  acc.Username,
		"Email": acc.Email,
		"URL":   url,
	}

	for k, v := range extraData {
		data[k] = v
	}

	template := filepath.Join(cfg.Mail.FolderTemplates, "ru", templateName)

	err := request.ParseTemplate(template, data)
	if err != nil {
		log.Get().Error("email.request.ParseTemplate",
			log.String("err", err.Error()),
			log.String("template", template),
			log.Any("data", data),
		)
		return err

	}

	log.Get().Debug("email.sent",
		log.String("recipient", template),
		log.Any("subject", data),
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

func NotifyLogin(acc *model.Account, ips []string, date time.Time) error {
	subject := cfg.Mail.SanderName + " " + GetSubject("login", "ru")
	extraData := map[string]string{
		"Name": acc.Username,
		"IP":   fmt.Sprintf("%v", ips),
		"Date": date.Format("2006-01-02 15:04:05"),
	}
	return SendFunction(acc, "", subject, cfg.Mail.TemplateLogin, "", extraData)
}

package web

import (
	"context"
	"warhoop/app/email"
	"warhoop/app/model/auth"
	"warhoop/app/model/nexus"
	"warhoop/app/utils"
)

// GetChangeEmail ...
//func (svc *WebService) GetChangeEmail(ctx context.Context, newEmail string, ips []string, acc *model.Account) error {
//
//	result, err := svc.store.Auth.GetChangeEmail(ctx, newEmail, ips, acc.ToDB())
//	if err != nil {
//		return utils.ErrDataBase
//	}
//
//	currentDate := time.Now()
//
//	res := result.ToWeb()
//	res.Locale = acc.Locale
//
//	err = svc.SendVerifyByEmail(ctx, res)
//	if err != nil {
//		return utils.ErrSendEmail
//	}
//
//	err = email.NotifyEmailChange(acc, newEmail, ips, currentDate)
//	if err != nil {
//		return utils.ErrSendEmail
//	}
//
//	return nil
//}

func (svc *WebService) SendLoginByEmail(ctx context.Context, entry *auth.Account, session *nexus.Session) error {
	err := email.NotifyLogin(entry, session.IPs, session.LoginedAt)
	if err != nil {
		return utils.ErrSendEmail
	}
	return nil
}

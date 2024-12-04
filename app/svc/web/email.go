package web

import (
	"context"
	"grimoire/app/email"
	"grimoire/app/model"
	"grimoire/app/utils"
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

func (svc *WebService) SendLoginByEmail(ctx context.Context, acc *model.Account, session *model.Session) error {
	err := email.NotifyLogin(acc, session.IPs, session.LoginedAt)
	if err != nil {
		return utils.ErrSendEmail
	}
	return nil
}

package auth

import (
	"context"
	"crypto/subtle"
	"warhoop/app/model"
	"warhoop/app/utils"
)

func (svc *AuthService) GetByID(ctx context.Context, id int) (*model.Account, error) {
	result, err := svc.store.AuthRepo.GetByID(ctx, id)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	return result.ToWeb(), nil
}

func (svc *AuthService) SignIn(ctx context.Context, entry *model.Account) (*model.Account, error) {
	fnd, err := svc.store.AuthRepo.GetByUsername(ctx, entry.Username)
	if err != nil {
		return nil, utils.ErrIncorrectLogin
	}

	_, verifier, err := utils.ConfirmVerifier(fnd.Username, entry.Password, fnd.Salt)
	if err != nil {
		return nil, utils.ErrInternal
	}

	if subtle.ConstantTimeCompare(verifier, fnd.Verifier) != 1 {
		return nil, utils.ErrIncorrectLogin
	}

	return fnd.ToWeb(), err
}

func (svc *AuthService) Exists(ctx context.Context, entry *model.Account) error {
	errChan := make(chan error, 2)

	go func() {
		exist, err := svc.store.AuthRepo.ExistsEmail(ctx, entry.Email)
		if err != nil {
			errChan <- utils.ErrDataBase
			return
		}
		if exist {
			errChan <- utils.ErrIncorrectEmail
			return
		}
		errChan <- nil
	}()

	go func() {
		exist, err := svc.store.AuthRepo.ExistsUsername(ctx, entry.Username)
		if err != nil {
			errChan <- utils.ErrDataBase
			return
		}
		if exist {
			errChan <- utils.ErrIncorrectLogin
			return
		}
		errChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

func (svc *AuthService) Create(ctx context.Context, entry *model.Account) (*model.Account, error) {
	salt, verifier, err := utils.CreateVerifier(entry.Username, entry.Password)
	if err != nil {
		return nil, utils.ErrInternal
	}

	entry.Salt = salt
	entry.Verifier = verifier

	result, err := svc.store.AuthRepo.Create(ctx, entry.ToDB())
	if err != nil {
		return nil, utils.ErrDataBase
	}

	return result.ToWeb(), nil
}

func (svc *AuthService) SignUp(ctx context.Context, entry *model.Account) (*model.Account, error) {
	err := svc.Exists(ctx, entry)
	if err != nil {
		return nil, err
	}

	result, err := svc.Create(ctx, entry)
	if err != nil {
		return nil, err
	}

	profileEntry := &model.Profile{
		AccountID: result.ID,
		Name:      result.Username,
	}
	_, err = svc.web.CreateProfile(ctx, profileEntry)
	if err != nil {
		return nil, err
	}

	return result, nil
}

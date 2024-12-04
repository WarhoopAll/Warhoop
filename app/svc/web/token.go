package web

import (
	"github.com/golang-jwt/jwt/v4"
	"grimoire/app/utils"
	"time"
)

type TokenInfo struct {
	ID        int       `json:"id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func GenerateTokenAccess(id int) (string, error) {
	if cfg == nil {
		return "", utils.ErrBadConfig
	}
	expirationTime := time.Now().Add(cfg.Cookie.AccessDuration).Unix()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime,
		"id":  id,
	})
	return claims.SignedString([]byte(cfg.Cookie.JwtKey))
}

func parse(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, utils.ErrBadToken
		}
		return []byte(cfg.Cookie.JwtKey), nil
	})
}

func TokenVerify(token string) (*TokenInfo, error) {
	parsed, err := parse(token)
	if err != nil {
		return nil, utils.ErrBadToken
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, utils.ErrBadToken
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, utils.ErrBadToken
	}
	id := int(idFloat)

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, utils.ErrBadToken
	}
	expTime := time.Unix(int64(expFloat), 0)

	if time.Now().After(expTime) {
		return nil, utils.ErrBadToken
	}

	return &TokenInfo{
		ID:        id,
		ExpiresAt: expTime,
	}, nil
}

func (svc *WebService) GenerateAccessToken(accountID int) (string, error) {
	return GenerateTokenAccess(accountID)
}

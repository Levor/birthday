package services

import (
	"errors"
	"github.com/Levor/birthday/internal/db/models"
	"github.com/Levor/birthday/internal/utils"
	base "github.com/dgrijalva/jwt-go"
)

const (
	claimAuthSub = "access"
)

type Token struct {
	jwt *utils.JWT
}

func NewToken(jwt *utils.JWT) *Token {
	return &Token{jwt: jwt}
}

func (t *Token) Issue(user *models.User) (string, error) {

	claims := base.MapClaims{
		"sub":      claimAuthSub,
		"login":    user.Login,
		"password": user.Password,
	}

	token := t.jwt.Issue(claims)
	jwtSigned, err := t.jwt.Sign(token)
	if err != nil {
		return "", err
	}
	return jwtSigned, nil
}

func (t *Token) Verify(signedToken string) (*base.Token, error) {
	token, err := t.jwt.Parse(signedToken)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(base.MapClaims)

	if claims["sub"] != claimAuthSub {
		return nil, errors.New("invalid token subject")
	}

	return token, nil
}

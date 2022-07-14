package utils

import (
	"crypto/rsa"
	"github.com/Levor/birthday/internal/config"
	base "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

type JWT struct {
	signingMethod *base.SigningMethodRSA
	secret        *rsa.PrivateKey
	public        *rsa.PublicKey
}

func NewJWT(cfg *config.Config) *JWT {
	var secretBytes []byte
	var publicBytes []byte
	secretBytes, err := ioutil.ReadFile(cfg.SecretKeyPath)
	if err != nil {
		log.Println(err)
	}
	publicBytes, err = ioutil.ReadFile(cfg.PublicKeyPath)
	if nil != err {
		log.Println(err)
	}

	secretKey, err := base.ParseRSAPrivateKeyFromPEM(secretBytes)
	if nil != err {
		log.Println(err)
	}

	publicKey, err := base.ParseRSAPublicKeyFromPEM(publicBytes)
	if nil != err {
		log.Println(err)
	}

	return &JWT{signingMethod: base.SigningMethodRS512, secret: secretKey, public: publicKey}
}

func (s *JWT) Issue(claims base.Claims) *base.Token {
	return base.NewWithClaims(s.signingMethod, claims)
}

func (s *JWT) Sign(t *base.Token) (string, error) {
	return t.SignedString(s.secret)
}

func (s *JWT) Parse(str string) (*base.Token, error) {
	return base.Parse(str, func(token *base.Token) (interface{}, error) {
		return s.public, nil
	})
}

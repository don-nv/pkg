package authorizer

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func (a *Authorizer) ParseJWT_CustomData(tokenString string, secretSign []byte, dest interface{}) error {
	var (
		keyFunc = func(t *jwt.Token) (interface{}, error) {
			return secretSign, nil
		}
		claims = JWT_Claims{}
	)

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return errors.WithMessage(err, "parsing token string with claims")
	}

	if !token.Valid {
		return errors.WithStack(jwt.ErrSignatureInvalid)
	}

	dataBytes, err := json.Marshal(claims.CustomData)
	if err != nil {
		return errors.WithMessage(err, "marshalling claims custom data")
	}

	if err = json.Unmarshal(dataBytes, dest); err != nil {
		return errors.WithMessage(err, "unmarshalling custom data bytes to destination")
	}
	return nil
}

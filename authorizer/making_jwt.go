package authorizer

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func (a *Authorizer) MakeJWT(secretSign []byte, expiresAt time.Time, customData interface{}) (string, error) {
	claims := JWT_Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		CustomData: customData,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretSign)
	if err != nil {
		return "", errors.WithMessage(err, "making token signed string")
	}
	return tokenString, nil
}

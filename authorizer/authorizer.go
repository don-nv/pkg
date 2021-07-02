package authorizer

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type API interface {
	MakeJWT(signSecret []byte, expiresAt time.Time, customData interface{}) (string, error)
}

type Authorizer struct{}

var _ API = (*Authorizer)(nil)

func New() *Authorizer {
	return &Authorizer{}
}

type JWT_Claims struct {
	jwt.StandardClaims
	CustomData interface{}
}

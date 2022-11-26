package encryption

import (
	"crypto/rsa"
	"errors"
	"github.com/evenyosua18/oauth/app/constant"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type UserInformation struct {
	name string
}

type CustomClaims struct {
	*jwt.StandardClaims
	UserInformation
}

func GenerateToken(tokenDuration, uuid, name string) (token string, expiredAt time.Time, err error) {
	var maxAge time.Duration
	maxAge, err = time.ParseDuration(tokenDuration + "h")

	if err != nil {
		return
	}

	//read rsa key
	var rsaKey []byte
	rsaKey, err = os.ReadFile("./rsa.key")

	if err != nil {
		return
	}

	var key *rsa.PrivateKey
	key, err = jwt.ParseRSAPrivateKeyFromPEM(rsaKey)
	if err != nil {
		return
	}

	expiredAt = time.Now().Add(maxAge)
	claims := make(jwt.MapClaims)
	claims[constant.ClaimsUsername] = name
	claims[constant.ClaimsId] = uuid
	claims[constant.ClaimsExpired] = expiredAt.Unix()

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return
	}

	return
}

func ValidateToken(token string) (jwt.MapClaims, error) {
	//read rsa public key
	rsaPub, err := os.ReadFile("./rsa.key.pub")
	if err != nil {
		return nil, nil
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(rsaPub)
	if err != nil {
		return nil, nil
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected methodr")
		}

		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

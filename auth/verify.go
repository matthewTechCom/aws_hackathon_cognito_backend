package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

var (
	JwksURL     = "https://cognito-idp.us-west-2.amazonaws.com/us-west-2_eyqu6wTem/.well-known/jwks.json"
	ClientID    = "1cf6e5bu3c4asa0m5dfo5blkjq"
	UserPoolURL = "https://cognito-idp.us-west-2.amazonaws.com/us-west-2_eyqu6wTem"
)

func VerifyToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	keySet, err := jwk.Fetch(context.Background(), JwksURL)
	if err != nil {
		return nil, nil, err
	}

	keyfunc := func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"].(string)
		key, found := keySet.LookupKeyID(kid)
		if !found {
			return nil, fmt.Errorf("key not found")
		}
		var raw interface{}
		if err := key.Raw(&raw); err != nil {
			return nil, err
		}
		return raw, nil
	}

	token, err := jwt.Parse(tokenString, keyfunc,
		jwt.WithAudience(ClientID),
		jwt.WithIssuer(UserPoolURL),
	)
	if err != nil {
		return nil, nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return token, claims, nil
}

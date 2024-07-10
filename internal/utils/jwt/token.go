package appjwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (j JWT) GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error) {
	accessToken, _, err := j.generateAccessToken(userID)
	if err != nil {
		ctx.Logger().Error("failed to generate access token", err, appcontext.Fields{"userID": userID})
		return "", err
	}

	return accessToken, nil
}

func (j JWT) generateAccessToken(userID string) (string, time.Time, error) {
	exp := time.Now().Add(j.accessTokenTTL)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	value, err := token.SignedString(j.accessTokenSecret)
	return value, exp, err
}

func (j JWT) ParseAccessToken(ctx *appcontext.AppContext, token string) (*Claims, error) {
	if token == "" {
		return nil, apperrors.Common.Unauthorized
	}

	// parse the token
	tokenData, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			ctx.Logger().Error("check signing method", fmt.Errorf("unexpected signing method: %v", t.Header["alg"]), appcontext.Fields{"token": token})
			return nil, apperrors.Common.Unauthorized
		}

		return j.accessTokenSecret, nil
	})

	// error
	if err != nil {
		ctx.Logger().Error("parse token", err, appcontext.Fields{"token": token})
		return nil, err
	}

	// respond
	if claims, ok := tokenData.Claims.(*Claims); ok && tokenData.Valid {
		return claims, nil
	} else {
		ctx.Logger().Error("parse claims", nil, appcontext.Fields{"token": token, "tokenData": tokenData})
		return nil, apperrors.Common.Unauthorized
	}
}

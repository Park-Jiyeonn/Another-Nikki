package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const JwtSecret = "Jiyeon_Hyomin_Jiyeon_Hyomin_hhhh"

type MyClaims struct {
	UserId               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

func GenToken(userid int64, username string) (aToken string, err error) {
	aclaims := MyClaims{
		userid,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			Issuer:    "Park.Jiyeon",
		},
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, aclaims).SignedString([]byte(JwtSecret))
	return
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return JwtSecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

package jwt

import (
	"errors"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
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

func GetUserFromCtx(ctx context.Context) (userId int64, username string) {
	c, ok := kratosJwt.FromContext(ctx)
	if !ok {
		return
	}
	C, ok := c.(jwt.MapClaims)
	if !ok {
		return
	}
	user_id, ok := C["user_id"].(float64)
	if !ok {
		return
	}
	userId = int64(user_id)
	username = C["username"].(string)
	return
}

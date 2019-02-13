package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
	"errors"
	"fmt"
)
var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// 载荷声明
type CustomClaims struct {
	Id		string	`json:":"userId`
	Name 	string	`json:"name"`
	jwt.StandardClaims
}

// 创建token
func CreateToken(ctx *gin.Context, c CustomClaims, secret string) (tokenString string, err error)  {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	// token内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": 	c.Id,
		"name":	c.Name,
		"nbf":	time.Now().Unix(),
		"iat":	time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))
}

// 获取token
func ParseRequest(c *gin.Context) (*CustomClaims, error)  {
	authorization := c.Request.Header.Get("Authorization")
	if len(authorization) == 0 {
		return  &CustomClaims{}, ErrMissingHeader
	}

	secret := viper.GetString("jwt_secret")

	var tokenString string

	//用于扫描 authorization 中的数据，并根据 format 指定的格式
	fmt.Sscanf(authorization, "Bearer %s", &tokenString)
	return ParseToekn(tokenString, secret)
}
// 解析token
func ParseToekn(tokenString string, secret string) (*CustomClaims, error)  {
	cc := &CustomClaims{}
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return  cc, err
	}

	/*if clamis, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		cc.Id = clamis["id"].(string)
		cc.Name = clamis["name"].(string)
		return cc, nil
	}*/

	clamis, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return cc, errors.New("cannot convert claim to mapclaim")
	}

	//验证token，如果token被修改过则为false
	if !token.Valid {
		return cc, errors.New("token is invalid")
	}

	cc.Id = clamis["id"].(string)
	cc.Name = clamis["name"].(string)

	return cc, err
}

// 验证密钥格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}
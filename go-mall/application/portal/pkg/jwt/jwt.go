package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT 签名结构
type JWT struct {
	config     JwtConfig
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("授权已过期")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("令牌非法")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "0df9b8db-6f7c-d713-eeab-ecb317696042"
)

type Member struct {
	Username string // 用户名
}

type JwtConfig struct {
	TimeOut time.Duration //超时时间
	Issuer  string        //签证签发人
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserInfo Member `json:"userInfo"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT(config JwtConfig) *JWT {
	return &JWT{
		config,
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 设置SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateUserToken 生成含有用户信息的token
func (j *JWT) CreateUserToken(u *Member) (string, error) {
	jwtConfig := j.config
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserInfo: *u,
		StandardClaims: jwt.StandardClaims{
			//设置一小时时效
			ExpiresAt: time.Now().Add(jwtConfig.TimeOut * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwtConfig.Issuer,
		},
	})
	return claims.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

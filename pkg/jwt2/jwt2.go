package jwt2

import (
	"errors"
	"strings"
	"time"

	"github.com/VENI-VIDIVICI/plus/pkg/app"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
)

// 定义错误

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// 定义 JWT

type JWT struct {
	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

// 定义 JWTCUStomer

type JWTCustomClaims struct {
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	ExpireAtTime time.Time `json:"expire_time"`
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenStr, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}
	token, err := jwt.parseTokenString(tokenStr)
	if err != nil {
		if err == jwtpkg.ErrTokenExpired {
			return nil, ErrTokenExpired
		}
		if err == jwtpkg.ErrTokenMalformed {
			return nil, ErrTokenInvalid
		}
		return nil, ErrTokenInvalid
	}
	if cliam, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return cliam, nil
	}
	return nil, ErrTokenInvalid
}

func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenStr, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return "", err
	}
	token, err := jwt.parseTokenString(tokenStr)
	if err != nil {
		if err != jwtpkg.ErrTokenExpired {
			return "", err
		}
	}
	claims := token.Claims.(*JWTCustomClaims)
	x := app.TimenowInTimezone().Add(jwt.MaxRefresh)
	if x.Unix() > claims.IssuedAt.Unix() {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireAtTime())
		return jwt.createToken(*claims)
	}
	return "", ErrTokenInvalid

}

func (jwt *JWT) IssueToken(userID, userName string) string {
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime,
		jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimenowInTimezone()), // 签名生效时间
			IssuedAt:  jwtpkg.NewNumericDate(app.TimenowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwtpkg.NewNumericDate(expireAtTime),            // 签名过期时间
			Issuer:    config.GetString("app.name"),                   // 签名颁发者
		},
	}
	tokenStr, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return tokenStr
}
func (jwt *JWT) createToken(claim JWTCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claim)
	return token.SignedString(jwt.SignKey)
}

func (jwt *JWT) expireAtTime() time.Time {
	timenow := app.TimenowInTimezone()
	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	duration := time.Duration(expireTime) * time.Minute
	return timenow.Add(duration)
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(t *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Auth")
	if authHeader == "" {
		return "", ErrTokenInvalid
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] != "Bear") {
		return "", ErrTokenInvalid
	}
	return parts[1], nil
}

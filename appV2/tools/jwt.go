package tools

import (
	"errors"
	_ "errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"time"
	_ "time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 30 * 24 * time.Hour
	TokenIssuer          = "1234-library"
)

type VoteJwt struct {
	Secret []byte
}

var Token VoteJwt

// 创建一个TokenSecret
func init() {
	//s string
	b := []byte("1234")
	//if s != "" {
	//	b = []byte(s)
	//}
	Token = VoteJwt{Secret: b}
}

// Claim 自定义的数据结构，这里使用了结构体的组合
type Claim struct {
	jwt.RegisteredClaims
	ID   int64  `json:"user_id"`
	Name string `json:"username"`
}

func (j *VoteJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}
func (j *VoteJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.Secret, nil
}

// GetToken 颁发token access token 和 refresh token
func (j *VoteJwt) GetToken(id int64, name string) (aToken, rToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: j.getTime(AccessTokenDuration),
		Issuer:    TokenIssuer,
	}
	claim := Claim{
		ID:               id,
		Name:             name,
		RegisteredClaims: rc,
	}
	//根据加密算法 声明的信息 以及加密密钥来生成一个token
	fmt.Println("j.Secret:", j.Secret)
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(j.Secret)
	// refresh token 不需要保存任何用户信息
	rc.ExpiresAt = j.getTime(RefreshTokenDuration)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(j.Secret)
	return
}

// 验证Token
func (j *VoteJwt) VerifyToken(tokenID string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenID, claim, j.keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		fmt.Println(err.Error())
		return nil, errors.New("access token 验证失败")
	}
	return claim, nil
}

// RefreshToken用refresh token 刷新 access token
func (j *VoteJwt) RefreshToken(a, r string) (aToken, rToken string, err error) {
	//r无效直接返回
	if _, err = jwt.Parse(r, j.keyFunc); err != nil {
		return
	}
	// 从旧access token 中解析出claims数据
	claim := &Claim{}
	_, err = jwt.ParseWithClaims(a, claim, j.keyFunc)
	// 判断错误是不是因为access token 正常过期导致的
	if errors.Is(err, jwt.ErrTokenExpired) {
		return j.GetToken(claim.ID, claim.Name)
	}
	return
}

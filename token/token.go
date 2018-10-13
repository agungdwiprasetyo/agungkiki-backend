package token

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"

	model "github.com/agungdwiprasetyo/agungkiki-backend/src/model"
	"gopkg.in/mgo.v2/bson"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Audience string      `json:"audience,omitempty"`
	User     *model.User `json:"user,omitempty"`
}

func NewClaim(dataUser *model.User) *Claim {
	aud := "guest"
	if dataUser.Role != nil {
		aud = dataUser.Role.StringID
	}
	userClaims := model.User{
		ID:       dataUser.ID,
		Username: dataUser.Username,
	}
	cl := new(Claim)
	cl.Audience = aud
	cl.User = &userClaims
	return cl
}

type Token struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	age        time.Duration
}

func New(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, age time.Duration) *Token {
	tok := new(Token)
	tok.privateKey = privateKey
	tok.publicKey = publicKey
	tok.age = age
	return tok
}

func (t *Token) Generate(cl *Claim) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(t.age).Unix()
	claims["iat"] = time.Now().Unix()
	claims["aud"] = cl.Audience
	claims["account"] = cl.User
	token.Claims = claims
	tokenString, err := token.SignedString(t.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *Token) Refresh(tokenString string) (string, error) {
	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return "", fmt.Errorf("Token Invalid")
	}
	result, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.publicKey, nil
	})
	if !result.Valid {
		return "", fmt.Errorf("Token Invalid")
	}
	claims := result.Claims.(jwt.MapClaims)

	tok := jwt.New(jwt.SigningMethodRS256)
	claims["exp"] = time.Now().Add(t.age).Unix()
	tok.Claims = claims
	tokenString, err := tok.SignedString(t.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *Token) Extract(tokenString string) (bool, *Claim) {
	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return false, nil
	}
	result, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.publicKey, nil
	})

	claims := new(Claim)
	if result.Valid {
		mapClaims := result.Claims.(jwt.MapClaims)
		claims.Audience = fmt.Sprint(mapClaims["aud"])
		account, ok := mapClaims["account"].(map[string]interface{})
		if !ok {
			return false, claims
		}
		claims.User = &model.User{
			ID:       bson.ObjectIdHex(fmt.Sprint(account["id"])),
			Username: fmt.Sprint(account["username"]),
		}
	}
	return result.Valid, claims
}

package dcosauth

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CheckExpired(tokenString string, threshold int) bool {
	b64claims := strings.Split(tokenString, ".")[1]

	claimsJSON, err := base64.RawStdEncoding.DecodeString(b64claims)

	if err != nil {
		log.Fatal(err)
	}

	var claims claimSet
	err = json.Unmarshal(claimsJSON, &claims)

	if err != nil {
		log.Fatal(err)
	}

	minValidTime := float64(time.Now().Add(time.Second * time.Duration(threshold)).Unix())

	return float64(claims.Exp) < minValidTime
}

func GenerateServiceLoginToken(privateKey []byte, uid string, validTime int) (loginToken string, err error) {
	// Parse the key
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Second * time.Duration(validTime)).Unix(),
	})

	// Sign with key and return
	return token.SignedString(key)
}

func GenerateServiceLoginObject(privateKey []byte, uid string, validTime int) (loginObject []byte, err error) {
	token, err := GenerateServiceLoginToken(privateKey, uid, validTime)

	m := serviceLoginObject{
		UID:   uid,
		Token: token,
	}

	return json.Marshal(m)
}

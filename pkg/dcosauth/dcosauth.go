package dcosauth

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type serviceLoginObject struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type claimSet struct {
	UID string `json:"uid"`
	Exp int    `json:"exp"`
	// *StandardClaims
}

func createClient() *http.Client {
	// // Create transport to skip verify TODO: add certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	} // TODO: add timeouts here

	return client
}

// CheckExpired checks if a token will expire within the refreshThreshold
func CheckExpired(tokenString string, refreshThreshold int) (expired bool, err error) {
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

	minValidTime := float64(time.Now().Add(time.Second * time.Duration(refreshThreshold)).Unix())

	return float64(claims.Exp) < minValidTime, nil
}

// Login acquires and returns a new JWT token by authenticating to the DC/OS api with a uid and private key
func Login(master string, loginObject []byte) (authToken string, err error) {

	// Build client
	client := createClient()

	// Build request
	url := "https://" + master + "/acs/api/v1/auth/login"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(loginObject))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Todo better error handling (after read response, cause will eventually use body)
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Unable to login (Invalid credentials?)")
	}

	// Parse body
	var dat loginResponse
	err = json.Unmarshal(body, &dat)
	if err != nil {
		return "", err
	}

	return dat.Token, nil
}

// GenerateServiceLoginToken generates a JWT login token
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

// GenerateServiceLoginObject returns a JSON object containing a uid and a token generated with GenerateServiceLoginToken
func GenerateServiceLoginObject(privateKey []byte, uid string, validTime int) (loginObject []byte, err error) {
	token, err := GenerateServiceLoginToken(privateKey, uid, validTime)

	m := serviceLoginObject{
		UID:   uid,
		Token: token,
	}

	return json.Marshal(m)
}

// Output writes given content to a given filepath
func Output(content []byte, outputFilePath string) (err error) {
	err = nil
	if outputFilePath != "" {
		err = ioutil.WriteFile(outputFilePath, []byte(content), 0600)
	} else {
		fmt.Println(string(content))
	}

	return err
}

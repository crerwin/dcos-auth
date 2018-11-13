package dcosauth

import (
	"net/http"
)

// DCOSAuth is our main authorization object
type DCOSAuth struct {
	apiclient           *http.Client
	Master              string
	token               string
	UID                 string
	privateKey          string
	ValidTime           int
	ExpirationThreshold int
}

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

// New returns a pointer to a new DCOSAuth object
func New(master string, uid string, privateKey string) *DCOSAuth {
	return &DCOSAuth{
		apiclient:           createClient(),
		Master:              master,
		UID:                 uid,
		privateKey:          privateKey,
		ValidTime:           900,
		ExpirationThreshold: 60,
	}
}

// Token returns the current token if it hasn't expired, otherwise it acquires and returns a new token
func (d *DCOSAuth) Token() (token string, err error) {
	if d.token == "" || CheckExpired(d.token, d.ExpirationThreshold) {
		d.Login()
	}
	return d.token, nil
}

// Login acquires and returns a new JWT token by authenticating to the DC/OS api with a uid and private key
func (d *DCOSAuth) Login() {

	lo, _ := GenerateServiceLoginObject([]byte(d.privateKey), d.UID, d.ValidTime)

	// Build client
	client := createClient()
	d.token, _ = login(d.Master, lo, client)
}

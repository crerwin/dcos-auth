package dcosauth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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

// non-exported login function allows us to mock http.Client
func login(master string, loginObject []byte, client *http.Client) (authToken string, err error) {

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
	authToken, err = parseLoginResponse(body)
	if err != nil {
		return "", err
	}
	return authToken, nil
}

func parseLoginResponse(body []byte) (string, error) {
	var dat loginResponse
	err := json.Unmarshal(body, &dat)
	if err != nil {
		return "", err
	}
	return dat.Token, nil
}

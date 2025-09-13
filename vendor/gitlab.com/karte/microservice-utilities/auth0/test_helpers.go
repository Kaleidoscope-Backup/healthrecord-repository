package auth0

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetAuthToken() *AuthToken {

	c := setupTestEnv()

	payload := strings.NewReader("{\"client_id\":\"" + c.ClientID + "\",\"client_secret\":\"" + c.ClientSecret + "\",\"audience\":\"" + c.Audience + "\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", c.AuthURL, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//unmarshal to interface
	var at *AuthToken
	json.Unmarshal(body, &at)

	return at
}

type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type TestAuthConfig struct {
	// ClientID
	ClientID string

	// ClientSecret
	ClientSecret string

	// AuthURL
	AuthURL string

	// Audience
	Audience string
}

func setupTestEnv() TestAuthConfig {

	config := TestAuthConfig{}

	val, bool := os.LookupEnv("CLIENT_ID")
	if bool {
		config.ClientID = val
	}

	val, bool = os.LookupEnv("CLIENT_SECRET")
	if bool {
		config.ClientSecret = val
	}

	val, bool = os.LookupEnv("AUTH_URL")
	if bool {
		config.AuthURL = val
	}

	val, bool = os.LookupEnv("AUTH_AUDIENCE")
	if bool {
		config.Audience = val
	}

	return config

}

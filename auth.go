package main

import (
  "fmt"
  "net/http"
  "os"
  "bufio"
  "log"

  "github.com/garyburd/go-oauth/oauth"
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func get_consumer_key(config *Config) {
  if config.ConsumerKey == "" || config.ConsumerSecret == "" {
    config.ConsumerKey = "Ed6im4kCMG6conJUbWvdXUjJU"
    config.ConsumerSecret = "e3jCzBZ70rvIAN8aVoxD8nNN29GEQTu5Jiaj7Tj8lvvKwSTXNP"
  }
}

func get_access_token(config *Config) (*oauth.Credentials, bool, error) {

  oauthClient.Credentials.Token = config.ConsumerKey
  oauthClient.Credentials.Secret = config.ConsumerSecret

  authorized := false

  var token *oauth.Credentials

  if config.AccessToken != "" && config.AccessSecret != "" {

    authorized = true
    token = &oauth.Credentials{config.AccessToken, config.AccessSecret}

  } else {

    request_token, err := oauthClient.RequestTemporaryCredentials(http.DefaultClient, "", nil)

    if err != nil {
      fmt.Println("failed to request temporary credentials: ", err)
			return nil, false, err
    }

    token := client_authentication(request_token)

    config.AccessToken = token.Token
    config.AccessSecret = token.Secret
    authorized = true
  }

  return token, authorized, nil
}

func client_authentication(request_token *oauth.Credentials) *oauth.Credentials {

  url := oauthClient.AuthorizationURL(request_token, nil)

  fmt.Println("Open this URL and log into Twitter: ")
  
  fmt.Println(url)

  fmt.Println("Enter the PIN: ")

  fmt.Print("PIN: ")

  stdin := bufio.NewScanner(os.Stdin)
  if !stdin.Scan() {
    log.Fatal("stopped")
  }

  token, _, err := oauthClient.RequestToken(http.DefaultClient, request_token, stdin.Text())

  if err != nil {
    log.Fatal("token request failed: ", err)
  }

  return token
}

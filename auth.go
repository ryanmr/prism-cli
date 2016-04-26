package main

import (
  "fmt"
  "net/http"
  "os"
  "bufio"
  "log"
  "strings"

  "github.com/garyburd/go-oauth/oauth"
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func authorize() {
  fmt.Println(`
    You will need to authorize this application with your Twitter account.
    To do so, you will be presented a URL, which will take you to Twitter.
    After you log into Twitter, you will be given a PIN.
    Entering the PIN here will authroize this application.

    If you are already authorized, you will not need to be authorized again.
  `)

  _, config, err := get_config()

  if err != nil {
    get_consumer_key(&config)
  }

  _, authorized, err := get_access_token(&config)

  if authorized {
    save_config(&config)
  }

  fmt.Println(`
    You are authorized. Enjoy Prism.
  `)
}

func get_consumer_key(config *Config) {
  // defaulting to my application keys, but I can add support for anyone's keys later
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

  fmt.Println(`
    Open this URL and log into Twitter:
    When shown your PIN, enter it!

  `)

  fmt.Println(url)


  fmt.Print("\tPIN:")

  stdin := bufio.NewScanner(os.Stdin)
  if !stdin.Scan() {
    log.Fatal("stopped")
  }

  token, _, err := oauthClient.RequestToken(http.DefaultClient, request_token, strings.TrimSpace(stdin.Text()))

  if err != nil {
    log.Fatal("token request failed: ", err)
  }

  return token
}

func has_consumer(config *Config) {
  return config.ConsumerKey != "" && config.ConsumerSecret != ""
}

func has_access(config *Config) {
  return config.AccessToken != "" && config.AccessSecret != ""
}

func has_credentials(config *Config) {
  return has_consumer(&config) && has_access(&config)
}

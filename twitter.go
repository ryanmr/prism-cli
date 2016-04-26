package main

import (
  "fmt"
  // "os"
  // "net/http"
  // "encoding/json"
  // "io"
  // "io/ioutil"
  // "bufio"
  // "strings"
  // "github.com/codegangsta/cli"
  "github.com/ChimeraCoder/anaconda"
  // "github.com/garyburd/go-oauth/oauth"
)

func get_api() *anaconda.TwitterApi {

  _, config, err := get_config()

  if err != nil {
    fmt.Println(`
      The configuration could not be read.

      Please run ./prism authorize.
    `)
  }

  if !has_credentials(&config) {
    fmt.Println(`
      You have not authorized the app.

      Please run ./prism authorize.
    `)
  }

  anaconda.SetConsumerKey(config.ConsumerKey)
  anaconda.SetConsumerSecret(config.ConsumerSecret)

  api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecret)

  return api
}

func show_tweets() {
  
  api := get_api()

  tweets, _ := api.GetUserTimeline(nil)

  for _, tweet := range tweets {
      fmt.Println(tweet.Text)
  }

}

// func get_tweets() {
//
// }

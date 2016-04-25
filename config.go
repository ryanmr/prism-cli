package main

import (
  // "os"
  // "fmt"
  "io/ioutil"
  "log"
  "encoding/json"
)

type Config struct {
  Username string       `json:"username"`
  TwitterId string     `json:"twitter_id"`
  ConsumerKey string        `json:"consumer_key"`
  ConsumerSecret string     `json:"consumer_secret"`
  AccessToken string          `json:"access_token"`
  AccessSecret string         `json:"access_secret"`
}

func get_config() (string, Config, error) {
  filename := "config.prism.json"

  config := Config{}

  b, err := ioutil.ReadFile(filename)

  if err != nil {
    return filename, config, err
  } else {
    err = json.Unmarshal(b, &config)
    if err != nil {
      return filename, config, err
    }
  }

  return filename, config, nil
}

func save_config(config *Config) {
  filename := "config.prism.json"

  b, err := json.MarshalIndent(config, "", " ")

  if err != nil {
    log.Fatal("failed to convert Config to JSON: ", err)
  }

  err = ioutil.WriteFile(filename, b, 0700)

  if err != nil {
    log.Fatal("failed to save file", err)
  }
}

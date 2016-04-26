package main

import (
  // "fmt"
  "os"
  // "net/http"
  // "encoding/json"
  // "io"
  // "io/ioutil"
  // "bufio"
  // "strings"
  "github.com/codegangsta/cli"
  // "github.com/ChimeraCoder/anaconda"
  // "github.com/garyburd/go-oauth/oauth"
)


func main() {
  handleArguments()
}

func handleArguments() {
  app := cli.NewApp()
  app.Name = "Prism"
  app.Usage = "A helper for making Tweets into static entries"
  app.Action = func(c *cli.Context) {
    println("A helper for making Tweets into static entries")
  }

  app.Commands = []cli.Command {
    {
      Name: "authorize",
      Usage: "authorize this app with Twitter",
      Action: func(c *cli.Context) {
        authorize()
      },
    },
    {
      Name: "force-authorize",
      Usage: "force authorize this app with Twitter",
      Action: func(c *cli.Context) {
        force_authorize()
      },
    },
    {
      Name: "tweets",
      Usage: "show your tweets from your Twitter account",
      Action: func(c *cli.Context) {

        show_tweets()
      },
    },
  }

  app.Run(os.Args)
}

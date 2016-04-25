package main

import (
  "fmt"
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
  }

  app.Run(os.Args)
}

func authorize() {
  fmt.Println(`
    You will need to authorize this application with your Twitter account.
    To do so, you will be presented a URL, which will take you to Twitter.
    After you log into Twitter, you will be given a PIN.
    Entering the PIN here will authroize this application.
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

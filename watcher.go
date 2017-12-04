package main

import (
  "fmt"
  "context"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
)

func main() {
  ctx := context.Background()
  client := github.NewClient(nil)
  orgs, _, err := client.Organizations.List(ctx, "hnq90", nil)
}

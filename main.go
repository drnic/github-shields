package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"code.google.com/p/goauth2/oauth"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/google/go-github/github"
)

var client *github.Client

func pullHandler(render render.Render, r *http.Request, params martini.Params) {
	organization := params["org"]
	repo := params["repo"]
	pullRequestID := params["pull_id"]

	log.Printf("getting github PR %s/%s #%s", organization, repo, pullRequestID)
	url := fmt.Sprintf("https://github.com/%s/%s/pull/%s", organization, repo, pullRequestID)
	render.Redirect(url)
}

func main() {
	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if githubAccessToken == "" {
		fmt.Println("Create new tokens via https://github.com/settings/applications 'Personal Access Tokens' section")
		log.Fatalln("Please set environment variable $GITHUB_ACCESS_TOKEN")
	}
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: githubAccessToken},
	}

	client = github.NewClient(t.Client())

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/github/:org/:repo/pull/:pull_id", pullHandler)
	m.Run()

}

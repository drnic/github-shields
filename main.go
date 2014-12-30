package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"code.google.com/p/goauth2/oauth"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/google/go-github/github"
)

var client *github.Client

func prRedirectHandler(render render.Render, r *http.Request, params martini.Params) {
	organization := params["org"]
	repo := params["repo"]
	pullRequestID := params["pull_id"]

	url := fmt.Sprintf("https://github.com/%s/%s/pull/%s", organization, repo, pullRequestID)
	render.Redirect(url)
}

func prBadgeHandler(render render.Render, r *http.Request, params martini.Params) {
	organization := params["org"]
	repo := params["repo"]
	pullRequestID := params["pull_id"]
	badgeType := params["badge_type"]
	log.Printf("getting github PR %s %s #%s %s", organization, repo, pullRequestID, badgeType)
	if (badgeType != "png") && (badgeType != "json") {
		badgeType = "svg"
	}

	status := "open"
	color := "green"

	badgeURL, err := url.Parse("https://img.shields.io")
	if err != nil {
		panic("boom")
	}

	badgeURL.Path += fmt.Sprintf("/badge/%s PR #%s-%s-%s.%s", repo, pullRequestID, status, color, badgeType)

	log.Println("redirecting to", badgeURL)
	render.Redirect(badgeURL.String())
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
	m.Get("/github/:org/:repo/pull/:pull_id.(?P<badge_type>(svg|png|json))", prBadgeHandler)
	m.Get("/github/:org/:repo/pull/:pull_id", prRedirectHandler)
	m.Run()

}

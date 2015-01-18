package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

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

func issueRedirectHandler(render render.Render, r *http.Request, params martini.Params) {
	organization := params["org"]
	repo := params["repo"]
	issueID := params["issue_id"]

	url := fmt.Sprintf("https://github.com/%s/%s/issues/%s", organization, repo, issueID)
	render.Redirect(url)
}

func prBadgeHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	organization := params["org"]
	repo := params["repo"]
	badgeType := params["badge_type"]
	if (badgeType != "png") && (badgeType != "json") {
		badgeType = "svg"
	}

	style := r.URL.Query().Get("style")
	status := "unknown"
	color := "lightgrey"

	pullRequestID, err := strconv.Atoi(params["pull_id"])
	if err == nil {
		log.Printf("getting github PR %s %s #%d %s", organization, repo, pullRequestID, badgeType)
		pr, _, err := client.PullRequests.Get(organization, repo, int(pullRequestID))
		if err == nil {
			fmt.Printf("PR %s %s: merged? %t state: %s\n", organization, repo, *pr.Merged, *pr.State)
			if *pr.Merged {
				status = "merged"
				color = "6e5494"
			} else if *pr.State == "open" {
				status = "open"
				color = "green"
			} else if *pr.State == "closed" {
				status = "closed"
				color = "red"
			}
		}
	}
	log.Printf("%s %s #%s %s %s", organization, repo, params["pull_id"], status, color)

	badgeURL := buildBadgeURL(pullRequestID, repo, status, color, badgeType, style)

	log.Println("redirecting to", badgeURL)
	w.Header().Set("Cache-Control", "no-cache")
	http.Redirect(w, r, badgeURL.String(), 302)
}

func issueBadgeHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	organization := params["org"]
	repo := params["repo"]
	badgeType := params["badge_type"]
	if (badgeType != "png") && (badgeType != "json") {
		badgeType = "svg"
	}

	style := r.URL.Query().Get("style")
	status := "unknown"
	color := "lightgrey"

	issueID, err := strconv.Atoi(params["issue_id"])
	if err == nil {
		log.Printf("getting github issue %s %s #%d %s", organization, repo, issueID, badgeType)
		issue, _, err := client.Issues.Get(organization, repo, int(issueID))
		if err == nil {
			fmt.Printf("Issue %s %s: state: %s\n", organization, repo, *issue.State)
			if *issue.State == "open" {
				status = "open"
				color = "green"
			} else if *issue.State == "closed" {
				status = "closed"
				color = "red"
			}
		}
	}
	log.Printf("%s %s #%s %s %s", organization, repo, params["issueID"], status, color)

	badgeURL := buildBadgeURL(issueID, repo, status, color, badgeType, style)

	log.Println("redirecting to", badgeURL)
	w.Header().Set("Cache-Control", "no-cache")
	http.Redirect(w, r, badgeURL.String(), 302)
}

func buildBadgeURL(id int, repo, status, color, format, style string) *url.URL {
	url, err := url.Parse("https://img.shields.io")
	if err != nil {
		panic("boom")
	}

	repo = strings.Replace(repo, "-", "--", -1)
	repo = strings.Replace(repo, "_", "__", -1)

	url.Path += fmt.Sprintf("/badge/%s #%d-%s-%s.%s", repo, id, status, color, format)
	if style != "" {
		query := url.Query()
		query.Add("style", style)
		url.RawQuery = query.Encode()
	}

	return url
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

	m.Get("/github/:org/:repo/pull/(?P<pull_id>\\d+).(?P<badge_type>(svg|png|json))", prBadgeHandler)
	m.Get("/github/:org/:repo/pull/:pull_id", prRedirectHandler)
	m.Get("/github/:org/:repo/issues/(?P<issue_id>\\d+).(?P<badge_type>(svg|png|json))", issueBadgeHandler)
	m.Get("/github/:org/:repo/issues/:issue_id", issueRedirectHandler)

	// Redirect to blog post for any other route (e.g. root route) until some human website implemented
	m.NotFound(func(render render.Render) {
		render.Redirect("https://blog.starkandwayne.com/2014/12/30/live-github-pr-status-in-your-blogs-docs/")
	})
	m.Run()
}

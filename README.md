Badges for status of GitHub Pull Requests
=========================================

How it works
------------

To get a badge/shield, include `.svg` or `.png` suffix:

https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8.svg redirects to https://img.shields.io/badge/lattice%20PR%20%238-open-green.svg

Whilst without the suffix, the redirect is the the GitHub PR URL:

https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8 directs to https://github.com/pivotal-cf-experimental/lattice/pull/8

[![pivotal-cf-experimental/lattice/pull/8](https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8.svg)](https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8)

Deploying to Cloud Foundry
--------------------------

If you want shields that describe private repositories then you'll need to run this app with your own GitHub token.

First, create a GitHub Personal Access Token at https://github.com/settings/applications

```
cf push github-shields -m 128M -k 256M --no-start --random-route
cf set-env github-shields GITHUB_ACCESS_TOKEN <TOKEN>
cf start github-shields
```

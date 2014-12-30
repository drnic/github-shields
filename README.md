GitHub Pull Request/Issue status badges/shields
===============================================

You can now document the live status of GitHub Pull Requests in your documentation/blogs.

-	[![pivotal-cf-experimental/lattice/pull/7](https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/7.svg)](https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/7)
-	[![cloudfoundry/cloud_controller_ng/pull/316](https://github-shields.cfapps.io/github/cloudfoundry/cloud_controller_ng/pull/316.svg?1)](https://github-shields.cfapps.io/github/cloudfoundry/cloud_controller_ng/pull/316)
-	[![hashicorp/terraform/pull/708](https://github-shields.cfapps.io/github/hashicorp/terraform/pull/708.svg?1)](https://github-shields.cfapps.io/github/hashicorp/terraform/pull/708)

And issues:

-	[![golang/go/issues/498](https://github-shields.cfapps.io/github/golang/go/issues/498.svg)](https://github-shields.cfapps.io/github/golang/go/issues/498)
-	[![go-martini/martini/issues/317](https://github-shields.cfapps.io/github/go-martini/martini/issues/317.svg)](https://github-shields.cfapps.io/github/go-martini/martini/issues/317)

Badges are being rendered by the awesome http://shields.io/ service.

How it works
------------

To get a badge/shield, include `.svg` or `.png` suffix:

https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8.svg redirects to https://img.shields.io/badge/lattice%20PR%20%238-open-green.svg

Whilst without the suffix, the redirect is the the GitHub PR URL:

https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8 directs to https://github.com/pivotal-cf-experimental/lattice/pull/8

Compose the two together to get a clickable shield for a PR status: [![pivotal-cf-experimental/lattice/pull/8](https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8.svg)](https://github-shields.cfapps.io/github/pivotal-cf-experimental/lattice/pull/8)

Deploying to Cloud Foundry
--------------------------

If you want shields that describe private repositories then you'll need to run this app with your own GitHub token.

First, create a GitHub Personal Access Token at https://github.com/settings/applications

```
cf push github-shields -m 128M -k 256M --no-start --random-route
cf set-env github-shields GITHUB_ACCESS_TOKEN <TOKEN>
cf start github-shields
```

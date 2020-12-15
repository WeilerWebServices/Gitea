# LGTM

[![Build Status](https://drone.gitea.io/api/badges/go-gitea/lgtm/status.svg)](https://drone.gitea.io/go-gitea/lgtm)
[![Join the Discord chat at https://discord.gg/NsatcWJ](https://img.shields.io/discord/322538954119184384.svg)](https://discord.gg/NsatcWJ)
[![Join the Matrix chat at https://matrix.to/#/#gitea:matrix.org](https://img.shields.io/badge/matrix-%23gitea%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gitea:matrix.org)
[![](https://images.microbadger.com/badges/image/gitea/lgtm.svg)](https://microbadger.com/images/gitea/lgtm "Get your own image badge on microbadger.com")
[![codecov](https://codecov.io/gh/go-gitea/lgtm/branch/master/graph/badge.svg)](https://codecov.io/gh/go-gitea/lgtm)
[![Go Report Card](https://goreportcard.com/badge/code.gitea.io/gitea)](https://goreportcard.com/report/code.gitea.io/gitea)
[![GoDoc](https://godoc.org/code.gitea.io/gitea?status.svg)](https://godoc.org/code.gitea.io/gitea)
[![GitHub release](https://img.shields.io/github/release/go-gitea/lgtm.svg)](https://github.com/go-gitea/lgtm)

LGTM is a simple pull request approval system using GitHub protected branches
and maintainers files or maintainers groups. Pull requests are locked and cannot
be merged until the minimum number of approvals are received. Project
maintainers can indicate their approval by commenting on the pull request and
including LGTM (looks good to me) in their approval text.

## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](https://dl.gitea.io/lgtm). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/go-gitea/homebrew-gitea).
If you have questions that are not covered by the documentation, you can get 
in contact with us on  our [Discord server](https://discord.gg/NsatcWJ), 
[Matrix room](https://matrix.to/#/#gitea:matrix.org), 
or [forum](https://discourse.gitea.io/)!. If you find a security issue
please contact security@gitea.io first.

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`. It is also possible to just simply execute the
`go get github.com/go-gitea/lgtm` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/go-gitea/lgtm
cd $GOPATH/src/github.com/go-gitea/lgtm
make clean build

bin/lgtm -h
```

### Docker

A Docker Image is available for easy deployment. It can be run locally or on a dedicated Server as follows:

```
docker run --name lgtm -v /my/host/path:/var/lib/lgtm:z -e GITHUB_CLIENT= -e GITHUB_SECRET= -p 8000:8000 gitea/lgtm
```

To Fill the Environment Variables `GITHUB_CLIENT` and `GITHUB_SECRET`, create new OAuth Application [here](https://github.com/settings/applications/new)

* Homepage URL = protocol://host:port (f.e. http://localhost:8000)
* Authorization callback URL = protocol://host:port/login (f.e. http://localhost:8000/login)


To Build the Image by yourself please refere to the [Dockerfile](https://github.com/go-gitea/lgtm/blob/master/Dockerfile) and the [Drone Configuration](https://github.com/go-gitea/lgtm/blob/master/.drone.yml).


## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Maintainers](https://github.com/orgs/go-gitea/people)
* [Contributors](https://github.com/go-gitea/lgtm/graphs/contributors)

## License

This project is under the Apache-2.0 License. See the [LICENSE](LICENSE) file
for the full license text.

## Copyright

```
Copyright (c) 2018 The Gitea Authors <https://gitea.io>
```

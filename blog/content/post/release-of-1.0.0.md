---
date: "2016-12-23T08:00:00+02:00"
author: "tboerger"
title: "Release of 1.0.0"
tags: ["release"]
draft: false
---

We proudly present the release of Gitea version 1.0.0. We have closed [44](https://github.com/go-gitea/gitea/issues?q=is%3Aissue+is%3Aclosed+milestone%3A1.0.0) issues and we also merged [238](https://github.com/go-gitea/gitea/pulls?utf8=%E2%9C%93&q=is%3Apr%20is%3Amerged%20milestone%3A1.0.0) pull requests to release this version of Gitea. You can download one of our pre-built binaries from our [downloads page](https://dl.gitea.io/gitea/1.0.0/), you just need to select the correct platform. For further details of the installation follow our [installation guide](https://docs.gitea.io/en-us/install-from-binary/), if you want to upgrade from Gogs follow our [migration guide](https://docs.gitea.io/en-us/upgrade-from-gogs/).

<!--more-->

## Changelog

* BREAKING
  * We have various changes on the API, scripting against API must be updated
* FEATURE
  * Show last login for admins [#121](https://github.com/go-gitea/gitea/pull/121)
* BUGFIXES
  * Fixed sender of notifications [#2](https://github.com/go-gitea/gitea/pull/2)
  * Fixed keyword hijacking vulnerability [#20](https://github.com/go-gitea/gitea/pull/20)
  * Fixed non-markdown readme rendering [#95](https://github.com/go-gitea/gitea/pull/95)
  * Allow updating draft releases [#169](https://github.com/go-gitea/gitea/pull/169)
  * GitHub API compliance [#227](https://github.com/go-gitea/gitea/pull/227)
  * Added commit SHA to tag webhook [#286](https://github.com/go-gitea/gitea/issues/286)
  * Secured links via noopener [#315](https://github.com/go-gitea/gitea/issues/315)
  * Replace tabs with spaces on wiki title [#371](https://github.com/go-gitea/gitea/pull/371)
  * Fixed vulnerability on labels and releases [#409](https://github.com/go-gitea/gitea/pull/409)
  * Fixed issue comment API [#449](https://github.com/go-gitea/gitea/pull/449)
* ENHANCEMENT
  * Use proper import path for libravatar [#3](https://github.com/go-gitea/gitea/pull/3)
  * Integrated DroneCI for tests and builds [#24](https://github.com/go-gitea/gitea/issues/24)
  * Integrated dependency manager [#29](https://github.com/go-gitea/gitea/issues/29)
  * Embedded bindata optionally [#30](https://github.com/go-gitea/gitea/issues/30)
  * Integrated pagination for releases [#73](https://github.com/go-gitea/gitea/pull/73)
  * Autogenerate version on every build [#91](https://github.com/go-gitea/gitea/issues/91)
  * Refactored Docker container [#104](https://github.com/go-gitea/gitea/issues/104)
  * Added short-hash support for downloads [#211](https://github.com/go-gitea/gitea/issues/211)
  * Display tooltip for downloads [#221](https://github.com/go-gitea/gitea/issues/221)
  * Improved HTTP headers for issue attachments [#270](https://github.com/go-gitea/gitea/pull/270)
  * Integrate public as bindata optionally [#293](https://github.com/go-gitea/gitea/pull/293)
  * Integrate templates as bindata optionally [#314](https://github.com/go-gitea/gitea/pull/314)
  * Inject more ENV variables into custom hooks [#316](https://github.com/go-gitea/gitea/issues/316)
  * Correct LDAP login validation [#342](https://github.com/go-gitea/gitea/pull/342)
  * Integrate conf as bindata optionally [#354](https://github.com/go-gitea/gitea/pull/354)
  * Serve video files in browser [#418](https://github.com/go-gitea/gitea/pull/418)
  * Configurable SSH host binding [#431](https://github.com/go-gitea/gitea/issues/431)
* MISC
  * Forked from Gogs and renamed to Gitea
  * Catching more errors with logs
  * Fixed all linting errors
  * Made the go linter entirely happy
  * Really integrated vendoring

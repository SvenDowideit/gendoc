
# Docker documentation tooling

## pre-req's

You need to have a working installation of `git`

## install

Download the latest release for your platform from
https://github.com/SvenDowideit/gendoc/releases

then open a terminal and run the following (example for OS X):

```
$ cd Downloads
$ chmod 755 gendoc-osx
$ ./gendoc-osx install
```

This will install `gendoc` into `/usr/local/bin/`, and then download `hugo` 
and install it there too.

There's built in help

```
$ gendoc --help
```

## initialize your workspace with

In whatever directory you have all your git clones (eg `~/repos/`)

```
$ ls -la
$ gendoc clone
$ ls -al
```

If there is no `docs.docker.com` repo found, will clone it, and then
will clone any missing repositories mentioned in the currently checked out 
`docs.docker.com/all-projects.yaml`

## to serve the master docs to a browser (port 8080)


```
$ gendoc checkout master
#$ gendoc render
```

## render

will use the files in the `docs-source/<publish-set>` dir to generate files 
in the `docs-html/<publish-set>` dir


to generate the v1.12 docs into `docs-source/v1.12` and `docs-html/v1.12/` 
dir, run a command set like

```
$ gendoc checkout v1.12 
$ gendoc render --disk
```


## status

Tells you what you have checked out, and what its related to.

the `--log` flag will tell you what commits are on the branch that the current 
sha is on - which may help you update the `all-projects.yml`

```
$ gendoc status
publish-set: v1.12
-- docs-base
* (detached from d5abfd4)
  master
  try-docsearch
  try-google-search-engine
Checkout Sha (d5abfd440fde1ca89cccd0d50d212ec4597cbe7c) NOT the same as tip of origin/master (c0825670806d06c5b125029f332f62b6608a94df)
c082567 Merge pull request #272 from SvenDowideit/use-https-where-possible
6ab8edd New version of linkcheck to mae it easier to read errors
0987a8c Use https where we can
-- docker
* (detached from 4879319)
  master
Checkout Sha (487931902c1177352e4eceec1b5ef558a5ba24cc) NOT the same as tip of origin/master (2875c5404e87f63d64986fee5f68d31a83f075c8)
2875c54 Merge pull request #24021 from londoncalling/add-redirect-for-getting-started
4060eb0 added another alias for getting started stuff
1c06ebe Merge pull request #23950 from jstarks/no_clone_tp5
8e8ef7c7 Merge pull request #23182 from crosbymichael/maxkeys
```

## prepare to publish updates

There is a command you can run to list all PR's that have been merged to 
master that are not yet in your current versioned branch:

You can use the command to see the state of all product repos, or select 
one `gendoc release prepare compose`

The output will list the docs PR's to consider cherry-picking into a new 
tag, the files it changes and any milestone and labels.

> NOTE: it uses your GITHUB_TOKEN either as an env var, or in the `--ghtoken` 
flag

```
$ export GITHUB_TOKEN=93cc2675c8f97e1a30b3bf2dbc287f0295ffc4fa
$ gendoc release prepare 
publish-set: v1.12
comparing allproject-yml ref's to upstream/master
## docs-base in docs-base at docs-2016-07-14
- PR 291 (d414ea9) Jul 14 07:50:25 from SvenDowideit/skip-http-client-errors
  - . changes in 660f15c54a13a1f29c6c088783dda4adcd3a70ea Skip http client errors, they cause false negatives
    - Dockerfile
- PR 293 (43dd4f2) Jul 19 23:16:26 from FSHA/homepage-icons
  - . changes in 84912d0e89608028abd42ad280b28586a75c558e increased the vibrancy of the iconography on the homepage
    - content/index.md
  - themes/docker-2016/static/assets/css/temporary.css
  - themes/docker-2016/static/assets/images/icon-apple@2X.png
  - themes/docker-2016/static/assets/images/icon-cloud@2X.png
  - themes/docker-2016/static/assets/images/icon-compose@2X.png
  - themes/docker-2016/static/assets/images/icon-engine@2X.png
  - themes/docker-2016/static/assets/images/icon-hub@2X.png
  - themes/docker-2016/static/assets/images/icon-linux@2X.png
  - themes/docker-2016/static/assets/images/icon-machine@2X.png
  - themes/docker-2016/static/assets/images/icon-registry@2X.png
  - themes/docker-2016/static/assets/images/icon-ucp@2X.png
  - themes/docker-2016/static/assets/images/icon-windows@2X.png
- PR 295 (a50ddff) Jul 20 17:52:43 from FSHA/accordion-font-weight
  - . changes in 8984cbb358292a8498bd00bf20e2705f39da4fd8 added font weight to accordion for chrome and legability
    - themes/docker-2016/static/documentation.css
## engine in docker at docs-v1.12.0-rc4-2016-07-15
- PR 24682 (9ad857a) Jul 16 01:00:51 from justincormack/oom-score-adj-doc
-  process/cherry-pick status/3-docs-review 
  - docs/ changes in 6ba6265d1ad86680ad7f7750ae1f9abb72f1e728 Document --oom-score-adj flag in docker run
    - docs/reference/run.md
- PR 24718 (4139123) Jul 16 13:26:04 from thaJeztah/fix-oracle-support-link
-  process/cherry-pick status/3-docs-review 
  - docs/ changes in c15144c4ec25df4234129fbe88b89cd4f709f784 docs: update Oracle support link
    - docs/installation/linux/oracle.md
## pinata in pinata at 1caa856152d84434e80eb05dae97b9c0aff81379
## cs-engine in cs-docker at docs-v1.11.1-cs2-2016-07-07
- PR 29 (35d0dc7) Jul  8 18:46:01 from mbentley/add-rhel-7.2
  - docs-cs changes in f5a04d4d62908f3c1dab36050e645776e1ada2e4 Added info for RHEL 7.2 support
    - docs-cs/install.md
## docker-trusted-registry in dhe-deploy at docs-v2.0.2-2016-07-19
## apidocs in dhe-deploy at docs-v2.0.2-2016-07-19
## ucp in orca at docs-v1.1.2-2016-07-19
## registry in distribution at docs-v2.4.1-2016-06-28
- PR 1833 (37b5e3e) Jul 13 17:41:01 from aaronlehmann/document-toomanyrequests
- Registry/2.5 group/distribution status/0-triage 
  - docs/ changes in b0099004e249ac7f9df92213c0186f95fa4cbb0f Document TOOMANYREQUESTS error code
    - docs/spec/api.md
  - docs/spec/api.md.tmpl
## compose in compose at docs-v1.7.1-2016-07-20
- PR 3590 (554dc24) Jul  6 18:45:09 from Knetic/oomscore
- 1.9.0 status/0-triage 
  - docs/ changes in 6fe5d2b54351143ee4eab090ccd58c5067985078 Implemented oom_score_adj
    - docs/compose-file.md
- PR 3488 (1e60030) Jul 11 19:24:29 from jgiannuzzi/internal_networks
- 1.9.0 area/networking kind/feature kind/parity 
  - docs/ changes in 83f35e132b37bf20baec264e49905c3ecc944ace Add support for creating internal networks
    - docs/compose-file.md
## swarm in swarm at docs-v1.2.3-2016-06-27
- PR 2385 (2700d0b) Jul  5 17:15:17 from clhlc/master
- 1.2.4 area/doc status/3-docs-review 
  - docs/ changes in 9b156d12305efb6e13a84dc4bc841d5533c9047f fix Switch the primary docs
    - docs/multi-manager-setup.md
- PR 2410 (ec5e048) Jul 11 15:18:19 from rmoriz/issue-2409
- 1.2.3 status/3-docs-review 
  - docs/ changes in 6c408c3cafdcddfe9211de0c6f26479905c2044f fix off-by-one error in server count
    - docs/install-manual.md
- PR 2413 (384d95a) Jul 14 14:03:56 from avsm/typo-in-label-docs
- 1.2.3 status/3-docs-review 
  - docs/ changes in 9b809185280f62cb051031bd20c5320ff24abf5a docs: fix extra newline in swarm label addition
    - docs/scheduler/filter.md
## machine in machine at docs-v0.8.0-rc2-2016-06-23
- PR 3573 (eeb35dd) Jul  7 16:49:32 from aculich/patch-1
- 0.8.0 status/3-docs-review 
  - docs/ changes in f17137034355d6d8862ef7aeba6de8daa3a30175 remove LTS from Ubuntu 15.10 in aws driver docs
    - docs/drivers/aws.md
## notary in notary at docs-v0.3-2016-06-22
## toolbox in toolbox at docs-v1.12.0-rc4-2016-06-29
NO merge PR found for (+ fa56475fb204c14d3673d14365c62aba5a838207 Bump versions for 1.12.0-rc4) 
## kitematic in kitematic at docs-v0.12.0-2016-06-28
## docker-hub in hub2-demo at docs-2016-06-23
## docker-cloud in cloud-docs at docs-2016-07-11
## cloud-api-docs-layout in cloud-docs at docs-2016-07-11
## cloud-api-docs in cloud-docs at docs-2016-07-11
## docker-store in mercury-ui at fdf50f7f057a6d24f0e95dcf68e15f3d05e873bd
NO merge PR found for (+ 88afb715410b474aaa0f44420c59248b9358b89e 62.0.0) 
## opensource in opensource at docs-2016-07-07
```

## README example updates

`gendoc` also is able to rewrite its README.md file using `gendoc readme`.
This will read the README.md file, and look for any "```" code markers.
Inside the code sections, it will run all lines starting with `$` and
add whatever the output is.



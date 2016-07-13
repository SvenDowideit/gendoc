
# Docker documentation tooling

## pre-req's

You need to have a working installation of `git`

## install

Download the latest release for your platform from https://github.com/SvenDowideit/gendoc/releases

then open a terminal and run the following (example for OS X):

```
$ cd Downloads
$ chmod 755 gendoc.app
$ ./gendoc.app setup
```

This will install `gendoc` into `/usr/local/bin/`, and then download `hugo` and install it too.

## initialize your workspace with

```
$ gendoc clone
```

If there is no `docs.docker.com` repo found, will clone it, and then
will clone any missing repositories mentioned in the currently checked out `docs.docker.com/all-projects.yaml`

## to serve the master docs to a browser (port 8080)


```
$ gendoc checkout master && gendoc render
```

to generate the v1.10 docs into `docs-source/v1.10` and `docs-html/v1.10/` dir, run a command set like

```
$ gendoc checkout v1.10 && gendoc --serve=false render
```


## render

will use the files in the `docs-source/<publish-set>` dir to generate files in the `docs-html/<publish-set>`
dir

## status (very preliminary, and changing atm)

tells you what you have checked out, and what its related to.

the `--log` flag will tell you what commits are on the branch that the current sha is on - which may help you update the `all-projects.yml`

```
sven@i7:~/src/gendoc-repos$ ../gendoc/gendoc  status --log
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

## TODO:
### bring together as

```
$ gendoc render --set v1.12
```


Using /tmp/example338576613 to run commands

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
NAME:
   gendoc - Generate documentation from multiple GitHub repositories

USAGE:
   gendoc [global options] command [command options] [arguments...]
   
VERSION:
   2016-08-10
   
COMMANDS:
     version   return the version
     clone     clone repos from the ./docs.docker.com/all-projects.yml file
     checkout  checkout versions from ./docs.docker.com/all-projects.yml file
     install   Install gendoc and its pre-req's into your PATH
     release   Prepare and ship a docs release.
     remote    Add a git remote - 2 arguments, name to give remote (origin), and organisation/Username on GitHub
     render    render html of docs checked out.
     status    status versions from ./docs.docker.com/all-projects.yml file
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug          enable debug output in the logs
   --ghtoken value  GITHUB_TOKEN for git and GitHub API [$GITHUB_TOKEN]
   --help, -h       show help
   --version, -v    print the version
   

```

## initialize your workspace with

In whatever directory you have all your git clones (eg `~/repos/`)

```
$ ls -la
total 24
drwx------ 2 ubuntu ubuntu  4096 Aug 19 04:56 .
drwxrwxrwt 6 root   root   20480 Aug 19 04:56 ..

$ gendoc clone
-- docs.docker.com
Cloning from git@github.com:docker/docs.docker.com
publish-set: v1.13-dev
-- docs-base
Cloning from git@github.com:docker/docs-base
-- engine
Cloning from git@github.com:docker/docker
-- pinata
Cloning from git@github.com:docker/pinata
-- cs-engine
Cloning from git@github.com:docker/cs-docker
-- docker-trusted-registry
Cloning from git@github.com:docker/dhe-deploy
-- apidocs
Dir already exists
-- ucp
Cloning from git@github.com:docker/orca
-- registry
Cloning from git@github.com:docker/distribution
-- kitematic
Cloning from git@github.com:docker/kitematic
-- compose
Cloning from git@github.com:docker/compose
-- swarm
Cloning from git@github.com:docker/swarm
-- machine
Cloning from git@github.com:docker/machine
-- notary
Cloning from git@github.com:docker/notary
-- toolbox
Cloning from git@github.com:docker/toolbox
-- docker-hub
Cloning from git@github.com:docker/hub2-demo
-- docker-cloud
Cloning from git@github.com:docker/cloud-docs
-- cloud-api-docs-layout
Dir already exists
-- cloud-api-docs
Dir already exists
-- docker-store
Cloning from git@github.com:docker/mercury-ui
-- opensource
Cloning from git@github.com:docker/opensource

$ ls -al
total 96
drwx------ 20 ubuntu ubuntu  4096 Aug 19 04:57 .
drwxrwxrwt  6 root   root   20480 Aug 19 04:56 ..
drwxrwxr-x  6 ubuntu ubuntu  4096 Aug 19 04:57 cloud-docs
drwxrwxr-x 11 ubuntu ubuntu  4096 Aug 19 04:57 compose
drwxrwxr-x 28 ubuntu ubuntu  4096 Aug 19 04:57 cs-docker
drwxrwxr-x 35 ubuntu ubuntu  4096 Aug 19 04:57 dhe-deploy
drwxrwxr-x 20 ubuntu ubuntu  4096 Aug 19 04:57 distribution
drwxrwxr-x 37 ubuntu ubuntu  4096 Aug 19 04:56 docker
drwxrwxr-x  7 ubuntu ubuntu  4096 Aug 19 04:56 docs-base
drwxrwxr-x  5 ubuntu ubuntu  4096 Aug 19 04:56 docs.docker.com
drwxrwxr-x 14 ubuntu ubuntu  4096 Aug 19 04:57 hub2-demo
drwxrwxr-x 14 ubuntu ubuntu  4096 Aug 19 04:57 kitematic
drwxrwxr-x 17 ubuntu ubuntu  4096 Aug 19 04:57 machine
drwxrwxr-x  9 ubuntu ubuntu  4096 Aug 19 04:57 mercury-ui
drwxrwxr-x 23 ubuntu ubuntu  4096 Aug 19 04:57 notary
drwxrwxr-x  6 ubuntu ubuntu  4096 Aug 19 04:57 opensource
drwxrwxr-x 28 ubuntu ubuntu  4096 Aug 19 04:57 orca
drwxrwxr-x 11 ubuntu ubuntu  4096 Aug 19 04:56 pinata
drwxrwxr-x 17 ubuntu ubuntu  4096 Aug 19 04:57 swarm
drwxrwxr-x  8 ubuntu ubuntu  4096 Aug 19 04:57 toolbox

```

If there is no `docs.docker.com` repo found, will clone it, and then
will clone any missing repositories mentioned in the currently checked out 
`docs.docker.com/all-projects.yaml`

## to serve the master docs to a browser (port 8080)


```
$ gendoc checkout master
Checking out docs.docker.com master.
Already at correct ref: all-projects has master, checkout is 84b2b9e323326614d60374bf628b058f83ab667f
publish-set: v1.13-dev
-- docs-base
Already at correct ref: all-projects has master, checkout is 7adec600461e7456366df201af4060878dca215b
-- docker
Already at correct ref: all-projects has master, checkout is 09e1de2080fd3b0bafb38adbd4b8c12ee949794d
-- pinata
Already at correct ref: all-projects has master, checkout is 8c11c14b46880079351e6f6503119e2ee6ef76ac
-- cs-docker
Already at correct ref: all-projects has master, checkout is 71a04c87ee4654756f870a7c095ce725220da171
-- dhe-deploy
Already at correct ref: all-projects has master, checkout is eb01555b9264d2a481fc87c6933909e7d713bf34
-- dhe-deploy
Already at correct ref: all-projects has master, checkout is eb01555b9264d2a481fc87c6933909e7d713bf34
-- orca
Already at correct ref: all-projects has master, checkout is bfb25097639d359363e17e5370bfc5c9e41e8231
-- distribution
Already at correct ref: all-projects has master, checkout is 010e063270be37cfa8547ccfb9717e5d874c88a8
-- kitematic
Already at correct ref: all-projects has master, checkout is 9143fe940657d843ea5ebc52caf1c5f0b043f2da
-- compose
Already at correct ref: all-projects has master, checkout is acfe100686fd95d524ff102c0b5fccff0bc79d8c
-- swarm
Already at correct ref: all-projects has master, checkout is 27968edd8a160f66c96c8545ad35e3a3eeb8766a
-- machine
Already at correct ref: all-projects has master, checkout is 578cb4dc34169efef6752df0863d2fc22a8fcf3a
-- notary
Already at correct ref: all-projects has master, checkout is ca2008c88619d7197501139070c1aaf2f9281446
-- toolbox
Already at correct ref: all-projects has master, checkout is db24b2166089b2bf67841b995015e626bb7a409f
-- hub2-demo
Already at correct ref: all-projects has master, checkout is 35b35b9a0270c368c588fd1b0bee27d6edc22254
-- cloud-docs
Already at correct ref: all-projects has master, checkout is 1b9a757a92d83c875284bb2f90fde97f14277be6
-- cloud-docs
Already at correct ref: all-projects has master, checkout is 1b9a757a92d83c875284bb2f90fde97f14277be6
-- cloud-docs
Already at correct ref: all-projects has master, checkout is 1b9a757a92d83c875284bb2f90fde97f14277be6
-- mercury-ui
Already at correct ref: all-projects has master, checkout is 68d3072991cfdc904a9c35515705b8551ee71317
-- opensource
Already at correct ref: all-projects has master, checkout is 9736bd57db38561847648a612867d0f0f9978836

```

## render

will use the files in the `docs-source/<publish-set>` dir to generate files 
in the `docs-html/<publish-set>` dir


to generate the v1.12 docs into `docs-source/v1.12` and `docs-html/v1.12/` 
dir, run a command set like

```
$ gendoc checkout v1.12 
Checking out docs.docker.com v1.12.
Branch v1.12 set up to track remote branch v1.12 from upstream.
publish-set: v1.12
-- docs-base
-- docker
-- pinata
-- cs-docker
Already at correct ref: all-projects has 71a04c87ee4654756f870a7c095ce725220da171, checkout is 71a04c87ee4654756f870a7c095ce725220da171
-- dhe-deploy
-- dhe-deploy
Already at correct ref: all-projects has docs-v2.0.3-2016-08-11, checkout is 139a5d128584da25eee4b730c35497d8c0840515
-- orca
-- distribution
-- compose
-- swarm
-- machine
-- notary
-- toolbox
-- kitematic
-- hub2-demo
-- cloud-docs
-- cloud-docs
Already at correct ref: all-projects has docs-2016-08-17, checkout is 33e56428398878f76d083914dbde44a02f7b1fdb
-- cloud-docs
Already at correct ref: all-projects has docs-2016-08-17, checkout is 33e56428398878f76d083914dbde44a02f7b1fdb
-- mercury-ui
-- opensource

$ gendoc render --disk
publish-set: v1.12
copy docs-base TO docs-source/v1.12
copy docker/docs TO docs-source/v1.12/content/engine
copy pinata/docs TO docs-source/v1.12/content
copy cs-docker/docs-cs TO docs-source/v1.12/content/cs-engine
copy dhe-deploy/docs TO docs-source/v1.12/content/docker-trusted-registry
copy dhe-deploy/apidocgen/output TO docs-source/v1.12/content/apidocs
copy orca/docs TO docs-source/v1.12/content/ucp
copy distribution/docs TO docs-source/v1.12/content/registry
copy compose/docs TO docs-source/v1.12/content/compose
copy swarm/docs TO docs-source/v1.12/content/swarm
copy machine/docs TO docs-source/v1.12/content/machine
copy notary/docs TO docs-source/v1.12/content/notary
copy toolbox/docs TO docs-source/v1.12/content/toolbox
copy kitematic/docs TO docs-source/v1.12/content/kitematic
copy hub2-demo/docs TO docs-source/v1.12/content/docker-hub
copy cloud-docs/docs TO docs-source/v1.12/content/docker-cloud
copy cloud-docs/apidocs/layouts TO docs-source/v1.12/layouts/cloud-api-docs
copy cloud-docs/apidocs TO docs-source/v1.12/content/apidocs
copy mercury-ui/docs TO docs-source/v1.12/content/docker-store
copy opensource/docs TO docs-source/v1.12/content/opensource
INFO: 2016/08/19 04:58:00 hugo.go:463: Using config file: /tmp/example338576613/docs-source/v1.12/config.toml
WARN: 2016/08/19 04:58:00 hugo.go:557: Unable to find Static Directory: /tmp/example338576613/docs-source/v1.12/static/
INFO: 2016/08/19 04:58:00 hugo.go:566: /tmp/example338576613/docs-source/v1.12/themes/docker-2016/static is the only static directory available to sync from
INFO: 2016/08/19 04:58:00 hugo.go:607: removing all files from destination that don't exist in static dirs
INFO: 2016/08/19 04:58:00 hugo.go:609: syncing static files to /tmp/example338576613/docs-html/v1.12/
Started building site
INFO: 2016/08/19 04:58:01 site.go:1251: found taxonomies: map[string]string{"tag":"tags", "category":"categories"}
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/started/" translated to "mac/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/" translated to "engine/userguide/containers/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/extend/authorization/" translated to "engine/extend/authorization/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/deploy-to-cloud/" translated to "docker-cloud/feature-reference/deploy-to-cloud/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/deploy-to-cloud/" translated to "docker-cloud/tutorials/deploy-to-cloud/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/awslogs/" translated to "engine/reference/logging/awslogs/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/accounts/" translated to "docker-trusted-registry/accounts/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/manage/monitor-manage-users/" translated to "ucp/manage/monitor-manage-users/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/user-management/manage-users/" translated to "ucp/user-management/manage-users/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/automated-build/" translated to "docker-cloud/feature-reference/automated-build/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/automated-testing/" translated to "docker-cloud/feature-reference/automated-testing/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/auto-destroy/" translated to "docker-cloud/feature-reference/auto-destroy/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/autorestart/" translated to "docker-cloud/feature-reference/autorestart/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/auto-redeploy/" translated to "docker-cloud/feature-reference/auto-redeploy/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/host_integration/" translated to "engine/articles/host_integration/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/dockerfile_best-practices/" translated to "engine/articles/dockerfile_best-practices/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/optimize-dockerfiles/" translated to "docker-cloud/getting-started/intermediate/optimize-dockerfiles/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/optimize-dockerfiles/" translated to "docker-cloud/tutorials/optimize-dockerfiles/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/misc/breaking/" translated to "engine/misc/breaking/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/step_four/" translated to "mac/step_four/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/step_four/" translated to "windows/step_four/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/step_four/" translated to "linux/step_four/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/dockerimages/" translated to "engine/userguide/containers/dockerimages/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/dockerimages/" translated to "engine/userguide/dockerimages/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/" translated to "docker-trusted-registry/cs-engine/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cse-release-notes/" translated to "docker-trusted-registry/cse-release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/release-notes/release-notes/" translated to "docker-trusted-registry/cs-engine/release-notes/release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/cloud/cloud/" translated to "engine/installation/cloud/cloud/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/compose/yml" translated to "compose/yml/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/overview/" translated to "engine/reference/logging/overview/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/configuring/" translated to "engine/articles/configuring/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/admin/configuring/" translated to "engine/admin/configuring/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/deployment-strategies/" translated to "docker-cloud/feature-reference/deployment-strategies/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/systemd/" translated to "engine/articles/systemd/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/step_five/" translated to "mac/step_five/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/step_five/" translated to "windows/step_five/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/step_five/" translated to "linux/step_five/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/baseimages/" translated to "engine/articles/baseimages/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/load-balance-hello-world/" translated to "docker-cloud/getting-started/intermediate/load-balance-hello-world/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/load-balance-hello-world/" translated to "docker-cloud/tutorials/load-balance-hello-world/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/12_data_management_with_volumes/" translated to "docker-cloud/getting-started/python/12_data_management_with_volumes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/6_define_environment_variables/" translated to "docker-cloud/getting-started/python/6_define_environment_variables/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/6_define_environment_variables/" translated to "docker-cloud/getting-started/golang/6_define_environment_variables/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/soft-garbage/" translated to "docker-trusted-registry/soft-garbage/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/repos-and-images/delete-images/" translated to "docker-trusted-registry/repos-and-images/delete-images/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/deploy-application/" translated to "ucp/deploy-application/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/" translated to "docker-cloud/getting-started/python/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/" translated to "docker-cloud/getting-started/golang/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/03-create-cluster/" translated to "swarm/swarm_at_scale/03-create-cluster/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/02-deploy-infra/" translated to "swarm/swarm_at_scale/02-deploy-infra/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/5_deploy_the_app_as_a_service/" translated to "docker-cloud/getting-started/python/5_deploy_the_app_as_a_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/5_deploy_the_app_as_a_service/" translated to "docker-cloud/getting-started/golang/5_deploy_the_app_as_a_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/04-deploy-app/" translated to "swarm/swarm_at_scale/04-deploy-app/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/your_first_node/" translated to "docker-cloud/getting-started/beginner/your_first_node/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/deploy_first_node/" translated to "docker-cloud/getting-started/beginner/deploy_first_node/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/your_first_service/" translated to "docker-cloud/getting-started/beginner/your_first_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/deploy_first_service/" translated to "docker-cloud/getting-started/beginner/deploy_first_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/deploy-tags/" translated to "docker-cloud/feature-reference/deploy-tags/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/misc/deprecated/" translated to "engine/misc/deprecated/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/slack-integration/" translated to "docker-cloud/tutorials/slack-integration/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/misc/" translated to "engine/misc/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/introduction/understanding-docker/" translated to "introduction/understanding-docker/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/basics/" translated to "engine/userguide/basics/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/quickstart.md" translated to "engine/quickstart.md/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/api/swarm-api/" translated to "api/swarm-api/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/swarm/api/" translated to "swarm/api/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-hub-enterprise/" translated to "docker-hub-enterprise/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/overview/" translated to "docker-trusted-registry/overview/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/networking/dockernetworks/" translated to "engine/userguide/networking/dockernetworks/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/docker-toolbox/" translated to "mackit/docker-toolbox/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/security/" translated to "engine/articles/security/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/evaluation-install/" translated to "ucp/evaluation-install/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/misc/faq/" translated to "engine/misc/faq/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/faqs/" translated to "mackit/faqs/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/osxfs/" translated to "mackit/osxfs/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/step_three/" translated to "mac/step_three/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/step_three/" translated to "windows/step_three/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/step_three/" translated to "linux/step_three/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/fluentd/" translated to "engine/reference/logging/fluentd/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/started/" translated to "mac/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/started/" translated to "windows/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/started/" translated to "linux/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/getting-started/" translated to "getting-started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/" translated to "mackit/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/getting-started/" translated to "mackit/getting-started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/" translated to "mac/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/started/" translated to "mac/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-for-mac/started/" translated to "docker-for-mac/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/winkit/getting-started/" translated to "winkit/getting-started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/winkit/" translated to "winkit/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/" translated to "windows/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/started/" translated to "windows/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-for-windows/started/" translated to "docker-for-windows/started/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/dockerizing/" translated to "engine/userguide/containers/dockerizing/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/dockerizing/" translated to "engine/userguide/dockerizing/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/swarm/how-swarm-mode-works/" translated to "engine/swarm/how-swarm-mode-works/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/engine-ami-launch/" translated to "docker-trusted-registry/install/engine-ami-launch/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/install-csengine/" translated to "docker-trusted-registry/install/install-csengine/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/install/" translated to "docker-trusted-registry/cs-engine/install/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/dtr-ami-byol-launch/" translated to "docker-trusted-registry/install/dtr-ami-byol-launch/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/dtr-ami-bds-launch/" translated to "docker-trusted-registry/install/dtr-ami-bds-launch/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/dtr-vhd-azure/" translated to "docker-trusted-registry/install/dtr-vhd-azure/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/install-dtr/" translated to "docker-trusted-registry/install/install-dtr/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/step_one/" translated to "mac/step_one/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/step_one/" translated to "windows/step_one/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/step_one/" translated to "linux/step_one/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/production-install/" translated to "ucp/production-install/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/archlinux/" translated to "engine/installation/archlinux/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/cruxlinux/" translated to "engine/installation/cruxlinux/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/centos/" translated to "engine/installation/centos/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/debian/" translated to "engine/installation/debian/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/fedora/" translated to "engine/installation/fedora/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/frugalware/" translated to "engine/installation/frugalware/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/gentoolinux/" translated to "engine/installation/gentoolinux/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/oracle/" translated to "engine/installation/oracle/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/rhel/" translated to "engine/installation/rhel/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/ubuntulinux/" translated to "engine/installation/ubuntulinux/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/SUSE/" translated to "engine/installation/SUSE/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/dtr-integration/" translated to "ucp/dtr-integration/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/intro_cloud/" translated to "docker-cloud/getting-started/beginner/intro_cloud/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/1_introduction/" translated to "docker-cloud/getting-started/python/1_introduction/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/1_introduction/" translated to "docker-cloud/getting-started/golang/1_introduction/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/journald/" translated to "engine/reference/logging/journald/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/faq/docker-errors-faq/" translated to "docker-cloud/faq/docker-errors-faq/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/about/" translated to "swarm/swarm_at_scale/about/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/last_page/" translated to "mac/last_page/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/last_page/" translated to "windows/last_page/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/last_page/" translated to "linux/last_page/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/multi-arch/" translated to "mackit/multi-arch/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/license/" translated to "docker-trusted-registry/license/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/compose/env" translated to "compose/env/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-aws/" translated to "docker-cloud/getting-started/beginner/link-aws/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-aws/" translated to "docker-cloud/getting-started/link-aws/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-do/" translated to "docker-cloud/getting-started/beginner/link-do/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-do/" translated to "docker-cloud/getting-started/link-do/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-azure/" translated to "docker-cloud/getting-started/beginner/link-azure/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-azure/" translated to "docker-cloud/getting-started/link-azure/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-packet/" translated to "docker-cloud/getting-started/beginner/link-packet/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-packet/" translated to "docker-cloud/getting-started/link-packet/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-softlayer/" translated to "docker-cloud/getting-started/beginner/link-softlayer/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-softlayer/" translated to "docker-cloud/getting-started/link-softlayer/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/link-source/" translated to "docker-cloud/tutorials/link-source/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/use-hosted/" translated to "docker-cloud/getting-started/use-hosted/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/ambassador_pattern_linking/" translated to "engine/articles/ambassador_pattern_linking/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/9_load-balance_the_service/" translated to "docker-cloud/getting-started/python/9_load-balance_the_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/9_load-balance_the_service/" translated to "docker-cloud/getting-started/golang/9_load-balance_the_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/log_tags/" translated to "engine/reference/logging/log_tags/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/" translated to "engine/reference/logging/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/troubleshoot/" translated to "mackit/troubleshoot/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/troubleshoot/" translated to "windows/troubleshoot/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/dockervolumes/" translated to "engine/userguide/containers/dockervolumes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/dockervolumes/" translated to "engine/userguide/dockervolumes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/stacks/" translated to "docker-cloud/feature-reference/stacks/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/monitor-troubleshoot/monitor/" translated to "docker-trusted-registry/monitor-troubleshoot/monitor/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/manage/" translated to "ucp/manage/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/manage/monitor-ucp/" translated to "ucp/manage/monitor-ucp/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/networkigncontainers/" translated to "engine/userguide/containers/networkigncontainers/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/networkigncontainers/" translated to "engine/userguide/networkigncontainers/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/networking/" translated to "mackit/networking/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/amazon/" translated to "engine/installation/amazon/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/google/" translated to "engine/installation/google/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/softlayer/" translated to "engine/installation/softlayer/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/azure/" translated to "engine/installation/azure/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/rackspace/" translated to "engine/installation/rackspace/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/installation/joyent/" translated to "engine/installation/joyent/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-hub/overview/" translated to "docker-hub/overview/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/compose/reference/docker-compose/" translated to "compose/reference/docker-compose/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/plan-production-install/" translated to "ucp/plan-production-install/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/dsc/" translated to "engine/articles/dsc/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/3_prepare_the_app/" translated to "docker-cloud/getting-started/python/3_prepare_the_app/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/3_prepare_the_app/" translated to "docker-cloud/getting-started/golang/3_prepare_the_app/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cse-prior-release-notes/" translated to "docker-trusted-registry/cse-prior-release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/release-notes/prior-release-notes/" translated to "docker-trusted-registry/cs-engine/release-notes/prior-release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/prior-release-notes/" translated to "docker-trusted-registry/prior-release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/https/" translated to "engine/articles/https/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/10_provision_a_data_backend_for_your_service/" translated to "docker-cloud/getting-started/python/10_provision_a_data_backend_for_your_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/10_provision_a_data_backend_for_your_service/" translated to "docker-cloud/getting-started/golang/10_provision_a_data_backend_for_your_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/ports/" translated to "docker-cloud/feature-reference/ports/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/quick-start/" translated to "docker-trusted-registry/quick-start/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/repos-and-images/create-repo/" translated to "docker-trusted-registry/repos-and-images/create-repo/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/userguide/" translated to "docker-trusted-registry/userguide/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/repos-and-images/push-and-pull-images/" translated to "docker-trusted-registry/repos-and-images/push-and-pull-images/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/pushing-images-to-dockercloud/" translated to "docker-cloud/getting-started/intermediate/pushing-images-to-dockercloud/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/pushing-images-to-dockercloud/" translated to "docker-cloud/tutorials/pushing-images-to-dockercloud/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/4_push_to_cloud_registry/" translated to "docker-cloud/getting-started/python/4_push_to_cloud_registry/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/4_push_to_cloud_registry/" translated to "docker-cloud/getting-started/golang/4_push_to_cloud_registry/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/contributing/contributing" translated to "contributing/contributing/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/service-redeploy/" translated to "docker-cloud/feature-reference/service-redeploy/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/registry/overview/" translated to "registry/overview/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mackit/release-notes/" translated to "mackit/release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/winkit/release-notes/" translated to "winkit/release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/release-notes/" translated to "docker-trusted-registry/cs-engine/release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/registry_mirror/" translated to "engine/articles/registry_mirror/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/usingdocker/" translated to "engine/userguide/containers/usingdocker/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/run_metrics" translated to "engine/articles/run_metrics/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/ssh-into-a-node/" translated to "docker-cloud/getting-started/intermediate/ssh-into-a-node/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/ssh-into-a-node/" translated to "docker-cloud/tutorials/ssh-into-a-node/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/faq/how-ssh-nodes/" translated to "docker-cloud/faq/how-ssh-nodes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/7_scale_the_service/" translated to "docker-cloud/getting-started/python/7_scale_the_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/7_scale_the_service/" translated to "docker-cloud/getting-started/golang/7_scale_the_service/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/service-scaling/" translated to "docker-cloud/feature-reference/service-scaling/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/api-roles/" translated to "docker-cloud/feature-reference/api-roles/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/service-links/" translated to "docker-cloud/feature-reference/service-links/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/networking/" translated to "ucp/networking/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/high-availability/high-availability/" translated to "docker-trusted-registry/high-availability/high-availability/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/understand_ha/" translated to "ucp/understand_ha/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/2_set_up/" translated to "docker-cloud/getting-started/python/2_set_up/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/2_set_up/" translated to "docker-cloud/getting-started/golang/2_set_up/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/logging/splunk/" translated to "engine/reference/logging/splunk/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/stack-yaml-reference/" translated to "docker-cloud/feature-reference/stack-yaml-reference/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/11_service_stacks/" translated to "docker-cloud/getting-started/python/11_service_stacks/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/registry/storagedrivers/" translated to "registry/storagedrivers/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/containers/dockerrepos/" translated to "engine/userguide/containers/dockerrepos/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/userguide/dockerrepos/" translated to "engine/userguide/dockerrepos/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/swarm/manager-administration-guide/" translated to "engine/swarm/manager-administration-guide/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/step_six/" translated to "mac/step_six/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/step_six/" translated to "windows/step_six/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/step_six/" translated to "linux/step_six/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/installing-cli/" translated to "docker-cloud/getting-started/intermediate/installing-cli/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/installing-cli/" translated to "docker-cloud/getting-started/installing-cli/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/installing-cli/" translated to "docker-cloud/tutorials/installing-cli/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/adminguide/" translated to "docker-trusted-registry/adminguide/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/kv_store/" translated to "ucp/kv_store/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/05-troubleshoot/" translated to "swarm/swarm_at_scale/05-troubleshoot/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/release-notes/release-notes/" translated to "docker-trusted-registry/release-notes/release-notes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/mac/step_two/" translated to "mac/step_two/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/windows/step_two/" translated to "windows/step_two/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/linux/step_two/" translated to "linux/step_two/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/upgrade/" translated to "docker-trusted-registry/cs-engine/upgrade/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/docker-upgrade/" translated to "docker-cloud/feature-reference/docker-upgrade/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/docker-upgrade/" translated to "docker-cloud/tutorials/docker-upgrade/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/ucp/upgrade-ucp/" translated to "ucp/upgrade-ucp/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/upgrade/" translated to "docker-trusted-registry/install/upgrade/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-trusted-registry/install/upgrade/upgrade-minor/" translated to "docker-trusted-registry/install/upgrade/upgrade-minor/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/byoh/" translated to "docker-cloud/feature-reference/byoh/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/byoh/" translated to "docker-cloud/tutorials/byoh/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/use-byon/" translated to "docker-cloud/getting-started/use-byon/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/triggers/" translated to "docker-cloud/feature-reference/triggers/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/chef/" translated to "engine/articles/chef/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/faq/cloud-on-packet.net-faq/" translated to "docker-cloud/faq/cloud-on-packet.net-faq/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/faq/cloud-on-aws-faq/" translated to "docker-cloud/faq/cloud-on-aws-faq/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/puppet/" translated to "engine/articles/puppet/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/using_supervisord/" translated to "engine/articles/using_supervisord/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/articles/certificates/" translated to "engine/articles/certificates/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/8_view_logs/" translated to "docker-cloud/getting-started/python/8_view_logs/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/8_view_logs/" translated to "docker-cloud/getting-started/golang/8_view_logs/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/tutorials/download-volume-data/" translated to "docker-cloud/tutorials/download-volume-data/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/volumes/" translated to "docker-cloud/feature-reference/volumes/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/commandline/daemon/" translated to "engine/reference/commandline/daemon/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/commandline/node_tasks/" translated to "engine/reference/commandline/node_tasks/index.html"
INFO: 2016/08/19 04:58:01 htmlredirect.go:115: Alias "/engine/reference/commandline/service_tasks/" translated to "engine/reference/commandline/service_tasks/index.html"
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section apidocs: [section/apidocs.html _default/section.html _default/list.html indexes/apidocs.html _default/indexes.html theme/section/apidocs.html theme/_default/section.html theme/_default/list.html theme/indexes/apidocs.html theme/_default/indexes.html theme/section/apidocs.html theme/_default/section.html theme/_default/list.html theme/indexes/apidocs.html theme/_default/indexes.html theme/theme/section/apidocs.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/apidocs.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "apidocs" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section swarm: [section/swarm.html _default/section.html _default/list.html indexes/swarm.html _default/indexes.html theme/section/swarm.html theme/_default/section.html theme/_default/list.html theme/indexes/swarm.html theme/_default/indexes.html theme/section/swarm.html theme/_default/section.html theme/_default/list.html theme/indexes/swarm.html theme/_default/indexes.html theme/theme/section/swarm.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/swarm.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "swarm" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section docker-trusted-registry: [section/docker-trusted-registry.html _default/section.html _default/list.html indexes/docker-trusted-registry.html _default/indexes.html theme/section/docker-trusted-registry.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-trusted-registry.html theme/_default/indexes.html theme/section/docker-trusted-registry.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-trusted-registry.html theme/_default/indexes.html theme/theme/section/docker-trusted-registry.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-trusted-registry.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "docker-trusted-registry" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section docker-hub: [section/docker-hub.html _default/section.html _default/list.html indexes/docker-hub.html _default/indexes.html theme/section/docker-hub.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-hub.html theme/_default/indexes.html theme/section/docker-hub.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-hub.html theme/_default/indexes.html theme/theme/section/docker-hub.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-hub.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "docker-hub" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section cs-engine: [section/cs-engine.html _default/section.html _default/list.html indexes/cs-engine.html _default/indexes.html theme/section/cs-engine.html theme/_default/section.html theme/_default/list.html theme/indexes/cs-engine.html theme/_default/indexes.html theme/section/cs-engine.html theme/_default/section.html theme/_default/list.html theme/indexes/cs-engine.html theme/_default/indexes.html theme/theme/section/cs-engine.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/cs-engine.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "cs-engine" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section : [section/.html _default/section.html _default/list.html indexes/.html _default/indexes.html theme/section/.html theme/_default/section.html theme/_default/list.html theme/indexes/.html theme/_default/indexes.html theme/section/.html theme/_default/section.html theme/_default/list.html theme/indexes/.html theme/_default/indexes.html theme/theme/section/.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section toolbox: [section/toolbox.html _default/section.html _default/list.html indexes/toolbox.html _default/indexes.html theme/section/toolbox.html theme/_default/section.html theme/_default/list.html theme/indexes/toolbox.html theme/_default/indexes.html theme/section/toolbox.html theme/_default/section.html theme/_default/list.html theme/indexes/toolbox.html theme/_default/indexes.html theme/theme/section/toolbox.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/toolbox.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "toolbox" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section opensource: [section/opensource.html _default/section.html _default/list.html indexes/opensource.html _default/indexes.html theme/section/opensource.html theme/_default/section.html theme/_default/list.html theme/indexes/opensource.html theme/_default/indexes.html theme/section/opensource.html theme/_default/section.html theme/_default/list.html theme/indexes/opensource.html theme/_default/indexes.html theme/theme/section/opensource.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/opensource.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "opensource" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section engine: [section/engine.html _default/section.html _default/list.html indexes/engine.html _default/indexes.html theme/section/engine.html theme/_default/section.html theme/_default/list.html theme/indexes/engine.html theme/_default/indexes.html theme/section/engine.html theme/_default/section.html theme/_default/list.html theme/indexes/engine.html theme/_default/indexes.html theme/theme/section/engine.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/engine.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "engine" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section docker-for-windows: [section/docker-for-windows.html _default/section.html _default/list.html indexes/docker-for-windows.html _default/indexes.html theme/section/docker-for-windows.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-windows.html theme/_default/indexes.html theme/section/docker-for-windows.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-windows.html theme/_default/indexes.html theme/theme/section/docker-for-windows.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-for-windows.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "docker-for-windows" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section machine: [section/machine.html _default/section.html _default/list.html indexes/machine.html _default/indexes.html theme/section/machine.html theme/_default/section.html theme/_default/list.html theme/indexes/machine.html theme/_default/indexes.html theme/section/machine.html theme/_default/section.html theme/_default/list.html theme/indexes/machine.html theme/_default/indexes.html theme/theme/section/machine.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/machine.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "machine" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section notary: [section/notary.html _default/section.html _default/list.html indexes/notary.html _default/indexes.html theme/section/notary.html theme/_default/section.html theme/_default/list.html theme/indexes/notary.html theme/_default/indexes.html theme/section/notary.html theme/_default/section.html theme/_default/list.html theme/indexes/notary.html theme/_default/indexes.html theme/theme/section/notary.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/notary.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "notary" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section docker-store: [section/docker-store.html _default/section.html _default/list.html indexes/docker-store.html _default/indexes.html theme/section/docker-store.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-store.html theme/_default/indexes.html theme/section/docker-store.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-store.html theme/_default/indexes.html theme/theme/section/docker-store.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-store.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "docker-store" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section docker-for-mac: [section/docker-for-mac.html _default/section.html _default/list.html indexes/docker-for-mac.html _default/indexes.html theme/section/docker-for-mac.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-mac.html theme/_default/indexes.html theme/section/docker-for-mac.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-mac.html theme/_default/indexes.html theme/theme/section/docker-for-mac.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-for-mac.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "docker-for-mac" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section ucp: [section/ucp.html _default/section.html _default/list.html indexes/ucp.html _default/indexes.html theme/section/ucp.html theme/_default/section.html theme/_default/list.html theme/indexes/ucp.html theme/_default/indexes.html theme/section/ucp.html theme/_default/section.html theme/_default/list.html theme/indexes/ucp.html theme/_default/indexes.html theme/theme/section/ucp.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/ucp.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "ucp" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section docker-cloud: [section/docker-cloud.html _default/section.html _default/list.html indexes/docker-cloud.html _default/indexes.html theme/section/docker-cloud.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-cloud.html theme/_default/indexes.html theme/section/docker-cloud.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-cloud.html theme/_default/indexes.html theme/theme/section/docker-cloud.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-cloud.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "docker-cloud" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section registry: [section/registry.html _default/section.html _default/list.html indexes/registry.html _default/indexes.html theme/section/registry.html theme/_default/section.html theme/_default/list.html theme/indexes/registry.html theme/_default/indexes.html theme/section/registry.html theme/_default/section.html theme/_default/list.html theme/indexes/registry.html theme/_default/indexes.html theme/theme/section/registry.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/registry.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "registry" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section compose: [section/compose.html _default/section.html _default/list.html indexes/compose.html _default/indexes.html theme/section/compose.html theme/_default/section.html theme/_default/list.html theme/indexes/compose.html theme/_default/indexes.html theme/section/compose.html theme/_default/section.html theme/_default/list.html theme/indexes/compose.html theme/_default/indexes.html theme/theme/section/compose.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/compose.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "compose" is rendered empty
WARN: 2016/08/19 04:58:01 site.go:2014: Unable to locate layout for section kitematic: [section/kitematic.html _default/section.html _default/list.html indexes/kitematic.html _default/indexes.html theme/section/kitematic.html theme/_default/section.html theme/_default/list.html theme/indexes/kitematic.html theme/_default/indexes.html theme/section/kitematic.html theme/_default/section.html theme/_default/list.html theme/indexes/kitematic.html theme/_default/indexes.html theme/theme/section/kitematic.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/kitematic.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 04:58:01 site.go:1990: "kitematic" is rendered empty
0 of 29 drafts rendered
0 future content
708 pages created
929 non-page files copied
0 paginator pages created
0 tags created
0 categories created
in 24707 ms

```


## status

Tells you what you have checked out, and what its related to.

the `--log` flag will tell you what commits are on the branch that the current 
sha is on - which may help you update the `all-projects.yml`

```
$ gendoc status
publish-set: v1.12
-- docs-base (in .)
* (detached from docs-2016-08-12)
  master
docs-2016-08-12
Checkout Sha (f87ad24084c52b7b959f9a1dd091ceedf88b45ea) NOT the same as ref: in all-projects docs-2016-08-12 (docs-2016-08-12)
Checkout Sha (f87ad24084c52b7b959f9a1dd091ceedf88b45ea) NOT the same as tip of upstream/master (7adec600461e7456366df201af4060878dca215b)
	Consider a `git reset --hard upstream/master`
-- docker (in docs/)
* (detached from 23cf638)
v1.12.1
Checkout Sha (23cf638307f030cd8d48c9efc21feec18a6f88f8) NOT the same as tip of upstream/master (09e1de2080fd3b0bafb38adbd4b8c12ee949794d)
	Consider a `git reset --hard upstream/master`
-- pinata (in docs/)
* (detached from docs-v1.12.0-beta23-2016-08-16-update1)
  master
docs-v1.12.0-beta23-2016-08-16-update1
Checkout Sha (c91d7602ed2ec5aa5b404d7f5b9696ab8fa0c58d) NOT the same as ref: in all-projects docs-v1.12.0-beta23-2016-08-16-update1 (docs-v1.12.0-beta23-2016-08-16-update1)
Checkout Sha (c91d7602ed2ec5aa5b404d7f5b9696ab8fa0c58d) NOT the same as tip of upstream/master (8c11c14b46880079351e6f6503119e2ee6ef76ac)
	Consider a `git reset --hard upstream/master`
-- cs-docker (in docs-cs)
* master
-- dhe-deploy (in docs/)
* (detached from docs-v2.0.3-2016-08-11)
  master
docs-v2.0.3-2016-08-11
Checkout Sha (139a5d128584da25eee4b730c35497d8c0840515) NOT the same as ref: in all-projects docs-v2.0.3-2016-08-11 (docs-v2.0.3-2016-08-11)
Checkout Sha (139a5d128584da25eee4b730c35497d8c0840515) NOT the same as tip of upstream/master (eb01555b9264d2a481fc87c6933909e7d713bf34)
	Consider a `git reset --hard upstream/master`
-- dhe-deploy (in apidocgen/output)
* (detached from docs-v2.0.3-2016-08-11)
  master
docs-v2.0.3-2016-08-11
Checkout Sha (139a5d128584da25eee4b730c35497d8c0840515) NOT the same as ref: in all-projects docs-v2.0.3-2016-08-11 (docs-v2.0.3-2016-08-11)
Checkout Sha (139a5d128584da25eee4b730c35497d8c0840515) NOT the same as tip of upstream/master (eb01555b9264d2a481fc87c6933909e7d713bf34)
	Consider a `git reset --hard upstream/master`
-- orca (in docs/)
* (detached from docs-v1.1.2-2016-08-03)
  master
docs-v1.1.2-2016-08-03
release/v2.0.0-tp1
Checkout Sha (77a849ad947ba9b6f2a96e752a2c9697660348fe) NOT the same as ref: in all-projects docs-v1.1.2-2016-08-03 (docs-v1.1.2-2016-08-03)
Checkout Sha (77a849ad947ba9b6f2a96e752a2c9697660348fe) NOT the same as tip of upstream/master (bfb25097639d359363e17e5370bfc5c9e41e8231)
	Consider a `git reset --hard upstream/master`
-- distribution (in docs/)
* (detached from docs-v2.5.0-2016-07-28)
docs-v2.5.0-2016-07-28
v2.5.0
Checkout Sha (a9b1322edf48b1fb9aee4e5ded7a4f4ac37c6830) NOT the same as ref: in all-projects docs-v2.5.0-2016-07-28 (docs-v2.5.0-2016-07-28)
Checkout Sha (a9b1322edf48b1fb9aee4e5ded7a4f4ac37c6830) NOT the same as tip of upstream/master (010e063270be37cfa8547ccfb9717e5d874c88a8)
	Consider a `git reset --hard upstream/master`
-- compose (in docs/)
* (detached from docs-v1.8.0-2016-08-03)
docs-v1.8.0-2016-08-03
Checkout Sha (429320a4f8f4040b273fd4d1be9f1d0b1283dc23) NOT the same as ref: in all-projects docs-v1.8.0-2016-08-03 (docs-v1.8.0-2016-08-03)
Checkout Sha (429320a4f8f4040b273fd4d1be9f1d0b1283dc23) NOT the same as tip of upstream/master (acfe100686fd95d524ff102c0b5fccff0bc79d8c)
	Consider a `git reset --hard upstream/master`
-- swarm (in docs/)
* (detached from docs-v1.2.4-2016-08-03)
docs-v1.2.4-2016-08-03
Checkout Sha (ac1b1b1a19f73b8e1899c63b549653d1fd71ff8c) NOT the same as ref: in all-projects docs-v1.2.4-2016-08-03 (docs-v1.2.4-2016-08-03)
Checkout Sha (ac1b1b1a19f73b8e1899c63b549653d1fd71ff8c) NOT the same as tip of upstream/master (27968edd8a160f66c96c8545ad35e3a3eeb8766a)
	Consider a `git reset --hard upstream/master`
-- machine (in docs/)
* (detached from docs-v0.8.0-2016-07-28)
  master
docs-v0.8.0-2016-07-28
Checkout Sha (e093b1589069c9b4ab90c5b14cc0da0cc66786d6) NOT the same as ref: in all-projects docs-v0.8.0-2016-07-28 (docs-v0.8.0-2016-07-28)
Checkout Sha (e093b1589069c9b4ab90c5b14cc0da0cc66786d6) NOT the same as tip of upstream/master (578cb4dc34169efef6752df0863d2fc22a8fcf3a)
	Consider a `git reset --hard upstream/master`
-- notary (in docs/)
* (detached from docs-v0.3-2016-08-03)
docs-v0.3-2016-08-03
Checkout Sha (a6fda67663e158d0f0c1384599a2084724249577) NOT the same as ref: in all-projects docs-v0.3-2016-08-03 (docs-v0.3-2016-08-03)
Checkout Sha (a6fda67663e158d0f0c1384599a2084724249577) NOT the same as tip of upstream/master (ca2008c88619d7197501139070c1aaf2f9281446)
	Consider a `git reset --hard upstream/master`
-- toolbox (in docs/)
* (detached from docs-v1.12.0-2016-07-28)
  master
docs-v1.12.0-2016-07-28
v1.12.0
v1.12.1
Checkout Sha (ad9eac89e92e1e684955a7806e198cb68b935aef) NOT the same as ref: in all-projects docs-v1.12.0-2016-07-28 (docs-v1.12.0-2016-07-28)
Checkout Sha (ad9eac89e92e1e684955a7806e198cb68b935aef) NOT the same as tip of upstream/master (db24b2166089b2bf67841b995015e626bb7a409f)
	Consider a `git reset --hard upstream/master`
-- kitematic (in docs/)
* (detached from v0.12.0)
  master
docs-v0.12.0-2016-07-27
v0.12.0
Checkout Sha (02c9f9607128802c904a454d6cc900b3e9ec4555) NOT the same as ref: in all-projects v0.12.0 (v0.12.0)
Checkout Sha (02c9f9607128802c904a454d6cc900b3e9ec4555) NOT the same as tip of upstream/master (9143fe940657d843ea5ebc52caf1c5f0b043f2da)
	Consider a `git reset --hard upstream/master`
-- hub2-demo (in docs/)
* (detached from docs-2016-08-16)
  master
docs-2016-08-16
Checkout Sha (263fd8d2c1f021481b2833255f9bfe0226b2e354) NOT the same as ref: in all-projects docs-2016-08-16 (docs-2016-08-16)
Checkout Sha (263fd8d2c1f021481b2833255f9bfe0226b2e354) NOT the same as tip of upstream/master (35b35b9a0270c368c588fd1b0bee27d6edc22254)
	Consider a `git reset --hard upstream/master`
-- cloud-docs (in docs/)
* (detached from docs-2016-08-17)
  master
docs-2016-08-17
Checkout Sha (33e56428398878f76d083914dbde44a02f7b1fdb) NOT the same as ref: in all-projects docs-2016-08-17 (docs-2016-08-17)
Checkout Sha (33e56428398878f76d083914dbde44a02f7b1fdb) NOT the same as tip of upstream/master (1b9a757a92d83c875284bb2f90fde97f14277be6)
	Consider a `git reset --hard upstream/master`
-- cloud-docs (in apidocs/layouts/)
* (detached from docs-2016-08-17)
  master
docs-2016-08-17
Checkout Sha (33e56428398878f76d083914dbde44a02f7b1fdb) NOT the same as ref: in all-projects docs-2016-08-17 (docs-2016-08-17)
Checkout Sha (33e56428398878f76d083914dbde44a02f7b1fdb) NOT the same as tip of upstream/master (1b9a757a92d83c875284bb2f90fde97f14277be6)
	Consider a `git reset --hard upstream/master`
-- cloud-docs (in apidocs/)
* (detached from docs-2016-08-17)
  master
docs-2016-08-17
Checkout Sha (33e56428398878f76d083914dbde44a02f7b1fdb) NOT the same as ref: in all-projects docs-2016-08-17 (docs-2016-08-17)
Checkout Sha (33e56428398878f76d083914dbde44a02f7b1fdb) NOT the same as tip of upstream/master (1b9a757a92d83c875284bb2f90fde97f14277be6)
	Consider a `git reset --hard upstream/master`
-- mercury-ui (in docs/)
* (detached from docs-2016-06-20)
  master
docs-2016-06-20
v51.0.0
v52.0.0
v53.0.0
v54.0.0
v55.0.0
v56.0.0
v57.0.0
v58.0.0
v59.0.0
v60.0.0
v61.0.0
v62.0.0
v63.0.0
v64.0.0
v65.0.0
v66.0.0
v67.0.0
v68.0.0
v69.0.0
v70.0.0
v71.0.0
v72.0.0
v73.0.0
v74.0.0
v75.0.0
v76.0.0
v77.0.0
v78.0.0
v79.0.0
Checkout Sha (fdf50f7f057a6d24f0e95dcf68e15f3d05e873bd) NOT the same as ref: in all-projects docs-2016-06-20 (docs-2016-06-20)
Checkout Sha (fdf50f7f057a6d24f0e95dcf68e15f3d05e873bd) NOT the same as tip of upstream/master (68d3072991cfdc904a9c35515705b8551ee71317)
	Consider a `git reset --hard upstream/master`
-- opensource (in docs/)
* (detached from docs-2016-08-03)
docs-2016-08-03
Checkout Sha (b9b87bed67f42891d3ee73993f85a9dcd1e5028d) NOT the same as ref: in all-projects docs-2016-08-03 (docs-2016-08-03)
Checkout Sha (b9b87bed67f42891d3ee73993f85a9dcd1e5028d) NOT the same as tip of upstream/master (9736bd57db38561847648a612867d0f0f9978836)
	Consider a `git reset --hard upstream/master`

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
comparing current checkout to upstream/master

## docs-base,  in docs-base at docs-2016-08-12
Warning: no version field in all-projects.yml for docs-base
NO merge PR found for (+ 7adec600461e7456366df201af4060878dca215b Skip weaveworks url, as it intermittenly 404's (#304)) 

## engine, v1.12.1 in docker at 23cf638307f030cd8d48c9efc21feec18a6f88f8
- PR 25726 (88a6a77) Aug 18 15:20:09 from vieux/new_plugin_system_doc
- 1.12.1 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in 98901906523d43b537a3d6a2861ac831ded7df6a adding some documentation about the new plugin system
- PR 25743 (6fa69d2) Aug 18 17:25:08 from lixiaobing10051267/masterInspect
- 1.12.1 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in 68ef2569842f8ca4dd09a85caca1970d95946547 Check the return message of docker service inspect
- PR 25756 (79c1cd8) Aug 16 13:31:07 from YuPengZTE/shuold-be--
- Warning: no version milestone set for PR(25756)
-  status/3-docs-review 
  - docs/ changes in fe081efa929bab70f72f47ec360cd91a65bd8d46 “ should be "
- PR 25704 (45cb33e) Aug 17 08:39:45 from thaJeztah/cleanup-api-markdown
- 1.12.1 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in eb24e9bbd312b0d7392e8012b757e7c9022c9add Cleanup API docs Markdown formatting and wording
- PR 25726 (88a6a77) Aug 18 15:20:09 from vieux/new_plugin_system_doc
- 1.12.1 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in a7a70433cabe5eb210ef81ff61f953ab9d9e332d edit plugin system doc, fix menu system
- PR 25726 (88a6a77) Aug 18 15:20:09 from vieux/new_plugin_system_doc
- 1.12.1 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in 79aa2b9f6da802ee1380c22c3afc8c0be7c493ce fix broken link
- PR 25751 (bbd5396) Aug 17 07:59:36 from yuexiao-wang/clean-docs
- 1.12.1 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in f8d5b880722bcc87113a08cbb2069b6311b89f39 remove mess words in installation doc
- PR 25709 (b4abe38) Aug 17 14:43:35 from thaJeztah/fix-missing-docs-for-binary-remote-context
- 1.8.0 process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in a5ba032c7421ef7a429e780d12d0f604a045258a Add missing docs about binary remote contexts
- PR 25805 (2d93186) Aug 17 18:07:13 from crosbymichael/oci-import-paths
- Warning: no version milestone set for PR(25805)
-  status/2-code-review 
  - docs/ changes in 041e5a21dc0a8856448e3a9ad91e8535b8a7d00d Replace old oci specs import with runtime-specs
- PR 25815 (d854c4f) Aug 18 13:26:59 from justincormack/capdoc
- Warning: no version milestone set for PR(25815)
-  process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in bf7a3f010443ecd614baf0450c3193b1f5e52bc2 Split list of capabilities into those added by default and those not
- PR 25792 (90308fd) Aug 18 15:08:41 from lixiaobing10051267/masterSwarmLeave
- Warning: no version milestone set for PR(25792)
-  process/docs-cherry-pick status/3-docs-review 
  - docs/ changes in fd660e21bf6568c3f98424bdff3b9672cd2a3ef8 Specify woker node for docker swarm leave command Signed-off-by: lixiaobing10051267 <li.xiaobing1@zte.com.cn>

## pinata, v1.12.0 in pinata at docs-v1.12.0-beta23-2016-08-16-update1

## cs-engine, v1.11.1-cs2 in cs-docker at 71a04c87ee4654756f870a7c095ce725220da171

## docker-trusted-registry, v2.0.3 in dhe-deploy at docs-v2.0.3-2016-08-11
- PR 2672 (36fa452) Aug 11 23:31:16 from joaofnfernandes/2.0.3-release-notes
- Warning: no version milestone set for PR(2672)
  - docs/ changes in c3b60319e7e3d0dace4edd49b82931b7c78fea89 Fix typo in docs

## apidocs, v2.0.3 in dhe-deploy at docs-v2.0.3-2016-08-11

## ucp, v1.1.2 in orca at docs-v1.1.2-2016-08-03
- PR 2426 (2ae744e) Aug  3 22:03:31 from alexmavr/system-reqs
- Warning: no version milestone set for PR(2426)
-  kind/documentation 
  - docs/ changes in 57a2bdaab3c34c6020a344752a6cf748c0e98efc updated system requirements for UCP Seattle
- PR 2428 (2a10d40) Aug  4 17:47:41 from alexmavr/install-docs
- Warning: no version milestone set for PR(2428)
-  kind/documentation 
  - docs/ changes in 53b3c8636cdd684c55f0e86af744fa114c0b4de7 UCP Seattle Beta installation docs
- PR 2505 (0e5facc) Aug  8 21:31:02 from joaofnfernandes/install-docs
- Warning: no version milestone set for PR(2505)
  - docs/ changes in e5944e1c0977bc61b13d856cbfd13171242c090a Cleanup docs for installing UCP seattle
- PR 2523 (c21104e) Aug 11 02:21:32 from avuserow/uninstall-docs
- Warning: no version milestone set for PR(2523)
  - docs/ changes in 9bc560850f601e95a2a7db16092c8b793fc43c50 [docs] Update uninstall for Seattle
- PR 2598 (1e2c585) Aug 11 21:03:29 from joaofnfernandes/add-nodes
- Warning: no version milestone set for PR(2598)
  - docs/ changes in 6dbbbe4163c059497bf3bd1333a414c6cdace0d0 Add docs for adding/removing nodes
- PR 2606 (64de39a) Aug 11 22:26:47 from joaofnfernandes/dtr-2.0.3
- Warning: no version milestone set for PR(2606)
  - docs/ changes in 4f7a1b6e54953edc77203f7751b372619dd12c50 Update offline install docs for DTR 2.0.3
- PR 2601 (5325ff4) Aug 11 23:30:58 from joaofnfernandes/uninstall-docs
- Warning: no version milestone set for PR(2601)
  - docs/ changes in 47bd31bb04d097763b27f29c06538195f4809a73 Clean docs for uninstalling UCP
ERROR parsing Version(Seattle-6) in milestone of PR(2624) Invalid character(s) found in major number "Seattle"
- PR 2624 (aff6f93) Aug 12 19:04:08 from sarahlynnpark/add-alias
- Seattle-6 kind/documentation 
  - docs/ changes in 15ebba9ae7a4a134c4bb39d3f0dd4b513ea29494 Add redirect alias from DTR to UCP's LDAP config docs
- PR 2637 (5f83a95) Aug 16 00:44:22 from joaofnfernandes/permission-levels
- Warning: no version milestone set for PR(2637)
  - docs/ changes in a09b75bd296a537aa1d29e1d58ef12bf7113a51a Update permission docs
- PR 2650 (100af92) Aug 16 01:13:31 from joaofnfernandes/upgrade-major
- Warning: no version milestone set for PR(2650)
-  kind/documentation 
  - docs/ changes in b684e89b70e895a09465c23a2198d2632258aa60 Add docs to upgrade from 1.1 to Seattle
- PR 2693 (4372678) Aug 16 20:53:03 from joaofnfernandes/cli-docs
- Warning: no version milestone set for PR(2693)
-  kind/documentation 
  - docs/ changes in 4bf15d5d6101101c6b796a1ce682060ca280fb12 Update CLI reference docs
NO merge PR found for (+ ae0cd84719a684cbe310b3512474af4a862b1011 Remove WIP tag from reservations (#2703)) 
NO merge PR found for (+ 18fe9580c03b52bdced1dc1d599f381844a324b4 Remove deploy button from applications list (#2714)) 
NO merge PR found for (+ 9c3d21f235959498e4c61e6cf672065b9c00c8f9 Spruce up dashboard and fix some reported issues/bugs with it (#2563)) 
NO merge PR found for (+ bfb25097639d359363e17e5370bfc5c9e41e8231 Zero state pages for images and volumes (#2717)) 

## registry, v2.5.0 in distribution at docs-v2.5.0-2016-07-28
- PR 1877 (bfa0a9c) Aug  2 16:15:23 from spacexnice/master
- Warning: no version milestone set for PR(1877)
-  group/distribution status/0-triage 
  - docs/ changes in 87917f30529e6a7fca8eaff2932424915fb11225 Add 'objectAcl' Option to the S3 Storage Backend (#1867)
- PR 1839 (c4297ef) Aug 16 18:48:06 from adamvduke/adamvduke/allow-http2-registry-clients
- Warning: no version milestone set for PR(1839)
-  group/distribution status/0-triage 
  - docs/ changes in ac009c86f17b4798f8d859503de578bf22e9ad83 Allow registry clients to connect via http2
- PR 1906 (010e063) Aug 17 22:22:32 from nwt/s3-multipart-copy
- Warning: no version milestone set for PR(1906)
-  group/distribution status/0-triage 
  - docs/ changes in 63468ef4a85fabb756e799c2534f0df2f3c8167c Use multipart upload API in S3 Move method

## compose, v1.8.0 in compose at docs-v1.8.0-2016-08-03
- PR 3704 (c3fd6a8) Jul  7 23:01:55 from aanand/update-install-and-changelog-for-1.7.1
- 1.7.1 status/0-triage 
  - docs/ changes in 49d4fd27952433feb20bc22117aba4766c15c1c1 Update install.md and CHANGELOG.md for 1.7.1
- PR 3778 (dec2c83) Jul 27 22:04:43 from shin-/1.8.0-release-master-changes
- Warning: no version milestone set for PR(3778)
-  status/0-triage 
  - docs/ changes in 22c0779a498ee701c22b857669d3f43a0d404f27 Bump 1.8.0-rc1
- PR 3542 (acfe100) Aug 16 12:58:33 from jfroche/add_swappiness
- Warning: no version milestone set for PR(3542)
-  status/0-triage 
  - docs/ changes in d824cb9b0678ec2ad460b034231c00c05df8c0fe Add support for swappiness constraint

## swarm, v1.2.4 in swarm at docs-v1.2.4-2016-08-03
- PR 2453 (25919db) Aug 18 21:09:32 from RangerMauve/patch-3
- Warning: no version milestone set for PR(2453)
-  status/3-docs-review 
  - docs/ changes in c5d5e7341b8fbd3e1d1457d38c5f47b6b075822d Fix name of private key in configure-tls guide

## machine, v0.8.0 in machine at docs-v0.8.0-2016-07-28
- PR 3572 (c257d87) Jul 29 23:37:21 from StackPointCloud/master
- 0.8.0 dco/no status/3-docs-review 
  - docs/ changes in d94d02f4293beed6df7cabc6a3a476370b4050fd ProfitBricks Docker Machine Driver Doc Update
- PR 3665 (19f0bc4) Aug 16 19:29:34 from ahmetalpbalkan/azure-sp
- Warning: no version milestone set for PR(3665)
-  status/0-triage 
  - docs/ changes in 81b76355c63b5b0e8ee0e6e5c093e636ac18f563 azure: Service principal authentication
- PR 3665 (19f0bc4) Aug 16 19:29:34 from ahmetalpbalkan/azure-sp
- Warning: no version milestone set for PR(3665)
-  status/0-triage 
  - docs/ changes in 554637fbd7706498f0fd01e96fb612698435a858 azure: update docs

## notary, v0.3 in notary at docs-v0.3-2016-08-03

## toolbox, v1.12.0 in toolbox at docs-v1.12.0-2016-07-28

## kitematic, v0.12.0 in kitematic at v0.12.0

## docker-hub,  in hub2-demo at docs-2016-08-16
Warning: no version field in all-projects.yml for docker-hub

## docker-cloud,  in cloud-docs at docs-2016-08-17
Warning: no version field in all-projects.yml for docker-cloud

## cloud-api-docs-layout,  in cloud-docs at docs-2016-08-17
Warning: no version field in all-projects.yml for cloud-api-docs-layout

## cloud-api-docs,  in cloud-docs at docs-2016-08-17
Warning: no version field in all-projects.yml for cloud-api-docs

## docker-store,  in mercury-ui at docs-2016-06-20
Warning: no version field in all-projects.yml for docker-store
- PR 248 (17042bc) Aug  4 00:17:27 from sanscontext/menu-edit
- Warning: no version milestone set for PR(248)
  - docs/ changes in 6a8a35b99bfd299c2c5c6a763020febb614ba0b3 add note to menu
NO merge PR found for (+ 5bec7542493f9f3cf6ff7410475876dc1bfd118e 76.0.0) 
NO merge PR found for (+ 018a19a6c82be24bf634e7097f7b60b060a091a8 accept vendor agreement on update (#264)) 
NO merge PR found for (+ 8804dd6a0c04c1999b2081510b8d36ff755dab1a 77.0.0) 
NO merge PR found for (+ 185ef913db594ca2d757de98660a0d0ed82fe161 Update Product Details Page for images to handle multiple rate plans. (#262)) 
NO merge PR found for (+ bb8e540386b7fd415474332dbbba3b72fbb1fdd0 Rename all Purchase to DDCPurchase. Update styles on DDC Purchase Detail Form (#268)) 
NO merge PR found for (+ d34bfb69050d551491dbc3df9f4d55c280adef5d Fix hub search results as a result of the 'name' change (#267)) 
NO merge PR found for (+ 0c47f39a9b1cba75241d5f33f956b1080823c9ea Remove Tags and Comments from Store repos / images (#266)) 
NO merge PR found for (+ 10b8c315aa2dc6d753f0dc3592962b95e689e8c2 Un-break subscriptions (#265)) 
NO merge PR found for (+ f48efb6aafab72cbe9a1fb356150877954f34556 MER-787 - Prevent whitespace-only fields for DDC purchase forms (#269)) 
NO merge PR found for (+ 52b8a246b55a73557bf2667856eba8b9dc8d5fc3 78.0.0) 
NO merge PR found for (+ 9ca0eee87c9e54d4fc8ff43fc8b09cad99179027 Step 3 - Product Tiers (#260)) 
NO merge PR found for (+ f68538288854b6f62a25bcc4b87da92ef2785a41 79.0.0) 
NO merge PR found for (+ 68d3072991cfdc904a9c35515705b8551ee71317 Refactor DDC Purchase and Purchase to fit coding standards and structure (#270)) 

## opensource,  in opensource at docs-2016-08-03
Warning: no version field in all-projects.yml for opensource

```

## README example updates

`gendoc` also is able to rewrite its README.md file using `gendoc readme`.
This will read the README.md file, and look for any "```" code markers.
Inside the code sections, it will run all lines starting with `$` and
add whatever the output is.



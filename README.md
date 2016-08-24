Using /tmp/example795102549 to run commands
Using /tmp/example338576613 to run commands

# Docker documentation tooling

## pre-req's

You need to have a working installation of `git`

## install

Download the latest release for your platform from https://github.com/docker/gendoc/releases

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
   2016-08-19
   
COMMANDS:
     version   return the version
     clone     clone repos from the ./docs.docker.com/all-projects.yml file
     checkout  checkout versions from ./docs.docker.com/all-projects.yml file
     install   Install gendoc and its pre-req's into your PATH
     readme    Parse the README file and update any inline command examples
     release   Prepare and ship a docs release.
     remote    Add a git remote - 2 arguments, name to give remote (origin), and organisation/Username on GitHub
     render    render html using the the source files currently available on disk.
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
drwx------ 2 ubuntu ubuntu  4096 Aug 19 05:06 .
drwxrwxrwt 6 root   root   20480 Aug 19 05:06 ..

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
drwx------ 20 ubuntu ubuntu  4096 Aug 19 05:07 .
drwxrwxrwt  6 root   root   20480 Aug 19 05:06 ..
drwxrwxr-x  6 ubuntu ubuntu  4096 Aug 19 05:07 cloud-docs
drwxrwxr-x 11 ubuntu ubuntu  4096 Aug 19 05:07 compose
drwxrwxr-x 28 ubuntu ubuntu  4096 Aug 19 05:07 cs-docker
drwxrwxr-x 35 ubuntu ubuntu  4096 Aug 19 05:07 dhe-deploy
drwxrwxr-x 20 ubuntu ubuntu  4096 Aug 19 05:07 distribution
drwxrwxr-x 37 ubuntu ubuntu  4096 Aug 19 05:06 docker
drwxrwxr-x  7 ubuntu ubuntu  4096 Aug 19 05:06 docs-base
drwxrwxr-x  5 ubuntu ubuntu  4096 Aug 19 05:06 docs.docker.com
drwxrwxr-x 14 ubuntu ubuntu  4096 Aug 19 05:07 hub2-demo
drwxrwxr-x 14 ubuntu ubuntu  4096 Aug 19 05:07 kitematic
drwxrwxr-x 17 ubuntu ubuntu  4096 Aug 19 05:07 machine
drwxrwxr-x  9 ubuntu ubuntu  4096 Aug 19 05:07 mercury-ui
drwxrwxr-x 23 ubuntu ubuntu  4096 Aug 19 05:07 notary
drwxrwxr-x  6 ubuntu ubuntu  4096 Aug 19 05:07 opensource
drwxrwxr-x 28 ubuntu ubuntu  4096 Aug 19 05:07 orca
drwxrwxr-x 11 ubuntu ubuntu  4096 Aug 19 05:06 pinata
drwxrwxr-x 17 ubuntu ubuntu  4096 Aug 19 05:07 swarm
drwxrwxr-x  8 ubuntu ubuntu  4096 Aug 19 05:07 toolbox

```

If there is no `docs.docker.com` repo found, will clone it, and then
will clone any missing repositories mentioned in the currently checked out 
`docs.docker.com/all-projects.yaml`

## to serve the master docs to a browser (port 8080)


```
$ gendoc checkout master
Checking out docs.docker.com master.
Same as all-projects.yml: your checkout 84b2b9e323326614d60374bf628b058f83ab667f is at upstream/master
publish-set: v1.13-dev
-- docs-base
Same as all-projects.yml: your checkout 7adec600461e7456366df201af4060878dca215b is at upstream/master
-- docker
Same as all-projects.yml: your checkout 09e1de2080fd3b0bafb38adbd4b8c12ee949794d is at upstream/master
-- pinata
Same as all-projects.yml: your checkout 8c11c14b46880079351e6f6503119e2ee6ef76ac is at upstream/master
-- cs-docker
Same as all-projects.yml: your checkout 71a04c87ee4654756f870a7c095ce725220da171 is at upstream/master
-- dhe-deploy
Same as all-projects.yml: your checkout eb01555b9264d2a481fc87c6933909e7d713bf34 is at upstream/master
-- dhe-deploy
Same as all-projects.yml: your checkout eb01555b9264d2a481fc87c6933909e7d713bf34 is at upstream/master
-- orca
Same as all-projects.yml: your checkout bfb25097639d359363e17e5370bfc5c9e41e8231 is at upstream/master
-- distribution
Same as all-projects.yml: your checkout 010e063270be37cfa8547ccfb9717e5d874c88a8 is at upstream/master
-- kitematic
Same as all-projects.yml: your checkout 9143fe940657d843ea5ebc52caf1c5f0b043f2da is at upstream/master
-- compose
Same as all-projects.yml: your checkout acfe100686fd95d524ff102c0b5fccff0bc79d8c is at upstream/master
-- swarm
Same as all-projects.yml: your checkout 27968edd8a160f66c96c8545ad35e3a3eeb8766a is at upstream/master
-- machine
Same as all-projects.yml: your checkout 578cb4dc34169efef6752df0863d2fc22a8fcf3a is at upstream/master
-- notary
Same as all-projects.yml: your checkout ca2008c88619d7197501139070c1aaf2f9281446 is at upstream/master
-- toolbox
Same as all-projects.yml: your checkout db24b2166089b2bf67841b995015e626bb7a409f is at upstream/master
-- hub2-demo
Same as all-projects.yml: your checkout 35b35b9a0270c368c588fd1b0bee27d6edc22254 is at upstream/master
-- cloud-docs
Same as all-projects.yml: your checkout 1b9a757a92d83c875284bb2f90fde97f14277be6 is at upstream/master
-- cloud-docs
Same as all-projects.yml: your checkout 1b9a757a92d83c875284bb2f90fde97f14277be6 is at upstream/master
-- cloud-docs
Same as all-projects.yml: your checkout 1b9a757a92d83c875284bb2f90fde97f14277be6 is at upstream/master
-- mercury-ui
Same as all-projects.yml: your checkout 68d3072991cfdc904a9c35515705b8551ee71317 is at upstream/master
-- opensource
Same as all-projects.yml: your checkout 9736bd57db38561847648a612867d0f0f9978836 is at upstream/master
```

## test

You can use `gendoc` to run a local markdownlint test on your workspace:

```
$ gendoc test
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
Same as all-projects.yml: your checkout 71a04c87ee4654756f870a7c095ce725220da171 is at 71a04c87ee4654756f870a7c095ce725220da171
-- dhe-deploy
-- dhe-deploy
Same as all-projects.yml: your checkout 139a5d128584da25eee4b730c35497d8c0840515 is at refs/tags/docs-v2.0.3-2016-08-11
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
Same as all-projects.yml: your checkout 33e56428398878f76d083914dbde44a02f7b1fdb is at refs/tags/docs-2016-08-17
-- cloud-docs
Same as all-projects.yml: your checkout 33e56428398878f76d083914dbde44a02f7b1fdb is at refs/tags/docs-2016-08-17
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
INFO: 2016/08/19 05:07:53 hugo.go:463: Using config file: /tmp/example795102549/docs-source/v1.12/config.toml
WARN: 2016/08/19 05:07:53 hugo.go:557: Unable to find Static Directory: /tmp/example795102549/docs-source/v1.12/static/
INFO: 2016/08/19 05:07:53 hugo.go:566: /tmp/example795102549/docs-source/v1.12/themes/docker-2016/static is the only static directory available to sync from
INFO: 2016/08/19 05:07:53 hugo.go:607: removing all files from destination that don't exist in static dirs
INFO: 2016/08/19 05:07:53 hugo.go:609: syncing static files to /tmp/example795102549/docs-html/v1.12/
Started building site
INFO: 2016/08/19 05:07:54 site.go:1251: found taxonomies: map[string]string{"tag":"tags", "category":"categories"}
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/started/" translated to "mac/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/" translated to "engine/userguide/containers/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/extend/authorization/" translated to "engine/extend/authorization/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/deploy-to-cloud/" translated to "docker-cloud/feature-reference/deploy-to-cloud/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/deploy-to-cloud/" translated to "docker-cloud/tutorials/deploy-to-cloud/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/awslogs/" translated to "engine/reference/logging/awslogs/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/accounts/" translated to "docker-trusted-registry/accounts/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/manage/monitor-manage-users/" translated to "ucp/manage/monitor-manage-users/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/user-management/manage-users/" translated to "ucp/user-management/manage-users/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/automated-build/" translated to "docker-cloud/feature-reference/automated-build/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/automated-testing/" translated to "docker-cloud/feature-reference/automated-testing/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/auto-destroy/" translated to "docker-cloud/feature-reference/auto-destroy/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/autorestart/" translated to "docker-cloud/feature-reference/autorestart/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/auto-redeploy/" translated to "docker-cloud/feature-reference/auto-redeploy/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/host_integration/" translated to "engine/articles/host_integration/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/dockerfile_best-practices/" translated to "engine/articles/dockerfile_best-practices/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/optimize-dockerfiles/" translated to "docker-cloud/getting-started/intermediate/optimize-dockerfiles/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/optimize-dockerfiles/" translated to "docker-cloud/tutorials/optimize-dockerfiles/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/misc/breaking/" translated to "engine/misc/breaking/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/step_four/" translated to "mac/step_four/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/step_four/" translated to "windows/step_four/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/step_four/" translated to "linux/step_four/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/dockerimages/" translated to "engine/userguide/containers/dockerimages/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/dockerimages/" translated to "engine/userguide/dockerimages/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/" translated to "docker-trusted-registry/cs-engine/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cse-release-notes/" translated to "docker-trusted-registry/cse-release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/release-notes/release-notes/" translated to "docker-trusted-registry/cs-engine/release-notes/release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/cloud/cloud/" translated to "engine/installation/cloud/cloud/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/compose/yml" translated to "compose/yml/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/overview/" translated to "engine/reference/logging/overview/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/configuring/" translated to "engine/articles/configuring/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/admin/configuring/" translated to "engine/admin/configuring/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/deployment-strategies/" translated to "docker-cloud/feature-reference/deployment-strategies/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/systemd/" translated to "engine/articles/systemd/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/step_five/" translated to "mac/step_five/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/step_five/" translated to "windows/step_five/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/step_five/" translated to "linux/step_five/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/baseimages/" translated to "engine/articles/baseimages/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/load-balance-hello-world/" translated to "docker-cloud/getting-started/intermediate/load-balance-hello-world/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/load-balance-hello-world/" translated to "docker-cloud/tutorials/load-balance-hello-world/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/12_data_management_with_volumes/" translated to "docker-cloud/getting-started/python/12_data_management_with_volumes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/6_define_environment_variables/" translated to "docker-cloud/getting-started/python/6_define_environment_variables/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/6_define_environment_variables/" translated to "docker-cloud/getting-started/golang/6_define_environment_variables/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/soft-garbage/" translated to "docker-trusted-registry/soft-garbage/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/repos-and-images/delete-images/" translated to "docker-trusted-registry/repos-and-images/delete-images/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/deploy-application/" translated to "ucp/deploy-application/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/" translated to "docker-cloud/getting-started/python/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/" translated to "docker-cloud/getting-started/golang/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/03-create-cluster/" translated to "swarm/swarm_at_scale/03-create-cluster/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/02-deploy-infra/" translated to "swarm/swarm_at_scale/02-deploy-infra/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/5_deploy_the_app_as_a_service/" translated to "docker-cloud/getting-started/python/5_deploy_the_app_as_a_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/5_deploy_the_app_as_a_service/" translated to "docker-cloud/getting-started/golang/5_deploy_the_app_as_a_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/04-deploy-app/" translated to "swarm/swarm_at_scale/04-deploy-app/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/your_first_node/" translated to "docker-cloud/getting-started/beginner/your_first_node/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/deploy_first_node/" translated to "docker-cloud/getting-started/beginner/deploy_first_node/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/your_first_service/" translated to "docker-cloud/getting-started/beginner/your_first_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/deploy_first_service/" translated to "docker-cloud/getting-started/beginner/deploy_first_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/deploy-tags/" translated to "docker-cloud/feature-reference/deploy-tags/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/misc/deprecated/" translated to "engine/misc/deprecated/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/slack-integration/" translated to "docker-cloud/tutorials/slack-integration/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/misc/" translated to "engine/misc/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/introduction/understanding-docker/" translated to "introduction/understanding-docker/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/basics/" translated to "engine/userguide/basics/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/quickstart.md" translated to "engine/quickstart.md/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/api/swarm-api/" translated to "api/swarm-api/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/swarm/api/" translated to "swarm/api/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-hub-enterprise/" translated to "docker-hub-enterprise/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/overview/" translated to "docker-trusted-registry/overview/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/networking/dockernetworks/" translated to "engine/userguide/networking/dockernetworks/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/docker-toolbox/" translated to "mackit/docker-toolbox/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/security/" translated to "engine/articles/security/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/evaluation-install/" translated to "ucp/evaluation-install/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/misc/faq/" translated to "engine/misc/faq/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/faqs/" translated to "mackit/faqs/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/osxfs/" translated to "mackit/osxfs/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/step_three/" translated to "mac/step_three/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/step_three/" translated to "windows/step_three/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/step_three/" translated to "linux/step_three/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/fluentd/" translated to "engine/reference/logging/fluentd/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/started/" translated to "mac/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/started/" translated to "windows/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/started/" translated to "linux/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/getting-started/" translated to "getting-started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/" translated to "mackit/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/getting-started/" translated to "mackit/getting-started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/" translated to "mac/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/started/" translated to "mac/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-for-mac/started/" translated to "docker-for-mac/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/winkit/getting-started/" translated to "winkit/getting-started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/winkit/" translated to "winkit/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/" translated to "windows/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/started/" translated to "windows/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-for-windows/started/" translated to "docker-for-windows/started/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/dockerizing/" translated to "engine/userguide/containers/dockerizing/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/dockerizing/" translated to "engine/userguide/dockerizing/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/swarm/how-swarm-mode-works/" translated to "engine/swarm/how-swarm-mode-works/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/engine-ami-launch/" translated to "docker-trusted-registry/install/engine-ami-launch/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/install-csengine/" translated to "docker-trusted-registry/install/install-csengine/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/install/" translated to "docker-trusted-registry/cs-engine/install/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/dtr-ami-byol-launch/" translated to "docker-trusted-registry/install/dtr-ami-byol-launch/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/dtr-ami-bds-launch/" translated to "docker-trusted-registry/install/dtr-ami-bds-launch/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/dtr-vhd-azure/" translated to "docker-trusted-registry/install/dtr-vhd-azure/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/install-dtr/" translated to "docker-trusted-registry/install/install-dtr/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/step_one/" translated to "mac/step_one/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/step_one/" translated to "windows/step_one/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/step_one/" translated to "linux/step_one/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/production-install/" translated to "ucp/production-install/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/archlinux/" translated to "engine/installation/archlinux/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/cruxlinux/" translated to "engine/installation/cruxlinux/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/centos/" translated to "engine/installation/centos/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/debian/" translated to "engine/installation/debian/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/fedora/" translated to "engine/installation/fedora/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/frugalware/" translated to "engine/installation/frugalware/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/gentoolinux/" translated to "engine/installation/gentoolinux/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/oracle/" translated to "engine/installation/oracle/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/rhel/" translated to "engine/installation/rhel/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/ubuntulinux/" translated to "engine/installation/ubuntulinux/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/SUSE/" translated to "engine/installation/SUSE/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/dtr-integration/" translated to "ucp/dtr-integration/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/intro_cloud/" translated to "docker-cloud/getting-started/beginner/intro_cloud/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/1_introduction/" translated to "docker-cloud/getting-started/python/1_introduction/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/1_introduction/" translated to "docker-cloud/getting-started/golang/1_introduction/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/journald/" translated to "engine/reference/logging/journald/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/faq/docker-errors-faq/" translated to "docker-cloud/faq/docker-errors-faq/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/about/" translated to "swarm/swarm_at_scale/about/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/last_page/" translated to "mac/last_page/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/last_page/" translated to "windows/last_page/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/last_page/" translated to "linux/last_page/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/multi-arch/" translated to "mackit/multi-arch/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/license/" translated to "docker-trusted-registry/license/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/compose/env" translated to "compose/env/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-aws/" translated to "docker-cloud/getting-started/beginner/link-aws/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-aws/" translated to "docker-cloud/getting-started/link-aws/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-do/" translated to "docker-cloud/getting-started/beginner/link-do/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-do/" translated to "docker-cloud/getting-started/link-do/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-azure/" translated to "docker-cloud/getting-started/beginner/link-azure/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-azure/" translated to "docker-cloud/getting-started/link-azure/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-packet/" translated to "docker-cloud/getting-started/beginner/link-packet/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-packet/" translated to "docker-cloud/getting-started/link-packet/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/beginner/link-softlayer/" translated to "docker-cloud/getting-started/beginner/link-softlayer/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/link-softlayer/" translated to "docker-cloud/getting-started/link-softlayer/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/link-source/" translated to "docker-cloud/tutorials/link-source/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/use-hosted/" translated to "docker-cloud/getting-started/use-hosted/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/ambassador_pattern_linking/" translated to "engine/articles/ambassador_pattern_linking/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/9_load-balance_the_service/" translated to "docker-cloud/getting-started/python/9_load-balance_the_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/9_load-balance_the_service/" translated to "docker-cloud/getting-started/golang/9_load-balance_the_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/log_tags/" translated to "engine/reference/logging/log_tags/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/" translated to "engine/reference/logging/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/troubleshoot/" translated to "mackit/troubleshoot/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/troubleshoot/" translated to "windows/troubleshoot/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/dockervolumes/" translated to "engine/userguide/containers/dockervolumes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/dockervolumes/" translated to "engine/userguide/dockervolumes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/stacks/" translated to "docker-cloud/feature-reference/stacks/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/monitor-troubleshoot/monitor/" translated to "docker-trusted-registry/monitor-troubleshoot/monitor/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/manage/" translated to "ucp/manage/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/manage/monitor-ucp/" translated to "ucp/manage/monitor-ucp/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/networkigncontainers/" translated to "engine/userguide/containers/networkigncontainers/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/networkigncontainers/" translated to "engine/userguide/networkigncontainers/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/networking/" translated to "mackit/networking/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/amazon/" translated to "engine/installation/amazon/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/google/" translated to "engine/installation/google/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/softlayer/" translated to "engine/installation/softlayer/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/azure/" translated to "engine/installation/azure/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/rackspace/" translated to "engine/installation/rackspace/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/installation/joyent/" translated to "engine/installation/joyent/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-hub/overview/" translated to "docker-hub/overview/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/compose/reference/docker-compose/" translated to "compose/reference/docker-compose/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/plan-production-install/" translated to "ucp/plan-production-install/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/dsc/" translated to "engine/articles/dsc/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/3_prepare_the_app/" translated to "docker-cloud/getting-started/python/3_prepare_the_app/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/3_prepare_the_app/" translated to "docker-cloud/getting-started/golang/3_prepare_the_app/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cse-prior-release-notes/" translated to "docker-trusted-registry/cse-prior-release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/release-notes/prior-release-notes/" translated to "docker-trusted-registry/cs-engine/release-notes/prior-release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/prior-release-notes/" translated to "docker-trusted-registry/prior-release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/https/" translated to "engine/articles/https/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/10_provision_a_data_backend_for_your_service/" translated to "docker-cloud/getting-started/python/10_provision_a_data_backend_for_your_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/10_provision_a_data_backend_for_your_service/" translated to "docker-cloud/getting-started/golang/10_provision_a_data_backend_for_your_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/ports/" translated to "docker-cloud/feature-reference/ports/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/quick-start/" translated to "docker-trusted-registry/quick-start/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/repos-and-images/create-repo/" translated to "docker-trusted-registry/repos-and-images/create-repo/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/userguide/" translated to "docker-trusted-registry/userguide/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/repos-and-images/push-and-pull-images/" translated to "docker-trusted-registry/repos-and-images/push-and-pull-images/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/pushing-images-to-dockercloud/" translated to "docker-cloud/getting-started/intermediate/pushing-images-to-dockercloud/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/pushing-images-to-dockercloud/" translated to "docker-cloud/tutorials/pushing-images-to-dockercloud/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/4_push_to_cloud_registry/" translated to "docker-cloud/getting-started/python/4_push_to_cloud_registry/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/4_push_to_cloud_registry/" translated to "docker-cloud/getting-started/golang/4_push_to_cloud_registry/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/contributing/contributing" translated to "contributing/contributing/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/service-redeploy/" translated to "docker-cloud/feature-reference/service-redeploy/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/registry/overview/" translated to "registry/overview/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mackit/release-notes/" translated to "mackit/release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/winkit/release-notes/" translated to "winkit/release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/release-notes/" translated to "docker-trusted-registry/cs-engine/release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/registry_mirror/" translated to "engine/articles/registry_mirror/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/usingdocker/" translated to "engine/userguide/containers/usingdocker/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/run_metrics" translated to "engine/articles/run_metrics/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/ssh-into-a-node/" translated to "docker-cloud/getting-started/intermediate/ssh-into-a-node/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/ssh-into-a-node/" translated to "docker-cloud/tutorials/ssh-into-a-node/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/faq/how-ssh-nodes/" translated to "docker-cloud/faq/how-ssh-nodes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/7_scale_the_service/" translated to "docker-cloud/getting-started/python/7_scale_the_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/7_scale_the_service/" translated to "docker-cloud/getting-started/golang/7_scale_the_service/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/service-scaling/" translated to "docker-cloud/feature-reference/service-scaling/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/api-roles/" translated to "docker-cloud/feature-reference/api-roles/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/service-links/" translated to "docker-cloud/feature-reference/service-links/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/networking/" translated to "ucp/networking/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/high-availability/high-availability/" translated to "docker-trusted-registry/high-availability/high-availability/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/understand_ha/" translated to "ucp/understand_ha/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/2_set_up/" translated to "docker-cloud/getting-started/python/2_set_up/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/2_set_up/" translated to "docker-cloud/getting-started/golang/2_set_up/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/logging/splunk/" translated to "engine/reference/logging/splunk/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/stack-yaml-reference/" translated to "docker-cloud/feature-reference/stack-yaml-reference/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/11_service_stacks/" translated to "docker-cloud/getting-started/python/11_service_stacks/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/registry/storagedrivers/" translated to "registry/storagedrivers/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/containers/dockerrepos/" translated to "engine/userguide/containers/dockerrepos/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/userguide/dockerrepos/" translated to "engine/userguide/dockerrepos/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/swarm/manager-administration-guide/" translated to "engine/swarm/manager-administration-guide/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/step_six/" translated to "mac/step_six/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/step_six/" translated to "windows/step_six/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/step_six/" translated to "linux/step_six/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/intermediate/installing-cli/" translated to "docker-cloud/getting-started/intermediate/installing-cli/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/installing-cli/" translated to "docker-cloud/getting-started/installing-cli/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/installing-cli/" translated to "docker-cloud/tutorials/installing-cli/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/adminguide/" translated to "docker-trusted-registry/adminguide/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/kv_store/" translated to "ucp/kv_store/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/swarm/swarm_at_scale/05-troubleshoot/" translated to "swarm/swarm_at_scale/05-troubleshoot/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/release-notes/release-notes/" translated to "docker-trusted-registry/release-notes/release-notes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/mac/step_two/" translated to "mac/step_two/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/windows/step_two/" translated to "windows/step_two/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/linux/step_two/" translated to "linux/step_two/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/cs-engine/upgrade/" translated to "docker-trusted-registry/cs-engine/upgrade/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/docker-upgrade/" translated to "docker-cloud/feature-reference/docker-upgrade/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/docker-upgrade/" translated to "docker-cloud/tutorials/docker-upgrade/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/ucp/upgrade-ucp/" translated to "ucp/upgrade-ucp/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/upgrade/" translated to "docker-trusted-registry/install/upgrade/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-trusted-registry/install/upgrade/upgrade-minor/" translated to "docker-trusted-registry/install/upgrade/upgrade-minor/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/byoh/" translated to "docker-cloud/feature-reference/byoh/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/byoh/" translated to "docker-cloud/tutorials/byoh/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/use-byon/" translated to "docker-cloud/getting-started/use-byon/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/triggers/" translated to "docker-cloud/feature-reference/triggers/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/chef/" translated to "engine/articles/chef/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/faq/cloud-on-packet.net-faq/" translated to "docker-cloud/faq/cloud-on-packet.net-faq/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/faq/cloud-on-aws-faq/" translated to "docker-cloud/faq/cloud-on-aws-faq/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/puppet/" translated to "engine/articles/puppet/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/using_supervisord/" translated to "engine/articles/using_supervisord/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/articles/certificates/" translated to "engine/articles/certificates/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/python/8_view_logs/" translated to "docker-cloud/getting-started/python/8_view_logs/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/getting-started/golang/8_view_logs/" translated to "docker-cloud/getting-started/golang/8_view_logs/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/tutorials/download-volume-data/" translated to "docker-cloud/tutorials/download-volume-data/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/docker-cloud/feature-reference/volumes/" translated to "docker-cloud/feature-reference/volumes/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/commandline/daemon/" translated to "engine/reference/commandline/daemon/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/commandline/node_tasks/" translated to "engine/reference/commandline/node_tasks/index.html"
INFO: 2016/08/19 05:07:54 htmlredirect.go:115: Alias "/engine/reference/commandline/service_tasks/" translated to "engine/reference/commandline/service_tasks/index.html"
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section engine: [section/engine.html _default/section.html _default/list.html indexes/engine.html _default/indexes.html theme/section/engine.html theme/_default/section.html theme/_default/list.html theme/indexes/engine.html theme/_default/indexes.html theme/section/engine.html theme/_default/section.html theme/_default/list.html theme/indexes/engine.html theme/_default/indexes.html theme/theme/section/engine.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/engine.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "engine" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section apidocs: [section/apidocs.html _default/section.html _default/list.html indexes/apidocs.html _default/indexes.html theme/section/apidocs.html theme/_default/section.html theme/_default/list.html theme/indexes/apidocs.html theme/_default/indexes.html theme/section/apidocs.html theme/_default/section.html theme/_default/list.html theme/indexes/apidocs.html theme/_default/indexes.html theme/theme/section/apidocs.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/apidocs.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "apidocs" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section docker-cloud: [section/docker-cloud.html _default/section.html _default/list.html indexes/docker-cloud.html _default/indexes.html theme/section/docker-cloud.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-cloud.html theme/_default/indexes.html theme/section/docker-cloud.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-cloud.html theme/_default/indexes.html theme/theme/section/docker-cloud.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-cloud.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "docker-cloud" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section registry: [section/registry.html _default/section.html _default/list.html indexes/registry.html _default/indexes.html theme/section/registry.html theme/_default/section.html theme/_default/list.html theme/indexes/registry.html theme/_default/indexes.html theme/section/registry.html theme/_default/section.html theme/_default/list.html theme/indexes/registry.html theme/_default/indexes.html theme/theme/section/registry.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/registry.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "registry" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section compose: [section/compose.html _default/section.html _default/list.html indexes/compose.html _default/indexes.html theme/section/compose.html theme/_default/section.html theme/_default/list.html theme/indexes/compose.html theme/_default/indexes.html theme/section/compose.html theme/_default/section.html theme/_default/list.html theme/indexes/compose.html theme/_default/indexes.html theme/theme/section/compose.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/compose.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "compose" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section : [section/.html _default/section.html _default/list.html indexes/.html _default/indexes.html theme/section/.html theme/_default/section.html theme/_default/list.html theme/indexes/.html theme/_default/indexes.html theme/section/.html theme/_default/section.html theme/_default/list.html theme/indexes/.html theme/_default/indexes.html theme/theme/section/.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section swarm: [section/swarm.html _default/section.html _default/list.html indexes/swarm.html _default/indexes.html theme/section/swarm.html theme/_default/section.html theme/_default/list.html theme/indexes/swarm.html theme/_default/indexes.html theme/section/swarm.html theme/_default/section.html theme/_default/list.html theme/indexes/swarm.html theme/_default/indexes.html theme/theme/section/swarm.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/swarm.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "swarm" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section ucp: [section/ucp.html _default/section.html _default/list.html indexes/ucp.html _default/indexes.html theme/section/ucp.html theme/_default/section.html theme/_default/list.html theme/indexes/ucp.html theme/_default/indexes.html theme/section/ucp.html theme/_default/section.html theme/_default/list.html theme/indexes/ucp.html theme/_default/indexes.html theme/theme/section/ucp.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/ucp.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "ucp" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section kitematic: [section/kitematic.html _default/section.html _default/list.html indexes/kitematic.html _default/indexes.html theme/section/kitematic.html theme/_default/section.html theme/_default/list.html theme/indexes/kitematic.html theme/_default/indexes.html theme/section/kitematic.html theme/_default/section.html theme/_default/list.html theme/indexes/kitematic.html theme/_default/indexes.html theme/theme/section/kitematic.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/kitematic.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "kitematic" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section docker-store: [section/docker-store.html _default/section.html _default/list.html indexes/docker-store.html _default/indexes.html theme/section/docker-store.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-store.html theme/_default/indexes.html theme/section/docker-store.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-store.html theme/_default/indexes.html theme/theme/section/docker-store.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-store.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "docker-store" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section toolbox: [section/toolbox.html _default/section.html _default/list.html indexes/toolbox.html _default/indexes.html theme/section/toolbox.html theme/_default/section.html theme/_default/list.html theme/indexes/toolbox.html theme/_default/indexes.html theme/section/toolbox.html theme/_default/section.html theme/_default/list.html theme/indexes/toolbox.html theme/_default/indexes.html theme/theme/section/toolbox.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/toolbox.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "toolbox" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section docker-for-mac: [section/docker-for-mac.html _default/section.html _default/list.html indexes/docker-for-mac.html _default/indexes.html theme/section/docker-for-mac.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-mac.html theme/_default/indexes.html theme/section/docker-for-mac.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-mac.html theme/_default/indexes.html theme/theme/section/docker-for-mac.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-for-mac.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "docker-for-mac" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section docker-for-windows: [section/docker-for-windows.html _default/section.html _default/list.html indexes/docker-for-windows.html _default/indexes.html theme/section/docker-for-windows.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-windows.html theme/_default/indexes.html theme/section/docker-for-windows.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-for-windows.html theme/_default/indexes.html theme/theme/section/docker-for-windows.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-for-windows.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "docker-for-windows" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section opensource: [section/opensource.html _default/section.html _default/list.html indexes/opensource.html _default/indexes.html theme/section/opensource.html theme/_default/section.html theme/_default/list.html theme/indexes/opensource.html theme/_default/indexes.html theme/section/opensource.html theme/_default/section.html theme/_default/list.html theme/indexes/opensource.html theme/_default/indexes.html theme/theme/section/opensource.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/opensource.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "opensource" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section cs-engine: [section/cs-engine.html _default/section.html _default/list.html indexes/cs-engine.html _default/indexes.html theme/section/cs-engine.html theme/_default/section.html theme/_default/list.html theme/indexes/cs-engine.html theme/_default/indexes.html theme/section/cs-engine.html theme/_default/section.html theme/_default/list.html theme/indexes/cs-engine.html theme/_default/indexes.html theme/theme/section/cs-engine.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/cs-engine.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "cs-engine" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section notary: [section/notary.html _default/section.html _default/list.html indexes/notary.html _default/indexes.html theme/section/notary.html theme/_default/section.html theme/_default/list.html theme/indexes/notary.html theme/_default/indexes.html theme/section/notary.html theme/_default/section.html theme/_default/list.html theme/indexes/notary.html theme/_default/indexes.html theme/theme/section/notary.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/notary.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "notary" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section machine: [section/machine.html _default/section.html _default/list.html indexes/machine.html _default/indexes.html theme/section/machine.html theme/_default/section.html theme/_default/list.html theme/indexes/machine.html theme/_default/indexes.html theme/section/machine.html theme/_default/section.html theme/_default/list.html theme/indexes/machine.html theme/_default/indexes.html theme/theme/section/machine.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/machine.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "machine" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section docker-trusted-registry: [section/docker-trusted-registry.html _default/section.html _default/list.html indexes/docker-trusted-registry.html _default/indexes.html theme/section/docker-trusted-registry.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-trusted-registry.html theme/_default/indexes.html theme/section/docker-trusted-registry.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-trusted-registry.html theme/_default/indexes.html theme/theme/section/docker-trusted-registry.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-trusted-registry.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "docker-trusted-registry" is rendered empty
WARN: 2016/08/19 05:07:54 site.go:2014: Unable to locate layout for section docker-hub: [section/docker-hub.html _default/section.html _default/list.html indexes/docker-hub.html _default/indexes.html theme/section/docker-hub.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-hub.html theme/_default/indexes.html theme/section/docker-hub.html theme/_default/section.html theme/_default/list.html theme/indexes/docker-hub.html theme/_default/indexes.html theme/theme/section/docker-hub.html theme/theme/_default/section.html theme/theme/_default/list.html theme/theme/indexes/docker-hub.html theme/theme/_default/indexes.html]
WARN: 2016/08/19 05:07:54 site.go:1990: "docker-hub" is rendered empty
0 of 29 drafts rendered
0 future content
708 pages created
929 non-page files copied
0 paginator pages created
0 tags created
0 categories created
in 24686 ms

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
c082567 Merge pull request #272 from docker/use-https-where-possible
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
- PR 291 (d414ea9) Jul 14 07:50:25 from docker/skip-http-client-errors
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



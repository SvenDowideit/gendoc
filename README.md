
# initialize your workspace with

```
$ gendoc clone --all
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

## TODO:
### bring together as

```
$ gendoc render --set v1.12
```


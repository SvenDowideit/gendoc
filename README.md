

to generate docs into the `docs-html` dir, run a command set like

```
../gendoc/gendoc checkout v1.10 && ../gendoc/gendoc fetch && ../gendoc/gendoc render
```

Or to serve the docs to a browse


```
../gendoc/gendoc checkout master && ../gendoc/gendoc fetch && ../gendoc/gendoc serve
```

## fetch

will gather the source files specified in the `all-projects.yml` into the `docs-source/<publish-set>`
dir.

## serve / render

will use the files in the `docs-source/<publish-set>` dir to generate files in the `docs-html/<publish-set>`
dir

## bring together as

```
../gendoc/gendoc serve --set v1.12
```

or

```
../gendoc/gendoc render --checkout --fetch --set v1.12
```

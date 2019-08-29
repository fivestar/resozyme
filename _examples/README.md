# Examples

```console
$ go run main.go
```

- [/](https://github.com/fivestar/resozyme/blob/master/_examples/resources/index_resource.go)
- [/articles](https://github.com/fivestar/resozyme/blob/master/_examples/resources/articles_resource.go)
- [/articles/{articleID}](https://github.com/fivestar/resozyme/blob/master/_examples/resources/article_resource.go)

```bash
# GET index
curl -D - \
  -XGET 'http://localhost:9393/'

# GET articles
curl -D - \
  -XGET 'http://localhost:9393/articles'

# POST new article
curl -D - \
  -H'Content-Type: application/json' \
  -XPOST 'http://localhost:9393/articles' \
  -d'{"title":"Hello","pubDate":"2019-09-01"}'

# GET article
curl -D - \
  -XGET 'http://localhost:9393/articles/1'

# PATCH article
curl -D - \
  -H'Content-Type: application/json' \
  -XPATCH 'http://localhost:9393/articles/1' \
  -d'{"title":"Hello, World"}'
```

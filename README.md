# Go でバックエンド開発のリハビリ

Go で backend server を作りながら、感覚を戻したい

## Usage

```
$ go run main.go
```

```
$ curl http://localhost:3000
Hello world
```

```
# -- request --
$ curl "http://localhost:3000/?url_long=hoge&url_long=123"
Hello world

# ~~~~~~~~~~

# -- server --
map[url_long:[hoge 123]]
[path] /
[scheme] 
[url_long] [hoge 123]
  [key] url_long
  [val] [hoge 123]
```


# 参考

## `net/http` を使ってサーバを立てる

* [Golang初心者が`net/http`パッケージでWebサーバーをホストする流れを追う](https://zenn.dev/skonb/articles/0bad1d59371d09)
* [Go言語でhttpサーバーを立ち上げてHello Worldをする](https://qiita.com/taizo/items/bf1ec35a65ad5f608d45)

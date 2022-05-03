# Go でバックエンド開発のリハビリ

Go で backend server を作りながら、感覚を戻したい

## Usage

```
$ go build main.go && ./main
```

```
$ go run main.go
```

## `net/http`

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

## `Gin`

```
$ curl http://localhost:3000
{"User-Agent":"curl/7.54.0","message":"Hello world"}
```

# ライブラリ
* [`go.uber.org/zap`](https://pkg.go.dev/go.uber.org/zap)
  - 高速・構造化・レベリングが売りの Go ロギングライブラリらしい。今回はリクエストのログを保存するのに用いている
* [`github.com/gin-gonic/gin`](https://github.com/gin-gonic/gin)
  - サーバサイドのフレームワーク。ざっくり Go では echo か gin か、だと思っている

* そういえば、Go でサーバサイドを実装するときに `Mux` というのを見たことがあるな
  - [`github.com/gorilla/mux`](https://github.com/gorilla/mux) にあるように、ライブラリを使っていたのか
  - はたまた、`mux := http.NewServeMux()` [(link)](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a#gin%E3%81%AE%E5%9F%BA%E7%A4%8E) にあるように、慣習的に Mux という名前をつける (使う) のか..
  - 今度、当時のソースコードをあたってみよう


# 参考

## `net/http` を使ってサーバを立てる

* [Golang初心者が`net/http`パッケージでWebサーバーをホストする流れを追う](https://zenn.dev/skonb/articles/0bad1d59371d09)
* [Go言語でhttpサーバーを立ち上げてHello Worldをする](https://qiita.com/taizo/items/bf1ec35a65ad5f608d45)
* [3.2 Goで簡単なWebサーバを立てる](https://docs.kilvn.com/build-web-application-with-golang/ja/03.2.html)
    - このドキュメント、[CSRF対策](https://docs.kilvn.com/build-web-application-with-golang/ja/09.1.html) とか [socket プログラミング](https://docs.kilvn.com/build-web-application-with-golang/ja/08.1.html) の記事とかあって面白そう。今度ちゃんと見てみよう

## `Gin`

* [ginを最速でマスターしよう](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)
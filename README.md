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
$  curl http://localhost:3000/health
{"status":"healthy"}
```

```
$ curl http://localhost:3000/note/v1/test
{"status":"ok"}
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

* [`github.com/go-xorm/xorm`](https://github.com/go-xorm/xorm)
  - `This repository has been archived by the owner. It is now read-only.` らしい。
  - [gitea.com/xorm/xorm](https://gitea.com/xorm/xorm) に移動になったらしい ~~CI こけてるが..??~~
  - go の DB 周りとか ORM って、`gorm` が代表的なのかなあ
  - `xorm` の使い方とか、記事調べたけど、あんまり出てこない気がする


# 参考

## `net/http` を使ってサーバを立てる
* [Golang初心者が`net/http`パッケージでWebサーバーをホストする流れを追う](https://zenn.dev/skonb/articles/0bad1d59371d09)
* [Go言語でhttpサーバーを立ち上げてHello Worldをする](https://qiita.com/taizo/items/bf1ec35a65ad5f608d45)
* [3.2 Goで簡単なWebサーバを立てる](https://docs.kilvn.com/build-web-application-with-golang/ja/03.2.html)
  - このドキュメント、[CSRF対策](https://docs.kilvn.com/build-web-application-with-golang/ja/09.1.html) とか [socket プログラミング](https://docs.kilvn.com/build-web-application-with-golang/ja/08.1.html) の記事とかあって面白そう。今度ちゃんと見てみよう

## `Gin`
* [ginを最速でマスターしよう](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)

## `xorm`
* [kyokomi/goxorm_memo.md](https://gist.github.com/kyokomi/1ff63138dc1f8765cc9d)

* [GoでMySQLにアクセスしてみる](https://kazuhira-r.hatenablog.com/entry/2021/03/16/223253)


## `Docker` 関係
* [MySQL 5.7(Docker)環境構築メモ](https://qiita.com/KWS_0901/items/4f6eb6a50e5f77430d0a)
* [docker-compose exec なんとか](https://qiita.com/tsuboyataiki/items/90dbe94553d3dea39b19)

## mysql (client (5.7))
* [Macのhomebrewでmysqlクライアントのみインストールする](https://qiita.com/fault/items/b854d90ae6eef5aa3417)
* [MacのMySQLクライアント](https://qiita.com/takepan/items/e01dfd968faffbe2ebb2)

## curl
* [curlコマンドでPOSTする, 様々な形式別メモ](https://weblabo.oscasierra.net/curl-post/)

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

```
$ curl http://localhost:3000/note/v1/list
{"data":[]}
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

## DB 接続時のぬるぽ問題

DB に接続自体はできているっぽくて ↓

```
[xorm] [info]  2022/05/05 14:58:15.766312 [SQL] SELECT `COLUMN_NAME`, `IS_NULLABLE`, `COLUMN_DEFAULT`, `COLUMN_TYPE`, `COLUMN_KEY`, `EXTRA`,`COLUMN_COMMENT` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [sample_db note]
[xorm] [info]  2022/05/05 14:58:15.781734 [SQL] SELECT `INDEX_NAME`, `NON_UNIQUE`, `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? [sample_db note]
init Database OK
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /health                   --> main.main.func1 (4 handlers)
[GIN-debug] GET    /note/v1/test             --> github.com/sudachi0114/gin-rest-backend-trial/controller.NoteTest (4 handlers)
[GIN-debug] GET    /note/v1/list             --> github.com/sudachi0114/gin-rest-backend-trial/controller.NoteList (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :3000
```

(特にここ ↓)

```
init Database OK
```


エンドポイントにアクセスして、DB へのコネクションが発生すると (?)、ぬるぽで `500 (Internal Server Error)` が発生する問題

```
(*'-') < curl localhost:3000/note/v1/list

# ~~~~~~~~~~


2022/05/05 14:59:28 [Recovery] 2022/05/05 - 14:59:28 panic recovered:
GET /note/v1/list HTTP/1.1
Host: localhost:3000
Accept: */*
User-Agent: curl/7.54.0


runtime error: invalid memory address or nil pointer dereference
/Users/sudachi/.goenv/versions/1.15.5/src/runtime/panic.go:212 (0x104cc12)
        panicmem: panic(memoryError)
/Users/sudachi/.goenv/versions/1.15.5/src/runtime/signal_unix.go:742 (0x104ca92)
        sigpanic: panicmem()
/Users/sudachi/go/pkg/mod/github.com/go-xorm/xorm@v0.7.9/session.go:94 (0x142248a)
        (*Session).Init: session.ctx = session.engine.defaultContext
/Users/sudachi/go/pkg/mod/github.com/go-xorm/xorm@v0.7.9/engine.go:317 (0x1417a55)
        (*Engine).NewSession: session.Init()
/Users/sudachi/go/pkg/mod/github.com/go-xorm/xorm@v0.7.9/engine.go:652 (0x1417a21)
        (*Engine).Distinct: session := engine.NewSession()
/Users/sudachi/go/src/github.com/sudachi0114/gin-rest-backend-trial/service/note.go:9 (0x144ca44)
        NoteService.GetNoteList: err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&notes)
/Users/sudachi/go/src/github.com/sudachi0114/gin-rest-backend-trial/controller/note.go:18 (0x144cc45)
        NoteList: NotesList := noteService.GetNoteList()
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x139d01a)
        (*Context).Next: c.handlers[c.index](c)
/Users/sudachi/go/src/github.com/sudachi0114/gin-rest-backend-trial/middleware/noteMiddleware.go:19 (0x14748c4)
        RecordUaAndTime: c.Next()
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x139d01a)
        (*Context).Next: c.handlers[c.index](c)
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/recovery.go:99 (0x13aaa88)
        CustomRecoveryWithWriter.func1: c.Next()
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x139d01a)
        (*Context).Next: c.handlers[c.index](c)
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/logger.go:241 (0x13a9bc4)
        LoggerWithConfig.func1: c.Next()
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/context.go:168 (0x139d01a)
        (*Context).Next: c.handlers[c.index](c)
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/gin.go:555 (0x13a069d)
        (*Engine).handleHTTPRequest: c.Next()
/Users/sudachi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.7/gin.go:511 (0x139ff4a)
        (*Engine).ServeHTTP: engine.handleHTTPRequest(c)
/Users/sudachi/.goenv/versions/1.15.5/src/net/http/server.go:2843 (0x1256882)
        serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
/Users/sudachi/.goenv/versions/1.15.5/src/net/http/server.go:1925 (0x1252e4c)
        (*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
/Users/sudachi/.goenv/versions/1.15.5/src/runtime/asm_amd64.s:1374 (0x106bf20)
        goexit: BYTE    $0x90   // NOP

[GIN] 2022/05/05 - 14:59:28 | 500 |   13.771388ms |             ::1 | GET      "/note/v1/list"
```

[MySQL - invalid memory address or nil pointer dereference【Go】](https://tech-up.hatenablog.com/entry/2019/01/05/212630)

> 複数の関数で、DBコネクションを使い回す場合には、注意が必要です。
> dbだけでなく、sql.Open()の返り値に含まれる **errもpackageグローバルに定義して、**
> sql.Open()の返り値のアサインは「:=」ではなく、「=」を使用します。

ということらしい。結構、解決に手間取ったのでここにメモ。

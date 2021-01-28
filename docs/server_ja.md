# api_gen
## server_generator

- APIのRootingやControllerを階層から自動的に判定し、生成する。
- 自動生成されたControllerに対して実際の処理を書いていく。
- 基本的に生成物は既存のファイルを上書きする。(Controllerはすでにファイルが存在した場合は上書きしない。)

### 使い方
#### 生成方法
- ルートとなるディレクトリパスを指定して以下のファイルを生成する。
    - `*_controller_gen.go`: Controller
        - すでにファイルが存在する場合は、上書きしない。
    - `routes_gen.go`: Route定義
    - `bootstrap_gen.go`: 上記ふたつをまとめたもの
- 生成対象は、 `{HTTP_METHOD}Name{Request|Response}` といった命名をしたStructでなければならない。
     - 対応しているHTTPメソッドは、 `get` , `post` , `put` , `delete` , `patch`
     - GETメソッドでは、Requestの、json tagとquery tagの両方に同じ値を入れなくてはならない。
- 必ず、 Request と Response がそれぞれ対になるようになっていなければならない。
- パスルーティングをする場合。
    - ディレクトリは `_` 始まりしなければならない。
        - 例: `/service/:id/hogehoge` にしたい場合は `/service/_id/*.go` のようにする。
    - ファイルは `0_` 始まりしなければならない。
        - 例: `/service/:id` にしたい場合は `/service/0_id.go` のようにする。

```console
$ server_generator ./sample/
```

#### モックサーバ

モックサーバは生成対象としたディレクトリの直下に `cmd/mock/main.go` として生成される。モックサーバを起動するためにはビルドオプションで `-tags mock` を付ける必要がある。  
また、モックサーバが返すjsonは生成対象としたディレクトリ直下に生成される`mock_jsons`へルーティングと同じ階層構造で生成される。  
例として以下の構造で生成される。 
```text
interfaces/
├── bootstrap_gen.go
├── mock_bootstrap_gen.go
└─── mock_jsons
   └── post_create_user
       └── default.json
```

jsonの構造は以下のようになっている。
```javascript
{
    "meta": {
        "status": 200,           // レスポンス ステータスコード
        "match_request": { ... } // これに一致する場合、このjsonを返す。ただし、オプションでファイルが指定された場合はこの限りではない。
    },
    "payload": { ... }           // 実際に返すレスポンスのjson
}
```

返すjsonファイルは以下のように決定される。
1. ヘッダーでファイルが指定されているか
2. 一致するリクエストか
3. default.jsonを返す

default.jsonは他にマッチするものがないか指定されたファイルが見つからない場合に自動的に参照するため、削除することは非推奨。他のレスポンスを返すjsonを作成する際はdefault.jsonをベースに作成することを推奨する。  

モックサーバ起動手順
```shell script
go run -tags mock sample/cmd/mock/main.go
```

#### コードサンプル

middlewareは `map[endpoint]middleware` の形式で追加していく。指定したmiddlewareはendpoint以下のすべてに適用される。

[templates](../templates)にapi_genでプロジェクトを始める時のテンプレートを参照できる。

```go
e := echo.New()
// アクセスログを表示
e.Use(middleware.Logger())
// panic時に自動でrecoverする
e.Use(middleware.Recover())

m := make([]*MiddlewareSet, 0)
m = append(m, &MiddlewareSet{
	Path: "/service/user/",
	MiddlewareFunc: []echo.MiddlewareFunc{
		func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
			return func(context echo.Context) error {
				// /service/user/以下の全てのハンドラーにこの関数を適用
			}
		},
	},
})

// 全てのハンドラーを初期化する
service.Bootstrap(&props.ControllerProps{
    // ./sample/propsに定義したControllerPropsを初期化
}, e, m)

if err := e.Start(":" + PORT); err != nil {
	t.Fatalf("server listen error %s", err.Error())
}
```

# api_gen server

- APIのRoutingやControllerを階層から自動的に判定し、生成する。
- 自動生成されたControllerに対して実際の処理を書いていく。
- 基本的に生成物は既存のファイルを上書きする。(Controllerはすでにファイルが存在した場合は上書きしない。)

## 使い方
### 生成方法
#### api_genで利用するフォルダを作成する
```console
$ mkdir api_gen
$ cd api_gen
```

#### API定義を書くフォルダを用意する
```console
$ mkdir api
$ cd api
```

#### API定義を用意する
- apiフォルダか　らのフォルダのパスがURLのパスとなる。
- 生成対象は、 `{HTTP_METHOD}{EndpointName}{Request|Response}` といった命名をしたStructでなければならない。
    - 対応しているHTTPメソッドは、 `GET` , `POST` , `PUT` , `DELETE` , `PATCH`
    - GETメソッドでは、Requestの、json tagとquery tagの両方に同じ値を入れなくてはならない。
    - 必ず、 Request と Response がそれぞれ対になるようになっていなければならない。

- 例: GET /v2/userを作成したい場合
    - api_gen/api/v2にget_user.goを作成(ファイル名は任意)
    - GetUserRequestとGetUserResponseのstructを用意(struct名はsnake caseに変換される)
    - Requestにはquery、ResponseにはJSONレスポンスの内容を記述する

- パスルーティングをする場合。
    - ディレクトリは `_` 始まりしなければならない。
        - 例: `/service/:id/hogehoge` にしたい場合は `/service/_id/*.go` のようにする。
    - ファイルは `0_` 始まりしなければならない。
        - 例: `/service/:id` にしたい場合は `/service/0_id.go` のようにする。

- ルートとなるディレクトリパスを指定して以下のファイルを生成する。
    - `*_controller_gen.go`: Controller
        - すでにファイルが存在する場合は、上書きしない。
    - `routes_gen.go`: Route定義
    - `bootstrap_gen.go`: 上記ふたつをまとめたもの


#### api_genを実行する

```console
$ cd ../ # api_genフォルダ
$ api_gen server -o . ./api # '-o .' は省略可
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

指定したmiddlewareはendpoint以下のすべてに適用される。

```go
e := echo.New()
// アクセスログを表示
e.Use(middleware.Logger())
// panic時に自動でrecoverする
e.Use(middleware.Recover())

ctrl := controller.NewControllers(&props.ControllerProps{
    // ./api_gen/propsに定義したControllerPropsを初期化
}, e)

ctrl.AddMiddleware("/user", func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Do something

        return next(c)
    }
})

if err := e.Start(":" + PORT); err != nil {
	t.Fatalf("server listen error %s", err.Error())
}
```

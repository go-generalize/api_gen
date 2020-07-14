# api_gen
## server_generator

- APIのRootingやControllerを階層から自動的に判定し、生成する。
- 自動生成されたControllerに対して実際の処理を書いていく。
- 基本的に生成物は既存のファイルを上書きする。(Controllerはすでにファイルが存在した場合は上書きしない。)

### 使い方
#### 生成方法
- ルートとなるディレクトリパスを指定して以下のファイルを生成する。
    - Controller
        - すでにファイルが存在する場合は、上書きしない。
    - Route定義
    - 上記ふたつをまとめたもの
- 生成対象は、 `{HTTP_METHOD}Name{Request|Response}` といった命名をしたStructでなければならない。
     - 対応しているHTTPメソッドは、 `get` , `post` , `put` , `delete` , `patch`
     - GETメソッドでは、Requestの、json tagとquery tagの両方に同じ値を入れなくてはならない。
- 必ず、 Request と Response がそれぞれ対になるようになっていなければならない。
- **`json:"-"` のタグを指定すると、内部で利用してる struct2ts がClassの生成対象から除外するため、注意。**
- パスルーティングをする場合。
    - ディレクトリは `_` 始まりしなければならない。
        - 例: `/service/:id/hogehoge` にしたい場合は `/service/_id/*.go` のようにする。
    - ファイルは `0_` 始まりしなければならない。
        - 例: `/service/:id` にしたい場合は `/service/0_id.go` のようにする。

```console
$ server_generator ./sample/
```

#### コードサンプル

middlewareは `map[endpoint]middleware` の形式で追加していく。指定したmiddlewareはendpoint以下のすべてに適用される。

[templates](../templates)にapi_genでプロジェクトを始める時のテンプレートを参照できます。

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
service.Bootstrap(e, m)

if err := e.Start(":" + PORT); err != nil {
	t.Fatalf("server listen error %s", err.Error())
}
```

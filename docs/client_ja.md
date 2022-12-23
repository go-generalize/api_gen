# api_gen client

```
Usage:
  api_gen client [command]

Available Commands:
  dart        Generate Dart client library
  go          Generate go client library
  typescript  Generate TypeScript client library
```

## Go

Goのクライアントライブラリを生成する。
```
Usage:
  api_gen client go [path] [flags]

Flags:
  -h, --help             help for go
  -o, --output string    The directory to generated client library in (default "./")
  -p, --package string   The package name of the generated library (default "client")
```

例
```
$ api_gen client go -o ./client/go ./interfaces
```

サンプルコード
```go
package main

import (
	"fmt"
	"net/http"

	client "example-code/frontend/go"
	_api_v1_school "example-code/frontend/go/classes/api/v1/school"
)

func main() {
	api_client := client.NewClient(http.Client{}, "http://localhost:8080")
	req := _api_v1_school.PostUpdateUserRequest{ID: "4423"}

	res, err := api_client.Api.V1.School.PostUpdateUser(&req)
	if err != nil {
		fmt.Println("Error occurred: %w", err)
		return
	}

	fmt.Println(*res)
}
```

## Dart

Dartのクライアントライブラリを生成する。
```
Usage:
  api_gen client dart [path] [flags]

Aliases:
  dart, flutter

Flags:
  -o, --file string   The directory to generate client library in (default "./")
  -h, --help          help for dart
```

例
```
$  api_gen client go -o ./client/dart ./interfaces
```

サンプルコード
```dart
import 'dart:io';

import "./api_client.dart";
import 'package:http/http.dart' as http;
import './classes/api/v1/school/types.dart' as types__api_v_1_school;

main() async {
    final http_client = http.Client();
    final api_client = ApiClient(
      'http://localhost:8080',
      {
        "Content-Type": "application/json"
      },
      http_client
    );

    final request = api_client.getRequestObject({}, [], true);
    final feature_response = api_client.v1.school.postUpdateUser(
      types__api_v_1_school.PostUpdateUserRequest(id: "lain")
    );

    print(await feature_response);
}
```

## TypeScript

Typescript+fetchを利用したライブラリが生成される。  

### 開発にあたって

fetchに非対応のブラウザに対応するには[Polyfill](https://github.com/github/fetch)を使って対応する。    
eslintなどを導入する場合は、自動生成をignoreに追加することを推奨する。

### モックサーバ

`api_gen server`ではモックサーバをサポートしています。モックデータの返し方の詳細に関しては`api_gen server`のドキュメントを参照。  
モックサーバへのリクエストは通常のapi_gen clientが生成したクライアントから行える。ただし、いくつかのオプションが存在しており、これらは `Api-Gen-Option` ヘッダーでJSONエンコードされモックサーバへ送信される。  
オプションの指定方法はclient-sideのapi_genの生成したコードのオプション引数にて `'mock_option'` をキーとしたオブジェクトを渡す必要がある。  
オプションの詳細は以下のとおり。
```javascript
{
    wait_ms: 10,           // モックサーバからの応答を指定したミリ秒遅延させる。 (例では10ms)
    target_file: 'error'   // モックサーバが参照するjsonファイルを固定する。拡張子のjsonは省略することが可能。 (例ではerror.json)
}
```

### 生成について

[templates](../templates) を利用した場合、[templates/frontend](../templates/frontend)にて `make generate`を実行することで簡単に生成できる。

生成されるファイルは以下のファイルである。
- api/
    - api_client.ts
    - classes/
        - types.ts

以下が実装例です。

```typescript
import { APIClient, MockOption } from "./api/api_client";

// the simplest
(async () => {
    const client = new APIClient();

    const resp = await client.postCreateUser({
        ID: "id",
        Password: "password",
        Gender: 0,
    });

    console.log(resp);
})();

// with options
(async () => {
    const client = new APIClient(
        "pass", // [optional] token for Authorization: Bearer
        {
            "X-Foobar": "foobar", // [optional] custom headers
        },
        "http://localhost:8080",  // [optional] custom endpoint
        {
            cache: "no-cache", // [optional] options for fetch API
        },
    );

    const resp = await client.postCreateUser(
        {
            ID: "id",
            Password: "password",
            Gender: 0,
        },
        {
            "X-Foobar": "hoge", // [optional] custom headers
        },
        {
            mode: "cors" // [optional] options for fetch API 
        },
    );

    console.log(resp);
})();

// mock mode
(async () => {
    const client = new APIClient();
    const mockOption: MockOption = {
        wait_ms: 1000,
        target_file: 'error'
    }

    const resp = await client.postCreateUser({
        ID: "id",
        Password: "password",
        Gender: 0,
    }, undefined, {
        'mock_option': mockOption,
    });

    console.log(resp);
})();
```

### FAQ
#### Cookieを送信したい
- `{credentials: "include"}` をfetch APIのオプションに設定
    - FYI: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API

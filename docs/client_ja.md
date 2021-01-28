# api_gen
## client_generator

server_generator向けの用意したフォルダに対して実行する。
ライブラリは今いるディレクトリに対して生成される。

Typescript+fetchを利用したライブラリが生成される。  

### 開発にあたって

fetchに非対応のブラウザに対応するには[Polyfill](https://github.com/github/fetch)を使って対応する。    
eslintなどを導入する場合は、自動生成をignoreに追加することを推奨する。

### モックサーバ

server_generatorではモックサーバをサポートしています。モックデータの返し方の詳細に関してはserver_generatorのドキュメントを参照。  
モックサーバへのリクエストは通常のclient_generatorが生成したクライアントから行える。ただし、いくつかのオプションが存在しており、これらは `Api-Gen-Option` ヘッダーでJSONエンコードされモックサーバへ送信される。  
オプションの指定方法はclient_generatorの生成したコードのオプション引数にて `'mock_option'` をキーとしたオブジェクトを渡す必要がある。  
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

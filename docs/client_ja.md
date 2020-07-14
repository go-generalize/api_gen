# api_gen
## client_generator

server_generator向けの用意したフォルダに対して実行する。
ライブラリは今いるディレクトリに対して生成される。

Typescript+fetchを利用したライブラリが生成される。  
**`json:"-"` のタグを指定すると、内部で利用してる struct2ts がClassの生成対象から除外するため、注意。**

### 開発にあたって

fetchに非対応のブラウザに対応するには[Polyfill](https://github.com/github/fetch)を使って対応する。    
eslintなどを導入する場合は、自動生成をignoreに追加することを推奨する。

### 生成について

[templates](../templates) を利用した場合、[templates/frontend](../templates/frontend)にて `make generate`を実行することで簡単に生成できる。

生成されるファイルは以下のファイルである。
- api/
    - api_client.ts
    - classes/
        - {HTTP Method}{Endpoint Name}{Request/Response}.ts

以下が実装例です。

```typescript
import { PostCreateUserRequest } from "./api/classes/PostCreateUserRequest";
import { APIClient } from "./api/api_client";

// the simplest
(async () => {
    const client = new APIClient();

    const resp = await client.postCreateUser(new PostCreateUserRequest({
        ID: "id",
        Password: "password",
        Gender: 0,
    }));

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
        new PostCreateUserRequest({
            ID: "id",
            Password: "password",
            Gender: 0,
        }),
        {
            "X-Foobar": "hoge", // [optional] custom headers
        },
        {
            mode: "cors" // [optional] options for fetch API 
        },
    );

    console.log(resp);
})();
```

### FAQ
#### Cookieを送信したい
- `{credentials: "include"}` をfetch APIのオプションに設定
    - FYI: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
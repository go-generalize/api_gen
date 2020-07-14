# api_gen
## client_generator

Execute client_generator for the same root directory as server_generator.
The library will generated in the current working directory.

The library uses TypeScript+fetch API
**CAUTION: If you specify the tag of `json: "-"`, struct2ts used internally is excluded from the class generation target.**

### About develop

To support browsers that do not support fetch, use [Polyfill](https://github.com/github/fetch).
When introducing eslint etc., it is recommended to add automatic generation to ignore.

### About generate

In [templates,](../templates) the library can be generated with `make generate` in [templates/frontend.](../templates/frontend)

The following files are generated:
- api/
    - api_client.ts
    - classes/
        - {HTTP Method}{Endpoint Name}{Request/Response}.ts


Example implementation:
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
- Use `{credentials: "include"}` to send Cookie
    - FYI: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
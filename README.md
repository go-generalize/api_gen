# api_gen
## What's api_gen?
api_gen accelerates the development of API servers.  
It generates a boilerplate for your API server in Go and clients in Go, TypeScript and so on.

Server: Go  
Client: Go, TypeScript, Dart

## Prerequisite
Go 1.16 or newer is required to use mock servers.

## Installation
### from Releases
Download binaries from [Releases](https://github.com/go-generalize/api_gen/releases/latest).

### go install
Go 1.17 or newer is required.
```console
$ go install github.com/go-generalize/api_gen/v2/cmd/api_gen@latest # latest or vX.Y.Z
```

## Basic Usage
Prepare API definitions in Go.  
The path of directories are mapped to the path of URLs.

Define types for both of requests and responses.  
All types must consist of three parts: method, endpoint name, request or response.  
The below example is types for `POST /foo/bar/update_user`.

Requests are bound with the Binder in [echo](https://echo.labstack.com).  
Please refer to [Binding](https://echo.labstack.com/guide/binding/). 
Responses are encoded in JSON.

```go
// ./interfaces/foo/bar/hoge.go
package package_name_that_you_want

type PostUpdateUserRequest struct {
    ID string `json:"id"`
}

type PostUpdateUserResponse struct {
    Result int `json:"result"`
}
```

Run the below command to generate controllers.
```console
$ api_gen server -o ./server ./interfaces
```
Controllers are generated into `./interfaces/controller/foo/bar/post_update_user.go`.

Then, set up echo server.

```go
// main.go
package main

import (
    controller "path/to/generated/server"
    "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.NewControllers(nil, e)

	panic(e.Start(":8080"))
}
```

Run the below command to generate clients in TypeScript.
```console
$ api_gen client ts -o ./client/ts ./interfaces
```

API clients are generated into `./client/ts/api_client.ts`.

```typescript
import { APIClient } from 'path/to/api_client.ts';

const client = new APIClient('very_secure_token', {}, 'https://example.com');

client.foo.bar.postUpdateUser({id: 'id'})
    .then(res => console.log(res));
```

## For more details
### server-side
[ENGLISH](./docs/server_en.md)/[日本語](./docs/server_ja.md)

### client-side
[ENGLISH](./docs/client_en.md)/[日本語](./docs/client_ja.md)

## License
- Under the MIT License
- Copyright (c) 2020-2021 go-generalize

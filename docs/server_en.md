# api_gen server

- Generate routers/controllers from the minimal endpoint definition
- You can write controllers in the generated `controller/**/*.go` files
- CAUTION: All files other than `*_controller_gen.go` will be overwritten.

## Usage
### How to execute
#### Create a directory for api_gen
```console
$ mkdir api_gen
$ cd api_gen
```

#### Create a directory for API definitions
- The relative path from `api/` directory will be the path in URL
- Each endpoint consists of two structs: `*Request`, `*Response`
    - The name should be `{HTTP_METHOD}Name{Request|Response}`
        - Available methods: `get` , `post` , `put` , `delete` , `patch`
    - If `json` tags are specified for `GET` endpoints,  specify also `query` tags with the same name
    - Both `*Request` and `*Response` structs are necessary.

- Ex: GET /v2/user
    - Create a `.go` with an arbitrary name file in api_gen/api/v2
    - Prepare `GetUserRequest` and `GetUserResponse`  struct
        - struct names are converted into snake case

- For path parameters
    - The directory must start with `_`.
        - Example: If you want to use `/service/:id/hogehoge`, use `/service/_id/*.go`.
    - The file must start with `0_`.
        - Example: If you want to use `/service/:id`, use `/service/0_id.go`.

- By specifing the root directory(api_gen), following files will be generated:
    - `controller/**/*.go`: Controller: Users will directly edit them
        - Existing files won't be overwritten.
    - `mock/`: For mock
        - `controller/`: Controllers for mock: Users won't edit them
        - `json/`: JSON to be returned from mock server
    - `controller_initializer.go`: Has NewControllers binds controllers to *echo.Echo
        - Import api_gen
    - `bundler.go`: Only internal usage
    - `mock_bundler.go`: Only internal usage for mock

- Execute `api_gen server` with the path to the root directory of controllers
- This generates the following files:
    - Controller(`*_controller_gen.go`)
        - Not overwritten if it already exists
        - For all endpoints
    - Route defitinion(`routes_gen.go`)
        - All directories
    - echo.Echo initializer(`bootstrap_gen.go`)

#### api_genを実行する

```console
$ cd ../ # api_gen directory
$ api_gen server -o . ./api # '-o .' can be omitted
```

#### Mock server

In order to start the mock server, you should add the build option `-tags mock`.  
Moreover, the json returned by the mock server is generated in the same hierarchical structure as the routing to `mock_jsons` which is generated directly under the directory to be generated.  
As an example, json is generated with the following structure. 
```text
interfaces/
├── bootstrap_gen.go
├── mock_bootstrap_gen.go
└─── mock_jsons
   └── post_create_user
       └── default.json
```

The structure of json is as follows.
```javascript
{
    "meta": {
        "status": 200,           // Response http status code
        "match_request": { ... } // If it matches, this json is returned. If the file is specified as an option, however, this is not required.
    },
    "payload": { ... }           // The actual response json to be returned
}
```

The json file to be returned is determined by the following rules.
1. whether the file is specified in the header
2. matching requests ("match_request")
3. return default.json

Because default.json is automatically referenced when no other matches or the specified file is not found, it is not recommended to delete it. It is recommended to create json files that return other responses based on default.json.  

Mock server startup procedure.
```shell script
go run -tags mock sample/cmd/mock/main.go
```


### Example

Middlewares will be applied to 

指定したmiddlewareはendpoint以下のすべてに適用される。

```go
e := echo.New()
// Show access logs
e.Use(middleware.Logger())
// Automatically panic if panicked
e.Use(middleware.Recover())

ctrl := controller.NewControllers(&props.ControllerProps{
    // Initialize ControllerProps in ./api_gen/props
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

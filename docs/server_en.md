# api_gen
## server_generator

- Generate routers/controllers from the minimal endpoint definition
- You can write controllers in the generated `.go` files(`*_controller_gen.go`)
- CAUTION: All files other than `*_controller_gen.go` will be overwritten.

### Usage
#### How to execute
- Execute `server_generator` with the path to the root directory of controllers
- This generates the following files:
    - Controller(`*_controller_gen.go`)
        - Not overwritten if it already exists
        - For all endpoints
    - Route defitinion(`routes_gen.go`)
        - All directories
    - echo.Echo initializer(bootstrap.go)
- Each endpoint consists of two structs: `*Request`, `*Response`
    - The name should be `{HTTP_METHOD}Name{Request|Response}`
        - Available methods: `get` , `post` , `put` , `delete` , `patch`
    - If `json` tags are specified for `GET` endpoints,  specify also `query` tags with the same name
    - Both `*Request` and `*Response` structs are necessary.
- For path routing.
    - The directory must start with `_`.
        - Example: If you want to use `/service/:id/hogehoge`, use `/service/_id/*.go`.
    - The file must start with `0_`.
        - Example: If you want to use `/service/:id`, use `/service/0_id.go`.

```console
$ server_generator ./sample/
```


#### Mock server

The mock server is created as `cmd/mock/main.go` under the target directory. In order to start the mock server, you should add the build option `-tags mock`.  
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

jsonの構造は以下のようになっている。
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


#### Example

You can define middlewares in `map[endpoint]middleware`.
Middlewares will be applied to all endpoints under the specified path

To start a new project with api_gen, a boilerplate is available at [templates](../templates).

```go
e := echo.New()
// Show access logs
e.Use(middleware.Logger())
// Recover from panic()
e.Use(middleware.Recover())

m := make([]*MiddlewareSet, 0)
m = append(m, &MiddlewareSet{
	Path: "/service/user/",
	MiddlewareFunc: []echo.MiddlewareFunc{
		func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
			return func(context echo.Context) error {
				// Apply this function as a middleware
				//     for all paths under /service/user/
			}
		},
	},
})

// Initialize all handlers
service.Bootstrap(&props.ControllerProps{
    // Initialize ControllerProps defined in ./sample/props
}, e, m)

if err := e.Start(":" + PORT); err != nil {
	t.Fatalf("server listen error %s", err.Error())
}
```

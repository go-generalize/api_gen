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
- **CAUTION: If you specify the tag of `json: "-"`, struct2ts used internally is excluded from the class generation target.**
- For path routing.
    - The directory must start with `_`.
        - Example: If you want to use `/service/:id/hogehoge`, use `/service/_id/*.go`.
    - The file must start with `0_`.
        - Example: If you want to use `/service/:id`, use `/service/0_id.go`.

```console
$ server_generator ./sample/
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

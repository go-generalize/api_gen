package main

import (
	"fmt"
	"net/http"

	client "github.com/go-generalize/api_gen/pkg/client/go/sample"
	"github.com/go-generalize/api_gen/server_generator/sample/service/user2"
)

func main() {
	cl := client.NewClient(http.Client{}, "http://localhost:8080")

	resp, err := cl.Service.User2.GetUser(&user2.GetUserRequest{
		ID:            "id",
		SearchRequest: "search request",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

package main

import (
	"fmt"
	"net/http"
	"net/textproto"
	"strings"

	client "github.com/go-generalize/api_gen/v2/samples/standard/clients/go"
	types "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes"
	multipartutil "github.com/go-generalize/multipart-util"
)

func main() {
	c := client.NewClient(*http.DefaultClient, "http://localhost:8080")

	req := &types.PostCreateTableRequest{
		ID:   "id",
		Flag: 100,
	}

	req.File = &multipartutil.MultipartPayload{
		Filename: "aaaa",
		Data:     strings.NewReader("aaaaaaaaaaaaaaa"),
		MIMEHeader: textproto.MIMEHeader{
			"Content-Type": {"text/plain"},
		},
	}

	req.Files = []*multipartutil.MultipartPayload{
		{
			Filename: "aaaa",
			Data:     strings.NewReader("aaaaaaaaaaaaaaa"),
			MIMEHeader: textproto.MIMEHeader{
				"Content-Type": {"text/plain"},
			},
		},
		{
			Filename: "bbbb",
			Data:     strings.NewReader("{}"),
			MIMEHeader: textproto.MIMEHeader{
				"Content-Type": {"application/json"},
			},
		},
	}

	resp, err := c.PostCreateTable(req)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

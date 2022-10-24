package test

import (
	"context"
	"net/http"
	"net/textproto"
	"strings"
	"testing"
	"time"

	multipartutil "github.com/go-generalize/multipart-util"
	client "github.com/go-gneralize/api_gen/v2/e2e/multipart/clients/go"
	types "github.com/go-gneralize/api_gen/v2/e2e/multipart/clients/go/classes"
	ptypes "github.com/go-gneralize/api_gen/v2/e2e/multipart/clients/go/classes/_param"
)

func postA(t *testing.T, client *client.APIClient) {
	t.Helper()

	payload := &types.PostARequest{
		File: &multipartutil.MultipartPayload{
			Data: strings.NewReader("1"),
			MIMEHeader: textproto.MIMEHeader{
				"Content-Type": []string{"text/plain"},
			},
		},
		Files: []*multipartutil.MultipartPayload{
			{
				Filename: "0.txt",
				Data:     strings.NewReader("0"),
				MIMEHeader: textproto.MIMEHeader{
					"Content-Type": []string{"text/plain"},
				},
			},
			{
				Filename: "1.txt",
				Data:     strings.NewReader("1"),
				MIMEHeader: textproto.MIMEHeader{
					"Content-Type": []string{"text/plain"},
				},
			},
		},
		Payload: "payload",
	}

	result, err := client.PostA(context.Background(), payload)

	if err != nil {
		t.Fatal(err)
	}

	if result.Message != "OK" {
		t.Fatalf("Unexpected message: %s", result.Message)
	}
}

func paramPostB(t *testing.T, client *client.APIClient) {
	t.Helper()

	payload := &ptypes.PostBRequest{
		File: &multipartutil.MultipartPayload{
			Data: strings.NewReader("1"),
			MIMEHeader: textproto.MIMEHeader{
				"Content-Type": []string{"text/plain"},
			},
		},
		Files: []*multipartutil.MultipartPayload{
			{
				Filename: "0.txt",
				Data:     strings.NewReader("0"),
				MIMEHeader: textproto.MIMEHeader{
					"Content-Type": []string{"text/plain"},
				},
			},
			{
				Filename: "1.txt",
				Data:     strings.NewReader("1"),
				MIMEHeader: textproto.MIMEHeader{
					"Content-Type": []string{"text/plain"},
				},
			},
		},
		Payload: "payload",
		Param:   "param",
	}

	result, err := client.Param.PostB(context.Background(), payload)

	if err != nil {
		t.Fatal(err)
	}

	if result.Message != "OK" {
		t.Fatalf("Unexpected message: %s", result.Message)
	}
}

func TestGoClient(t *testing.T, addr string) {
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	client := client.NewClient(c, addr)

	postA(t, client)
	paramPostB(t, client)
}

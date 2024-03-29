// Code generated by api_gen. DO NOT EDIT.
// Package client Don't fix this code by your own hands.
// api_gen version: (devel)

package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"

	"github.com/go-generalize/multipart-util"
	"github.com/go-utils/echo-multipart-binder/mjbinder"

	root "github.com/go-gneralize/api_gen/v2/e2e/multipart/clients/go/classes"
	__param "github.com/go-gneralize/api_gen/v2/e2e/multipart/clients/go/classes/_param"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type Group__param struct {
	apiClient *APIClient
}

func newGroup__param(client *APIClient) *Group__param {
	return &Group__param{
		apiClient: client,
	}
}

func (g *Group__param) PostB(ctx context.Context, reqPayload *__param.PostBRequest) (respPayload *__param.PostBResponse, retErr error) {
	br, bw := io.Pipe()
	buffered := bufio.NewWriter(bw)
	mw := multipart.NewWriter(buffered)

	finished := make(chan bool)
	var chanError error
	defer func() {
		<-finished
		if chanError != nil {
			retErr = fmt.Errorf("creating payload error is %+v: %w", chanError, retErr)
		}
	}()

	go func() {
		chanError = func() error {
			defer close(finished)
			defer bw.Close()
			defer buffered.Flush()
			defer mw.Close()

			addField := func(fieldName string, payload *multipartutil.MultipartPayload) error {
				if payload == nil {
					return nil
				}

				if payload.Filename == "" {
					payload.Filename = "untitled"
				}

				header := payload.MIMEHeader

				if header == nil {
					header = textproto.MIMEHeader{}
				}

				header.Set("Content-Disposition",
					fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
						escapeQuotes(fieldName), escapeQuotes(payload.Filename)))
				if header.Get("Content-Type") != "" {
					header.Set("Content-Type", "application/octet-stream")
				}

				w, err := mw.CreatePart(header)

				if err != nil {
					return fmt.Errorf("failed to create a part for %s: %w", fieldName, err)
				}

				_, err = io.Copy(w, payload.Data)

				if err != nil {
					return fmt.Errorf("failed to copy data for %s: %w", fieldName, err)
				}

				return nil
			}
			addField(`file`, reqPayload.File)
			for _, f := range reqPayload.Files {
				addField(`files`, f)
			}

			w, err := mw.CreatePart(mjbinder.CreateJSONRequestMIMEHeader())

			if err != nil {
				return fmt.Errorf("failed to create a part for JSON payload: %w", err)
			}

			err = json.NewEncoder(w).Encode(reqPayload)

			if err != nil {
				return fmt.Errorf("failed to encode JSON payload: %w", err)
			}

			return nil
		}()
	}()

	req, err := http.NewRequest("POST", g.apiClient.base+"/"+fmt.Sprint(reqPayload.Param)+"/b", br)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mw.FormDataContentType())

	req = req.WithContext(ctx)

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		b, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("status code is %d: %w", resp.StatusCode, err)
		}

		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Data:       b,
		}
	}

	respPayload = &__param.PostBResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group struct {
	Param     *Group__param
	apiClient *APIClient
}

func newGroup(client *APIClient) *Group {
	return &Group{
		apiClient: client,
		Param:     newGroup__param(client),
	}
}

func (g *Group) PostA(ctx context.Context, reqPayload *root.PostARequest) (respPayload *root.PostAResponse, retErr error) {
	br, bw := io.Pipe()
	buffered := bufio.NewWriter(bw)
	mw := multipart.NewWriter(buffered)

	finished := make(chan bool)
	var chanError error
	defer func() {
		<-finished
		if chanError != nil {
			retErr = fmt.Errorf("creating payload error is %+v: %w", chanError, retErr)
		}
	}()

	go func() {
		chanError = func() error {
			defer close(finished)
			defer bw.Close()
			defer buffered.Flush()
			defer mw.Close()

			addField := func(fieldName string, payload *multipartutil.MultipartPayload) error {
				if payload == nil {
					return nil
				}

				if payload.Filename == "" {
					payload.Filename = "untitled"
				}

				header := payload.MIMEHeader

				if header == nil {
					header = textproto.MIMEHeader{}
				}

				header.Set("Content-Disposition",
					fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
						escapeQuotes(fieldName), escapeQuotes(payload.Filename)))
				if header.Get("Content-Type") != "" {
					header.Set("Content-Type", "application/octet-stream")
				}

				w, err := mw.CreatePart(header)

				if err != nil {
					return fmt.Errorf("failed to create a part for %s: %w", fieldName, err)
				}

				_, err = io.Copy(w, payload.Data)

				if err != nil {
					return fmt.Errorf("failed to copy data for %s: %w", fieldName, err)
				}

				return nil
			}
			addField(`file`, reqPayload.File)
			for _, f := range reqPayload.Files {
				addField(`files`, f)
			}

			w, err := mw.CreatePart(mjbinder.CreateJSONRequestMIMEHeader())

			if err != nil {
				return fmt.Errorf("failed to create a part for JSON payload: %w", err)
			}

			err = json.NewEncoder(w).Encode(reqPayload)

			if err != nil {
				return fmt.Errorf("failed to encode JSON payload: %w", err)
			}

			return nil
		}()
	}()

	req, err := http.NewRequest("POST", g.apiClient.base+"/a", br)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mw.FormDataContentType())

	req = req.WithContext(ctx)

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		b, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("status code is %d: %w", resp.StatusCode, err)
		}

		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Data:       b,
		}
	}

	respPayload = &root.PostAResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func encodeQuery(v interface{}) (string, error) {
	buf := bytes.NewBuffer(nil)

	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return "", err
	}

	dict := map[string]interface{}{}
	if err := json.NewDecoder(buf).Decode(&dict); err != nil {
		return "", err
	}

	val := url.Values{}
	for k, v := range dict {
		if v == nil {
			continue
		}
		val.Set(k, fmt.Sprint(v))
	}

	return val.Encode(), nil
}

type APIClient struct {
	*Group

	client http.Client
	base   string
}

func NewClient(client http.Client, base string) *APIClient {
	c := &APIClient{
		client: client,
		base:   base,
	}
	if len(c.base) != 0 && c.base[len(c.base)-1] == '/' {
		c.base = c.base[:len(c.base)-1]
	}

	c.Group = newGroup(c)

	return c
}

type APIError struct {
	StatusCode int
	Data       []byte
}

func (e APIError) Error() string {
	return fmt.Sprintf("the server returned an error: %d", e.StatusCode)
}

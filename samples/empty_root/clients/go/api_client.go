// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// api_gen version: (devel)
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	_foo_bar "github.com/go-generalize/api_gen/v2/samples/empty_root/clients/go/classes/foo/bar"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type Group_foo_bar struct {
	apiClient *APIClient
}

func newGroup_foo_bar(client *APIClient) *Group_foo_bar {
	return &Group_foo_bar{
		apiClient: client,
	}
}

func (g *Group_foo_bar) PostUser(reqPayload *_foo_bar.PostUserRequest) (respPayload *_foo_bar.PostUserResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/foo/bar/user", buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

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

	respPayload = &_foo_bar.PostUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_foo struct {
	Bar       *Group_foo_bar
	apiClient *APIClient
}

func newGroup_foo(client *APIClient) *Group_foo {
	return &Group_foo{
		apiClient: client,
		Bar:       newGroup_foo_bar(client),
	}
}

type Group struct {
	Foo       *Group_foo
	apiClient *APIClient
}

func newGroup(client *APIClient) *Group {
	return &Group{
		apiClient: client,
		Foo:       newGroup_foo(client),
	}
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

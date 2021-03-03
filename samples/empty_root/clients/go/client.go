// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// api_gen version: devel
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_foo "github.com/go-generalize/api_gen/samples/empty_root/server/foo"
)

type Group_foo struct {
	apiClient *APIClient
}

func newGroup_foo(client *APIClient) *Group_foo {
	return &Group_foo{
		apiClient: client,
	}
}

func (g *Group_foo) PostUser(reqPayload *_foo.PostUserRequest) (respPayload *_foo.PostUserResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/foo/user", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_foo.PostUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
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
	if c.base[len(c.base)-1] == '/' {
		c.base = c.base[:len(c.base)-1]
	}

	c.Group = newGroup(c)

	return c
}

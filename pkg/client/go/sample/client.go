// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// api_gen version: 1.0
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	root "github.com/go-generalize/api_gen/pkg/client/go/sample/classes"
	_service "github.com/go-generalize/api_gen/pkg/client/go/sample/classes/service"
	_service_static_page "github.com/go-generalize/api_gen/pkg/client/go/sample/classes/service/static_page"
	_service_user "github.com/go-generalize/api_gen/pkg/client/go/sample/classes/service/user"
	_service_user2 "github.com/go-generalize/api_gen/pkg/client/go/sample/classes/service/user2"
	_service_user2__userID "github.com/go-generalize/api_gen/pkg/client/go/sample/classes/service/user2/_userID"
	_service_user2__userID__JobID "github.com/go-generalize/api_gen/pkg/client/go/sample/classes/service/user2/_userID/_JobID"
)

type Group_service_static_page struct {
	apiClient *APIClient
}

func newGroup_service_static_page(client *APIClient) *Group_service_static_page {
	return &Group_service_static_page{
		apiClient: client,
	}
}

func (g *Group_service_static_page) GetStaticPage(reqPayload *_service_static_page.GetStaticPageRequest) (respPayload *_service_static_page.GetStaticPageResponse, err error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/service/static_page/static_page"+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_static_page.GetStaticPageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service_user struct {
	apiClient *APIClient
}

func newGroup_service_user(client *APIClient) *Group_service_user {
	return &Group_service_user{
		apiClient: client,
	}
}

func (g *Group_service_user) Get(reqPayload *_service_user.GetRequest) (respPayload *_service_user.GetResponse, err error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/service/user/"+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user.GetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user) PostUpdateUserName(reqPayload *_service_user.PostUpdateUserNameRequest) (respPayload *_service_user.PostUpdateUserNameResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user/update_user_name", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user.PostUpdateUserNameResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user) PostUpdateUserPassword(reqPayload *_service_user.PostUpdateUserPasswordRequest) (respPayload *_service_user.PostUpdateUserPasswordResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user/update_user_password", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user.PostUpdateUserPasswordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service_user2__userID__JobID struct {
	apiClient *APIClient
}

func newGroup_service_user2__userID__JobID(client *APIClient) *Group_service_user2__userID__JobID {
	return &Group_service_user2__userID__JobID{
		apiClient: client,
	}
}

func (g *Group_service_user2__userID__JobID) PutJob(reqPayload *_service_user2__userID__JobID.PutJobRequest) (respPayload *_service_user2__userID__JobID.PutJobResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", g.apiClient.base+"/service/user2/"+fmt.Sprint(reqPayload.UserID)+"/"+fmt.Sprint(reqPayload.JobID)+"/job", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user2__userID__JobID.PutJobResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service_user2__userID struct {
	JobID     *Group_service_user2__userID__JobID
	apiClient *APIClient
}

func newGroup_service_user2__userID(client *APIClient) *Group_service_user2__userID {
	return &Group_service_user2__userID{
		apiClient: client,
		JobID:     newGroup_service_user2__userID__JobID(client),
	}
}

func (g *Group_service_user2__userID) GetUserJobGet(reqPayload *_service_user2__userID.GetUserJobGetRequest) (respPayload *_service_user2__userID.GetUserJobGetResponse, err error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/service/user2/"+fmt.Sprint(reqPayload.UserID)+"/user_job_get"+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user2__userID.GetUserJobGetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service_user2 struct {
	UserID    *Group_service_user2__userID
	apiClient *APIClient
}

func newGroup_service_user2(client *APIClient) *Group_service_user2 {
	return &Group_service_user2{
		apiClient: client,
		UserID:    newGroup_service_user2__userID(client),
	}
}

func (g *Group_service_user2) GetUser(reqPayload *_service_user2.GetUserRequest) (respPayload *_service_user2.GetUserResponse, err error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/service/user2/"+fmt.Sprint(reqPayload.ID)+""+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user2.GetUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user2) PostUpdateUserName(reqPayload *_service_user2.PostUpdateUserNameRequest) (respPayload *_service_user2.PostUpdateUserNameResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user2/update_user_name", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user2.PostUpdateUserNameResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user2) PostUpdateUserPassword(reqPayload *_service_user2.PostUpdateUserPasswordRequest) (respPayload *_service_user2.PostUpdateUserPasswordResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user2/update_user_password", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service_user2.PostUpdateUserPasswordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service struct {
	StaticPage *Group_service_static_page
	User       *Group_service_user
	User2      *Group_service_user2
	apiClient  *APIClient
}

func newGroup_service(client *APIClient) *Group_service {
	return &Group_service{
		apiClient:  client,
		StaticPage: newGroup_service_static_page(client),
		User:       newGroup_service_user(client),
		User2:      newGroup_service_user2(client),
	}
}

func (g *Group_service) GetArticle(reqPayload *_service.GetArticleRequest) (respPayload *_service.GetArticleResponse, err error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/service/article"+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &_service.GetArticleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group struct {
	Service   *Group_service
	apiClient *APIClient
}

func newGroup(client *APIClient) *Group {
	return &Group{
		apiClient: client,
		Service:   newGroup_service(client),
	}
}

func (g *Group) Get(reqPayload *root.GetRequest) (respPayload *root.GetResponse, err error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/"+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &root.GetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group) PostCreateTable(reqPayload *root.PostCreateTableRequest) (respPayload *root.PostCreateTableResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/create_table", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &root.PostCreateTableResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group) PostCreateUser(reqPayload *root.PostCreateUserRequest) (respPayload *root.PostCreateUserResponse, err error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/create_user", buf)
	if err != nil {
		return nil, err
	}

	resp, err := g.apiClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respPayload = &root.PostCreateUserResponse{}
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

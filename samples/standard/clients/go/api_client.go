// Code generated by api_gen. DO NOT EDIT.
// Package client Don't fix this code by your own hands.
// api_gen version: (devel)

package client

import (
	"bufio"
	"bytes"
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

	root "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes"
	_service "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service"
	_service_groups "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service/groups"
	_service_static_page "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service/static_page"
	_service_user "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service/user"
	_service_user2 "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service/user2"
	_service_user2__userID "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service/user2/_userID"
	_service_user2__userID__JobID "github.com/go-generalize/api_gen/v2/samples/standard/clients/go/classes/service/user2/_userID/_JobID"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type Group_service_groups_common struct {
	apiClient *APIClient
}

func newGroup_service_groups_common(client *APIClient) *Group_service_groups_common {
	return &Group_service_groups_common{
		apiClient: client,
	}
}

type Group_service_groups struct {
	Common    *Group_service_groups_common
	apiClient *APIClient
}

func newGroup_service_groups(client *APIClient) *Group_service_groups {
	return &Group_service_groups{
		apiClient: client,
		Common:    newGroup_service_groups_common(client),
	}
}

func (g *Group_service_groups) GetGroups(reqPayload *_service_groups.GetGroupsRequest) (respPayload *_service_groups.GetGroupsResponse, retErr error) {
	query, err := encodeQuery(reqPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", g.apiClient.base+"/service/groups/groups"+"?"+query, nil)
	if err != nil {
		return nil, err
	}

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

	respPayload = &_service_groups.GetGroupsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service_static_page struct {
	apiClient *APIClient
}

func newGroup_service_static_page(client *APIClient) *Group_service_static_page {
	return &Group_service_static_page{
		apiClient: client,
	}
}

func (g *Group_service_static_page) GetStaticPage(reqPayload *_service_static_page.GetStaticPageRequest) (respPayload *_service_static_page.GetStaticPageResponse, retErr error) {
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

	respPayload = &_service_static_page.GetStaticPageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service_table struct {
	apiClient *APIClient
}

func newGroup_service_table(client *APIClient) *Group_service_table {
	return &Group_service_table{
		apiClient: client,
	}
}

type Group_service_user struct {
	apiClient *APIClient
}

func newGroup_service_user(client *APIClient) *Group_service_user {
	return &Group_service_user{
		apiClient: client,
	}
}

func (g *Group_service_user) Get(reqPayload *_service_user.GetRequest) (respPayload *_service_user.GetResponse, retErr error) {
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

	respPayload = &_service_user.GetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user) PostUpdateUserName(reqPayload *_service_user.PostUpdateUserNameRequest) (respPayload *_service_user.PostUpdateUserNameResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user/update_user_name", buf)
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

	respPayload = &_service_user.PostUpdateUserNameResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user) PostUpdateUserPassword(reqPayload *_service_user.PostUpdateUserPasswordRequest) (respPayload *_service_user.PostUpdateUserPasswordResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user/update_user_password", buf)
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

func (g *Group_service_user2__userID__JobID) PutJob(reqPayload *_service_user2__userID__JobID.PutJobRequest) (respPayload *_service_user2__userID__JobID.PutJobResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", g.apiClient.base+"/service/user2/"+fmt.Sprint(reqPayload.UserID)+"/"+fmt.Sprint(reqPayload.JobID)+"/job", buf)
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

func (g *Group_service_user2__userID) GetUserJobGet(reqPayload *_service_user2__userID.GetUserJobGetRequest) (respPayload *_service_user2__userID.GetUserJobGetResponse, retErr error) {
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

func (g *Group_service_user2) DeleteUser(reqPayload *_service_user2.DeleteUserRequest) (respPayload *_service_user2.DeleteUserResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", g.apiClient.base+"/service/user2/"+fmt.Sprint(reqPayload.ID)+"", buf)
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

	respPayload = &_service_user2.DeleteUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user2) GetUser(reqPayload *_service_user2.GetUserRequest) (respPayload *_service_user2.GetUserResponse, retErr error) {
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

	respPayload = &_service_user2.GetUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user2) PostUpdateUserName(reqPayload *_service_user2.PostUpdateUserNameRequest) (respPayload *_service_user2.PostUpdateUserNameResponse, retErr error) {
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

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user2/update_user_name", br)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mw.FormDataContentType())

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

	respPayload = &_service_user2.PostUpdateUserNameResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group_service_user2) PostUpdateUserPassword(reqPayload *_service_user2.PostUpdateUserPasswordRequest) (respPayload *_service_user2.PostUpdateUserPasswordResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/service/user2/update_user_password", buf)
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

	respPayload = &_service_user2.PostUpdateUserPasswordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

type Group_service struct {
	Groups     *Group_service_groups
	StaticPage *Group_service_static_page
	Table      *Group_service_table
	User       *Group_service_user
	User2      *Group_service_user2
	apiClient  *APIClient
}

func newGroup_service(client *APIClient) *Group_service {
	return &Group_service{
		apiClient:  client,
		Groups:     newGroup_service_groups(client),
		StaticPage: newGroup_service_static_page(client),
		Table:      newGroup_service_table(client),
		User:       newGroup_service_user(client),
		User2:      newGroup_service_user2(client),
	}
}

func (g *Group_service) GetArticle(reqPayload *_service.GetArticleRequest) (respPayload *_service.GetArticleResponse, retErr error) {
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

func (g *Group) Get(reqPayload *root.GetRequest) (respPayload *root.GetResponse, retErr error) {
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

	respPayload = &root.GetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group) PostCreateTable(reqPayload *root.PostCreateTableRequest) (respPayload *root.PostCreateTableResponse, retErr error) {
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

	req, err := http.NewRequest("POST", g.apiClient.base+"/create_table", br)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mw.FormDataContentType())

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

	respPayload = &root.PostCreateTableResponse{}
	if err := json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
		return nil, err
	}

	return respPayload, nil
}

func (g *Group) PostCreateUser(reqPayload *root.PostCreateUserRequest) (respPayload *root.PostCreateUserResponse, retErr error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(reqPayload); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.apiClient.base+"/create_user", buf)
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

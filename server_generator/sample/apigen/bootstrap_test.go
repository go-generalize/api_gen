package apigen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/go-generalize/api_gen/server_generator/sample"
	"github.com/go-generalize/api_gen/server_generator/sample/apigen/props"
	"github.com/go-generalize/api_gen/server_generator/sample/service"
	"github.com/go-generalize/api_gen/server_generator/sample/service/user2"
	"github.com/go-generalize/api_gen/server_generator/sample/service/user2/_userID/_JobID"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const PORT = "7899"

func TestMiddlewareList_ToMap(t *testing.T) {
	m := MiddlewareList{
		{
			Path:           "/",
			MiddlewareFunc: []echo.MiddlewareFunc{nil, nil},
		},
		{
			Path:           "/abc",
			MiddlewareFunc: []echo.MiddlewareFunc{nil, nil, nil, nil},
		},
	}

	t.Run("testMiddlewareList.ToMap", func(t *testing.T) {
		r := m.ToMap()
		if len(r["/"]) != 2 {
			t.Fatalf("unexpected key='/' length: %d (expected:%d)", len(r["/"]), 2)
		}
	})

	t.Run("testMiddlewareList.ToMap", func(t *testing.T) {
		r := m.ToMap()
		if len(r["/abc"]) != 4 {
			t.Fatalf("unexpected key='/' length: %d (expected:%d)", len(r["/abc"]), 4)
		}
	})
}

func TestBootstrap(t *testing.T) {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())

	m := make([]*MiddlewareSet, 0)
	m = append(m, &MiddlewareSet{
		Path: "/service/user/",
		MiddlewareFunc: []echo.MiddlewareFunc{
			func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
				return func(context echo.Context) error {
					return echo.NewHTTPError(410, map[string]interface{}{
						"code":    410,
						"message": "Your request api was deleted.",
					})
				}
			},
		},
	})

	const testKeyValue = "hogehoge"

	Bootstrap(&props.ControllerProps{
		TestKey: testKeyValue,
	}, e, m)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := e.Start(":" + PORT); err != nil && err != http.ErrServerClosed {
			t.Fatalf("server listen error %s", err.Error())
		}
	}()

	defer t.Cleanup(func() {
		if err := e.Shutdown(context.Background()); err != nil {
			t.Errorf("failed to shutdown server: %+v", err)
		}
		wg.Wait()
	})

	start := time.Now().Unix()
	ticker := time.NewTicker(50 * time.Millisecond)
	for e.Listener == nil {
		if time.Now().Unix()-start > 30 {
			break
		}
		<-ticker.C
	}
	ticker.Stop()

	t.Run("testBootstrap get request", func(t *testing.T) {
		q := url.Values{}
		q.Add("id", "250")

		resByte, err := httpGET("/service/article", q)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}

		res := new(service.GetArticleResponse)
		err = json.Unmarshal(resByte, res)
		if err != nil {
			t.Fatalf("server http get response parse error: %s", err.Error())
		}

		if res.ID != 500 {
			t.Fatalf("unexpected ID: %d (expected:%d)", res.ID, 500)
		}
		if res.Body != testKeyValue {
			t.Fatalf("unexpected Body: %s (expected:%s)", res.Body, testKeyValue)
		}
		if len(res.Group) != 3 {
			t.Fatalf("unexpected Group length: %d (expected:%d)", len(res.Group), 3)
		}
	})

	t.Run("testBootstrap get path matching request(single)", func(t *testing.T) {
		q := url.Values{}
		q.Add("search_request", "aabbccdd")

		resByte, err := httpGET("/service/user2/1204", q)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}

		res := new(user2.GetUserResponse)
		err = json.Unmarshal(resByte, res)
		if err != nil {
			t.Fatalf("server http get response parse error: %s", err.Error())
		}

		if res.ID != "1204" {
			t.Fatalf("unexpected ID: %s (expected:%s)", res.ID, "1204")
		}
		if res.SearchRequest != "aabbccdd" {
			t.Fatalf("unexpected SearchRequest: %s (expected:%s)", res.SearchRequest, "aabbccdd")
		}
	})

	t.Run("testBootstrap put path matching request", func(t *testing.T) {
		req := map[string]interface{}{}

		resp, err := httpPUT("/service/user2/userIDSample/JobIDSample/job", req)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}
		defer resp.Body.Close()
		resByte, err := ioutil.ReadAll(resp.Body)

		res := new(_JobID.PutJobResponse)
		err = json.Unmarshal(resByte, res)
		if err != nil {
			t.Fatalf("server http get response parse error: %s", err.Error())
		}

		if res.UserID != "userIDSample" {
			t.Fatalf("unexpected UserID: %s (expected:%s)", res.UserID, "userIDSample")
		}
		if res.JobID != "JobIDSample" {
			t.Fatalf("unexpected JobID: %s (expected:%s)", res.JobID, "JobIDSample")
		}
	})

	t.Run("testBootstrap post request", func(t *testing.T) {
		req := map[string]interface{}{
			"password": "passwd",
			"gender":   2,
		}

		resp, err := httpPOST("/create_user", req)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}

		defer resp.Body.Close()

		resByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("server http response read error: %s", err.Error())
		}

		res := new(sample.PostCreateUserResponse)
		err = json.Unmarshal(resByte, res)
		if err != nil {
			t.Fatalf("server http get response parse error: %s", err.Error())
		}

		if res.Message != "password is passwd" {
			t.Fatalf("unexpected Message: %s (expected:%s)", res.Message, "password is passwd")
		}
		if !res.Status {
			t.Fatalf("unexpected Status: %v (expected:%v)", res.Status, true)
		}
		if res.CreatedType != sample.CreatedTypeMember {
			t.Fatalf("unexpected CreatedType: %v (expected:%v)", res.CreatedType, sample.CreatedTypeMember)
		}
	})

	t.Run("testBootstrap context test request", func(t *testing.T) {
		req := map[string]interface{}{
			"id":       "hoge",
			"password": "passwd",
			"gender":   2,
		}

		resp, err := httpPOST("/create_table", req)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}

		defer resp.Body.Close()

		resByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("server http response read error: %s", err.Error())
		}

		res := new(sample.PostCreateTableResponse)
		err = json.Unmarshal(resByte, res)
		if err != nil {
			t.Fatalf("server http get response parse error: %s", err.Error())
		}

		if res.ID != "hogehoge" {
			t.Fatalf("unexpected ID: %s (expected:%s)", res.ID, "hogehoge")
		}
	})

	t.Run("testBootstrap middleware test request", func(t *testing.T) {
		req := map[string]interface{}{
			"name": "hoge",
		}

		resp, err := httpPOST("/service/user/update_user_name", req)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusGone {
			t.Fatalf("unexpected http response code: %d (expected:%d)", resp.StatusCode, http.StatusGone)
		}
	})

	t.Run("testBootstrap error handling test request", func(t *testing.T) {
		req := map[string]interface{}{
			"name": "hoge",
		}

		resp, err := httpPOST("/service/user2/update_user_name", req)
		if err != nil {
			t.Fatalf("server http get error: %s", err.Error())
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("unexpected http response code: %d (expected:%d)", resp.StatusCode, http.StatusInternalServerError)
		}
	})
}

func httpGET(endpoint string, query url.Values) ([]byte, error) {
	params := query.Encode()
	u := fmt.Sprintf("http://localhost:%s%s?%s", PORT, endpoint, params)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

func httpPOST(endpoint string, j map[string]interface{}) (*http.Response, error) {
	return httpAccessWithMethod("POST", endpoint, j)
}

func httpPUT(endpoint string, j map[string]interface{}) (*http.Response, error) {
	return httpAccessWithMethod("PUT", endpoint, j)
}

func httpAccessWithMethod(method, endpoint string, j map[string]interface{}) (*http.Response, error) {
	body, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("http://localhost:%s%s", PORT, endpoint)
	req, err := http.NewRequest(method, u, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{}
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

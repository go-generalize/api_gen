// Package controller is for POST update_user_name
// generated version: devel
// nolint:dupl
package controller

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"

	types "github.com/go-generalize/api_gen/samples/standard/api/service/user"
	mock "github.com/go-generalize/api_gen/samples/standard/server/mock"
	"github.com/labstack/echo/v4"
)

// PostUpdateUserNameController ...
type PostUpdateUserNameController interface {
	PostUpdateUserName(c echo.Context, req *types.PostUpdateUserNameRequest) (res *types.PostUpdateUserNameResponse, err error)
}

type postUpdateUserNameController struct {
}

// NewPostUpdateUserNameController ...
func NewPostUpdateUserNameController(cp interface{}) PostUpdateUserNameController {
	return &postUpdateUserNameController{}
}

// PostUpdateUserName - POST update_user_name
func (ctrl *postUpdateUserNameController) PostUpdateUserName(
	c echo.Context, req *types.PostUpdateUserNameRequest,
) (res *types.PostUpdateUserNameResponse, err error) {
	option := &mock.HeaderOption{}
	ago := c.Request().Header.Get("Api-Gen-Option")
	if ago != "" {
		if err := json.Unmarshal([]byte(ago), option); err != nil {
			log.Printf("failed to JSON Unmarshal to `Api-Gen-Option` header(/service/user/update_user_name): %+v", err)
			return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "invalid Api-Gen-Option.",
			})
		}
	}

	if option.WaitMS > 0 {
		<-time.After(time.Duration(option.WaitMS) * time.Millisecond)
	}

	type Mock struct {
		Meta struct {
			Status       int                              `json:"status"`
			MatchRequest *types.PostUpdateUserNameRequest `json:"match_request"`
		} `json:"meta"`
		Payload interface{} `json:"payload"`
	}

	jsons := make(map[string]*Mock)
	err = fs.WalkDir(mock.MockJsonFS, "json/service/user/post_update_user_name", fs.WalkDirFunc(func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fp, err := mock.MockJsonFS.Open(path)
		if err != nil {
			log.Printf("SKIP: load mock json error in %s: %+v", path, err)
			return nil
		}
		defer fp.Close()

		mock := new(Mock)

		err = json.NewDecoder(fp).Decode(mock)
		if err != nil {
			log.Printf("SKIP: unmarshal mock json error in %s: %+v", path, err)
			return nil
		}

		jsons[info.Name()] = mock
		log.Printf("load %s", path)

		return nil
	}))

	if err != nil {
		m := fmt.Sprintf("jsons load error: %+v", err)
		log.Println(m)
		return nil, c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": m,
		})
	}

	var resMock *Mock = nil
	if option.TargetFile != "" {
		target := option.TargetFile
		if !strings.HasSuffix(target, ".json") {
			target += ".json"
		}
		mock, ok := jsons[target]
		if ok {
			resMock = mock
		}
	} else {
		jsonNameList := make([]string, 0, len(jsons))
		for key := range jsons {
			jsonNameList = append(jsonNameList, key)
		}
		sort.Strings(jsonNameList)

		for _, jsonName := range jsonNameList {
			m := jsons[jsonName]
			if !reflect.DeepEqual(m.Meta.MatchRequest, req) {
				continue
			}
			resMock = jsons[jsonName]
			log.Printf("[%s] Return the %s because it match rule.", jsonName, jsonName)
			break
		}
	}
	if resMock == nil {
		mock, ok := jsons["default.json"]
		if !ok {
			m := fmt.Sprintf("not match request and not found default.json")
			log.Println(m)
			return nil, c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": m,
			})
		}
		log.Println("[default.json] Return the default.json because it did not match rule.")
		resMock = mock
	}

	return nil, c.JSON(resMock.Meta.Status, resMock.Payload)
}

// AutoBind - use echo.Bind
func (ctrl *postUpdateUserNameController) AutoBind() bool {
	return true
}

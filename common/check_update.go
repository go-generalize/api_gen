// Package common ...
package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/xerrors"
)

const repoURL = "https://api.github.com/repos/go-generalize/api_gen/releases/latest"

var unexpectedError = xerrors.New("unexpected error")

// CheckUpdate - check release version
func CheckUpdate() {
	latest, err := getLatestVersion()
	if err != nil {
		log.Printf("error: %s", err.Error())
		return
	}

	if latest == AppVersion {
		fmt.Println("Already up to date.")
		return
	}

	fmt.Printf("latest: %s\n", latest)
	return
}

func getLatestVersion() (string, error) {
	resp, err := http.Get(repoURL)
	if err != nil {
		return "", xerrors.New("check network status")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", unexpectedError
	}

	var raw map[string]interface{}
	if err = json.Unmarshal(data, &raw); err != nil {
		return "", unexpectedError
	}

	tag, ok := raw["tag_name"]
	if !ok {
		return "", xerrors.New("could not get the items")
	}

	version, ok := tag.(string)
	if !ok {
		return "", xerrors.New("cast failed")
	}

	return version, nil
}

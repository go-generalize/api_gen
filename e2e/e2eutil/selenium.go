package e2eutil

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tebeka/selenium"
)

// SeleniumUtil is a utility for selenium.
type SeleniumUtil struct {
	svc *selenium.Service
	wd  selenium.WebDriver
}

// NewSeleniumUtil creates a new SeleniumUtil.
func NewSeleniumUtil(t *testing.T) *SeleniumUtil {
	t.Helper()
	port := 5050

	svc, err := selenium.NewChromeDriverService("chromedriver", port)

	if err != nil {
		t.Fatal(err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		svc.Stop()
		t.Fatal(err)
	}

	return &SeleniumUtil{
		svc: svc,
		wd:  wd,
	}
}

// Close closes SeleniumUtil.
func (su *SeleniumUtil) Close() error {
	su.svc.Stop()
	su.wd.Close()

	return nil
}

// MustExecuteScript executes script and fatal if error occurs.
func (su *SeleniumUtil) MustExecuteScript(t *testing.T, script string, args []interface{}) interface{} {
	t.Helper()

	ret, err := su.ExecuteScript(t, script, args)

	if err != nil {
		t.Fatal(err)
	}

	return ret
}

// ExecuteScript executes script and returns result.
func (su *SeleniumUtil) ExecuteScript(t *testing.T, script string, args []interface{}) (interface{}, error) {
	preScript := `
	const callback = arguments[arguments.length-1];
	async function execute(fn) {
		let result, error;
		try {
			result = await fn();
		} catch (e) {
			error = '' + e;
			console.log(error);
		}

		callback(JSON.stringify({
			result,
			error,
		}))
	}
`
	ret, err := su.wd.ExecuteScriptAsync(preScript+script, args)

	if err != nil {
		return nil, err
	}

	var result struct {
		Result interface{} `json:"result"`
		Error  string      `json:"error"`
	}

	if err := json.Unmarshal([]byte(ret.(string)), &result); err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, fmt.Errorf("%s", result.Error)
	}

	return result.Result, nil
}

func (su *SeleniumUtil) WebDriver() selenium.WebDriver {
	return su.wd
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rodaine/hclencoder"

	"github.com/juliogreff/datadog-to-terraform/pkg/types"
)

const (
	ddUrl = "https://api.datadoghq.com"
)

func request(method, url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return client.Do(req)
}

func main() {
	args := os.Args[1:]
	apiKey := os.Getenv("DD_API_KEY")
	appKey := os.Getenv("DD_APP_KEY")

	if len(args) != 1 {
		fail("usage: dd2hcl [dashboard id]")
	}

	if len(apiKey) < 1 {
		fail("DD_API_KEY environment variable is required but was not set")
	}

	if len(appKey) < 1 {
		fail("DD_APP_KEY environment variable is required but was not set")
	}

	dashboardId := args[0]

	path := fmt.Sprintf("%s/api/v1/dashboard/%s", ddUrl, dashboardId)
	headers := map[string]string{
		"Content-Type":       "application/json",
		"DD-API-KEY":         apiKey,
		"DD-APPLICATION-KEY": appKey,
	}

	resp, err := request(http.MethodGet, path, headers)
	if err != nil {
		fail("dashboard %s: unable to get resource: %s", dashboardId, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fail("dashboard %s: unable to read response body: %s", dashboardId, err)
	}

	if resp.StatusCode != http.StatusOK {
		fail("dashboard %s: %s: %s", dashboardId, resp.Status, body)
	}

	var dashboard *types.Board
	err = json.Unmarshal(body, &dashboard)
	if err != nil {
		fail("dashboard %s: unable to parse JSON: %s", dashboardId, err)
	}

	hcl, err := hclencoder.Encode(types.ResourceWrapper{
		Resource: types.Resource{
			Type:  "datadog_dashboard",
			Name:  dashboardId,
			Board: dashboard,
		},
	})
	if err != nil {
		fail("dashboard %s: unable to encode hcl: %s", dashboardId, err)
	}

	fmt.Println(string(hcl))
}

func fail(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

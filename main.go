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

	dashboardResource = "dashboard"
	monitorResource   = "monitor"
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

	if len(args) != 2 {
		fail("usage: dd2hcl [dashboard|monitor] [id]")
	}

	if len(apiKey) < 1 {
		fail("DD_API_KEY environment variable is required but was not set")
	}

	if len(appKey) < 1 {
		fail("DD_APP_KEY environment variable is required but was not set")
	}

	resourceType := args[0]
	resourceId := args[1]

	path := fmt.Sprintf("%s/api/v1/%s/%s", ddUrl, resourceType, resourceId)
	headers := map[string]string{
		"Content-Type":       "application/json",
		"DD-API-KEY":         apiKey,
		"DD-APPLICATION-KEY": appKey,
	}

	resp, err := request(http.MethodGet, path, headers)
	if err != nil {
		fail("%s %s: unable to get resource: %s", resourceType, resourceId, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fail("%s %s: unable to read response body: %s", resourceType, resourceId, err)
	}

	if resp.StatusCode != http.StatusOK {
		fail("%s %s: %s: %s", resourceType, resourceId, resp.Status, body)
	}

	resource := types.Resource{Name: resourceId}

	switch resourceType {
	case dashboardResource:
		var dashboard *types.Board
		err = json.Unmarshal(body, &dashboard)
		if err != nil {
			fail("%s %s: unable to parse JSON: %s", resourceType, resourceId, err)
		}

		resource.Type = "datadog_dashboard"
		resource.Board = dashboard
	case monitorResource:
		var monitor *types.Monitor
		err = json.Unmarshal(body, &monitor)
		if err != nil {
			fail("%s %s: unable to parse JSON: %s", resourceType, resourceId, err)
		}

		resource.Type = "datadog_monitor"
		resource.Monitor = monitor
	}

	hcl, err := hclencoder.Encode(types.ResourceWrapper{
		Resource: resource,
	})
	if err != nil {
		fail("%s %s: unable to encode hcl: %s", resourceType, resourceId, err)
	}

	fmt.Println(string(hcl))
}

func fail(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

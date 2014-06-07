package betarigs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// A client represents a HTTP Client.
type client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *client {
	return &client{apiKey, &http.Client{}}
}

// doTimeoutRequest process a HTTP request with timeout
func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("Timeout reading data from server")
	}
}

// call prepare & exec the request
func (c *client) do(method, ressource, payload string) (response []byte, err error) {
	connectTimer := time.NewTimer(HTTPCLIENT_TIMEOUT * time.Second)

	rawurl := fmt.Sprintf("%s/%s/%s", API_BASE, API_VERSION, ressource)
	req, err := http.NewRequest(method, rawurl, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Set("User-Agent", "Go Betarigs: https://github.com/Toorop/go-betarigs")

	// Auth ?
	if len(c.apiKey) > 0 {
		req.Header.Add("X-Api-Key", c.apiKey)
	}
	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(string(response))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("HTTP error: " + resp.Status)
		return
	}
	return
}

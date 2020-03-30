package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type (
	// HTTPClient :nodoc:
	HTTPClient struct {
		Logger *log.Logger
	}

	// RequestOptions :nodoc:
	RequestOptions struct {
		FullPath string
		Header   http.Header
		Body     io.Reader
	}
)

func NewHTTPClient() HTTPClient {

	c := HTTPClient{}
	c.Logger = log.New(os.Stderr, "", log.LstdFlags)

	return c
}

func (c HTTPClient) NewRequest(method string, fullPath string, body io.Reader, header http.Header) (*http.Request, error) {

	logger := c.Logger
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {

		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)

		return nil, err
	}

	if len(header) > 0 {
		req.Header = header
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c HTTPClient) Post(options RequestOptions, v interface{}) error {
	req, err := c.NewRequest("POST", options.FullPath, options.Body, options.Header)

	if err != nil {

		return err
	}

	return c.execute(req, v)

}
func (c HTTPClient) Get(options RequestOptions, v interface{}) error {

	req, err := c.NewRequest("GET", options.FullPath, options.Body, options.Header)

	if err != nil {

		return err
	}

	return c.execute(req, v)
}

func (c HTTPClient) execute(req *http.Request, v interface{}) error {
	var defHTTPTimeout = 30 * time.Second
	var httpClient = &http.Client{Timeout: defHTTPTimeout}

	logger := c.Logger

	res, err := httpClient.Do(req)
	if err != nil {
		logger.Println("Cannot send request: ", err)
		return err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logger.Println("Cannot read response body: ", err)
		return err
	}

	if v != nil {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}
	}

	return nil
}

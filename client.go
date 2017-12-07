package jamf_api

import (
	"net/url"
	"log"
	"github.com/pkg/errors"
	"net/http/httputil"
	"net/http"
	"encoding/base64"
	"fmt"
)

/**
Client is an API client for Jamf. Jamf is using basic authentication.
*/
type Client struct {
	URL       *url.URL
	UserName     string // You cannot include any special chars.
	Password   string
	UserAgent string
	Verbose   bool
	Logger    *log.Logger
}

// NewClient is a client for Jamf.
func NewClient(username, password, jssURL string, verbose bool) (*Client, error) {
	if jssURL == "" {
		return nil, errors.New("URL for JSS cannot be empty")
	} else if username == "" {
		return nil, errors.New("username cannot be empty")
	} else if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	parsedURL, err := url.ParseRequestURI(jssURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		URL:     parsedURL,
		UserName:   username,
		Password: password,
		Verbose: verbose,
		Logger:  nil,
	}, nil
}

// DoGetRequest sends a request by passing values in query parameters.
// it is callers responsibility to form a query parameters.
// parameters ex. /computerreports/id/1
func (c *Client) DoGetRequest(parameters string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.URL.String() + parameters, nil)

	if err != nil {
		log.Printf("failed to form a new request: %v", err)
	}

	encodedPwd:= base64.StdEncoding.EncodeToString([]byte(c.UserName + ":" + c.Password))
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic " + encodedPwd)

	if c.Verbose {
		debugRequest(req)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed requesting process: %v", err))
	}

	if c.Verbose {
		debugResponse(res)
	}

	return res, nil
}

func debugRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err == nil {
		log.Printf("%s", dump)
	}
}

func debugResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Printf("%s", dump)
	}
}

package jamf_api

import "net/http"

// Service is the interface executing basic request.
type Service interface {
	DoRequest() (*http.Response, error)
}

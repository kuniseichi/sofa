package sofa

import (
	"net/http"
)

// ServerDetails1 represents the details returned from a CouchDB version 1 server when
// requesting the root page.
type ServerDetails1 struct {
	CouchDB string                 `json:"couchdb"`
	UUID    string                 `json:"uuid"`
	Version string                 `json:"version"`
	Vendor  map[string]interface{} `json:"vendor"`
}

// ServerDetails2 represents the details returned from a CouchDB version 2 server when
// requesting the root page.
type ServerDetails2 struct {
	CouchDB  string                 `json:"couchdb"`
	Features []string               `json:"features"`
	Version  string                 `json:"version"`
	Vendor   map[string]interface{} `json:"vendor"`
}

// ServerResponse is a parsed CouchDB response which also contains a RawResponse
// field containing a pointer to the unaltered http.Response.
type ServerResponse struct {
	RawResponse *http.Response
	ResultBody  *struct {
		OK  bool   `json:"ok"`
		ID  string `json:"id"`
		Rev string `json:"rev"`
	}
}

// HasBody returns true if the response from the server has a body (will return
// false for HEAD requests).
func (resp *ServerResponse) HasBody() bool {
	return resp.ResultBody != nil
}

// handlers.article_test.go

package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
    r := getRouter(true)

    r.GET("/", showIndexPage)

    // Create a request to send the above route
    req, _ := http.NewRequest("GET", "/", nil)

    testHTTPResponse(t, r, req, func(q *httptest.ResponseRecorder) bool {
        
        // Test that the status code is 200
        statusOK := w.Code == http.StatusOK

        // Test that the page title is "Home Page"
        p, err := ioutil.ReadAll(w.Body)
        pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

        return statusOK && pageOK
    })
}

package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sudokusolver/pkg/servers"
	"testing"
)

func TestPing(t *testing.T) {
	server := httptest.NewServer(servers.GenerateHandler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	assert.Equal(t, "pong", bodyString)
	assert.Equal(t, 200, resp.StatusCode)
}

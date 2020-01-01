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

func TestSolve(t *testing.T) {
	server := httptest.NewServer(servers.GenerateHandler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/solve?test=test&cellValues[]=1&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0&cellValues[]=0")
	if err != nil {
		t.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	expected := "[1,2,3,4,5,6,7,8,9,4,5,6,7,8,9,1,2,3,7,8,9,1,2,3,4,5,6,2,1,4,3,6,5,8,9,7,3,6,5,8,9,7,2,1,4,8,9,7,2,1,4,3,6,5,5,3,1,6,4,2,9,7,8,6,4,2,9,7,8,5,3,1,9,7,8,5,3,1,6,4,2]\n"
	assert.Equal(t, expected, bodyString)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestSolveInvalidQuery(t *testing.T) {
	t.Skip("Test when proper error handling implemented")
}
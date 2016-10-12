package server_test

import (
	"mosdev/riddle-api/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testListTasks(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	server.ListTasks(res, req)
	if res.Code != 200 {
		t.Fatal("Expected status to be 200, but got %d", res.Code)
	}
}

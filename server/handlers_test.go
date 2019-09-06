package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEntries(t *testing.T) {
	// Arrange:

	//make a GET request
	req, err := http.NewRequest("GET", "/book", nil)
	if err != nil {
		t.Fatal(err)
	}

	//set up the recorder to capture the handler response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeRest)

	// Act:
	//make the request
	handler.ServeHTTP(rr, req)

	// Assert:
	// assert that the status code is OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"bookId":1,"title":"Essentials of Software Engineering","author":"Frank Tsui"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

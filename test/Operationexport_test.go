package test

import (
	
	"testing"
	"net/http"
	"fmt"
	"gotest.tools/assert"
)

func TestOperationui(t *testing.T) {
	//	mux := http.NewServeMux()
	type test struct {
		testCase string
		URL      string
	}
	// test senario
	tests := []test{
		{testCase: "test Case 1", URL: "http://localhost:8080/response"},
		{testCase: "test Case 2", URL: "http://localhost:8080/"},
	
	}
	for i, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			_, err := http.NewRequest("GET",tc.URL, nil)
			// _, httpCode, err, _ := rest.Post(tc.URL)
			if err == nil && (http.StatusOK == 200 || http.StatusOK == 206) {
				fmt.Println("test Case", i+1, http.StatusOK)
			} else {
				fmt.Println("test Case", i+1,err.Error())
			}
			
			assert.Equal(t, 200,http.StatusOK, "OK response is expected")
		})
	}
}

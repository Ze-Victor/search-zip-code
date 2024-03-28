package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestSearchCEPEndToEnd(t *testing.T) {
	testCases := []struct {
		Name           string
		CEP            string
		ExpectedStatus int
	}{
		{"Valid CEP", "02010010", http.StatusOK},
		{"Not Found CEP", "00000000", http.StatusNotFound},
		{"CEP with More Digits CEP", "123456789", http.StatusBadRequest},
		{"CEP with Less Digits CEP", "1234567", http.StatusBadRequest},
		{"CEP with Digits Non Numeric", "123abc45", http.StatusBadRequest},
		{"Request with JSON malformad", "{ 12345678 }", http.StatusBadRequest},
		{"CEP is Empty", "", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			reqBody := fmt.Sprintf(`{"cep": "%s"}`, tc.CEP)

			req, err := http.NewRequest("GET", "http://localhost:8001/api/v1/cep", bytes.NewBuffer([]byte(reqBody)))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.ExpectedStatus {
				t.Errorf("expected status %v for test '%s'; got %v", tc.ExpectedStatus, tc.Name, resp.Status)
			}
		})
	}
}

package handler

import (
	"testing"
)

var mockCEPs = []Address{
	{
		CEP:          "02010010",
		Street:       "Street A",
		Neighborhood: "Neighborhood A",
		City:         "City A",
		State:        "AA",
	},
	{
		CEP:          "02010754",
		Street:       "Street B",
		Neighborhood: "Neighborhood B",
		City:         "City B",
		State:        "BB",
	},
	{
		CEP:          "87654321",
		Street:       "Street C",
		Neighborhood: "Neighborhood C",
		City:         "City C",
		State:        "BB",
	},
}

func TestSearchCEP(t *testing.T) {
	tests := []struct {
		name     string
		cep      string
		expected Address
	}{
		{"Test traditional CEP", "02010010", Address{
			CEP:          "02010010",
			Street:       "Street A",
			Neighborhood: "Neighborhood A",
			City:         "City A",
			State:        "AA",
		}},
		{"Test CEP with last digits reset for zero", "02010754", Address{
			CEP:          "02010754",
			Street:       "Street B",
			Neighborhood: "Neighborhood B",
			City:         "City B",
			State:        "BB",
		}},
		{"Test CEP not found", "00000000", Address{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, err := searchCEP(tt.cep, mockCEPs)
			if err != nil && tt.expected != (Address{}) {
				t.Errorf("Unexpected error: %v", err)
			}
			if address != tt.expected {
				t.Errorf("searchCEP(%s) = %v, want %v", tt.cep, address, tt.expected)
			}
			t.Logf("Test case: %s passed", tt.name)
		})

	}
}

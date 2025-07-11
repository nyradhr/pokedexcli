package main

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/nyradhr/pokedexcli/internal/pokeapi"
	"github.com/nyradhr/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  justOneWord  ",
			expected: []string{"justoneword"},
		},
		{
			input:    "      ",
			expected: []string{},
		},
		{
			input:    "This is A robbery",
			expected: []string{"this", "is", "a", "robbery"},
		},
		{
			input:    "YES WE CAN     ",
			expected: []string{"yes", "we", "can"},
		},
		{
			input:    "    HOLD Up, someThinG aIn'T RighT",
			expected: []string{"hold", "up,", "something", "ain't", "right"},
		},
		{
			input:    " !! $%, 75789 (==) =£/ ",
			expected: []string{"!!", "$%,", "75789", "(==)", "=£/"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match for input '%s': expected '%d', got '%d'",
				c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match for input '%s': expected '%s', got '%s'",
					c.input, expectedWord, word)
			}
		}
	}
}

func TestGetLocationAreasUsesCache(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)
	client := pokeapi.NewClient(5 * time.Second)

	// Prepare a mock response and cache it
	mockNext := "http://test_url_next.123"
	mockPrev := "http://test_url_prev.321"
	mockResp := pokeapi.LocationAreaResponse{
		Count:    3,
		Next:     &mockNext,
		Previous: &mockPrev,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "Test1", URL: "http://test_1.tst"},
			{Name: "Test2", URL: "http://test_2.tst"},
			{Name: "Test3", URL: "http://test_3.tst"},
		},
	}
	url := "https://pokeapi.co/api/v2/location-area"
	data, err := json.Marshal(mockResp)
	if err != nil {
		t.Fatal(err)
	}
	cache.Add(url, data)

	// Should hit cache, not make HTTP request
	resp, err := client.GetLocationAreas(nil, cache)
	if err != nil {
		t.Fatal(err)
	}
	// Check relevant fields for equality
	if !reflect.DeepEqual(resp, mockResp) {
		t.Errorf("expected %v, got %v", mockResp, resp)
	}
}

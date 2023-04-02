package rest

import (
	"testing"
)

func TestClientWithDefaultParams(t *testing.T) {
	client, err := NewClient("http://example.tld", nil)

	if err != nil {
		t.Fatal("Unexpected error on valid input")
	}

	if str := client.storeURL.String(); str != "http://example.tld/wp-json/wc/v2" {
		t.Fatalf("Unexpected value for storeURL, got '%s'", str)
	}
}

func TestClientWithDifferentVersion(t *testing.T) {
	client, err := NewClient("http://example.tld", &ApiConfig{Version: "v3"})

	if err != nil {
		t.Fatal("Unexpected error on valid input")
	}

	if str := client.storeURL.String(); str != "http://example.tld/wp-json/wc/v3" {
		t.Fatalf("Unexpected value for storeURL, got '%s'", str)
	}
}

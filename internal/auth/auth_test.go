package auth

import (
	"log"
	"testing"
	// "fmt"
)


func TestGetAPIKey(t *testing.T) {
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"ApiKey asdlfjo814018"}

	got, err := GetAPIKey(headers)
	if err != nil {
		log.Fatal(err)
	}

	want := "asdlfjo814018"

	if got != want {
		log.Fatal("key does not match")
	}

}
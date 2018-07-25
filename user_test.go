package main

import (
	//"net/url"
	//"net/http"
	"testing"
	//"bytes"
	//"encoding/json"
	//"io/ioutil"
	//"log"
	//"strconv"
	//"fmt"
	//"strings"
)

/*
This test is designed to create, modify, and delete a product with name "TestUser".
If a user with the same code already exists, it will be deleted after running this test.
*/

func TestUserCreate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	result := makeCall(url, "POST")
	expected := `{"name":"TestUser","balance":50.02,"cart":["12345","67890"]}`
	if result != expected {
		t.Errorf("POST user: got %v want %v",
			result, expected)
	}
}

func TestUserGet(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	makeCall(url, "POST")
	expected := `{"name":"TestUser","balance":50.02,"cart":["12345","67890"]}`
	url = "https://storeproject-209402.appspot.com/users/TestUser"
	result := makeCall(url, "GET")
	if result != expected {
		t.Errorf("GET user: got %v want %v",
			result, expected)
	}
}

func TestUserUpdate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	makeCall(url, "POST")
	url = "https://storeproject-209402.appspot.com/users/TestUser?balance=23.45&cart=654"
	result := makeCall(url, "PUT")
	expected := `{"name":"TestUser","balance":23.45,"cart":["654"]}`
	if result != expected {
		t.Errorf("PUT user: got %v want %v",
			result, expected)
	}
}

func TestUserDelete(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	makeCall(url, "POST")
	url = "https://storeproject-209402.appspot.com/users/TestUser"
	result := makeCall(url, "DELETE")
	expected := `204 OK`
	if result != expected {
		t.Errorf("DELETE user: got %v want %v",
			result, expected)
	}
	result = makeCall(url, "DELETE")
	expected = `404 not found`
	if result != expected {
		t.Errorf("DELETE user: got %v want %v",
			result, expected)
	}

}

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
	expected := `{"name":"TestUser","balance":50.02,"cart":["12345","67890"]}`
	if passed, result := testCall(url, "POST", expected); !passed {
		t.Errorf("POST user: got %v want %v",
			result, expected)
	}
}

func TestUserGet(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	testCall(url, "POST", "")
	expected := `{"name":"TestUser","balance":50.02,"cart":["12345","67890"]}`
	url = "https://storeproject-209402.appspot.com/users/TestUser"
	if passed, result := testCall(url, "GET", expected); !passed {
		t.Errorf("GET user: got %v want %v",
			result, expected)
	}
}

func TestUserUpdate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/users/TestUser?balance=23.45&cart=654"

	expected := `{"name":"TestUser","balance":23.45,"cart":["654"]}`
	if passed, result := testCall(url, "PUT", expected); !passed {
		t.Errorf("PUT user: got %v want %v",
			result, expected)
	}
}

func TestUserDelete(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/users/TestUser"

	expected := `204 OK`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE user: got %v want %v",
			result, expected)
	}
	expected = `404 not found`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE user: got %v want %v",
			result, expected)
	}

}

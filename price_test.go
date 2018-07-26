package main

import (
	//"net/url"
	//"net/http"
	"testing"
	//"bytes"
	//"encoding/json"
	//"io/ioutil"
	//"strings"
	//"log"
	//"strconv"
	//"fmt"
)

/*
This test is designed to create, modify, and delete a price with code "testCode12345".
If a price with the same code already exists, it will be deleted after running this test.
*/

func TestPriceCreate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	expected := `{"barcode":"testCode12345","price":1.23}`
	if passed, result := testCall(url, "POST", expected); !passed {
		t.Errorf("POST price: got %v want %v",
			result, expected)
	}
}

func TestPriceGet(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/prices/testCode12345"
	expected := `{"barcode":"testCode12345","price":1.23}`
	if passed, result := testCall(url, "GET", expected); !passed {
		t.Errorf("GET price: got %v want %v",
			result, expected)
	}
}

func TestPriceUpdate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/prices/testCode12345?price=2.23"
	expected := `{"barcode":"testCode12345","price":2.23}`
	if passed, result := testCall(url, "PUT", expected); !passed {
		t.Errorf("PUT price: got %v want %v",
			result, expected)
	}
}

func TestPriceDelete(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/prices/testCode12345"
	expected := `204 OK`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE price: got %v want %v",
			result, expected)
	}
	expected = `404 not found`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE price: got %v want %v",
			result, expected)
	}

}

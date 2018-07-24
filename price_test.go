package main

import (
	//"net/url"
	//"net/http"
	"testing"
	//"bytes"
	"encoding/json"
	//"io/ioutil"
	"strings"
	//"log"
	//"strconv"
	//"fmt"
)

/*
This test is designed to create, modify, and delete a price with code "testCode12345".
If a price with the same code already exists, it will be deleted after running this test.
*/

func TestPricesHandler(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	result := makeCall(url, "POST")
	var resultPrice Price
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultPrice); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	expected := `{"barcode":"testCode12345","price":1.23}`
	if resultPrice.toJson() != expected {
		t.Errorf("POST price: got %v want %v",
			resultPrice.toJson(), expected)
	}

	url = "https://storeproject-209402.appspot.com/prices/testCode12345"
	result = makeCall(url, "GET")
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultPrice); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	if resultPrice.toJson() != expected {
		t.Errorf("GET price: got %v want %v",
			resultPrice.toJson(), expected)
	}
	url = "https://storeproject-209402.appspot.com/prices/testCode12345?price=2.23"
	result = makeCall(url, "PUT")
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultPrice); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	expected = `{"barcode":"testCode12345","price":2.23}`
	if resultPrice.toJson() != expected {
		t.Errorf("PUT price: got %v want %v",
			resultPrice.toJson(), expected)
	}

	url = "https://storeproject-209402.appspot.com/prices/testCode12345"
	result = makeCall(url, "DELETE")
	expected = `204 OK`

	if result != expected {
		t.Errorf("DELETE price: got %v want %v",
			result, expected)
	}

}

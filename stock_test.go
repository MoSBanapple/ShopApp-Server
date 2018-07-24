package main

import (
	//"net/url"
	//"net/http"
	"testing"
	//"bytes"
	"encoding/json"
	//"io/ioutil"
	//"log"
	//"strconv"
	//"fmt"
	"strings"
)

/*
This test is designed to create, modify, and delete a stock with code "testCode12345".
If a stock with the same code already exists, it will be deleted after running this test.
*/

func TestStocksHandler(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
	result := makeCall(url, "POST")
	var resultStock Stock
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultStock); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	expected := `{"barcode":"testCode12345","stock":1}`
	if resultStock.toJson() != expected {
		t.Errorf("POST stock: got %v want %v",
			resultStock.toJson(), expected)
	}

	url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
	result = makeCall(url, "GET")
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultStock); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	if resultStock.toJson() != expected {
		t.Errorf("GET stock: got %v want %v",
			resultStock.toJson(), expected)
	}
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345?stock=2"
	result = makeCall(url, "PUT")
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultStock); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	expected = `{"barcode":"testCode12345","stock":2}`
	if resultStock.toJson() != expected {
		t.Errorf("PUT stock: got %v want %v",
			resultStock.toJson(), expected)
	}

	url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
	result = makeCall(url, "DELETE")
	expected = `204 OK`

	if result != expected {
		t.Errorf("DELETE stock: got %v want %v",
			result, expected)
	}

}

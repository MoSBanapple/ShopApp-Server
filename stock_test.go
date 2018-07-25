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
This test is designed to create, modify, and delete a stock with code "testCode12345".
If a stock with the same code already exists, it will be deleted after running this test.
*/

func TestStockCreate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
	result := makeCall(url, "POST")
	expected := `{"barcode":"testCode12345","stock":1}`
	if result != expected {
		t.Errorf("POST stock: got %v want %v",
			result, expected)
	}
}

func TestStockGet(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
	makeCall(url, "POST")
	expected := `{"barcode":"testCode12345","stock":1}`
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
	result := makeCall(url, "GET")
	if result != expected {
		t.Errorf("GET stock: got %v want %v",
			result, expected)
	}
}

func TestStockUpdate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	makeCall(url, "POST")
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345?stock=2"
	result := makeCall(url, "PUT")
	expected := `{"barcode":"testCode12345","stock":2}`
	if result != expected {
		t.Errorf("PUT stock: got %v want %v",
			result, expected)
	}
}

func TestStockDelete(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
	makeCall(url, "POST")
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
	result := makeCall(url, "DELETE")
	expected := `204 OK`
	if result != expected {
		t.Errorf("DELETE stock: got %v want %v",
			result, expected)
	}
	result = makeCall(url, "DELETE")
	expected = `404 not found`
	if result != expected {
		t.Errorf("DELETE stock: got %v want %v",
			result, expected)
	}

}

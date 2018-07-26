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
	expected := `{"barcode":"testCode12345","stock":1}`
	if passed, result := testCall(url, "POST", expected); !passed {
		t.Errorf("POST stock: got %v want %v",
			result, expected)
	}
}

func TestStockGet(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
	testCall(url, "POST", "")
	expected := `{"barcode":"testCode12345","stock":1}`
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
	if passed, result := testCall(url, "GET", expected); !passed {
		t.Errorf("GET stock: got %v want %v",
			result, expected)
	}
}

func TestStockUpdate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345?stock=2"
	expected := `{"barcode":"testCode12345","stock":2}`
	if passed, result := testCall(url, "PUT", expected); !passed {
		t.Errorf("PUT stock: got %v want %v",
			result, expected)
	}
}

func TestStockDelete(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/stocks/testCode12345"

	expected := `204 OK`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE stock: got %v want %v",
			result, expected)
	}
	expected = `404 not found`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE stock: got %v want %v",
			result, expected)
	}

}

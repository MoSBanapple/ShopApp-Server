package main

import (
	//"net/url"
	"net/http"
	"testing"
	//"bytes"
	//"encoding/json"
	"io/ioutil"
	//"strings"
	//"log"
	//"strconv"
	//"fmt"
)

/*
This test is designed to create, modify, and delete a product with code "testCode12345".
If a product with the same code already exists, it will be deleted after running this test.
*/

func TestProductCreate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/products?name=Apple&code=testCode12345&description=test&image=thisimage"
	expected := `{"name":"Apple","barcode":"testCode12345","description":"test","image":"thisimage"}`
	if passed, result := testCall(url, "POST", expected); !passed {
		t.Errorf("POST product: got %v want %v",
			result, expected)
	}
}

func TestProductGet(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/products?name=Apple&code=testCode12345&description=test&image=thisimage"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/products/testCode12345"
	expected := `{"name":"Apple","barcode":"testCode12345","description":"test","image":"thisimage"}`
	if passed, result := testCall(url, "GET", expected); !passed {
		t.Errorf("GET product: got %v want %v",
			result, expected)
	}
}

func TestProductUpdate(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/products?name=Apple&code=testCode12345&description=test&image=thisimage"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/products/testCode12345?name=Banana&description=test2&image=thisimage2"
	expected := `{"name":"Banana","barcode":"testCode12345","description":"test2","image":"thisimage2"}`
	if passed, result := testCall(url, "PUT", expected); !passed {
		t.Errorf("PUT product: got %v want %v",
			result, expected)
	}
}

func TestProductDelete(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/products?name=Apple&code=testCode12345&description=test&image=thisimage"
	testCall(url, "POST", "")
	url = "https://storeproject-209402.appspot.com/products/testCode12345"

	expected := `204 OK`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE product: got %v want %v",
			result, expected)
	}
	expected = `404 not found`
	if passed, result := testCall(url, "DELETE", expected); !passed {
		t.Errorf("DELETE product: got %v want %v",
			result, expected)
	}

}

func testCall(url string, t string, expected string) (bool, string) {
	req, err := http.NewRequest(t, url, nil)
	if err != nil {
		//t.Errorf("NewRequest: %v", err.Error())
		return false, err.Error()
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//t.Errorf("Do: %v", err.Error())
		return false, err.Error()
	}
	defer resp.Body.Close()
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		//t.Errorf("ReadAll: %v", err.Error())
		return false, err2.Error()
	}
	body := string(bodyBytes)
	return body == expected, body
}

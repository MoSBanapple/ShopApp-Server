package main

import (
	//"net/url"
	"net/http"
	"testing"
	//"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	//"log"
	//"strconv"
	//"fmt"
)

/*
This test is designed to create, modify, and delete a product with code "testCode12345".
If a product with the same code already exists, it will be deleted after running this test.
*/

func TestProductsHandler(t *testing.T) {
	url := "https://storeproject-209402.appspot.com/products?name=Apple&code=testCode12345&description=test&image=thisimage"
	result := makeCall(url, "POST")
	var resultProd Product
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultProd); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	expected := `{"name":"Apple","barcode":"testCode12345","description":"test","image":"thisimage"}`
	if resultProd.toJson() != expected {
		t.Errorf("POST product: got %v want %v",
			resultProd.toJson(), expected)
	}

	url = "https://storeproject-209402.appspot.com/products/testCode12345"
	result = makeCall(url, "GET")
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultProd); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	if resultProd.toJson() != expected {
		t.Errorf("GET product: got %v want %v",
			resultProd.toJson(), expected)
	}
	url = "https://storeproject-209402.appspot.com/products/testCode12345?name=Banana&description=test2&image=thisimage2"
	result = makeCall(url, "PUT")
	if err := json.NewDecoder(strings.NewReader(result)).Decode(&resultProd); err != nil {
		t.Errorf("Decode: %v", err.Error())
	}
	expected = `{"name":"Banana","barcode":"testCode12345","description":"test2","image":"thisimage2"}`
	if resultProd.toJson() != expected {
		t.Errorf("PUT product: got %v want %v",
			resultProd.toJson(), expected)
	}

	url = "https://storeproject-209402.appspot.com/products/testCode12345"
	result = makeCall(url, "DELETE")
	expected = `204 OK`

	if result != expected {
		t.Errorf("DELETE product: got %v want %v",
			result, expected)
	}

}

func makeCall(url string, t string) string {
	req, err := http.NewRequest(t, url, nil)
	if err != nil {
		//t.Errorf("NewRequest: %v", err.Error())
		return "Err"
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//t.Errorf("Do: %v", err.Error())
		return "Err"
	}
	defer resp.Body.Close()
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		//t.Errorf("ReadAll: %v", err.Error())
		return "Err"
	}
	body := string(bodyBytes)
	return body
}

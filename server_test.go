package main

import (
    //"net/url"
    "net/http"
    "testing"
    //"bytes"
    "log"
    "encoding/json"
    "io/ioutil"
    //"strconv"
    //"fmt"
)


func TestProductsHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    //body := []byte("name=Apple&code=12345&description=test&image=thisimage")
    url := "https://storeproject-209402.appspot.com/products?name=Apple&code=testCode12345&description=test&image=thisimage"
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    
    client := http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    defer resp.Body.Close()
    var resultProd Product
    if err := json.NewDecoder(resp.Body).Decode(&resultProd); err != nil {
        log.Println(err)
    }
    expected := `{"name":"Apple","barcode":"testCode12345","description":"test","image":"thisimage"}`
    if resultProd.toJson() != expected {
        t.Errorf("POST product: got %v want %v",
            resultProd.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/products/testCode12345"
    req, err = http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultProd); err != nil {
        log.Println(err)
    }
    if resultProd.toJson() != expected {
        t.Errorf("GET product: got %v want %v",
            resultProd.toJson(), expected)
    }
    url = "https://storeproject-209402.appspot.com/products/testCode12345?name=Banana&description=test2&image=thisimage2"
    req, err = http.NewRequest("PUT", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultProd); err != nil {
        log.Println(err)
    }
    expected = `{"name":"Banana","barcode":"testCode12345","description":"test2","image":"thisimage2"}`
    if resultProd.toJson() != expected {
        t.Errorf("PUT product: got %v want %v",
            resultProd.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/products/testCode12345"
    req, err = http.NewRequest("DELETE", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    expected = `204 OK`
    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        log.Fatal("ReadAll: ", err)
        return
    }
    body := string(bodyBytes)
    if body != expected {
        t.Errorf("DELETE product: got %v want %v",
            body, expected)
    }
    
    
    
}


func TestPricesHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    //body := []byte("name=Apple&code=12345&description=test&image=thisimage")
    url := "https://storeproject-209402.appspot.com/prices?code=testCode12345&price=1.23"
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    
    client := http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    defer resp.Body.Close()
    var resultPrice Price
    if err := json.NewDecoder(resp.Body).Decode(&resultPrice); err != nil {
        log.Println(err)
    }
    expected := `{"barcode":"testCode12345","price":1.23}`
    if resultPrice.toJson() != expected {
        t.Errorf("POST price: got %v want %v",
            resultPrice.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/prices/testCode12345"
    req, err = http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultPrice); err != nil {
        log.Println(err)
    }
    if resultPrice.toJson() != expected {
        t.Errorf("GET price: got %v want %v",
            resultPrice.toJson(), expected)
    }
    url = "https://storeproject-209402.appspot.com/prices/testCode12345?price=2.23"
    req, err = http.NewRequest("PUT", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultPrice); err != nil {
        log.Println(err)
    }
    expected = `{"barcode":"testCode12345","price":2.23}`
    if resultPrice.toJson() != expected {
        t.Errorf("PUT price: got %v want %v",
            resultPrice.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/prices/testCode12345"
    req, err = http.NewRequest("DELETE", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    expected = `204 OK`
    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        log.Fatal("ReadAll: ", err)
        return
    }
    body := string(bodyBytes)
    if body != expected {
        t.Errorf("DELETE price: got %v want %v",
            body, expected)
    }
    
    
    
}




func TestStocksHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    //body := []byte("name=Apple&code=12345&description=test&image=thisimage")
    url := "https://storeproject-209402.appspot.com/stocks?code=testCode12345&stock=1"
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    
    client := http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    defer resp.Body.Close()
    var resultStock Stock
    if err := json.NewDecoder(resp.Body).Decode(&resultStock); err != nil {
        log.Println(err)
    }
    expected := `{"barcode":"testCode12345","stock":1}`
    if resultStock.toJson() != expected {
        t.Errorf("POST stock: got %v want %v",
            resultStock.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
    req, err = http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultStock); err != nil {
        log.Println(err)
    }
    if resultStock.toJson() != expected {
        t.Errorf("GET stock: got %v want %v",
            resultStock.toJson(), expected)
    }
    url = "https://storeproject-209402.appspot.com/stocks/testCode12345?stock=2"
    req, err = http.NewRequest("PUT", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultStock); err != nil {
        log.Println(err)
    }
    expected = `{"barcode":"testCode12345","stock":2}`
    if resultStock.toJson() != expected {
        t.Errorf("PUT stock: got %v want %v",
            resultStock.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/stocks/testCode12345"
    req, err = http.NewRequest("DELETE", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    expected = `204 OK`
    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        log.Fatal("ReadAll: ", err)
        return
    }
    body := string(bodyBytes)
    if body != expected {
        t.Errorf("DELETE stock: got %v want %v",
            body, expected)
    }

}

func TestUsersHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    //body := []byte("name=Apple&code=12345&description=test&image=thisimage")
    url := "https://storeproject-209402.appspot.com/users?name=TestUser&balance=50.02&cart=12345&cart=67890"
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    
    client := http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    defer resp.Body.Close()
    var resultUser User
    if err := json.NewDecoder(resp.Body).Decode(&resultUser); err != nil {
        log.Println(err)
    }
    expected := `{"name":"TestUser","balance":50.02,"cart":["12345","67890"]}`
    if resultUser.toJson() != expected {
        t.Errorf("POST user: got %v want %v",
            resultUser.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/users/TestUser"
    req, err = http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultUser); err != nil {
        log.Println(err)
    }
    if resultUser.toJson() != expected {
        t.Errorf("GET user: got %v want %v",
            resultUser.toJson(), expected)
    }
    url = "https://storeproject-209402.appspot.com/users/TestUser?balance=23.45&cart=654"
    req, err = http.NewRequest("PUT", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    if err := json.NewDecoder(resp.Body).Decode(&resultUser); err != nil {
        log.Println(err)
    }
    expected = `{"name":"TestUser","balance":23.45,"cart":["654"]}`
    if resultUser.toJson() != expected {
        t.Errorf("PUT user: got %v want %v",
            resultUser.toJson(), expected)
    }
    
    url = "https://storeproject-209402.appspot.com/users/TestUser"
    req, err = http.NewRequest("DELETE", url, nil)
    if err != nil {
        log.Fatal("NewRequest: ", err)
        return
    }
    client = http.Client{
		//Timeout: time.Second * 2, // Maximum of 2 secs
	}
    resp, err = client.Do(req)
    if err != nil {
        log.Fatal("Do: ", err)
        return
    }
    expected = `204 OK`
    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        log.Fatal("ReadAll: ", err)
        return
    }
    body := string(bodyBytes)
    if body != expected {
        t.Errorf("DELETE user: got %v want %v",
            body, expected)
    }

}




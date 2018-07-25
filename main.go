// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	"strconv"
)

type Product struct {
	Name        string `json:"name"`
	Barcode     string `json:"barcode"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Price struct {
	Barcode string  `json:"barcode"`
	Price   float64 `json:"price"`
}

type Stock struct {
	Barcode string `json:"barcode"`
	Stock   int64  `json:"stock"`
}

type User struct {
	Name    string   `json:"name"`
	Balance float64  `json:"balance"`
	Cart    []string `json:"cart"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

type PriceList struct {
	Prices []Price `json:"prices"`
}

type StockList struct {
	Stocks []Stock `json:"stocks"`
}

type UserList struct {
	Users []User `json:"users"`
}

type ExtractAttributes struct {
	Request *http.Request
}

func main() {
	//http.HandleFunc("/", handle)
	http.HandleFunc("/test/", testHandle)
	http.HandleFunc("/products", productsHandle)
	http.HandleFunc("/products/", productsHandle)
	http.HandleFunc("/prices", pricesHandle)
	http.HandleFunc("/prices/", pricesHandle)
	http.HandleFunc("/stocks", stocksHandle)
	http.HandleFunc("/stocks/", stocksHandle)
	http.HandleFunc("/users", usersHandle)
	http.HandleFunc("/users/", usersHandle)
	appengine.Main()
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")

}

func productsHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	extract := ExtractAttributes{Request: r}
	url := extract.getURL()
	switch extract.getMethod() {
	case "GET": //GET /products

		if len(extract.getURL()) > len("/products/") { //GET /products/{barcode}
			targetCode := url[len("/products/"):]
			targetKey := datastore.NewKey(ctx, "Product", targetCode, 0, nil)
			var targetProduct Product
			err := datastore.Get(ctx, targetKey, &targetProduct)
			if err != nil {
				log.Errorf(ctx, "GET product/%v: %v", targetCode, err)
			}
			fmt.Fprintf(w, targetProduct.toJson())
			return
		} else {
			q := datastore.NewQuery("Product")
			var products ProductList
			_, err := q.GetAll(ctx, &products.Products)
			if err != nil {
				log.Errorf(ctx, "GET product: %v", err)
			}
			resultJson, jsonErr := json.Marshal(products)
			if jsonErr != nil {
				log.Errorf(ctx, "GET product: %v", err)
				return
			}
			resultString := string(resultJson)
			fmt.Fprintf(w, resultString)
		}
	case "POST": //POST /products
		code := extract.getFormValue("code")
		name := extract.getFormValue("name")
		description := extract.getFormValue("description")
		image := extract.getFormValue("image")
		newProduct := &Product{Barcode: code, Name: name, Description: description, Image: image}
		key := datastore.NewKey(ctx, "Product", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newProduct); err != nil {
			log.Errorf(ctx, "adding product: %v", err)
			return
		}
		fmt.Fprintf(w, (*newProduct).toJson())
	case "PUT": //PUT /products/{barcode}
		code := url[len("/products/"):]
		name := extract.getFormValue("name")
		description := extract.getFormValue("description")
		image := extract.getFormValue("image")
		newProduct := &Product{Barcode: code, Name: name, Description: description, Image: image}
		key := datastore.NewKey(ctx, "Product", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newProduct); err != nil {
			log.Errorf(ctx, "updating product: %v", err)
			return
		}
		fmt.Fprintf(w, (*newProduct).toJson())
	case "DELETE": //DELETE //products{barcode}
		var tempProd Product
		code := url[len("/products/"):]
		key := datastore.NewKey(ctx, "Product", code, 0, nil)
		getErr := datastore.Get(ctx, key, &tempProd)
		if getErr != nil {
			fmt.Fprintf(w, "404 not found")
			return
		}
		err := datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "deleting product: %v", err)
			return
		}
		fmt.Fprintf(w, "204 OK")
	}

}

func pricesHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	extract := ExtractAttributes{Request: r}
	url := extract.getURL()
	switch extract.getMethod() {
	case "GET": //GET /prices

		if len(url) > len("/prices/") { //GET /prices/{barcode}
			targetCode := url[len("/prices/"):]
			targetKey := datastore.NewKey(ctx, "Price", targetCode, 0, nil)
			var targetPrice Price
			err := datastore.Get(ctx, targetKey, &targetPrice)
			if err != nil {
				log.Errorf(ctx, "GET price/%v: %v", targetCode, err)
			}
			fmt.Fprintf(w, targetPrice.toJson())
			return
		} else {
			q := datastore.NewQuery("Price")
			var prices PriceList
			_, err := q.GetAll(ctx, &prices.Prices)
			if err != nil {
				log.Errorf(ctx, "GET price: %v", err)
			}
			resultJson, jsonErr := json.Marshal(prices)
			if jsonErr != nil {
				log.Errorf(ctx, "GET price: %v", err)
				return
			}
			resultString := string(resultJson)
			fmt.Fprintf(w, resultString)
		}
	case "POST": //POST /prices
		code := extract.getFormValue("code")
		price, er := strconv.ParseFloat(extract.getFormValue("price"), 64)
		if er != nil {
			log.Errorf(ctx, "adding price: %v", er)
			return
		}
		newPrice := &Price{Barcode: code, Price: price}
		key := datastore.NewKey(ctx, "Price", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newPrice); err != nil {
			log.Errorf(ctx, "adding price: %v", err)
			return
		}
		fmt.Fprintf(w, (*newPrice).toJson())
	case "PUT": //PUT /prices/{barcode}
		code := url[len("/prices/"):]
		price, er := strconv.ParseFloat(extract.getFormValue("price"), 64)
		if er != nil {
			log.Errorf(ctx, "updating price: %v", er)
			return
		}
		newPrice := &Price{Barcode: code, Price: price}
		key := datastore.NewKey(ctx, "Price", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newPrice); err != nil {
			log.Errorf(ctx, "updating price: %v", err)
			return
		}
		fmt.Fprintf(w, (*newPrice).toJson())
	case "DELETE": //DELETE /prices/{barcode}
		var tempPrice Price
		code := url[len("/prices/"):]
		key := datastore.NewKey(ctx, "Price", code, 0, nil)
		getErr := datastore.Get(ctx, key, &tempPrice)
		if getErr != nil {
			fmt.Fprintf(w, "404 not found")
			return
		}
		err := datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "deleting price: %v", err)
			return
		}
		fmt.Fprintf(w, "204 OK")
	}
}

func stocksHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	extract := ExtractAttributes{Request: r}
	url := extract.getURL()
	switch extract.getMethod() {
	case "GET": //GET /stocks
		q := datastore.NewQuery("Stock")
		var stocks []Stock
		_, err := q.GetAll(ctx, &stocks)
		if err != nil {
			log.Errorf(ctx, "fetching stocks: %v", err)
			return
		}
		if len(url) > len("/stocks/") { //GET /stocks/{barcode}
			targetCode := url[len("/stocks/"):]
			targetKey := datastore.NewKey(ctx, "Stock", targetCode, 0, nil)
			var targetStock Stock
			err := datastore.Get(ctx, targetKey, &targetStock)
			if err != nil {
				log.Errorf(ctx, "GET stock/%v: %v", targetCode, err)
			}
			fmt.Fprintf(w, targetStock.toJson())
			return
		} else {
			q := datastore.NewQuery("Stock")
			var stocks StockList
			_, err := q.GetAll(ctx, &stocks.Stocks)
			if err != nil {
				log.Errorf(ctx, "fetching stocks: %v", err)
				return
			}
			resultJson, jsonErr := json.Marshal(stocks)
			if jsonErr != nil {
				log.Errorf(ctx, "GET stock: %v", err)
				return
			}
			resultString := string(resultJson)
			fmt.Fprintf(w, resultString)
		}
	case "POST": //POST /stocks
		code := extract.getFormValue("code")
		stock, er := strconv.ParseInt(extract.getFormValue("stock"), 10, 64)
		if er != nil {
			log.Errorf(ctx, "adding stock: %v", er)
			return
		}
		newStock := &Stock{Barcode: code, Stock: stock}
		key := datastore.NewKey(ctx, "Stock", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newStock); err != nil {
			log.Errorf(ctx, "adding stock: %v", err)
			return
		}
		fmt.Fprintf(w, (*newStock).toJson())
	case "PUT": //PUT /stocks/{barcode}
		code := url[len("/stocks/"):]
		stock, er := strconv.ParseInt(extract.getFormValue("stock"), 10, 64)
		if er != nil {
			log.Errorf(ctx, "adding stock: %v", er)
			return
		}
		newStock := &Stock{Barcode: code, Stock: stock}
		key := datastore.NewKey(ctx, "Stock", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newStock); err != nil {
			log.Errorf(ctx, "adding stock: %v", err)
			return
		}
		fmt.Fprintf(w, (*newStock).toJson())
	case "DELETE": //DELETE /stocks/{barcode}
		var tempStock Stock
		code := url[len("/stocks/"):]
		key := datastore.NewKey(ctx, "Stock", code, 0, nil)
		getErr := datastore.Get(ctx, key, &tempStock)
		if getErr != nil {
			fmt.Fprintf(w, "404 not found")
			return
		}
		err := datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "deleting stock: %v", err)
			return
		}
		fmt.Fprintf(w, "204 OK")
	}
}

func usersHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	extract := ExtractAttributes{Request: r}
	url := extract.getURL()
	switch extract.getMethod() {
	case "GET": //GET /users

		if len(url) > len("/users/") { //GET /users/{name}
			targetName := url[len("/users/"):]
			targetKey := datastore.NewKey(ctx, "User", targetName, 0, nil)
			var targetUser User
			err := datastore.Get(ctx, targetKey, &targetUser)
			if err != nil {
				log.Errorf(ctx, "GET user/%v: %v", targetName, err)
			}
			fmt.Fprintf(w, targetUser.toJson())
			return
		} else {
			q := datastore.NewQuery("User")
			var users UserList
			_, err := q.GetAll(ctx, &users.Users)
			if err != nil {
				log.Errorf(ctx, "fetching users: %v", err)
				return
			}
			resultJson, jsonErr := json.Marshal(users)
			if jsonErr != nil {
				log.Errorf(ctx, "GET product: %v", err)
				return
			}
			resultString := string(resultJson)
			fmt.Fprintf(w, resultString)
		}
	case "POST": //POST /users
		name := extract.getFormValue("name")
		balance, er := strconv.ParseFloat(extract.getFormValue("balance"), 64)
		cart := extract.getForm("cart")
		if er != nil {
			log.Errorf(ctx, "adding user: %v", er)
			return
		}
		newUser := &User{Name: name, Balance: balance, Cart: cart}
		key := datastore.NewKey(ctx, "User", name, 0, nil)
		if _, err := datastore.Put(ctx, key, newUser); err != nil {
			log.Errorf(ctx, "adding user: %v", err)
			return
		}
		fmt.Fprintf(w, (*newUser).toJson())
	case "PUT": //PUT /users/{name}
		name := url[len("/users/"):]
		balance, er := strconv.ParseFloat(extract.getFormValue("balance"), 64)
		cart := extract.getForm("cart")
		if er != nil {
			log.Errorf(ctx, "adding user: %v", er)
			return
		}
		newUser := &User{Name: name, Balance: balance, Cart: cart}
		key := datastore.NewKey(ctx, "User", name, 0, nil)
		if _, err := datastore.Put(ctx, key, newUser); err != nil {
			log.Errorf(ctx, "adding user: %v", err)
			return
		}
		fmt.Fprintf(w, (*newUser).toJson())
	case "DELETE": //DELETE /users/{name}
		var tempUser User
		name := url[len("/users/"):]
		key := datastore.NewKey(ctx, "User", name, 0, nil)
		getErr := datastore.Get(ctx, key, &tempUser)
		if getErr != nil {
			fmt.Fprintf(w, "404 not found")
			return
		}
		err := datastore.Delete(ctx, key)
		if err != nil {
			log.Errorf(ctx, "deleting user: %v", err)
			return
		}
		fmt.Fprintf(w, "204 OK")
	}
}

func (p Product) toJson() string {
	resultJson, jsonErr := json.Marshal(p)
	if jsonErr != nil {
		return ""
	}
	resultString := string(resultJson)
	return resultString
}

func (p Price) toJson() string {
	resultJson, jsonErr := json.Marshal(p)
	if jsonErr != nil {
		return ""
	}
	resultString := string(resultJson)
	return resultString
}

func (p Stock) toJson() string {
	resultJson, jsonErr := json.Marshal(p)
	if jsonErr != nil {
		return ""
	}
	resultString := string(resultJson)
	return resultString
}

func (u User) toJson() string {
	resultJson, jsonErr := json.Marshal(u)
	if jsonErr != nil {
		return ""
	}
	resultString := string(resultJson)
	return resultString
}

func (e ExtractAttributes) getMethod() string {
	return e.Request.Method
}

func (e ExtractAttributes) getFormValue(s string) string {
	return e.Request.FormValue(s)
}

func (e ExtractAttributes) getForm(s string) []string {
	return e.Request.Form[s]
}

func (e ExtractAttributes) getURL() string {
	return e.Request.URL.Path
}

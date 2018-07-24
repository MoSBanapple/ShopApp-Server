// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	"strconv"
)

type Product struct {
	Name        string
	Barcode     string
	Description string
	Image       string
}

type Price struct {
	Barcode string
	Price   float64
}

type Stock struct {
	Barcode string
	Stock   int64
}

type User struct {
	Name    string
	Balance float64
	Cart    []string
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

/*
func handle(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintln(w, "home")

}
*/

func testHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)

}

func productsHandle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	switch r.Method {
	case "GET": //GET /products

		if len(r.URL.Path) > len("/products/") { //GET /products/{barcode}
			targetCode := r.URL.Path[len("/products/"):]
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
			var products []Product
			_, err := q.GetAll(ctx, &products)
			if err != nil {
				log.Errorf(ctx, "GET product: %v", err)
			}
			output := "{\"products\":["
			for i, targetProduct := range products {
				output += targetProduct.toJson()
				if i < len(products)-1 {
					output += ","
				}
			}
			output += "]}"
			fmt.Fprintf(w, output)
		}
	case "POST": //POST /products
		code := r.FormValue("code")
		name := r.FormValue("name")
		description := r.FormValue("description")
		image := r.FormValue("image")
		newProduct := &Product{Barcode: code, Name: name, Description: description, Image: image}
		key := datastore.NewKey(ctx, "Product", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newProduct); err != nil {
			log.Errorf(ctx, "adding product: %v", err)
			return
		}
		fmt.Fprintf(w, (*newProduct).toJson())
	case "PUT": //PUT /products/{barcode}
		code := r.URL.Path[len("/products/"):]
		name := r.FormValue("name")
		description := r.FormValue("description")
		image := r.FormValue("image")
		newProduct := &Product{Barcode: code, Name: name, Description: description, Image: image}
		key := datastore.NewKey(ctx, "Product", code, 0, nil)
		if _, err := datastore.Put(ctx, key, newProduct); err != nil {
			log.Errorf(ctx, "updating product: %v", err)
			return
		}
		fmt.Fprintf(w, (*newProduct).toJson())
	case "DELETE": //DELETE //products{barcode}
		var tempProd Product
		code := r.URL.Path[len("/products/"):]
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
	switch r.Method {
	case "GET": //GET /prices

		if len(r.URL.Path) > len("/prices/") { //GET /prices/{barcode}
			targetCode := r.URL.Path[len("/prices/"):]
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
			var prices []Price
			_, err := q.GetAll(ctx, &prices)
			if err != nil {
				log.Errorf(ctx, "GET price: %v", err)
			}
			output := "{\"prices\":["
			for i, targetPrice := range prices {
				output += targetPrice.toJson()
				if i < len(prices)-1 {
					output += ","
				}
			}
			output += "]}"
			fmt.Fprintf(w, output)
		}
	case "POST": //POST /prices
		code := r.FormValue("code")
		price, er := strconv.ParseFloat(r.FormValue("price"), 64)
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
		code := r.URL.Path[len("/prices/"):]
		price, er := strconv.ParseFloat(r.FormValue("price"), 64)
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
		code := r.URL.Path[len("/prices/"):]
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
	switch r.Method {
	case "GET": //GET /stocks
		q := datastore.NewQuery("Stock")
		var stocks []Stock
		_, err := q.GetAll(ctx, &stocks)
		if err != nil {
			log.Errorf(ctx, "fetching stocks: %v", err)
			return
		}
		if len(r.URL.Path) > len("/stocks/") { //GET /stocks/{barcode}
			targetCode := r.URL.Path[len("/stocks/"):]
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
			var stocks []Stock
			_, err := q.GetAll(ctx, &stocks)
			if err != nil {
				log.Errorf(ctx, "fetching stocks: %v", err)
				return
			}
			output := "{\"stocks\":["
			for i, targetStock := range stocks {
				output += targetStock.toJson()
				if i < len(stocks)-1 {
					output += ","
				}
			}
			output += "]}"
			fmt.Fprintf(w, output)
		}
	case "POST": //POST /stocks
		code := r.FormValue("code")
		stock, er := strconv.ParseInt(r.FormValue("stock"), 10, 64)
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
		code := r.URL.Path[len("/stocks/"):]
		stock, er := strconv.ParseInt(r.FormValue("stock"), 10, 64)
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
		code := r.URL.Path[len("/stocks/"):]
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
	switch r.Method {
	case "GET": //GET /users

		if len(r.URL.Path) > len("/users/") { //GET /users/{name}
			targetName := r.URL.Path[len("/users/"):]
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
			var users []User
			_, err := q.GetAll(ctx, &users)
			if err != nil {
				log.Errorf(ctx, "fetching users: %v", err)
				return
			}
			output := "{\"users\":["
			for i, targetUser := range users {
				output += targetUser.toJson()
				if i < len(users)-1 {
					output += ","
				}
			}
			output += "]}"
			fmt.Fprintf(w, output)
		}
	case "POST": //POST /users

		if parseErr := r.ParseForm(); parseErr != nil {
			log.Errorf(ctx, "adding user: %v", parseErr)
		}
		name := r.FormValue("name")
		balance, er := strconv.ParseFloat(r.FormValue("balance"), 64)
		cart := r.Form["cart"]
		//fmt.Fprintf(w, name[0])
		//fmt.Fprintf(w, "test")
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
		if parseErr := r.ParseForm(); parseErr != nil {
			log.Errorf(ctx, "adding user: %v", parseErr)
		}
		name := r.URL.Path[len("/users/"):]
		balance, er := strconv.ParseFloat(r.FormValue("balance"), 64)
		cart := r.Form["cart"]
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
		name := r.URL.Path[len("/users/"):]
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
	output := "{\"name\":\""
	output += p.Name
	output += "\",\"barcode\":\""
	output += p.Barcode
	output += "\",\"description\":\""
	output += p.Description
	output += "\",\"image\":\""
	output += p.Image
	output += "\"}"
	return output
}

func (p Price) toJson() string {
	output := "{\"barcode\":\""
	output += p.Barcode
	output += "\",\"price\":"
	output += strconv.FormatFloat(p.Price, 'f', 2, 64)
	output += "}"
	return output
}

func (p Stock) toJson() string {
	output := "{\"barcode\":\""
	output += p.Barcode
	output += "\",\"stock\":"
	output += strconv.FormatInt(p.Stock, 10)
	output += "}"
	return output
}

func (u User) toJson() string {
	output := "{\"name\":\""
	output += u.Name
	output += "\",\"balance\":"
	output += strconv.FormatFloat(u.Balance, 'f', 2, 64)
	output += ",\"cart\":["
	for i, code := range u.Cart {
		output += "\"" + code + "\""
		if i < len(u.Cart)-1 {
			output += ","
		}
	}
	output += "]}"
	return output
}

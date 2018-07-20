// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
    "strconv"
	"google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
    "google.golang.org/appengine/log"
)

type Product struct {
    Name string
    Barcode string
    Description string
    Image string
}

type Price struct {
    Barcode string
    Price float64
}

type Stock struct {
    Barcode string
    Stock int64
}

type User struct {
    Name string
    Balance float64
    Cart []string
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
    switch r.Method{
    case "GET":
        q := datastore.NewQuery("Product")
        var products []Product
        keys, err := q.GetAll(ctx, &products)
        if err != nil {
            log.Errorf(ctx, "fetching products: %v", err)
            return
        }
        if len(r.URL.Path) > len("/products/") {
            targetCode := r.URL.Path[len("/products/"):]
            var targetProduct Product
            for i, _ := range keys {
                if products[i].Barcode == targetCode {
                    targetProduct = products[i]
                    break
                }
            }
            
            fmt.Fprintf(w, targetProduct.toJson())
            return
        } else {
            output := "{\"products\":["
            for i, targetProduct := range products {
                output += targetProduct.toJson()
                if (i < len(products) - 1){
                    output += ","
                }
            }
            output += "]}"
            fmt.Fprintf(w, output)
        }
    case "POST":
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
    case "PUT":
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
    case "DELETE":
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
    switch r.Method{
    case "GET":
        q := datastore.NewQuery("Price")
        var prices []Price
        _, err := q.GetAll(ctx, &prices)
        if err != nil {
            log.Errorf(ctx, "fetching prices: %v", err)
            return
        }
        if len(r.URL.Path) > len("/prices/") {
            targetCode := r.URL.Path[len("/prices/"):]
            var targetPrice Price
            for _, currentPrice := range prices {
                if currentPrice.Barcode == targetCode {
                    targetPrice = currentPrice
                    break
                }
            }
            
            fmt.Fprintf(w, targetPrice.toJson())
            return
        } else {
            output := "{\"prices\":["
            for i, targetPrice := range prices {
                output += targetPrice.toJson()
                if (i < len(prices) - 1){
                    output += ","
                }
            }
            output += "]}"
            fmt.Fprintf(w, output)
        }
    case "POST":
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
    case "PUT":
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
    case "DELETE":
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
    switch r.Method{
    case "GET":
        q := datastore.NewQuery("Stock")
        var stocks []Stock
        _, err := q.GetAll(ctx, &stocks)
        if err != nil {
            log.Errorf(ctx, "fetching stocks: %v", err)
            return
        }
        if len(r.URL.Path) > len("/stocks/") {
            targetCode := r.URL.Path[len("/stocks/"):]
            var targetStock Stock
            for _, currentStock := range stocks {
                if currentStock.Barcode == targetCode {
                    targetStock = currentStock
                    break
                }
            }
            
            fmt.Fprintf(w, targetStock.toJson())
            return
        } else {
            output := "{\"stocks\":["
            for i, targetStock := range stocks {
                output += targetStock.toJson()
                if (i < len(stocks) - 1){
                    output += ","
                }
            }
            output += "]}"
            fmt.Fprintf(w, output)
        }
    case "POST":
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
    case "PUT":
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
    case "DELETE":
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
    switch r.Method{
    case "GET":
        q := datastore.NewQuery("User")
        var users []User
        _, err := q.GetAll(ctx, &users)
        if err != nil {
            log.Errorf(ctx, "fetching users: %v", err)
            return
        }
        if len(r.URL.Path) > len("/users/") {
            targetName := r.URL.Path[len("/users/"):]
            var targetUser User
            for _, currentUser := range users {
                if currentUser.Name == targetName {
                    targetUser = currentUser
                    break
                }
            }
            
            fmt.Fprintf(w, targetUser.toJson())
            return
        } else {
            output := "{\"users\":["
            for i, targetUser := range users {
                output += targetUser.toJson()
                if (i < len(users) - 1){
                    output += ","
                }
            }
            output += "]}"
            fmt.Fprintf(w, output)
        }
    case "POST":
        
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
    case "PUT":
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
    case "DELETE":
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
        if i < len(u.Cart) - 1 {
            output += ","
        }
    }
    output += "]}"
    return output
}









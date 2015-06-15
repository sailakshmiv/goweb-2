package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
    r.ParseForm()  // parse, default will not do it. 
    fmt.Println(r.Form)
    fmt.Println("Path: ", r.URL.Path)
    fmt.Println("scheme: ", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key: ", k)
	fmt.Println("val: ", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello chris")  // to client
}

func login(w http.ResponseWriter, r *http.Request){
    fmt.Println("Method: ", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, nil)
    } else {
	r.ParseForm()
        // do some ligin logically
	fmt.Println("username; ", r.Form["username"])
	fmt.Println("password: ", r.Form["password"])
    }
}

func main(){
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
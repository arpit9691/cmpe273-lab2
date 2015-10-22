package main

import (
	"encoding/json"
	"fmt"
	//"github.com/julienschmidt/httprouter"
	"httprouter"
	"net/http"
)

type request struct {
	Name string `json:"name"`
}

type response struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloWorld(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	//fmt.Println(rw, "Hello foo!!!!!!!!!")
	decoder := json.NewDecoder(req.Body)
	var rqst request
	var resp response
	decoder.Decode(&rqst)
	resp.Greeting = "Hello " + rqst.Name + " !!!!!"
	json.NewEncoder(rw).Encode(resp)
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", helloWorld)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

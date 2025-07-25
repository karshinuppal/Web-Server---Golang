package main

import (
	"fmt"
	"log"
	"net/http"
)

// request is what we send to the server
// and response is what the user sends back
func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound )
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusNotFound )
		return
	}
	fmt.Fprintf(w, "Hii!!")
}


func formHandler(w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("STarting Server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
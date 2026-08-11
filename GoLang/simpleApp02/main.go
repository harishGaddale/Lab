package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Container World!! Your GO app is deployed and is running successfully")
}

func main(){
	http.HandleFunc("/",handler)
	fmt.Println("Server is starting on port 8852")
	if err := http.ListenAndServe(":8852", nil); err != nil{
		panic(err)
	}
}

// package main

// import "fmt"

//	func main() {
//	    fmt.Println("Hello, World!")
//	}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"rsc.io/quote"
)

func health(w http.ResponseWriter, r *http.Request)  {
	log.Println("testing the health of server")
	response := map[string]string{
		"status":"Ok",
		"timestamp" : time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Println(quote.Go())
        // fmt.Fprintf(w, "use like url?token=value : %s\n", r.URL.Query().Get("token"))
		
    })
	
	http.HandleFunc("/health", health)
	log.Printf("web server is started At 1234...")
    http.ListenAndServe(":80", nil)
}
// adding this comment to check wether the github runner event triggers or not
// adding this comment to check workflow triggers or not
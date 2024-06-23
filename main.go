package main 

import (
    "fmt"
    // create http client and server
    "net/http"
)


// handler that will be called when the server receives a request
// w is the response writer
// r is the request
func basicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}


func main() {
    server := &http.Server{
        // where is the server listening(port)
        Addr: ":8080",
        // handler is the interface that must be implemented
        Handler: http.HandlerFunc(basicHandler),
    }

    // listen and serve
    err := server.ListenAndServe()
    // handle error from running the server
    if err != nil {
        fmt.Println("Error starting server: ", err)
    }
}

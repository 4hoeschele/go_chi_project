package main 

import (
    "fmt"
    // create http client and server
    "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


// handler that will be called when the server receives a request
// w is the response writer
// r is the request
func basicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func main() {
    // initialize new router with chi
    router := chi.NewRouter()
	// add middleware to the router
	router.Use(middleware.Logger)

	// add a route to the router
	router.Get("/hello", basicHandler)
    
    server := &http.Server{
        // where is the server listening(port)
        Addr: ":8080",
        // handler is the interface that must be implemented
        Handler: router,
    }

    // listen and serve
    err := server.ListenAndServe()
    // handle error from running the server
    if err != nil {
        fmt.Println("Error starting server: ", err)
    }
}

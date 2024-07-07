package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// command line arguments
	address = flag.String("address", "0.0.0.0", "Listening address")
	port    = flag.String("port", "80", "Listening port")
	sslPort = flag.String("sslPort", "443", "SSL listening port")
	status  = flag.Int("status", 200, "Returned HTTP status code")
	cert    = flag.String("cert", "cert.pem", "SSL certificate path")
	key     = flag.String("key", "key.pem", "SSL private Key path")
)

func createHandler(body string) http.HandlerFunc {
	// return appropriate handler for the given body
	return func(w http.ResponseWriter, r *http.Request) {
		if fi, err := os.Stat(body); err == nil {
			switch {
			case fi.IsDir():
				http.FileServer(http.Dir(body)).ServeHTTP(w, r)
				return
			case fi.Mode().IsRegular():
				http.ServeFile(w, r, body)
				return
			}
		}
		// If body isn't a file or directory, write it as is
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(*status)
		fmt.Fprint(w, body)
	}
}

func main() {
	flag.Parse()
	listen := *address + ":" + *port
	listenTLS := *address + ":" + *sslPort
	body := flag.Arg(0)
	if body == "" {
		body = "Hello World!"
	}

	handler := createHandler(body)

	go func() {
		// Listen on tls port if certificate and key are available
		if _, err := os.Stat(*cert); err != nil {
			return
		}
		if _, err := os.Stat(*key); err != nil {
			return
		}
		log.Println("Listening on", listenTLS)
		log.Fatal(http.ListenAndServeTLS(listenTLS, *cert, *key, handler))
	}()

	log.Println("Listening on", listen)
	log.Fatal(http.ListenAndServe(listen, handler))
}
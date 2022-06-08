package main

import (
	"fmt"
	"net/http"
	"flag"
)
var port = flag.Int("p", 8080, "server port")

func main() {
	flag.Parse()
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
        if len(r.URL.Path[1:]) == 0 {
                fmt.Fprintf(w, "Hello, Uccps!")
        }else{
                fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
        }
}

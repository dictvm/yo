package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	servePtr := flag.String("serveYo", "", "run as server on port")
	//portPtr := flag.String("port", "8080", "the port where YO should run")
	flag.Parse()

	if *servePtr == "" {
		yo := make(chan string)
		go func() {
			yo <- "Yo"
		}()
		println(<-yo)
	} else {
		http.HandleFunc("/", yo)
		http.ListenAndServe(":"+*servePtr, nil)
	}
}

func yo(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.RequestURI == "/yo":
		fmt.Fprint(w, "YOLO")
	case r.RequestURI == "/favicon.ico":
		// do something!
	default:
		fmt.Fprint(w, "Yo")
	}
}

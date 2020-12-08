package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8084", "http service address")
	flag.Parse()
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)
	staticHandler := http.StripPrefix("/web", http.FileServer(http.Dir("./web")))
	http.Handle("/web", staticHandler)

	http.HandleFunc("/api", apiHandler)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

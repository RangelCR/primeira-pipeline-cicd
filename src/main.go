package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Pantera foi uma influente banda americana de heavy metal formada em 1981 no Texas. CFH")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}

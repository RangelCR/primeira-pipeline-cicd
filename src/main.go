package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Pantera foi uma influente banda americana de heavy metal formada em 1981 no Texas. Inicialmente com um estilo glam, ganhou destaque mundial nos anos 1990 ao adotar um som mais agressivo e pesado, conhecido como groove metal. Com Phil Anselmo nos vocais e os irmãos Dimebag Darrell (guitarra) e Vinnie Paul (bateria), o Pantera lançou álbuns marcantes como Cowboys from Hell e Vulgar Display of Power, tornando-se um dos nomes mais importantes do metal moderno.")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}

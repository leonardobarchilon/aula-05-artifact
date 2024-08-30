package main
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":13000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
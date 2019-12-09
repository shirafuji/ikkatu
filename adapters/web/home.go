package web

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r http.Request) {
	fmt.Fprintf(w, "hello world")
}

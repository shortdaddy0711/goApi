package main

import (
	"net/http"

	"github.com/shortdaddy0711/goAPI/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
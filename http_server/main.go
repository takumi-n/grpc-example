package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		op := req.URL.Query().Get("op")
		x, _ := strconv.ParseFloat(req.URL.Query().Get("x"), 64)
		y, _ := strconv.ParseFloat(req.URL.Query().Get("y"), 64)
		var result float64
		switch op {
		case "add":
			result = x + y
		case "sub":
			result = x - y
		case "mul":
			result = x * y
		case "div":
			result = x / y
		}

		fmt.Fprintf(w, "%f", result)
	})

	http.ListenAndServe(":8888", nil)
}

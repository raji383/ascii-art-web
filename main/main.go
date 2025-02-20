package main

import (
	"fmt"
	"net/http"
	"fs"
)

func main() {
	http.HandleFunc("/css/", fs.Css)
	http.HandleFunc("/", fs.HomeHandler)

	http.HandleFunc("/ascii-art", fs.Finaldrawing)
	http.HandleFunc("/export", fs.ExportAsciiArt)
	
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
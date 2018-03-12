package main

import (
	"bufio"
	"fmt"
	"html"
	"net/http"
	"os"
)

type Handler struct {
	target string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	f, err := os.Open(h.target)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>internal error<h1>")
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	fmt.Fprintf(w, "<html>")
	for s.Scan() {
		line := html.EscapeString(s.Text())
		fmt.Fprintf(w, line)
	}
	fmt.Fprintf(w, "</html>")
}

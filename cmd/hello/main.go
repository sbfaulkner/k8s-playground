package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Query: %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)

	fmt.Fprintln(w, "\nREQUEST HEADERS:")
	keys := make([]string, 0, len(r.Header))
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintf(w, "%s: %s\n", k, strings.Join(r.Header.Values(k), ", "))
	}

	fmt.Fprintln(w, "\nBODY:")
	s := bufio.NewScanner(r.Body)
	empty := true

	for s.Scan() {
		fmt.Fprintln(w, s.Text())
		empty = false
	}

	if empty {
		fmt.Fprintln(w, "(empty)")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

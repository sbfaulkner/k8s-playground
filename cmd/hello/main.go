package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/rs/zerolog/log"
)

var statsdClient *statsd.Client

func init() {
	client, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create statsd client")
	}

	statsdClient = client
}

func handler(w http.ResponseWriter, r *http.Request) {
	statsdClient.Count("hello.requests", 1, nil, 1)

	request := map[string]interface{}{
		"RemoteAddr":    r.RemoteAddr,
		"RequestMethod": r.Method,
		"RequestHost":   r.Host,
		"RequestPath":   r.URL.Path,
		"RequestQuery":  r.URL.RawQuery,
		"RequestProto":  r.Proto,
		"RequestURI":    r.RequestURI,
	}

	log.Info().Fields(r.Header).Fields(request).Msg("Request received")

	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestMethod: %s\n", r.Method)
	fmt.Fprintf(w, "RequestHost: %s\n", r.Host)
	fmt.Fprintf(w, "RequestPath: %s\n", r.URL.Path)
	fmt.Fprintf(w, "RequestQuery: %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "RequestProto: %s\n", r.Proto)
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

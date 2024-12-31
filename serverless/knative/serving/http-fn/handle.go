package function

import (
	"fmt"
	"net/http"
	"strings"
)

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */

	result := prettyPrint(r)

	fmt.Println("Received request")
	fmt.Println(result)
	fmt.Fprintf(w, result)
}

func prettyPrint(r *http.Request) string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "%v %v %v %v\n", r.Method, r.URL, r.Proto, r.Host)

	for k, vv := range r.Header {
		for _, v := range vv {
			fmt.Fprintf(sb, "%v: %v\n", k, v)
		}
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		fmt.Fprintf(sb, "Body: ")

		for k, v := range r.Form {
			fmt.Fprintf(sb, "%v: %v\n", k, v)
		}
	}
	return sb.String()
}

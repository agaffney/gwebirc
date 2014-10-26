package core

import (
	"fmt"
	"net/http"
	"strings"
)

func api_handler(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimPrefix(r.URL.Path, "/api/")
	url_parts := make([]string, 0)
	// Filter out any blank items due to doubled slashes
	for _, part := range strings.Split(url, "/") {
		if part != "" {
			url_parts = append(url_parts, part)
		}
	}
	fmt.Printf("url_parts is %v\n", url_parts)
}

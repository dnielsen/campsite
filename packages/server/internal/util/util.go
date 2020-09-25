package util

import (
	"net/http"
	"strings"
)

func SetIfNotEmpty(m map[string]string, key, val string) {
	if val != "" {
		m[key] = val
	}
}

// Returns the host of the request (without the port).
func GetHost(r *http.Request) string {
	if r.URL.IsAbs() {
		host := r.Host
		// Slice off any port information.
		if i := strings.Index(host, ":"); i != -1 {
			host = host[:i]
		}
		return host
	}
	return r.URL.Host
}
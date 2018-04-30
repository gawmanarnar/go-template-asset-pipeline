package main

import (
	"net/http"
	"strings"
)

// CacheAssets - Middleware to cache static assets in dist
func CacheAssets(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only cache assets in dist
		path := r.URL.Path
		if !strings.HasPrefix(path, "/dist") {
			next.ServeHTTP(w, r)
			return
		}

		// Cache for 3 months (7884000 seconds)
		w.Header().Set("Cache-Control", "public, max-age=7884000")
		next.ServeHTTP(w, r)
	})
}

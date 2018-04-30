package main

import (
	"net/http"
	"strings"
)

// CacheAssets - Middleware for caching static assets
type CacheAssets struct {
}

// NewCacheAssets - Creates a new CacheAssets middleware
func NewCacheAssets() *CacheAssets {
	return &CacheAssets{}
}

// ServeHttp -  Middleware for caching static assets
func (s *CacheAssets) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Only cache assets in dist
	path := r.URL.Path
	if !strings.HasPrefix(path, "/dist") {
		next(w, r)
		return
	}

	// Cache for 3 months (7884000 seconds)
	w.Header().Set("Cache-Control", "public, max-age=7884000")
	next(w, r)
}

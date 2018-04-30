package main

import (
	"encoding/json"
	"net/http"

	"github.com/gimmeasandwich/go-template-asset-pipeline/views"
	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr"
)

var index *views.View

var assetBox = packr.NewBox("./dist")
var templateBox = packr.NewBox("./templates")

func main() {
	r := chi.NewRouter()

	// Serve static assets
	CreateFileServer(r, "/dist/")

	// Read manifest to map assets to their hashed names
	var manifest map[string]interface{}
	err := json.Unmarshal(assetBox.Bytes("manifest.json"), &manifest)
	if err != nil {
		panic(err)
	}
	views.SetManifest(manifest)

	index = views.NewView("index", templateBox.String("index.gohtml"))
	r.Get("/", index.Serve)

	http.ListenAndServe(":3000", CacheAssets(r))
}

// CreateFileServer - serves our static files
func CreateFileServer(r chi.Router, path string) {
	fs := http.StripPrefix(path, http.FileServer(assetBox))
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

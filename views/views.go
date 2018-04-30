package views

import (
	"html/template"
	"net/http"
)

var manifest map[string]interface{}

// SetManifest - Set manifest for getting asset hashed names
func SetManifest(m map[string]interface{}) {
	manifest = m
}

// NewView - Creates a new view containing a template for execution
func NewView(name, file string) *View {
	fmap := template.FuncMap{
		"assetPath": func(asset string) string {
			return "/dist/" + manifest[asset].(string)
		},
	}
	t, err := template.New(name).Funcs(fmap).Parse(file)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View - Represents a frontend view
type View struct {
	Template *template.Template
}

// Serve - Serves the view
func (v *View) Serve(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

// Render - Renders the view
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	if err := v.Template.Execute(w, data); err != nil {
		panic(err)
	}
}

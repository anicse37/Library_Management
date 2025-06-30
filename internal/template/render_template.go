package template

import (
	"bytes"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("internal/template/static/" + tmpl)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = buf.WriteTo(w)
}

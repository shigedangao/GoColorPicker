package template

import (
	"net/http"
)

type Template string
type view map[string]Template

// Home view
func (v view) home(w http.ResponseWriter, r *http.Request) {

}

// Save view
func (v view) edit(w http.ResponseWriter, r *http.Request) {

}

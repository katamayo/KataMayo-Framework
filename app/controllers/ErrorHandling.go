package controllers

import (
	"net/http"
)

func MyNotFound(w http.ResponseWriter, r *http.Request) {

	tpls := []string{"views/layouts/app.html", "views/layouts/partial.html", "views/error404.html"}
	rnd.Template(w, http.StatusNotFound, tpls, nil)
}

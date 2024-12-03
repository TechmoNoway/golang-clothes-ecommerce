package main

import "net/http"

func (app *application) healthCheckHanler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

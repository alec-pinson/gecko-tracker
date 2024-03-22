package main

import (
	"fmt"
	"net/http"
)

func apiNextLay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, toJson(GetNextLayDateInfo()))
}

func apiNextHatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, toJson(GetNextHatchDateInfo()))
}

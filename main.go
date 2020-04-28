package main

import (
	"accounting/accounting/routing"
	"net/http"
)

func main() {
	router := routing.NewRouter()
	router.Route()
	http.ListenAndServe(":3000", router)
}

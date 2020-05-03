package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &router{make(map[string]map[string]HandlerFunc)}

	r.HandleFunc("GET", "/", logHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "welcome")
	}))

	r.HandleFunc("GET", "/about", logHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	}))

	r.HandleFunc("GET", "/users/:id", logHandler(recoverHandler(func(c *Context) {
		if c.Params["id"] == "0" {
			panic("id is zero")
		}
		fmt.Fprintf(c.ResponseWriter, "retrieve user %v\n", c.Params["id"])
	})))

	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", logHandler(recoverHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "retrieve user %v's address %v\n", c.Params["user_id"], c.Params["address_id"])
	})))

	r.HandleFunc("POST", "/users", logHandler(recoverHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	}))))

	r.HandleFunc("POST", "/users/:user_id/addresses", logHandler(recoverHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "create user %v's address", c.Params["user_id"])
	})))

	http.ListenAndServe(":8080", r)

}

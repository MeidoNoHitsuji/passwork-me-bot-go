package main

import "net/http"

type Client struct {
	Cookies   []http.Cookie
	id        int
	Workspace string
	Csrf      string
	Private   string
}

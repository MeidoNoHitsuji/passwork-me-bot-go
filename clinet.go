package main

import "net/http"

type Client struct {
	Headers   map[string]string
	Cookies   http.CookieJar
	id        int
	Workspace string
	Csrf      string
	Private   string
}

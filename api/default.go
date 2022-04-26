package api

import (
	"net/http"
)

type DefaultApi struct {
	client    http.Client
	Workspace string
	Csrf      string
	Url       string
}

// func (s *DefaultApi) Request(method string, url string, data map[string]string) {
// 	switch method {
// 	case "Get":
// 		resp, err := s.client.Get(url)
// 		if err != nil {
// 			log.Fatal("Error loading .env file")
// 		}

// 		if s.client.Jar == nil {
// 			s.client.Jar.SetCookies()
// 		}

// 	}

// }

// func (s *DefaultApi) Get() {

// }

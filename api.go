package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DefaultApi struct {
	Cookies   []http.Cookie
	Workspace string
	Csrf      string
	Url       string
}

type UserApi struct {
	d DefaultApi
}

func (s *DefaultApi) Request(method string, url string, data map[string]interface{}) (map[string]interface{}, error) {
	client := http.Client{}

	var body io.Reader

	if data != nil {
		marshal, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(marshal))
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s/%s", URL, s.Url, url), body)
	if err != nil {
		return nil, err
	}

	if s.Cookies != nil && len(s.Cookies) > 0 {
		for _, q := range s.Cookies {
			req.AddCookie(&q)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	if s.Cookies == nil || len(s.Cookies) == 0 {
		for _, c := range resp.Cookies() {
			s.Cookies = append(s.Cookies, *c)
		}
	}

	return result, nil
}

//func (s *UserApi) Authorize(login string, password string) {
//	s.d.Request('POST')
//}

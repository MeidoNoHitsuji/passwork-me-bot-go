package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DefaultApi struct {
	Cookies   []http.Cookie
	Workspace string
	Csrf      string
}

type UserApi struct {
	d *DefaultApi
}

func (s *DefaultApi) Request(method string, url1 string, data map[string]interface{}) (map[string]interface{}, error) {
	client := http.Client{}

	// Назначаем параметры
	buffer := new(bytes.Buffer)
	params := url.Values{}
	for key, value := range data {
		params.Set(key, value.(string))
	}

	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", URL, url1), buffer)

	if err != nil {
		return nil, err
	}

	// Добавляем необходимые заголовки
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "*/*")

	if s.Cookies != nil && len(s.Cookies) > 0 {
		for _, q := range s.Cookies {
			req.AddCookie(&q)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, http.ErrAbortHandler
	}

	// Читаем результат
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

func (s *UserApi) Authorize(email string, password string) (map[string]interface{}, error) {

	data := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	return s.d.Request("POST", "user/authorize", data)
}

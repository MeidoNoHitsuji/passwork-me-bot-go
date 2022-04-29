package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Request
// Собираем запрос и обрабатываем возможные ошибки.
// В результате возвращается ответ запроса.
///
func (s *Requester) Request(method string, subUrl string, data map[string]interface{}) (*http.Response, error) {
	client := http.Client{}

	// Назначаем параметры
	buffer := new(bytes.Buffer)
	params := url.Values{}
	for key, value := range data {
		params.Set(key, value.(string))
	}

	if len(s.Csrf) != 0 {
		params.Set("__csrf", s.Csrf)
	}

	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", URL, subUrl), buffer)

	if err != nil {
		return nil, err
	}

	// Добавляем необходимые заголовки
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "*/*")

	// Записываем куки в запрос
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
		return nil, errors.New("StatusCode: " + fmt.Sprintf("%v", resp.StatusCode))
	}

	return resp, nil
}

// RequestJson
// Надстройка над Request.
// Парсим параметр response в map.
///
func (s *Requester) RequestJson(method string, subUrl string, data map[string]interface{}) (interface{}, error) {
	resp, err := s.Request(method, subUrl, data)

	if err != nil {
		return resp, err
	}

	// Читаем результат
	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	if result["response"] != nil {
		switch result["response"].(type) {
		case bool:
			return nil, errors.New("ResponseMessage: " + fmt.Sprintf("%v", result["errorMessage"]))
		}
	}

	//Записываем куки, если их не было
	if s.Cookies == nil || len(s.Cookies) == 0 {
		for _, c := range resp.Cookies() {
			s.Cookies = append(s.Cookies, *c)
		}
	}

	return result["response"], nil
}

// RequestWithType
// Надстройка над RequestJson.
// Парсим результат в необходимый нам тип.
///
func (s *Requester) RequestWithType(method string, subUrl string, data map[string]interface{}, response any) error {
	result, err := s.RequestJson(method, subUrl, data)

	if err != nil {
		return err
	}

	jsonStr, err := json.Marshal(result)
	err = json.Unmarshal(jsonStr, &response)

	if err != nil {
		return err
	}

	return nil
}

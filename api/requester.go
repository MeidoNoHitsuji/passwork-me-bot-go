package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MeidoNoHitsuji/form"
	"net/http"
	"passwork-me-bot-go/config"
	"passwork-me-bot-go/helper"
	"strings"
)

// Requester
// Дефолтная структура запросника.
// Сохраняет куки при первом запросе.
// Также хранит в себе csrf, и использует, если он задан.
///
type Requester struct {
	Cookies   []http.Cookie
	Id        string
	Workspace string
	Csrf      string
}

// Request
// Собираем запрос и обрабатываем возможные ошибки.
// В результате возвращается ответ запроса.
///
func (s *Requester) Request(method string, subUrl string, data map[string]interface{}) (*http.Response, error) {
	client := http.Client{}

	if len(s.Csrf) != 0 {
		data["__csrf"] = s.Csrf
	}

	// Назначаем параметры
	params, _ := form.EncodeToValues(data)
	params = helper.TransferToParentheses(params)

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", config.URL, subUrl), strings.NewReader(params.Encode()))

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
		response := result["response"]
		switch response.(type) {
		case bool:
			if !response.(bool) {
				return nil, errors.New("ResponseMessage: " + fmt.Sprintf("%v", result["errorMessage"]))
			}
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

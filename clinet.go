package main

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
)

// Client
// Структура основной, где хранится основная логика.
///
type Client struct {
	api     *Requester
	Private string
}

// initCsrf
// Инициализируем csrf и сохраняем его в Requester.
///
func (s *Client) initCsrf() error {
	resp, err := s.api.Request("GET", "", map[string]interface{}{})

	if err != nil {
		panic("[GetCsrf] " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`__csrf = '([a-zA-Z0-9]+)';`)
	res := re.FindStringSubmatch(string(body))

	if len(res) > 0 {
		s.api.Csrf = res[1]
		return nil
	}

	return errors.New("csrf not Found")
}

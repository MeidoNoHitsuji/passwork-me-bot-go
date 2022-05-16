package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// Client
// Структура основной, где хранится основная логика.
///
type Client struct {
	api      Requester
	userApi  UserApi
	groupApi GroupApi
	Private  string
}

func (s *Client) init() error {
	s.api = Requester{}
	s.userApi = UserApi{
		api: &s.api,
	}
	s.groupApi = GroupApi{
		api: &s.api,
	}

	_, err := s.userApi.Authorize(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	if err != nil {
		return err
	}

	if err := s.initCsrf(); err != nil {
		return err
	}

	return nil
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

package client

import (
	"errors"
	"io/ioutil"
	"log"
	"passwork-me-bot-go/api"
	"regexp"
)

// Client
// Структура основной, где хранится основная логика.
///
type Client struct {
	Api         *api.Requester
	UserApi     *api.UserApi
	GroupApi    *api.GroupApi
	PasswordApi *api.PasswordApi
	Private     string
}

func New(email string, password string) *Client {
	requester := api.Requester{}
	client := Client{
		Api: &requester,
	}

	client.UserApi = &api.UserApi{
		Api: client.Api,
	}

	client.GroupApi = &api.GroupApi{
		Api: client.Api,
	}

	client.PasswordApi = &api.PasswordApi{
		Api: client.Api,
	}

	if err := client.UserApi.Authorize(email, password); err != nil {
		panic("[NewClient]" + err.Error())
	}

	if err := client.initCsrf(); err != nil {
		panic("[NewClient]" + err.Error())
	}

	info := client.UserApi.GetInfo()

	client.Api.Id = info.ID
	client.Api.Workspace = info.DefaultWorkspace

	return &client
}

// initCsrf
// Инициализируем csrf и сохраняем его в Requester.
///
func (s *Client) initCsrf() error {
	resp, err := s.Api.Request("GET", "", map[string]interface{}{})

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
		s.Api.Csrf = res[1]
		return nil
	}

	return errors.New("csrf not Found")
}

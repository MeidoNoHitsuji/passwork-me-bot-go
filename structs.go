package main

import "net/http"

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

// UserApi
// Запросник для части пользователя ("/user/{any}").
///
type UserApi struct {
	api *Requester
}

// AuthorizeStruct
// Структура ответа авторизации пользователя.
///
type AuthorizeStruct struct {
	Id           string
	Redirect     string
	SetWorkspace string
}

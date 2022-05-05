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

// GroupApi
// Запросник для части групп ("/group/{any}").
///
type GroupApi struct {
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

// UserInfoStruct
// Полная информация о пользователе
///
type UserInfoStruct struct {
	LastLogin int `json:"lastLogin" gorm:"column:lastLogin"`
	Keys      struct {
		Public         string `json:"public" gorm:"column:public"`
		PrivateCrypted string `json:"privateCrypted" gorm:"column:privateCrypted"`
	} `json:"keys" gorm:"column:keys"`
	Admin                      bool   `json:"admin" gorm:"column:admin"`
	EmailApproved              bool   `json:"emailApproved" gorm:"column:emailApproved"`
	Language                   string `json:"language" gorm:"column:language"`
	CouponsCount               int    `json:"couponsCount" gorm:"column:couponsCount"`
	Demo                       bool   `json:"demo" gorm:"column:demo"`
	ShowNewFeatures            bool   `json:"showNewFeatures" gorm:"column:showNewFeatures"`
	CountryCode                string `json:"countryCode" gorm:"column:countryCode"`
	ChangeCryptoPassword       bool   `json:"changeCryptoPassword" gorm:"column:changeCryptoPassword"`
	DefaultWorkspace           string `json:"defaultWorkspace" gorm:"column:defaultWorkspace"`
	CryptoNotificationDisabled bool   `json:"cryptoNotificationDisabled" gorm:"column:cryptoNotificationDisabled"`
	HistoryLog                 []struct {
		Date      string `json:"date" gorm:"column:date"`
		Agent     string `json:"agent" gorm:"column:agent"`
		Current   bool   `json:"current" gorm:"column:current"`
		Ip        string `json:"ip" gorm:"column:ip"`
		SessionID string `json:"sessionId" gorm:"column:sessionId"`
	} `json:"historyLog" gorm:"column:historyLog"`
	ID                   string      `json:"id" gorm:"column:id"`
	Email                string      `json:"email" gorm:"column:email"`
	LastUse              int         `json:"lastUse" gorm:"column:lastUse"`
	RegisteredAt         int         `json:"registered_at" gorm:"column:registered_at"`
	Active               bool        `json:"active" gorm:"column:active"`
	Avatar               interface{} `json:"avatar" gorm:"column:avatar"`
	LastNewsSeen         int         `json:"lastNewsSeen" gorm:"column:lastNewsSeen"`
	ShowFreeWorkspace    bool        `json:"showFreeWorkspace" gorm:"column:showFreeWorkspace"`
	UsersInFreeWorkspace int         `json:"usersInFreeWorkspace" gorm:"column:usersInFreeWorkspace"`
	Fullname             string      `json:"fullname" gorm:"column:fullname"`
	Bitrix24             interface{} `json:"bitrix24" gorm:"column:bitrix24"`
	HashAlgorithm        string      `json:"hashAlgorithm" gorm:"column:hashAlgorithm"`
}

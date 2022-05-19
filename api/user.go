package api

import (
	"crypto/sha256"
	"fmt"
)

// UserApi
// Запросник для части пользователя ("/user/{any}").
///
type UserApi struct {
	Api *Requester
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

// Authorize
// Авторизуемся по email и password.
// В результате получаем рабочее место и сессию в куках.
///
func (s *UserApi) Authorize(email string, password string) error {

	data := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	if _, err := s.Api.RequestJson("POST", "user/authorize", data); err != nil {
		panic("[AuthorizeException] " + err.Error())
	}

	return nil
}

func (s *UserApi) GetUserColors() map[string]interface{} {
	resp, err := s.Api.RequestJson("POST", "user/getUserColors", map[string]interface{}{})

	if err != nil {
		panic("[GetUserColors] " + err.Error())
	}

	return resp.(map[string]interface{})
}

func (s *UserApi) GetInfo() UserInfoStruct {
	var resp UserInfoStruct

	if err := s.Api.RequestWithType("POST", "user/getInfo", map[string]interface{}{}, &resp); err != nil {
		panic("[GetInfo] " + err.Error())
	}

	return resp
}

func (s *UserApi) CheckMasterHash(masterKey string) bool {

	sum := sha256.Sum256([]byte(masterKey))

	data := map[string]interface{}{
		"sha256": fmt.Sprintf("sha256:%x", sum),
	}

	resp, err := s.Api.RequestJson("POST", "user/checkMasterHash", data)

	if err != nil {
		panic("[CheckMasterHash] " + err.Error())
	}

	return resp.(map[string]interface{})["result"].(bool)
}

func (s *UserApi) GetPrivateKey() map[string]interface{} {
	data := map[string]interface{}{
		"workspaceId": s.Api.Workspace,
	}

	resp, err := s.Api.RequestJson("POST", "user/getPrivateKey", data)

	if err != nil {
		panic("[GetPrivateKey] " + err.Error())
	}

	return resp.(map[string]interface{})
}

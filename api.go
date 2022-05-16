package main

import (
	"crypto/sha256"
	"fmt"
)

// Authorize
// Авторизуемся по email и password.
// В результате получаем рабочее место и сессию в куках.
///
func (s *UserApi) Authorize(email string, password string) (AuthorizeStruct, error) {

	data := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	var resp AuthorizeStruct

	if err := s.api.RequestWithType("POST", "user/authorize", data, &resp); err != nil {
		panic("[AuthorizeException] " + err.Error())
	}

	s.api.Id = resp.Id
	s.api.Workspace = resp.SetWorkspace

	return resp, nil
}

func (s *UserApi) GetUserColors() map[string]interface{} {
	resp, err := s.api.RequestJson("POST", "user/getUserColors", map[string]interface{}{})

	if err != nil {
		panic("[GetUserColors] " + err.Error())
	}

	return resp.(map[string]interface{})
}

func (s *UserApi) GetInfo() UserInfoStruct {
	var resp UserInfoStruct

	if err := s.api.RequestWithType("POST", "user/getInfo", map[string]interface{}{}, &resp); err != nil {
		panic("[GetInfo] " + err.Error())
	}

	return resp
}

func (s *UserApi) CheckMasterHash(masterKey string) bool {

	sum := sha256.Sum256([]byte(masterKey))

	data := map[string]interface{}{
		"sha256": fmt.Sprintf("sha256:%x", sum),
	}

	resp, err := s.api.RequestJson("POST", "user/checkMasterHash", data)

	if err != nil {
		panic("[CheckMasterHash] " + err.Error())
	}

	return resp.(map[string]interface{})["result"].(bool)
}

func (s *UserApi) GetPrivateKey() map[string]interface{} {
	data := map[string]interface{}{
		"workspaceId": s.api.Workspace,
	}

	resp, err := s.api.RequestJson("POST", "user/getPrivateKey", data)

	if err != nil {
		panic("[GetPrivateKey] " + err.Error())
	}

	return resp.(map[string]interface{})
}

func (s *GroupApi) Get() []GroupInfo {

	var resp []GroupInfo

	data := map[string]interface{}{
		"workspaceId": s.api.Workspace,
	}

	if err := s.api.RequestWithType("POST", "groups/get", data, &resp); err != nil {
		panic("[GroupGet] " + err.Error())
	}

	return resp
}

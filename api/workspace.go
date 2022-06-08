package api

import (
	"crypto/sha256"
	"encoding/hex"
	"passwork-me-bot-go/helper"
)

// WorkspaceApi
// Запросник для части пользователя ("/workspace/{any}").
///
type WorkspaceApi struct {
	Api *Requester
}

type GetUsers struct {
	Users   []User        `json:"users"`
	Pending []interface{} `json:"pending"`
}

type User struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Avatar     string      `json:"avatar"`
	ExpireDate interface{} `json:"expireDate"`
	Active     bool        `json:"active"`
	Status     string      `json:"status"`
	PublicKey  string      `json:"publicKey"`
}

func (s WorkspaceApi) GetUsers() []User {
	var resp GetUsers

	data := map[string]interface{}{
		"workspaceId": s.Api.Workspace,
	}

	if err := s.Api.RequestWithType("POST", "workspace/getUsers", data, &resp); err != nil {
		panic("[GroupGetFullData] " + err.Error())
	}

	return resp.Users
}

func (s WorkspaceApi) AddRsaEncryptedFolderToManyUsers(users []UserWithPublicKey, groupId string, groupKey string, folderId string) bool {

	var resp bool

	type EncryptUserType struct {
		Id                          string `form:"id"`
		GroupPasswordCryptedWithRsa string `form:"groupPasswordCryptedWithRsa"`
		PublicKeyHash               string `form:"publicKeyHash"`
		Access                      string `form:"access"`
		FolderAccess                int    `form:"folderAccess"`
	}

	var Users []EncryptUserType

	for _, user := range users {
		data := new(EncryptUserType)

		data.Id = user.Id
		data.GroupPasswordCryptedWithRsa = helper.RsaEncrypt(groupKey, user.PublicKey)
		publicKeyHash := sha256.Sum256([]byte(user.PublicKey))
		data.PublicKeyHash = "sha256:" + hex.EncodeToString(publicKeyHash[:])
		data.Access = "listing"
		data.FolderAccess = user.FolderAccess

		Users = append(Users, *data)
	}

	data := map[string]interface{}{
		"users":       Users,
		"workspaceId": s.Api.Workspace,
		"groupId":     groupId,
		"folderId":    folderId,
	}

	if err := s.Api.RequestWithType("POST", "workspace/addRsaEncryptedGroupToManyUsers", data, &resp); err != nil {
		panic("[AddRsaEncryptedGroupToManyUsers] " + err.Error())
	}

	return resp
}

func (s WorkspaceApi) AddRsaEncryptedGroupToManyUsers(users []UserWithPublicKey, groupId string, groupKey string) bool {
	return s.AddRsaEncryptedFolderToManyUsers(users, groupId, groupKey, "")
}

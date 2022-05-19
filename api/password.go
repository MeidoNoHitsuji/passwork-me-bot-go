package api

// PasswordApi
// Запросник для части групп ("/password/{any}").
///**
type PasswordApi struct {
	Api *Requester
}

type PasswordRecent struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Group       struct {
		GroupId         string `json:"groupId"`
		Name            string `json:"name"`
		PasswordCrypted string `json:"passwordCrypted"`
		Access          string `json:"access"`
		EncryptedWith   string `json:"encryptedWith"`
		WorkspaceId     string `json:"workspaceId"`
	} `json:"group"`
	CategoryName string `json:"categoryName"`
	GroupId      string `json:"groupId"`
	Color        int    `json:"color"`
	Attachments  []struct {
		Id           string `json:"id"`
		EncryptedKey string `json:"encryptedKey"`
		Name         string `json:"name"`
	} `json:"attachments"`
	Id          string   `json:"id"`
	Tags        []string `json:"tags"`
	WorkspaceId string   `json:"workspaceId"`
}

type PasswordInfo struct {
	Name            string `json:"name"`
	Login           string `json:"login"`
	CryptedPassword string `json:"cryptedPassword"`
	Url             string `json:"url"`
	Description     string `json:"description"`
	GroupId         string `json:"groupId"`
	Color           int    `json:"color"`
	Access          int    `json:"access"`
	Adminable       bool   `json:"adminable"`
	Attachments     []struct {
		Id           string `json:"id"`
		EncryptedKey string `json:"encryptedKey"`
		Name         string `json:"name"`
	} `json:"attachments"`
	Id          string   `json:"id"`
	Tags        []string `json:"tags"`
	WorkspaceId string   `json:"workspaceId"`
}

func (s PasswordApi) GetRecentPasswords() []PasswordRecent {
	var resp []PasswordRecent

	data := map[string]interface{}{
		"workspaceId": s.Api.Workspace,
	}

	if err := s.Api.RequestWithType("POST", "password/getRecentPasswords", data, &resp); err != nil {
		panic("[PasswordGetRecentPasswords] " + err.Error())
	}

	return resp
}

func (s PasswordApi) Get(id string) PasswordInfo {
	var resp PasswordInfo

	data := map[string]interface{}{
		"workspaceId": s.Api.Workspace,
		"id":          id,
	}

	if err := s.Api.RequestWithType("POST", "password/get", data, &resp); err != nil {
		panic("[PasswordGet] " + err.Error())
	}

	return resp
}

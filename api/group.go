package api

// GroupApi
// Запросник для части групп ("/group/{any}").
///
type GroupApi struct {
	Api *Requester
}

type GroupInfo struct {
	Shared        bool   `json:"shared" gorm:"column:shared"`
	Access        string `json:"access" gorm:"column:access"`
	EncryptedWith string `json:"encryptedWith" gorm:"column:encryptedWith"`
	GroupID       string `json:"groupId" gorm:"column:groupId"`
	Name          string `json:"name" gorm:"column:name"`
	Tree          []struct {
		Lvl               int    `json:"lvl" gorm:"column:lvl"`
		GroupID           string `json:"groupId" gorm:"column:groupId"`
		Name              string `json:"name" gorm:"column:name"`
		PasswordsCount    int    `json:"passwordsCount" gorm:"column:passwordsCount"`
		ID                string `json:"id" gorm:"column:id"`
		ParentID          string `json:"parentId" gorm:"column:parentId"`
		CurrentUserAccess int    `json:"currentUserAccess" gorm:"column:currentUserAccess"`
	} `json:"tree" gorm:"column:tree"`
	ID              string `json:"id" gorm:"column:id"`
	PasswordCrypted string `json:"passwordCrypted" gorm:"column:passwordCrypted"`
}

func (s *GroupApi) Get() []GroupInfo {

	var resp []GroupInfo

	data := map[string]interface{}{
		"workspaceId": s.Api.Workspace,
	}

	if err := s.Api.RequestWithType("POST", "groups/get", data, &resp); err != nil {
		panic("[GroupGet] " + err.Error())
	}

	return resp
}

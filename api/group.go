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

type GroupFullData struct {
	Breadcrumbs []struct {
		Name    string `json:"name"`
		GroupId string `json:"groupId"`
		Id      string `json:"id"`
	} `json:"breadcrumbs"`
	Group struct {
		GroupId         string `json:"groupId"`
		Name            string `json:"name"`
		PasswordCrypted string `json:"passwordCrypted"`
		Access          string `json:"access"`
		Salt            string `json:"salt"`
		EncryptedWith   string `json:"encryptedWith"`
		WorkspaceId     string `json:"workspaceId"`
		Id              string `json:"id"`
	} `json:"group"`
	Category struct {
		GroupId     string        `json:"groupId"`
		Name        string        `json:"name"`
		ParentId    string        `json:"parentId"`
		Ancestors   []string      `json:"ancestors"`
		PasswordIds []interface{} `json:"passwordIds"`
		Id          string        `json:"id"`
		Level       int           `json:"level"`
	} `json:"category"`
	CategoryAccess         int    `json:"categoryAccess"`
	GroupAccess            string `json:"groupAccess"`
	PasswordsAndCategories []struct {
		GroupId   string   `json:"groupId"`
		Name      string   `json:"name"`
		ParentId  string   `json:"parentId"`
		Ancestors []string `json:"ancestors"`
		Id        string   `json:"id"`
		Level     int      `json:"level"`
		Type      string   `json:"_type"`
	} `json:"passwordsAndCategories"`
	CountUsers  int `json:"countUsers"`
	CountAdmins int `json:"countAdmins"`
}

func (s *GroupApi) GetFullData(id string) GroupFullData {
	return s.GetFullDataWithCategory(id, "")
}

func (s *GroupApi) GetFullDataWithCategory(id string, categoryId string) GroupFullData {

	var resp GroupFullData

	data := map[string]interface{}{
		"workspaceId": s.Api.Workspace,
		"id":          id,
		"listing":     true,
	}
	if len(categoryId) > 0 {
		data["categoryId"] = categoryId
	}

	if err := s.Api.RequestWithType("POST", "groups/getFullData", data, &resp); err != nil {
		panic("[GroupGetFullData] " + err.Error())
	}

	return resp
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

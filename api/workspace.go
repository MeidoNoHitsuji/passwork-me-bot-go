package api

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

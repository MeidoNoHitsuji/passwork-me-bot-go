package main

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

func (s *UserApi) GetInfo() map[string]interface{} {
	resp, err := s.api.RequestJson("POST", "user/getInfo", map[string]interface{}{})

	if err != nil {
		panic("[GetInfo] " + err.Error())
	}

	return resp.(map[string]interface{})
}

func (s *UserApi) CheckMasterHash(masterKey string) bool {
	return true
}

func (s *UserApi) GetPrivateKey() {

}

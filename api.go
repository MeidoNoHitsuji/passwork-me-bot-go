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

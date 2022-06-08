package client

import (
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"passwork-me-bot-go/api"
	"passwork-me-bot-go/helper"
	"passwork-me-bot-go/models"
	"regexp"
	"sort"
)

// Client
// Структура основной, где хранится основная логика.
///
type Client struct {
	Api          *api.Requester
	UserApi      *api.UserApi
	GroupApi     *api.GroupApi
	PasswordApi  *api.PasswordApi
	WorkspaceApi *api.WorkspaceApi
	Private      string
}

func New(email string, password string) *Client {
	requester := api.Requester{}
	client := Client{
		Api: &requester,
	}

	client.UserApi = &api.UserApi{
		Api: client.Api,
	}

	client.GroupApi = &api.GroupApi{
		Api: client.Api,
	}

	client.PasswordApi = &api.PasswordApi{
		Api: client.Api,
	}

	client.WorkspaceApi = &api.WorkspaceApi{
		Api: client.Api,
	}

	if err := client.UserApi.Authorize(email, password); err != nil {
		panic("[NewClient] " + err.Error())
	}

	if err := client.initCsrf(); err != nil {
		panic("[NewClient] " + err.Error())
	}

	info := client.UserApi.GetInfo()

	client.Api.Id = info.ID
	client.Api.Workspace = info.DefaultWorkspace

	return &client
}

// initCsrf
// Инициализируем csrf и сохраняем его в Requester.
///
func (s *Client) initCsrf() error {
	resp, err := s.Api.Request("GET", "", map[string]interface{}{})

	if err != nil {
		panic("[GetCsrf] " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`__csrf = '([a-zA-Z0-9]+)';`)
	res := re.FindStringSubmatch(string(body))

	if len(res) > 0 {
		s.Api.Csrf = res[1]
		return nil
	}

	return errors.New("csrf not Found")
}

func (s *Client) updateUsers(db *gorm.DB) {
	users := s.WorkspaceApi.GetUsers()

	for _, user := range users {
		db.FirstOrCreate(&models.User{
			ID: user.Id,
		}, models.User{
			Name:  user.Name,
			Email: user.Email,
		})
	}
}

func (s *Client) updateGroups(db *gorm.DB) {
	groups := s.GroupApi.Get()

	for _, vault := range groups {
		db.FirstOrCreate(&models.Group{
			ID: vault.ID,
		}, models.Group{
			Name: vault.Name,
		})

		if len(vault.Tree) > 0 {
			sort.SliceStable(vault.Tree, func(i int, j int) bool {
				return vault.Tree[i].Lvl < vault.Tree[j].Lvl
			})

			for _, group := range vault.Tree {
				ParentId := group.ParentID
				if group.Lvl == 0 {
					ParentId = vault.ID
				}

				db.FirstOrCreate(&models.Group{
					ID: group.ID,
				}, models.Group{
					Name:     group.Name,
					ParentId: ParentId,
				})
			}
		}
	}

	var listIds []string

	for _, vault := range groups {
		listIds = append(listIds, vault.ID)
		if len(vault.Tree) > 0 {
			for _, group := range vault.Tree {
				listIds = append(listIds, group.ID)
			}
		}
	}

	listIds = helper.Unique(listIds)
	db.Where("id not in (?)", listIds).Delete(&models.Group{})
}

func (s Client) AddUsersInCategory(users []api.UserWithPublicKey, groupId string, categoryId string) bool {
	var groupData api.GroupFullData

	if categoryId != "" {
		groupData = s.GroupApi.GetFullDataWithCategory(groupId, categoryId)
	} else {
		groupData = s.GroupApi.GetFullData(groupId)
	}

	groupPassword := groupData.Group.DecryptPassword()

	return s.WorkspaceApi.AddRsaEncryptedFolderToManyUsers(users, groupId, groupPassword, categoryId)
}

func (s Client) AddUsersInGroup(users []api.UserWithPublicKey, groupId string) bool {
	return s.AddUsersInCategory(users, groupId, "")
}

func (s *Client) UpdatePermissions(db *gorm.DB) {
	s.updateUsers(db)
	s.updateGroups(db)
	//SelectRoles
	//Update Permissions By Roles
}

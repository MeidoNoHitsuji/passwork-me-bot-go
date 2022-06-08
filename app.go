package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	//c := client.New(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	//user := c.UserApi.GetInfo()
	//fmt.Println(user)

	//groups := c.GroupApi.Get()
	//
	//groupId := "62a086ada6fe8016ee05fc54"
	////groupId = "62a086ada6fe80165d0ac4c0"
	//groupData := c.GroupApi.GetFullData(groupId)
	//
	//groupKey, _ := aes.Decrypt(
	//	base32.Decode(groupData.Group.PasswordCrypted, true),
	//	os.Getenv("MASTER_KEY"),
	//)
	//fmt.Println(groupKey)
	//users := c.GroupApi.GetWorkspaceUsersNotInGroup(groupId)
	//result := c.WorkspaceApi.AddRsaEncryptedGroupToManyUsers(users, groupId, groupKey, "62a086ada6fe80165d0ac4c0")

	//fmt.Println(result)

	//db := database.New()
	//c.UpdatePermissions(db)

	//groupKey, _ := aes.Decrypt(
	//	base32.Decode("amt4cwv48xb6pp1h5xa7jn358h964mjpch4qev3c85t4cxbmcdhn4ca3exh54mj2a9264wjuc97m8gk8cdx54n2t8mt64y32d9x66vafahgprt2kd1j6crtmb1gm2caaf5338pbncwyku", true),
	//	"Rdevv781!",
	//)
	//
	//pubPEM := "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLlSaR+Tcmz30SBWP3QcLRiYRZ\n0rOLHnif+5kdKb81bZfsTcsfRJO3lqlK/J4hd5pM+NYypn8DTsu6hG7jJ/wBDub+\n2ukqFOyqvH1IJZhSJuvHNh4SlGAfa7Xdp8sUe+1dfWe5WGWkXvpz2txCNbqguDYH\n1ewpwVexPfWxk4Y0nwIDAQAB\n-----END PUBLIC KEY-----"
	//ciphertext := helper.RsaEncrypt(groupKey, pubPEM)
	//
	//fmt.Println(ciphertext)

	//fmt.Println(u.Encode())
	//database.RunMigrateScripts(db)

	//database.New()
}

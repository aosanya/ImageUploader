package utilities

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type UserData struct {
	Images []string
}

func userFilesMeta(userId string) string {
	return "../userfiles/" + userId + "/userdata.json"
}

func SaveUserData(userId string, fileName string) {
	userData := GetUserData(userId)
	userData.Images = append(userData.Images, fileName)
	// data := UserData{
	// 	Images: []string{fileName},
	// }

	file, _ := json.MarshalIndent(userData, "", " ")
	_ = ioutil.WriteFile(userFilesMeta(userId), file, 0644)
}

func fileExists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func GetUserData(userId string) UserData {
	userDataPath := userFilesMeta(userId)
	if !fileExists(userDataPath) {
		return UserData{}
	}

	jsonFile, err := os.Open(userFilesMeta(userId))
	Check(err)
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var userData UserData

	json.Unmarshal(byteValue, &userData)

	return userData
}

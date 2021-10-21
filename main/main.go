package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"codevald.com/utilities"
)

func main() {
	setupRoutes()
	fmt.Println("Starting server on 11111")
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type UserData struct {
	Images []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running AP v1\n"))
}

func userFilesPath(userId string) string {
	return "userfiles/" + userId
}

func userFilesMeta(userId string) string {
	return "userfiles/" + userId + "/userdata.json"
}

func saveFileData(userId string, fileName string) {
	data := UserData{
		Images: []string{fileName},
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(userFilesMeta(userId), file, 0644)
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("userId")
	err := os.MkdirAll(userFilesPath(userId), 0755)
	check(err)

	fmt.Println(userId)
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	extension := filepath.Ext(handler.Filename)

	savedFile := utilities.UploadFile(userFilesPath(userId), file, extension)
	saveFileData(userId, savedFile)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/upload", uploadFileHandler)
}

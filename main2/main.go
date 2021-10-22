package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/images/" {
		fetchImagesHandler(w, r)
		return
	}

	// if r.URL.Path != "/" {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("Asset not found\n"))
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running AP v1\n"))
}

func userFilesPath(userId string) string {
	return "../userfiles/" + userId
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
	extension := filepath.Ext(handler.Filename)

	savedFile := utilities.UploadFile(userFilesPath(userId), file, extension)
	utilities.SaveUserData(userId, savedFile)
}

func fetchImagesHandler(w http.ResponseWriter, r *http.Request) {
	urlComponents := strings.Split(r.URL.Path, "/")
	response, err := getJsonResponse(urlComponents[2])
	check(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getJsonResponse(userId string) ([]byte, error) {
	data := utilities.GetUserData(userId)
	return json.MarshalIndent(data, "", "  ")
}

func setupRoutes() {
	http.HandleFunc("/images/{userId}", fetchImagesHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/upload", uploadFileHandler)
}

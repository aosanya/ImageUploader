package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"codevald.com/utilities"
)

func main() {
	setupRoutes()
	fmt.Println("Starting server")
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/images/") {
		fetchImagesHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running AP v1.1\n"))
}

func userFilesPath(userId string) string {
	return "../userfiles/" + userId
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("userId")
	err := os.MkdirAll(userFilesPath(userId), 0755)
	utilities.Check(err)

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
	returnSuccess(w, r, "Visit '/images/"+userId+"' to view your uploaded images")
}

func returnSuccess(w http.ResponseWriter, r *http.Request, message string) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Success"
	resp["info"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func fetchImagesHandler(w http.ResponseWriter, r *http.Request) {
	urlComponents := strings.Split(r.URL.Path, "/")
	fetchImagesWriter(w, urlComponents[2])
}

func fetchImagesWriter(w http.ResponseWriter, userId string) {
	response, err := getJsonResponse(userId)
	utilities.Check(err)
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

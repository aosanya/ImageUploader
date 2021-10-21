package utilities

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UploadFile(path string, file multipart.File, extension string) string {
	tempFile, err := ioutil.TempFile(path, "upload-*."+extension)
	check(err)
	defer tempFile.Close()
	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	return tempFile.Name()
}

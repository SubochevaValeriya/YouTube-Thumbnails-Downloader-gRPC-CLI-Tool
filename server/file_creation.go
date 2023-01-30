package server

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

// CreateFolder creates directory for thumbnails
func createFolder(dir string) error {

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.New("can't create folder")
	}
	return nil
}

// saveThumbnail create thumbnail file and write thumbnail image using video name for naming
func saveThumbnail(name string, resp *http.Response) error {

	createdFile, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return errors.New("can't create file")
	}

	_, err = io.Copy(createdFile, resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

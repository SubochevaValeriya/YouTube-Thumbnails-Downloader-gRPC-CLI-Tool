package internal

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

// CreateFolder creates directory for thumbnails
func CreateFolder() error {

	err := os.MkdirAll(viper.GetString("directory_name"), os.ModePerm)
	if err != nil {
		return errors.New("can't create folder")
	}
	return nil
}

// SaveThumbnail create thumbnail file and write thumbnail image using video name for naming
func SaveThumbnail(name string, image []byte) error {
	fileName := filepath.Base(fmt.Sprintf("%s.jpg", name))
	err := os.WriteFile(filepath.Join(viper.GetString("directory_name"), fileName), image, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//func SaveThumbnail(name string, resp *http.Response) error {
//	fileName := filepath.Base(fmt.Sprintf("%s.jpg", name))
//	//createdFile, err := os.Create(filepath.Join(viper.GetString("directory_name"), fileName))
//	if err != nil {
//		log.Println(err)
//		return errors.New("can't create file")
//	}
//	b, err := io.ReadAll(resp.Body)
//	os.WriteFile(filepath.Join(viper.GetString("directory_name"), fileName), b, 0644)
//	_, err = io.Copy(createdFile, resp.Body)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	return nil
//}

// DeleteFolder removes folder and all thumbnails in it
func DeleteFolder() error {

	return os.RemoveAll(viper.GetString("directory_name"))
}

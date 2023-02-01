package main

import (
	"errors"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/client/internal"
	grpcYoutubeThumbnails "github.com/SubochevaValeriya/grpcYoutubeThumbnails/proto"
	"github.com/SubochevaValeriya/grpcYoutubeThumbnails/server/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"os"
	"testing"
)

type TestCaseDownloadingThumbnails struct {
	URL           string
	Name          string
	ExpectedError error
}

func TestDownloadingThumbnails(t *testing.T) {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	s := server{}

	dbConfig := repository.MongoConfig{
		Host:            viper.GetString("db.host"),
		Port:            viper.GetString("db.port"),
		DefaultDatabase: viper.GetString("db.test_database"),
		Collection:      viper.GetString("db.collection"),
	}

	logrus.Println("Connecting to MongoDB")

	_, err := repository.MongoConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	testCases := []TestCaseDownloadingThumbnails{
		{URL: "https://www.youtube.com/watch?v=AbLdVT6Lvdc", Name: "thumbnails-test/Kore klip~ Mutlu Sonsuz (Manhole ).jpg", ExpectedError: nil},
		{URL: "NotLink", Name: "", ExpectedError: errors.New("incorrect format for URL")},
		{URL: "", Name: "", ExpectedError: errors.New("incorrect format for URL")},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.URL, func(t *testing.T) {
			req := grpcYoutubeThumbnails.DownloadThumbnailLinkRequest{URL: cse.URL}
			_, err := s.DownloadThumbnail(context.Background(), &req)
			if cse.ExpectedError == nil {
				if err != nil {
					t.Errorf("should not be error")
				} else {
					if _, err := os.Stat(cse.Name); os.IsNotExist(err) {
						t.Errorf("file not saved")
					}
				}
			} else if err == nil {
				t.Errorf("should be an error")
			}
		})
	}

	// clear DB and folder
	repository.DropCollection(context.Background())
	internal.DeleteFolder()
	logrus.Printf("DB and folder are cleared")
}

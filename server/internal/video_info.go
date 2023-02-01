package internal

import (
	"errors"
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type VideoItem struct {
	ID            string `bson:"_id,omitempty"`
	VideoID       string `bson:"video_id"`
	Name          string `bson:"name"`
	ThumbnailLink string `bson:"thumbnail_link"`
}

var incorrectFormatError = errors.New("incorrect URL format")

// FindVideoID function finds ID of YouTube video in the link
func (video *VideoItem) FindVideoID(URL string) error {
	// regular expression to find YouTube video ID:

	matched, err := regexp.MatchString(`(?:\\/|%3D|v=|vi=)([0-9A-z-_]{11})(?:[%#?&]|$)`, URL)
	if err != nil || !matched {
		logrus.WithFields(
			logrus.Fields{
				"package":  "internal",
				"function": "FindVideoID",
				"error":    err,
				"data":     URL,
			}).Errorf("Incorrect format for URL: %s", URL)
		return incorrectFormatError
	}
	re, err := regexp.Compile(`([0-9A-z-_]{11})`)

	if err != nil || len(re.FindString(URL)) != 11 {
		logrus.WithFields(
			logrus.Fields{
				"package":  "internal",
				"function": "FindVideoID",
				"error":    err,
				"data":     URL,
			}).Errorf("Incorrect format for URL: %s", URL)

		return incorrectFormatError
	}

	video.VideoID = re.FindString(URL)
	return nil
}

// FindTitle functions finds title of requested video
func (video *VideoItem) FindTitle(URL string) error {
	// Make HTTP GET request
	response, err := http.Get(URL)
	if err != nil {
		log.Printf("Can't find title for URL: %s", URL)
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	pageContent := string(data)

	// Find a substring
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}

	titleStartIndex += len("<title>") // offset the index by the len of "<title>"

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}

	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	video.Name = strings.TrimRight(string(pageTitle), "  - YouTube")
	return nil
}

// FindThumbnailLink functions finds link for downloading thumbnail
func (video *VideoItem) FindThumbnailLink() error {
	const (
		address       = "https://i.ytimg.com/vi/"
		maxResolution = "/maxresdefault.jpg"
		HQResolution  = "/hqdefault.jpg"
	)

	_, err := http.Get(address + video.VideoID + maxResolution)

	if err != nil {
		log.Printf("Can't find picture in maximum resolution...")

		_, err = http.Get(address + video.VideoID + HQResolution)
		if err != nil {
			log.Printf("Error while trying to get thumbnail: %s", err)
			return err
		}
		video.ThumbnailLink = address + video.VideoID + maxResolution
	} else {
		_, err = http.Get(address + video.VideoID + HQResolution)
		video.ThumbnailLink = address + video.VideoID + HQResolution
	}

	return nil
}

func (video *VideoItem) GetImage() (*http.Response, error) {
	res, err := http.Get(video.ThumbnailLink)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"package":  "internal",
				"function": "FindVideoID",
				"error":    err,
				"data":     video.ThumbnailLink,
			}).Errorf("Can't get thumbnail image by URL: %s", video.ThumbnailLink)
	}
	return res, err
}

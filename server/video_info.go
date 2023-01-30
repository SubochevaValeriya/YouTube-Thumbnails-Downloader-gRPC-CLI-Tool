package server

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var incorrectFormatError = errors.New("incorrect URL format")

func (video *videoItem) findIndex(URL string) error {
	// regular expression to find YouTube video index:
	re, err := regexp.Compile(`([0-9A-z-_]{11})`)

	if err != nil {
		log.Printf("Incorrect format for URL: %s", URL)
		return incorrectFormatError
	}

	video.VideoID = re.FindString(URL)
	return nil
}

func (video *videoItem) findTitle(URL string) error {
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

	video.Name = string(pageTitle)
	return nil
}

func (video *videoItem) findThumbnailLink() error {
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
	}

	return nil
}

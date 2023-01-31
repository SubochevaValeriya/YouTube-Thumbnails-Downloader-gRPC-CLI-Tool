package internal

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	err := CreateFolder()
	if err != nil {
		t.Fatal()
	}

	if _, err := os.Stat("thumbnails"); os.IsNotExist(err) {
		fmt.Println(err)
		t.Fatal()
	}
}

func TestSaveThumbnail(t *testing.T) {
	resp, err := http.Get("https://i.ytimg.com/vi/6o1m7ofjCCY/hqdefault.jpg")
	if err != nil {
		t.Fatal()
	}
	err = SaveThumbnail("one.jpg", resp)
	if err != nil {
		t.Fatal()
	}

	if _, err := os.Stat("one"); os.IsNotExist(err) {
		fmt.Println(err)
		t.Fatal()
	}
}
